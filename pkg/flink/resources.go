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
	"path"
	"regexp"
	"strings"

	corev1 "k8s.io/api/core/v1"

	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/utils"
	flinkOp "github.com/spotify/flink-on-k8s-operator/api/v1beta1"
	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	jobManagerVolumeClaim  = "pvc-jm"
	taskManagerVolumeClaim = "pvc-tm"
	volumeClaimMountPath   = "/flink-tmp"
	flinkIoTmpDirsProperty = "io.tmp.dirs"
	jarsVolumePath         = "/jars"
	gcsPrefix              = "gs://"
)

var (
	regexpFlinkClusterName = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`)
)

type FlinkCluster flinkOp.FlinkCluster

func getPersistentVolumeClaim(name string, pv *flinkIdl.Resource_PersistentVolume) corev1.PersistentVolumeClaim {
	storageClass := strings.ReplaceAll(strings.ToLower(pv.GetType().String()), "_", "-")
	if pv.GetSize() == nil {
		return corev1.PersistentVolumeClaim{}
	}

	storageSize := resource.MustParse(pv.GetSize().GetString_())

	return corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: storageSize,
				},
			},
			StorageClassName: &storageClass,
		},
	}
}

func (fc *FlinkCluster) updateJobManagerSpec(taskCtx FlinkTaskContext) {
	out := &fc.Spec.JobManager

	out.PodAnnotations = utils.UnionMaps(taskCtx.Annotations, out.PodAnnotations)
	out.PodLabels = utils.UnionMaps(taskCtx.Labels, out.PodLabels)

	jm := taskCtx.Job.JobManager

	if cpu := jm.GetResource().GetCpu(); cpu != nil {
		if quantity := resource.MustParse(cpu.GetString_()); !quantity.IsZero() {
			out.Resources.Limits[corev1.ResourceCPU] = quantity
		}
	}

	if memory := jm.GetResource().GetMemory(); memory != nil {
		if quantity := resource.MustParse(memory.GetString_()); !quantity.IsZero() {
			out.Resources.Limits[corev1.ResourceMemory] = quantity
		}
	}

	if pv := jm.GetResource().GetPersistentVolume(); pv != nil {
		claim := getPersistentVolumeClaim(jobManagerVolumeClaim, pv)

		out.VolumeClaimTemplates = []corev1.PersistentVolumeClaim{claim}
		out.VolumeMounts = append(out.VolumeMounts, corev1.VolumeMount{
			Name:      claim.Name,
			ReadOnly:  false,
			MountPath: volumeClaimMountPath,
		})

		fc.Spec.FlinkProperties[flinkIoTmpDirsProperty] = volumeClaimMountPath
	}
}

func (fc *FlinkCluster) updateTaskManagerSpec(taskCtx FlinkTaskContext) {
	out := &fc.Spec.TaskManager

	out.PodAnnotations = utils.UnionMaps(taskCtx.Annotations, out.PodAnnotations)
	out.PodLabels = utils.UnionMaps(taskCtx.Labels, out.PodLabels)

	tm := taskCtx.Job.TaskManager

	if cpu := tm.GetResource().GetCpu(); cpu != nil {
		if quantity := resource.MustParse(cpu.GetString_()); !quantity.IsZero() {
			out.Resources.Limits[corev1.ResourceCPU] = quantity
		}
	}

	if memory := tm.GetResource().GetMemory(); memory != nil {
		if quantity := resource.MustParse(memory.GetString_()); !quantity.IsZero() {
			out.Resources.Limits[corev1.ResourceMemory] = quantity
		}
	}

	if replicas := tm.GetReplicas(); replicas > 0 {
		out.Replicas = replicas
	}

	if pv := tm.GetResource().GetPersistentVolume(); pv != nil {
		claim := getPersistentVolumeClaim(taskManagerVolumeClaim, pv)

		out.VolumeClaimTemplates = []corev1.PersistentVolumeClaim{claim}
		out.VolumeMounts = append(out.VolumeMounts, corev1.VolumeMount{
			Name:      claim.Name,
			ReadOnly:  false,
			MountPath: volumeClaimMountPath,
		})

		fc.Spec.FlinkProperties[flinkIoTmpDirsProperty] = volumeClaimMountPath
	}
}

func (fc *FlinkCluster) updateJobSpec(taskCtx FlinkTaskContext, taskManagerReplicas, taskManagerTaskSlots int32) {
	if fc.Spec.Job == nil {
		fc.Spec.Job = &flinkOp.JobSpec{}
	}
	out := fc.Spec.Job

	out.PodAnnotations = utils.UnionMaps(taskCtx.Annotations, out.PodAnnotations)
	out.PodLabels = utils.UnionMaps(taskCtx.Labels, out.PodLabels)

	out.ClassName = &taskCtx.Job.MainClass
	out.Args = taskCtx.Job.Args

	parallelism := taskManagerReplicas * taskManagerTaskSlots
	out.Parallelism = &parallelism

	out.CleanupPolicy = &flinkOp.CleanupPolicy{
		AfterJobSucceeds:  flinkOp.CleanupActionDeleteCluster,
		AfterJobFails:     flinkOp.CleanupActionDeleteCluster,
		AfterJobCancelled: flinkOp.CleanupActionDeleteCluster,
	}

	urls := make([]string, len(taskCtx.Job.GetJarFiles())+len(taskCtx.Job.GetJflyte().GetArtifacts()))
	useGcs := true
	for i, s := range taskCtx.Job.GetJarFiles() {
		urls[i] = s
		if useGcs && !strings.HasPrefix(s, gcsPrefix) {
			useGcs = false
		}
	}
	for i, a := range taskCtx.Job.GetJflyte().GetArtifacts() {
		urls[i] = a.Location
		if useGcs && !strings.HasPrefix(a.Location, gcsPrefix) {
			useGcs = false
		}
	}

	// XXX(julient) I don't like that this would just silently fail if the condition is not satisfied
	if useGcs && len(urls) > 0 {
		//TODO(regadas): add job resources to the config
		resourceList := corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("1"),
			corev1.ResourceMemory: resource.MustParse("1Gi"),
		}

		commands := []string{"/bin/sh", "-c"}
		tmp := "/tmp/artifacts"
		jarPath := path.Join(jarsVolumePath, "job.jar")
		args := []string{
			fmt.Sprintf("mkdir %s/lib", tmp),
			fmt.Sprintf("gsutil cp %s %s/lib", strings.Join(urls[:], " "), tmp),
			fmt.Sprintf("$(cd %s && zip -r job.jar .)", tmp),
			fmt.Sprintf("cp /tmp/job.jar %s", jarPath),
		}

		//FIXME(regadas): this strategy will likely change
		container := corev1.Container{
			Name:      "gcs-downloader",
			Image:     "google/cloud-sdk",
			Command:   commands,
			Args:      args,
			Resources: corev1.ResourceRequirements{Limits: resourceList},
		}
		out.JarFile = jarPath
		out.InitContainers = append(out.InitContainers, container)

		volumeName := fmt.Sprintf("%s-jars", taskCtx.Name)
		out.Volumes = append(out.Volumes, corev1.Volume{Name: volumeName})
		out.VolumeMounts = append(out.VolumeMounts, corev1.VolumeMount{Name: volumeName, MountPath: jarsVolumePath})
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

	cluster.Spec.FlinkProperties = MergeProperties(
		cluster.Spec.FlinkProperties,
		taskCtx.Job.FlinkProperties,
		config.FlinkPropertiesOverride,
	)

	if image := taskCtx.Job.GetImage(); len(image) != 0 {
		cluster.Spec.Image.Name = image
	}

	if sa := taskCtx.Job.GetServiceAccount(); len(sa) != 0 {
		cluster.Spec.ServiceAccountName = &sa
	}

	cluster.updateJobManagerSpec(taskCtx)
	cluster.updateTaskManagerSpec(taskCtx)

	taskSlots, err := Properties(cluster.Spec.FlinkProperties).GetInt("taskmanager.numberOfTaskSlots")
	if err != nil {
		return nil, err
	}
	cluster.updateJobSpec(taskCtx, cluster.Spec.TaskManager.Replicas, int32(taskSlots))

	// fill in defaults
	resource := flinkOp.FlinkCluster(cluster)
	resource.Default()

	err = resource.ValidateCreate()
	if err != nil {
		return nil, err
	}

	return &resource, nil
}
