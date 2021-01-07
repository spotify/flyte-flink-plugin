package flink

import (
	"strconv"

	corev1 "k8s.io/api/core/v1"

	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	flinkOp "github.com/regadas/flink-on-k8s-operator/api/v1beta1"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func BuildJobManagerResource(flinkProperties FlinkProperties, annotations Annotations, labels Labels) flinkOp.JobManagerSpec {
	jobManagerCores := resource.MustParse(flinkProperties["kubernetes.jobmanager.cores"])
	jobManagerMemory := resource.MustParse(flinkProperties["kubernetes.jobmanager.memory"])

	return flinkOp.JobManagerSpec{
		PodAnnotations: annotations,
		PodLabels:      labels,
		Ports:          flinkOp.JobManagerPorts{UI: &jobManagerUIPort},
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
	taskManagerCores := resource.MustParse(flinkProperties["kubernetes.taskmanager.cores"])
	taskManagerMemory := resource.MustParse(flinkProperties["kubernetes.taskmanager.memory"])
	taskManagerReplicas, err := strconv.Atoi(flinkProperties["kubernetes.taskmanager.replicas"])
	if err != nil {
		taskManagerReplicas = defaultTaskManagerReplicas
	}

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
	taskSlots, err := strconv.Atoi(flinkProperties["taskmanager.numberOfTaskSlots"])
	if err != nil {
		taskSlots = 1
	}
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

func BuildFlinkClusterResource(flinkProperties FlinkProperties, annotations Annotations, labels Labels, jobManager flinkOp.JobManagerSpec, taskManager flinkOp.TaskManagerSpec, job flinkOp.JobSpec) flinkOp.FlinkCluster {
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
			ServiceAccountName: &serviceAccount,
			Image: flinkOp.ImageSpec{
				Name:       flinkImage,
				PullPolicy: corev1.PullAlways,
			},
			JobManager:      jobManager,
			TaskManager:     taskManager,
			Job:             &job,
			FlinkProperties: flinkProperties,
		},
	}
}
