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
	"bytes"
	"net/url"
	"strings"
	"text/template"

	corev1 "k8s.io/api/core/v1"

	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/utils"
	flinkOp "github.com/spotify/flink-on-k8s-operator/api/v1beta1"
	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	containerTmpl       = template.New("container-template").Funcs(template.FuncMap{"join": strings.Join})
	flinkPropertiesTmpl = template.New("flink-properties-template").Funcs(template.FuncMap{"join": strings.Join})
)

type ContainerTemplateData struct {
	ArtifactsByScheme map[string][]string
	Artifacts         []string
}

func NewContainerTemplateData(artifacts []string) *ContainerTemplateData {
	return &ContainerTemplateData{
		ArtifactsByScheme: GroupByScheme(artifacts),
		Artifacts:         artifacts,
	}
}

type FlinkPropertiesTemplateData struct {
	Namespace   string
	ClusterName ClusterName
	Labels      map[string]string
}

func NewFlinkPropertiesTemplateData(namespace string, clusterName ClusterName, labels map[string]string) *FlinkPropertiesTemplateData {
	return &FlinkPropertiesTemplateData{
		Namespace:   namespace,
		ClusterName: clusterName,
		Labels:      labels,
	}
}

func GroupByScheme(artifacts []string) map[string][]string {
	groupBy := make(map[string][]string)
	for _, artifact := range artifacts {
		url, _ := url.Parse(artifact)
		groupBy[url.Scheme] = append(groupBy[url.Scheme], artifact)
	}
	return groupBy
}

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

func addPersistentVolumeClaim(
	claims []corev1.PersistentVolumeClaim,
	volumeMounts []corev1.VolumeMount,
	claim corev1.PersistentVolumeClaim,
	mountPath string) ([]corev1.PersistentVolumeClaim, []corev1.VolumeMount) {

	claimsByName := make(map[string]corev1.PersistentVolumeClaim)
	for _, c := range claims {
		claimsByName[c.Name] = c
	}

	mounts := []corev1.VolumeMount{}
	for _, volumeMount := range volumeMounts {
		if volumeMount.MountPath != mountPath {
			mounts = append(mounts, volumeMount)
		} else {
			delete(claimsByName, volumeMount.Name)
		}
	}
	mounts = append(mounts, corev1.VolumeMount{
		Name:      claim.Name,
		ReadOnly:  false,
		MountPath: volumeClaimMountPath,
	})

	templates := []corev1.PersistentVolumeClaim{claim}
	for _, c := range claimsByName {
		templates = append(templates, c)
	}

	return templates, mounts
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
		out.VolumeClaimTemplates, out.VolumeMounts = addPersistentVolumeClaim(
			out.VolumeClaimTemplates,
			out.VolumeMounts,
			claim,
			volumeClaimMountPath,
		)
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
		out.VolumeClaimTemplates, out.VolumeMounts = addPersistentVolumeClaim(
			out.VolumeClaimTemplates,
			out.VolumeMounts,
			claim,
			volumeClaimMountPath,
		)
		fc.Spec.FlinkProperties[flinkIoTmpDirsProperty] = volumeClaimMountPath
	}
}

func (fc *FlinkCluster) updateJobSpec(taskCtx FlinkTaskContext) error {
	if fc.Spec.Job == nil {
		fc.Spec.Job = &flinkOp.JobSpec{}
	}
	out := fc.Spec.Job

	out.PodAnnotations = utils.UnionMaps(taskCtx.Annotations, out.PodAnnotations)
	out.PodLabels = utils.UnionMaps(taskCtx.Labels, out.PodLabels)

	out.ClassName = &taskCtx.Job.MainClass
	out.Args = taskCtx.Job.Args

	if taskCtx.Job.Parallelism != 0 {
		out.Parallelism = &taskCtx.Job.Parallelism
	}

	artifacts := taskCtx.Job.GetJarFiles()
	if len(artifacts) == 0 {
		// use jflyte artifacts as fallback only
		urls := make([]string, len(taskCtx.Job.GetJflyte().GetArtifacts()))
		for i, a := range taskCtx.Job.GetJflyte().GetArtifacts() {
			urls[i] = a.Location
		}
		artifacts = urls
	}

	if out.JarFile == nil && len(artifacts) == 1 {
		out.JarFile = &artifacts[0]
	} else {
		initContainers := []corev1.Container{}
		for _, container := range out.InitContainers {
			resultArgs := []string{}
			for _, arg := range container.Args {
				tmpl, err := containerTmpl.Parse(arg)
				if err != nil {
					return err
				}

				var tpl bytes.Buffer
				if err := tmpl.Execute(&tpl, NewContainerTemplateData(artifacts)); err != nil {
					return err
				}

				resultArgs = append(resultArgs, tpl.String())
			}
			container.Args = resultArgs
			initContainers = append(initContainers, container)
		}
		out.InitContainers = initContainers
	}

	return nil
}

func (fc *FlinkCluster) updateFlinkProperties(config *Config, taskCtx FlinkTaskContext) error {

	props := MergeProperties(
		fc.Spec.FlinkProperties,
		taskCtx.Job.FlinkProperties,
		config.FlinkPropertiesOverride,
	)

	result := make(map[string]string)
	for k, v := range props {
		tmpl, err := flinkPropertiesTmpl.Parse(v)
		if err != nil {
			return err
		}
		var tpl bytes.Buffer
		ft := NewFlinkPropertiesTemplateData(taskCtx.Namespace, taskCtx.ClusterName, taskCtx.Labels)
		if err := tmpl.Execute(&tpl, ft); err != nil {
			return err
		}
		result[k] = tpl.String()
	}

	fc.Spec.FlinkProperties = result
	return nil
}

func NewFlinkCluster(config *Config, taskCtx FlinkTaskContext) (*flinkOp.FlinkCluster, error) {
	cluster := FlinkCluster(*config.DefaultFlinkCluster.DeepCopy())
	cluster.ObjectMeta = metav1.ObjectMeta{
		Name:        taskCtx.ClusterName.String(),
		Namespace:   taskCtx.Namespace,
		Annotations: taskCtx.Annotations,
		Labels:      taskCtx.Labels,
	}
	cluster.TypeMeta = metav1.TypeMeta{
		Kind:       KindFlinkCluster,
		APIVersion: flinkOp.GroupVersion.String(),
	}

	cluster.updateFlinkProperties(config, taskCtx)

	if version := taskCtx.Job.GetFlinkVersion(); len(version) != 0 {
		cluster.Spec.FlinkVersion = version
	}

	if image := taskCtx.Job.GetImage(); len(image) != 0 {
		cluster.Spec.Image.Name = image
	}

	if sa := taskCtx.Job.GetServiceAccount(); len(sa) != 0 {
		cluster.Spec.ServiceAccountName = &sa
	}

	cluster.updateJobManagerSpec(taskCtx)
	cluster.updateTaskManagerSpec(taskCtx)
	if err := cluster.updateJobSpec(taskCtx); err != nil {
		return nil, err
	}

	// fill in defaults
	resource := flinkOp.FlinkCluster(cluster)
	resource.Default()

	err := resource.ValidateCreate()
	if err != nil {
		return nil, err
	}

	return &resource, nil
}
