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
	flinkImage                 = "gcr.io/esquilo/flink:1.10.1-scala_2.12-gcs"
	serviceAccount             = "ff-dev-workload-sa"
	cacheVolumes               = []corev1.Volume{{Name: "cache-volume"}}
	cacheVolumeMounts          = []corev1.VolumeMount{{Name: "cache-volume", MountPath: "/cache"}}
)
