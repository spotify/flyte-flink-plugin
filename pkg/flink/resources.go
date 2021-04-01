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
	"regexp"
	"strings"

	corev1 "k8s.io/api/core/v1"

	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/utils"
	flinkOp "github.com/spotify/flink-on-k8s-operator/api/v1beta1"
	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var regexpFlinkClusterName = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`)

type FlinkCluster flinkOp.FlinkCluster

func persistentVolumeTypeString(pdType flinkIdl.Resource_PersistentVolume_Type) string {
	return strings.ReplaceAll(strings.ToLower(pdType.String()), "_", "-")
}

func (fc *FlinkCluster) updateJobManagerSpec(taskCtx FlinkTaskContext) {
	out := &fc.Spec.JobManager

	out.PodAnnotations = utils.UnionMaps(taskCtx.Annotations, out.PodAnnotations)
	out.PodLabels = utils.UnionMaps(taskCtx.Labels, out.PodLabels)

	jm := taskCtx.Job.JobManager

	if cpu := jm.GetResource().GetCpu(); cpu != nil && !cpu.IsZero() {
		out.Resources.Limits[corev1.ResourceCPU] = *cpu
	}

	if memory := jm.GetResource().GetMemory(); memory != nil && !memory.IsZero() {
		out.Resources.Limits[corev1.ResourceMemory] = *memory
	}

	if pd := jm.GetResource().GetPersistentVolume(); pd != nil {
		storageClass := persistentVolumeTypeString(pd.GetType())
		storageSize := pd.GetSize()

		claim := corev1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("claim-jm-%s", fc.ObjectMeta.Name),
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
		out.VolumeClaimTemplates = []corev1.PersistentVolumeClaim{claim}

		claimVolume := corev1.Volume{
			Name: fmt.Sprintf("volume-%s", claim.Name),
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: claim.Name,
					ReadOnly:  false,
				},
			},
		}
		out.Volumes = append(out.Volumes, claimVolume)
		out.VolumeMounts = append(out.VolumeMounts, corev1.VolumeMount{
			Name:      claimVolume.Name,
			ReadOnly:  false,
			MountPath: "/data/flink",
		})
	}
}

func (fc *FlinkCluster) updateTaskManagerSpec(taskCtx FlinkTaskContext) {
	out := &fc.Spec.TaskManager

	out.PodAnnotations = utils.UnionMaps(taskCtx.Annotations, out.PodAnnotations)
	out.PodLabels = utils.UnionMaps(taskCtx.Labels, out.PodLabels)

	tm := taskCtx.Job.TaskManager

	if cpu := tm.GetResource().GetCpu(); cpu != nil && !cpu.IsZero() {
		out.Resources.Limits[corev1.ResourceCPU] = *cpu
	}

	if memory := tm.GetResource().GetMemory(); memory != nil && !memory.IsZero() {
		out.Resources.Limits[corev1.ResourceMemory] = *memory
	}

	if replicas := tm.GetReplicas(); replicas > 0 {
		out.Replicas = replicas
	}

	if pd := tm.GetResource().GetPersistentVolume(); pd != nil {
		storageClass := persistentVolumeTypeString(pd.GetType())
		storageSize := pd.GetSize()

		claim := corev1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("claim-tm-%s", fc.GetName()),
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
		out.VolumeClaimTemplates = []corev1.PersistentVolumeClaim{claim}

		claimVolume := corev1.Volume{
			Name: fmt.Sprintf("volume-%s", claim.Name),
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: claim.Name,
					ReadOnly:  false,
				},
			},
		}
		out.Volumes = append(out.Volumes, claimVolume)
		out.VolumeMounts = append(out.VolumeMounts, corev1.VolumeMount{
			Name:      claimVolume.Name,
			ReadOnly:  false,
			MountPath: "/data/flink",
		})
	}
}

func (fc *FlinkCluster) updateJobSpec(taskCtx FlinkTaskContext, taskManagerReplicas, taskManagerTaskSlots int32) {
	out := fc.Spec.Job
	if out == nil {
		out = &flinkOp.JobSpec{}
	}

	out.PodAnnotations = utils.UnionMaps(taskCtx.Annotations, out.PodAnnotations)
	out.PodLabels = utils.UnionMaps(taskCtx.Labels, out.PodLabels)

	out.JarFile = taskCtx.Job.JarFile
	out.ClassName = &taskCtx.Job.MainClass
	out.Args = taskCtx.Job.Args

	parallelism := taskManagerReplicas * int32(taskManagerTaskSlots)
	out.Parallelism = &parallelism

	out.CleanupPolicy = &flinkOp.CleanupPolicy{
		AfterJobSucceeds:  flinkOp.CleanupActionDeleteCluster,
		AfterJobFails:     flinkOp.CleanupActionDeleteCluster,
		AfterJobCancelled: flinkOp.CleanupActionDeleteCluster,
	}

	if strings.HasPrefix(out.JarFile, "gs://") {
		//TODO(regadas): add job resources to the config
		resourceList := corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("1"),
			corev1.ResourceMemory: resource.MustParse("1Gi"),
		}
		//FIXME(regadas): this strategy will likely change
		container := corev1.Container{
			Name:      "gcs-downloader",
			Image:     "google/cloud-sdk",
			Command:   []string{"gsutil"},
			Args:      []string{"cp", out.JarFile, "/cache/job.jar"},
			Resources: corev1.ResourceRequirements{Limits: resourceList},
		}
		out.JarFile = "/cache/job.jar"
		out.InitContainers = append(out.InitContainers, container)
	}
}

func NewFlinkCluster(config *Config, taskCtx FlinkTaskContext) (*flinkOp.FlinkCluster, error) {
	cluster := FlinkCluster(*config.DefaultFlinkCluster.DeepCopy())

	if err := validate(taskCtx.Name, regexpFlinkClusterName); err != nil {
		return nil, err
	}

	cluster.ObjectMeta = metav1.ObjectMeta{
		Name:        taskCtx.Name,
		Namespace:   taskCtx.Namespace,
		Annotations: taskCtx.Annotations,
		Labels:      taskCtx.Labels,
	}
	cluster.TypeMeta = metav1.TypeMeta{
		Kind:       KindFlinkCluster,
		APIVersion: flinkOp.GroupVersion.String(),
	}

	cluster.Spec.FlinkProperties = BuildFlinkProperties(config, taskCtx.Job)

	if image := taskCtx.Job.GetImage(); len(image) != 0 {
		cluster.Spec.Image.Name = image
	}

	if sa := taskCtx.Job.GetServiceAccount(); len(sa) != 0 {
		cluster.Spec.ServiceAccountName = &sa
	}

	cluster.updateJobManagerSpec(taskCtx)
	cluster.updateTaskManagerSpec(taskCtx)

	taskSlots := int32(FlinkProperties(cluster.Spec.FlinkProperties).GetInt("taskmanager.numberOfTaskSlots"))
	cluster.updateJobSpec(taskCtx, cluster.Spec.TaskManager.Replicas, taskSlots)

	// fill in defaults
	resource := flinkOp.FlinkCluster(cluster)
	resource.Default()

	err := resource.ValidateCreate()
	if err != nil {
		return nil, err
	}

	return &resource, nil
}
