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
	"regexp"

	pluginsConfig "github.com/flyteorg/flyteplugins/go/tasks/config"
	flinkOp "github.com/spotify/flink-on-k8s-operator/apis/flinkcluster/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/utils/pointer"
)

const (
	KindFlinkCluster = "FlinkCluster"
	// Flyte flink task type
	FlinkTaskType = "flink"

	// FlinkCluster resource default values
	jobManagerVolumeClaim  = "pvc-jm"
	taskManagerVolumeClaim = "pvc-tm"
	volumeClaimMountPath   = "/tmp"

	// Flink properties
	flinkIoTmpDirsProperty = "io.tmp.dirs"
)

var (
	regexpFlinkClusterName = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`)
	generatedNameMaxLength = 50
	nonRetryableExitCodes  = []int32{}
	defaultServiceAccount  = "default"
	defaultConfig          = Config{
		DefaultFlinkCluster: flinkOp.FlinkCluster{
			Spec: flinkOp.FlinkClusterSpec{
				ServiceAccountName: &defaultServiceAccount,
				JobManager: &flinkOp.JobManagerSpec{
					Resources: corev1.ResourceRequirements{
						Requests: map[corev1.ResourceName]resource.Quantity{
							corev1.ResourceCPU:    resource.MustParse("4"),
							corev1.ResourceMemory: resource.MustParse("4Gi"),
						},
					},
				},
				TaskManager: &flinkOp.TaskManagerSpec{
					Replicas: pointer.Int32(1),
					Resources: corev1.ResourceRequirements{
						Requests: map[corev1.ResourceName]resource.Quantity{
							corev1.ResourceCPU:    resource.MustParse("4"),
							corev1.ResourceMemory: resource.MustParse("4Gi"),
						},
					},
				},
			},
		},
		GeneratedNameMaxLength: &generatedNameMaxLength,
		NonRetryableExitCodes:  nonRetryableExitCodes,
	}

	flinkConfigSection = pluginsConfig.MustRegisterSubSection("flink", &defaultConfig)
)
