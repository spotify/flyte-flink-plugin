package flink

import (
	"strings"

	corev1 "k8s.io/api/core/v1"

	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	pluginsCore "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"
	flinkOp "github.com/regadas/flink-on-k8s-operator/api/v1beta1"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	cacheVolumes      = []corev1.Volume{{Name: "cache-volume"}}
	cacheVolumeMounts = []corev1.VolumeMount{{Name: "cache-volume", MountPath: "/cache"}}
)

func buildJobManagerSpec(jm *flinkIdl.JobManager, config *JobManagerConfig, annotations Annotations, labels Labels) flinkOp.JobManagerSpec {
	spec := flinkOp.JobManagerSpec{
		PodAnnotations: annotations,
		PodLabels:      labels,
		Volumes:        cacheVolumes,
		VolumeMounts:   cacheVolumeMounts,
	}

	resourceList := make(corev1.ResourceList)

	cpu := config.Cpu
	if jm.GetCpu() != nil {
		cpu = *jm.GetCpu()
	}
	if !cpu.IsZero() {
		resourceList[corev1.ResourceCPU] = cpu
	}

	memory := config.Memory
	if jm.GetMemory() != nil {
		memory = *jm.GetMemory()
	}
	if !memory.IsZero() {
		resourceList[corev1.ResourceMemory] = memory
	}

	spec.Resources.Limits = resourceList

	return spec
}

func buildTaskManagerSpec(tm *flinkIdl.TaskManager, config *TaskManagerConfig, annotations Annotations, labels Labels) flinkOp.TaskManagerSpec {
	spec := flinkOp.TaskManagerSpec{
		PodAnnotations: annotations,
		PodLabels:      labels,
		Volumes:        cacheVolumes,
		VolumeMounts:   cacheVolumeMounts,
	}

	resourceList := make(corev1.ResourceList)

	cpu := config.Cpu
	if tm.GetCpu() != nil {
		cpu = *tm.GetCpu()
	}
	if !cpu.IsZero() {
		resourceList[corev1.ResourceCPU] = cpu
	}

	memory := config.Memory
	if tm.GetMemory() != nil {
		memory = *tm.GetMemory()
	}
	if !memory.IsZero() {
		resourceList[corev1.ResourceMemory] = memory
	}

	spec.Resources.Limits = resourceList

	replicas := int32(config.Replicas)
	if tm.GetReplicas() > 0 {
		replicas = tm.GetReplicas()
	}

	if replicas > 0 {
		spec.Replicas = replicas
	}

	return spec
}

func buildJobSpec(job flinkIdl.FlinkJob, taskManager flinkOp.TaskManagerSpec, flinkProperties FlinkProperties) flinkOp.JobSpec {
	taskSlots := flinkProperties.GetInt("taskmanager.numberOfTaskSlots")
	parallelism := taskManager.Replicas * int32(taskSlots)

	//TODO(regadas): add job resources to the config
	resourceList := corev1.ResourceList{
		corev1.ResourceCPU:    resource.MustParse("1"),
		corev1.ResourceMemory: resource.MustParse("1Gi"),
	}

	spec := flinkOp.JobSpec{
		JarFile:      job.JarFile,
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
		Resources:      corev1.ResourceRequirements{Limits: resourceList},
		InitContainers: []corev1.Container{},
	}

	if strings.HasPrefix(job.JarFile, "gs://") {
		//FIXME(regadas): this strategy will likely change
		container := corev1.Container{
			Name:      "gcs-downloader",
			Image:     "google/cloud-sdk",
			Command:   []string{"gsutil"},
			Args:      []string{"cp", job.JarFile, "/cache/job.jar"},
			Resources: corev1.ResourceRequirements{Limits: resourceList},
		}
		spec.JarFile = "/cache/job.jar"
		spec.InitContainers = append(spec.InitContainers, container)
	}

	return spec
}

func buildFlinkClusterSpec(config *Config, jobManager flinkOp.JobManagerSpec, taskManager flinkOp.TaskManagerSpec, job flinkOp.JobSpec, flinkProperties FlinkProperties, annotations Annotations, labels Labels) flinkOp.FlinkCluster {
	return flinkOp.FlinkCluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindFlinkCluster,
			APIVersion: flinkOp.GroupVersion.String(),
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

func BuildFlinkClusterSpec(taskCtx pluginsCore.TaskExecutionMetadata, job flinkIdl.FlinkJob, config *Config) (*flinkOp.FlinkCluster, error) {
	annotations := GetDefaultAnnotations(taskCtx)
	labels := GetDefaultLabels(taskCtx)
	flinkProperties := BuildFlinkProperties(config, job)

	jobManagerSpec := buildJobManagerSpec(job.JobManager, &config.JobManager, annotations, labels)
	taskManagerSpec := buildTaskManagerSpec(job.TaskManager, &config.TaskManager, annotations, labels)
	jobSpec := buildJobSpec(job, taskManagerSpec, flinkProperties)
	flinkCluster := buildFlinkClusterSpec(config, jobManagerSpec, taskManagerSpec, jobSpec, flinkProperties, annotations, labels)

	flinkCluster.ObjectMeta = metav1.ObjectMeta{
		Name:        taskCtx.GetTaskExecutionID().GetGeneratedName(),
		Namespace:   taskCtx.GetNamespace(),
		Annotations: annotations,
		Labels:      labels,
	}

	// fill in defaults
	flinkCluster.Default()

	err := flinkCluster.ValidateCreate()
	if err != nil {
		return nil, err
	}

	return &flinkCluster, nil
}
