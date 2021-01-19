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

func BuildJobManagerSpec(jm *flinkIdl.JobManager, config *JobManagerConfig, annotations Annotations, labels Labels) flinkOp.JobManagerSpec {
	cpu := config.Cpu
	if jm.GetCpu() != nil {
		cpu = *jm.GetCpu()
	}
	memory := config.Memory
	if jm.GetMemory() != nil {
		memory = *jm.GetMemory()
	}

	return flinkOp.JobManagerSpec{
		PodAnnotations: annotations,
		PodLabels:      labels,
		Resources: corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:    cpu,
				corev1.ResourceMemory: memory,
			},
		},
		Volumes:      cacheVolumes,
		VolumeMounts: cacheVolumeMounts,
	}
}

func BuildTaskManagerSpec(tm *flinkIdl.TaskManager, config *TaskManagerConfig, annotations Annotations, labels Labels) flinkOp.TaskManagerSpec {
	cpu := config.Cpu
	if tm.GetCpu() != nil {
		cpu = *tm.GetCpu()
	}
	memory := config.Memory
	if tm.GetMemory() != nil {
		memory = *tm.GetMemory()
	}
	replicas := int32(config.Replicas)
	if tm.GetReplicas() > 0 {
		replicas = tm.GetReplicas()
	}

	return flinkOp.TaskManagerSpec{
		PodAnnotations: annotations,
		PodLabels:      labels,
		Replicas:       replicas,
		Volumes:        cacheVolumes,
		VolumeMounts:   cacheVolumeMounts,
		Resources: corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:    cpu,
				corev1.ResourceMemory: memory,
			},
		},
	}
}

func BuildJobSpec(job flinkIdl.FlinkJob, taskManager flinkOp.TaskManagerSpec, flinkProperties FlinkProperties) flinkOp.JobSpec {
	taskSlots := flinkProperties.GetInt("taskmanager.numberOfTaskSlots")
	parallelism := taskManager.Replicas * int32(taskSlots)

	return flinkOp.JobSpec{

		JarFile:      "/cache/job.jar",
		ClassName:    &job.MainClass,
		Args:         job.Args,
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
					job.JarFile,
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

func BuildFlinkClusterSpec(config *Config, jobManager flinkOp.JobManagerSpec, taskManager flinkOp.TaskManagerSpec, job flinkOp.JobSpec, flinkProperties FlinkProperties, annotations Annotations, labels Labels) flinkOp.FlinkCluster {
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
