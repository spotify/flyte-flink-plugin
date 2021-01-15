package flink

import (
	corev1 "k8s.io/api/core/v1"

	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	flinkOp "github.com/regadas/flink-on-k8s-operator/api/v1beta1"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	cacheVolumes      = []corev1.Volume{{Name: "cache-volume"}}
	cacheVolumeMounts = []corev1.VolumeMount{{Name: "cache-volume", MountPath: "/cache"}}
)

func BuildJobManagerResource(flinkProperties FlinkProperties, annotations Annotations, labels Labels) flinkOp.JobManagerSpec {
	jobManagerCores := flinkProperties.GetResourceQuantity("kubernetes.jobmanager.cores")
	jobManagerMemory := flinkProperties.GetResourceQuantity("kubernetes.jobmanager.memory")

	return flinkOp.JobManagerSpec{
		PodAnnotations: annotations,
		PodLabels:      labels,
		Resources: corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:    jobManagerCores,
				corev1.ResourceMemory: jobManagerMemory,
			},
		},
		Volumes:      cacheVolumes,
		VolumeMounts: cacheVolumeMounts,
	}
}

func BuildTaskManagerResource(flinkProperties FlinkProperties, annotations Annotations, labels Labels) flinkOp.TaskManagerSpec {
	taskManagerCores := flinkProperties.GetResourceQuantity("kubernetes.taskmanager.cores")
	taskManagerMemory := flinkProperties.GetResourceQuantity("kubernetes.taskmanager.memory")
	taskManagerReplicas := flinkProperties.GetInt("kubernetes.taskmanager.replicas")

	return flinkOp.TaskManagerSpec{
		PodAnnotations: annotations,
		PodLabels:      labels,
		Replicas:       int32(taskManagerReplicas),
		Volumes:        cacheVolumes,
		VolumeMounts:   cacheVolumeMounts,
		Resources: corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:    taskManagerCores,
				corev1.ResourceMemory: taskManagerMemory,
			},
		},
	}
}

func BuildJobResource(taskManager flinkOp.TaskManagerSpec, flinkProperties FlinkProperties, flinkJob flinkIdl.FlinkJob) flinkOp.JobSpec {
	taskSlots := flinkProperties.GetInt("taskmanager.numberOfTaskSlots")
	parallelism := taskManager.Replicas * int32(taskSlots)

	return flinkOp.JobSpec{

		JarFile:      "/cache/job.jar",
		ClassName:    &flinkJob.MainClass,
		Args:         flinkJob.Args,
		Parallelism:  &parallelism,
		Volumes:      cacheVolumes,
		VolumeMounts: cacheVolumeMounts,
		CleanupPolicy: &flinkOp.CleanupPolicy{
			AfterJobSucceeds:  flinkOp.CleanupActionDeleteCluster,
			AfterJobFails:     flinkOp.CleanupActionDeleteCluster,
			AfterJobCancelled: flinkOp.CleanupActionDeleteCluster,
		},
		Resources: corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("1Gi"),
			},
		},
		InitContainers: []corev1.Container{
			{
				Name:    "gcs-downloader",
				Image:   "google/cloud-sdk",
				Command: []string{"gsutil"},
				Args: []string{
					"cp",
					flinkJob.JarFile,
					"/cache/job.jar",
				},
				Resources: corev1.ResourceRequirements{
					Limits: corev1.ResourceList{
						corev1.ResourceCPU:    resource.MustParse("1"),
						corev1.ResourceMemory: resource.MustParse("1Gi"),
					},
				},
			},
		},
	}
}

func BuildFlinkClusterResource(config *Config, flinkProperties FlinkProperties, annotations Annotations, labels Labels, jobManager flinkOp.JobManagerSpec, taskManager flinkOp.TaskManagerSpec, job flinkOp.JobSpec) flinkOp.FlinkCluster {
	return flinkOp.FlinkCluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindFlinkCluster,
			APIVersion: flinkOp.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Annotations: annotations,
			Labels:      labels,
		},
		Spec: flinkOp.FlinkClusterSpec{
			ServiceAccountName: &config.ServiceAccount,
			Image: flinkOp.ImageSpec{
				Name:       config.Image,
				PullPolicy: corev1.PullAlways,
			},
			JobManager:      jobManager,
			TaskManager:     taskManager,
			Job:             &job,
			FlinkProperties: flinkProperties,
		},
	}
}
