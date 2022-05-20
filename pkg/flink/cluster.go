package flink

import (
	"bytes"
	"strings"

	flinkOp "github.com/spotify/flink-on-k8s-operator/apis/flinkcluster/v1beta1"
	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FlinkCluster struct {
	*flinkOp.FlinkCluster
}

func NewFlinkCluster(taskCtx *FlinkTaskContext) (*flinkOp.FlinkCluster, error) {
	cluster := FlinkCluster{&flinkOp.FlinkCluster{}}
	cluster.Default()

	if err := cluster.setObjectMeta(taskCtx); err != nil {
		return nil, err
	}
	if err := cluster.setTypeMeta(taskCtx); err != nil {
		return nil, err
	}
	if err := cluster.setSpec(taskCtx); err != nil {
		return nil, err
	}

	return cluster.FlinkCluster, nil
}

func (c *FlinkCluster) setObjectMeta(taskCtx *FlinkTaskContext) error {
	c.ObjectMeta = metav1.ObjectMeta{
		Name:        taskCtx.ClusterName.String(),
		Namespace:   taskCtx.Namespace,
		Annotations: taskCtx.Annotations,
		Labels:      taskCtx.Labels,
	}
	return nil
}

func (c *FlinkCluster) setTypeMeta(taskCtx *FlinkTaskContext) error {
	c.TypeMeta = metav1.TypeMeta{
		Kind:       KindFlinkCluster,
		APIVersion: flinkOp.GroupVersion.String(),
	}
	return nil
}

func (c *FlinkCluster) setSpec(taskCtx *FlinkTaskContext) error {
	c.Spec.EnvVars = append(c.Spec.EnvVars, corev1.EnvVar{
		Name:  stagedJarsEnvVarName,
		Value: strings.Join(getJobArtifacts(&taskCtx.Job), " "),
	})

	if err := c.setFlinkProperties(taskCtx); err != nil {
		return nil
	}

	if version := taskCtx.Job.GetFlinkVersion(); len(version) != 0 {
		c.Spec.FlinkVersion = version
	}

	if image := taskCtx.Job.GetImage(); len(image) != 0 {
		c.Spec.Image.Name = image
	}

	if sa := taskCtx.Job.GetServiceAccount(); len(sa) != 0 {
		c.Spec.ServiceAccountName = &sa
	}

	c.updateJobManagerSpec(taskCtx)
	c.updateTaskManagerSpec(taskCtx)
	if err := c.updateJobSpec(taskCtx); err != nil {
		return err
	}

	return nil
}

func (fc *FlinkCluster) setFlinkProperties(taskCtx *FlinkTaskContext) error {
	result := make(map[string]string)
	for k, v := range taskCtx.Job.FlinkProperties {
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

func (c *FlinkCluster) updateJobManagerSpec(taskCtx *FlinkTaskContext) {
	replicas := int32(1)
	jmSpec := &flinkOp.JobManagerSpec{
		PodAnnotations: taskCtx.Annotations,
		PodLabels:      taskCtx.Labels,
		Replicas:       &replicas,
		Resources: corev1.ResourceRequirements{
			Limits: corev1.ResourceList{},
		},
	}

	jm := taskCtx.Job.JobManager

	if cpu := jm.GetResource().GetCpu(); cpu != nil {
		if quantity := resource.MustParse(cpu.GetString_()); !quantity.IsZero() {
			jmSpec.Resources.Limits[corev1.ResourceCPU] = quantity
		}
	}

	if memory := jm.GetResource().GetMemory(); memory != nil {
		if quantity := resource.MustParse(memory.GetString_()); !quantity.IsZero() {
			jmSpec.Resources.Limits[corev1.ResourceMemory] = quantity
		}
	}

	if pv := jm.GetResource().GetPersistentVolume(); pv != nil {
		claim := getPersistentVolumeClaim(jobManagerVolumeClaim, pv)
		jmSpec.VolumeClaimTemplates, jmSpec.VolumeMounts = addPersistentVolumeClaim(
			jmSpec.VolumeClaimTemplates,
			jmSpec.VolumeMounts,
			claim,
			volumeClaimMountPath,
		)
		c.Spec.FlinkProperties[flinkIoTmpDirsProperty] = volumeClaimMountPath
	}

	c.Spec.JobManager = jmSpec
}

func (c *FlinkCluster) updateTaskManagerSpec(taskCtx *FlinkTaskContext) {
	tmSpec := &flinkOp.TaskManagerSpec{
		PodAnnotations: taskCtx.Annotations,
		PodLabels:      taskCtx.Labels,
		Resources: corev1.ResourceRequirements{
			Limits: corev1.ResourceList{},
		},
	}

	tm := taskCtx.Job.TaskManager

	if cpu := tm.GetResource().GetCpu(); cpu != nil {
		if quantity := resource.MustParse(cpu.GetString_()); !quantity.IsZero() {
			tmSpec.Resources.Limits[corev1.ResourceCPU] = quantity
		}
	}

	if memory := tm.GetResource().GetMemory(); memory != nil {
		if quantity := resource.MustParse(memory.GetString_()); !quantity.IsZero() {
			tmSpec.Resources.Limits[corev1.ResourceMemory] = quantity
		}
	}

	if replicas := tm.GetReplicas(); replicas > 0 {
		tmSpec.Replicas = &replicas
	}

	if pv := tm.GetResource().GetPersistentVolume(); pv != nil {
		claim := getPersistentVolumeClaim(taskManagerVolumeClaim, pv)
		tmSpec.VolumeClaimTemplates, tmSpec.VolumeMounts = addPersistentVolumeClaim(
			tmSpec.VolumeClaimTemplates,
			tmSpec.VolumeMounts,
			claim,
			volumeClaimMountPath,
		)
		c.Spec.FlinkProperties[flinkIoTmpDirsProperty] = volumeClaimMountPath
	}

	c.Spec.TaskManager = tmSpec
}

func (c *FlinkCluster) updateJobSpec(taskCtx *FlinkTaskContext) error {
	jobSpec := &flinkOp.JobSpec{
		PodAnnotations: taskCtx.Annotations,
		PodLabels:      taskCtx.Labels,
		ClassName:      &taskCtx.Job.MainClass,
		Args:           taskCtx.Job.Args,
	}

	if taskCtx.Job.Parallelism != 0 {
		jobSpec.Parallelism = &taskCtx.Job.Parallelism
	}

	artifacts := getJobArtifacts(&taskCtx.Job)

	if jobSpec.JarFile == nil && len(artifacts) == 1 {
		jobSpec.JarFile = &artifacts[0]
	} else {
		initContainers := []corev1.Container{}
		for _, container := range jobSpec.InitContainers {
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
		jobSpec.InitContainers = initContainers
	}

	c.Spec.Job = jobSpec

	return nil
}

func getJobArtifacts(job *flinkIdl.FlinkJob) []string {
	artifacts := job.GetJarFiles()
	if len(artifacts) == 0 {
		// use jflyte artifacts as fallback only
		urls := make([]string, len(job.GetJflyte().GetArtifacts()))
		for i, a := range job.GetJflyte().GetArtifacts() {
			urls[i] = a.Location
		}
		artifacts = urls
	}

	return artifacts
}
