// Copyright 2021 Spotify AB.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flink

import (
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"

	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	pluginsCore "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"
	flinkOp "github.com/spotify/flink-on-k8s-operator/api/v1beta1"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	cacheVolumes      = []corev1.Volume{{Name: "cache-volume"}}
	cacheVolumeMounts = []corev1.VolumeMount{{Name: "cache-volume", MountPath: "/cache"}}
)

func persistentVolumeTypeString(pdType flinkIdl.Resource_PersistentVolume_Type) string {
	return strings.ReplaceAll(strings.ToLower(pdType.String()), "_", "-")
}

func buildJobManagerSpec(jm *flinkIdl.JobManager, config *JobManagerConfig, objectMeta *metav1.ObjectMeta) flinkOp.JobManagerSpec {
	spec := flinkOp.JobManagerSpec{
		PodAnnotations: objectMeta.Annotations,
		PodLabels:      objectMeta.Labels,
		Volumes:        cacheVolumes,
		VolumeMounts:   cacheVolumeMounts,
	}

	resourceList := make(corev1.ResourceList)

	cpu := config.Cpu
	if jm.GetResource().GetCpu() != nil {
		cpu = *jm.GetResource().GetCpu()
	}
	if !cpu.IsZero() {
		resourceList[corev1.ResourceCPU] = cpu
	}

	memory := config.Memory
	if jm.GetResource().GetMemory() != nil {
		memory = *jm.GetResource().GetMemory()
	}
	if !memory.IsZero() {
		resourceList[corev1.ResourceMemory] = memory
	}

	spec.Resources.Limits = resourceList

	if pd := jm.GetResource().GetPersistentVolume(); pd != nil {
		storageClass := persistentVolumeTypeString(pd.GetType())
		storageSize := pd.GetSize()

		claim := corev1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("claim-jm-%s", objectMeta.Name),
			},
			Spec: corev1.PersistentVolumeClaimSpec{
				AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
				Resources: corev1.ResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceStorage: *storageSize,
					},
				},
				StorageClassName: &storageClass,
			},
		}
		spec.VolumeClaimTemplates = []corev1.PersistentVolumeClaim{claim}

		claimVolume := corev1.Volume{
			Name: fmt.Sprintf("volume-%s", claim.Name),
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: claim.Name,
					ReadOnly:  false,
				},
			},
		}
		spec.Volumes = append(spec.Volumes, claimVolume)

		spec.VolumeMounts = append(spec.VolumeMounts, corev1.VolumeMount{
			Name:      claimVolume.Name,
			ReadOnly:  false,
			MountPath: "/data/flink",
		})
	}

	return spec
}

func buildTaskManagerSpec(tm *flinkIdl.TaskManager, config *TaskManagerConfig, objectMeta *metav1.ObjectMeta) flinkOp.TaskManagerSpec {
	spec := flinkOp.TaskManagerSpec{
		PodAnnotations: objectMeta.Annotations,
		PodLabels:      objectMeta.Labels,
		Volumes:        cacheVolumes,
		VolumeMounts:   cacheVolumeMounts,
	}

	resourceList := make(corev1.ResourceList)

	cpu := config.Cpu
	if tm.GetResource().GetCpu() != nil {
		cpu = *tm.GetResource().GetCpu()
	}
	if !cpu.IsZero() {
		resourceList[corev1.ResourceCPU] = cpu
	}

	memory := config.Memory
	if tm.GetResource().GetMemory() != nil {
		memory = *tm.GetResource().GetMemory()
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

	if pd := tm.GetResource().GetPersistentVolume(); pd != nil {
		storageClass := persistentVolumeTypeString(pd.GetType())
		storageSize := pd.GetSize()

		claim := corev1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("claim-tm-%s", objectMeta.Name),
			},
			Spec: corev1.PersistentVolumeClaimSpec{
				AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
				Resources: corev1.ResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceStorage: *storageSize,
					},
				},
				StorageClassName: &storageClass,
			},
		}
		spec.VolumeClaimTemplates = []corev1.PersistentVolumeClaim{claim}

		claimVolume := corev1.Volume{
			Name: fmt.Sprintf("volume-%s", claim.Name),
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: claim.Name,
					ReadOnly:  false,
				},
			},
		}
		spec.Volumes = append(spec.Volumes, claimVolume)

		spec.VolumeMounts = append(spec.VolumeMounts, corev1.VolumeMount{
			Name:      claimVolume.Name,
			ReadOnly:  false,
			MountPath: "/data/flink",
		})
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

func buildFlinkClusterSpec(config *Config, job flinkIdl.FlinkJob, jobManager flinkOp.JobManagerSpec, taskManager flinkOp.TaskManagerSpec, jobSpec flinkOp.JobSpec, flinkProperties FlinkProperties, objectMeta *metav1.ObjectMeta) flinkOp.FlinkCluster {
	image := job.GetImage()
	if len(image) == 0 {
		image = config.Image
	}
	sa := job.GetServiceAccount()
	if len(sa) == 0 {
		sa = config.ServiceAccount
	}

	return flinkOp.FlinkCluster{
		ObjectMeta: *objectMeta,
		TypeMeta: metav1.TypeMeta{
			Kind:       KindFlinkCluster,
			APIVersion: flinkOp.GroupVersion.String(),
		},
		Spec: flinkOp.FlinkClusterSpec{
			ServiceAccountName: &sa,
			Image: flinkOp.ImageSpec{
				Name:       image,
				PullPolicy: corev1.PullIfNotPresent,
			},
			JobManager:      jobManager,
			TaskManager:     taskManager,
			Job:             &jobSpec,
			FlinkProperties: flinkProperties,
		},
	}
}

func BuildFlinkClusterSpec(taskCtx pluginsCore.TaskExecutionMetadata, job flinkIdl.FlinkJob, config *Config) (*flinkOp.FlinkCluster, error) {
	annotations := GetDefaultAnnotations(taskCtx)
	labels := GetDefaultLabels(taskCtx)
	objectMeta := &metav1.ObjectMeta{
		Name:        taskCtx.GetTaskExecutionID().GetGeneratedName(),
		Namespace:   taskCtx.GetNamespace(),
		Annotations: annotations,
		Labels:      labels,
	}
	flinkProperties := BuildFlinkProperties(config, job)

	jobManagerSpec := buildJobManagerSpec(job.JobManager, &config.JobManager, objectMeta)
	taskManagerSpec := buildTaskManagerSpec(job.TaskManager, &config.TaskManager, objectMeta)
	jobSpec := buildJobSpec(job, taskManagerSpec, flinkProperties)
	flinkCluster := buildFlinkClusterSpec(config, job, jobManagerSpec, taskManagerSpec, jobSpec, flinkProperties, objectMeta)

	// fill in defaults
	flinkCluster.Default()

	err := flinkCluster.ValidateCreate()
	if err != nil {
		return nil, err
	}

	return &flinkCluster, nil
}
