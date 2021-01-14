package flink

import (
	corev1 "k8s.io/api/core/v1"
)

const KindFlinkCluster = "FlinkCluster"

var (
	flinkTaskType              = "flink"
	jobManagerUIPort           = int32(8081)
	defaultTaskManagerReplicas = 1
	defaultJobParallelism      = 1
	cacheVolumes               = []corev1.Volume{{Name: "cache-volume"}}
	cacheVolumeMounts          = []corev1.VolumeMount{{Name: "cache-volume", MountPath: "/cache"}}
)
