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
	"net/url"
	"strings"

	corev1 "k8s.io/api/core/v1"

	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/utils"
	flinkOp "github.com/spotify/flink-on-k8s-operator/api/v1beta1"
	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Downloader interface {
	Container(artifacts []string) corev1.Container
}

type localDownloader struct{}

func (localDownloader) Container(artifacts []string) corev1.Container {
	cmd := strings.Join([]string{
		fmt.Sprintf("mkdir -p %s", defaultJarLibPath),
		fmt.Sprintf("cp %s %s", strings.Join(artifacts[:], " "), defaultJarLibPath),
	}, " && ")

	return corev1.Container{
		Name:      "local-downloader",
		Image:     "alpine",
		Command:   []string{"/bin/sh"},
		Args:      []string{"-c", cmd},
		Resources: defaultInitResources,
	}
}

type gcsDownloader struct{}

func (gcsDownloader) Container(artifacts []string) corev1.Container {
	cmd := strings.Join([]string{
		fmt.Sprintf("mkdir -p %s", defaultJarLibPath),
		`if [ -n "${GOOGLE_APPLICATION_CREDENTIALS}" ]; then gcloud auth activate-service-account --key-file $GOOGLE_APPLICATION_CREDENTIALS; fi`,
		fmt.Sprintf("gsutil -m cp %s %s", strings.Join(artifacts[:], " "), defaultJarLibPath),
	}, " && ")

	return corev1.Container{
		Name:      "gcs-downloader",
		Image:     "google/cloud-sdk",
		Command:   []string{"/bin/sh"},
		Args:      []string{"-c", cmd},
		Resources: defaultInitResources,
	}
}

type DownloaderRegistry map[string]Downloader

func (dr DownloaderRegistry) GetDownloader(scheme string) (Downloader, error) {
	d, ok := dr[scheme]
	if !ok {
		return nil, fmt.Errorf("downloader not implemented for scheme: %s", scheme)
	}

	return d, nil
}

var downloaderRegistry = DownloaderRegistry{
	"":   localDownloader{},
	"gs": gcsDownloader{},
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

	out.CleanupPolicy = &flinkOp.CleanupPolicy{
		AfterJobSucceeds:  flinkOp.CleanupActionDeleteCluster,
		AfterJobFails:     flinkOp.CleanupActionDeleteCluster,
		AfterJobCancelled: flinkOp.CleanupActionDeleteCluster,
	}

	groupBy := GroupByScheme(taskCtx.Job.GetJarFiles())
	if len(groupBy) == 0 {
		// use jflyte artifacts as fallback only
		urls := make([]string, len(taskCtx.Job.GetJflyte().GetArtifacts()))
		for i, a := range taskCtx.Job.GetJflyte().GetArtifacts() {
			urls[i] = a.Location
		}
		groupBy = GroupByScheme(urls)
	}

	if len(groupBy) == 0 {
		return fmt.Errorf("no artifacts provided")
	}

	volumeName := fmt.Sprintf("%s-jars", taskCtx.ClusterName.String())
	out.Volumes = append(out.Volumes, corev1.Volume{Name: volumeName})
	out.VolumeMounts = append(out.VolumeMounts, corev1.VolumeMount{Name: volumeName, MountPath: jarsVolumePath})
	out.JarFile = defaultJarFile

	for scheme, urls := range groupBy {
		downloader, err := downloaderRegistry.GetDownloader(scheme)
		if err != nil {
			return err
		}
		out.InitContainers = append(out.InitContainers, downloader.Container(urls))
	}
	out.InitContainers = append(out.InitContainers, artifactZip)

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

	cluster.Spec.FlinkProperties = MergeProperties(
		cluster.Spec.FlinkProperties,
		taskCtx.Job.FlinkProperties,
		config.FlinkPropertiesOverride,
	)

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
