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
	"path"
	"regexp"

	pluginsConfig "github.com/flyteorg/flyteplugins/go/tasks/config"
	flinkOp "github.com/spotify/flink-on-k8s-operator/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	KindFlinkCluster = "FlinkCluster"
	// Flyte flink task type
	FlinkTaskType = "flink"

	// FlinkCluster resource default values
	jobManagerVolumeClaim  = "pvc-jm"
	taskManagerVolumeClaim = "pvc-tm"
	volumeClaimMountPath   = "/flink-tmp"
	jarsVolumePath         = "/jars"

	// Flink properties
	flinkIoTmpDirsProperty = "io.tmp.dirs"
)

var (
	regexpFlinkClusterName = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`)

	defaultInitResources = corev1.ResourceRequirements{
		Limits: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("1"),
			corev1.ResourceMemory: resource.MustParse("512Mi"),
		},
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("100m"),
			corev1.ResourceMemory: resource.MustParse("100Mi"),
		},
	}
	defaultJarFile    = path.Join(jarsVolumePath, "job.jar")
	defaultJarLibPath = path.Join(jarsVolumePath, "lib")
	artifactZip       = corev1.Container{
		Name:       "zip",
		Image:      "alpine",
		Command:    []string{"/bin/sh"},
		Args:       []string{"-c", "apk add zip && zip -r job.jar ."},
		WorkingDir: jarsVolumePath,
		Resources:  defaultInitResources,
	}

	generatedNameMaxLength = 50
	defaultServiceAccount  = "default"
	defaultConfig          = &Config{
		DefaultFlinkCluster: flinkOp.FlinkCluster{
			Spec: flinkOp.FlinkClusterSpec{
				ServiceAccountName: &defaultServiceAccount,
				JobManager: flinkOp.JobManagerSpec{
					AccessScope: "ClusterIP",
					Resources: corev1.ResourceRequirements{
						Limits: map[corev1.ResourceName]resource.Quantity{
							corev1.ResourceCPU:    resource.MustParse("4"),
							corev1.ResourceMemory: resource.MustParse("4Gi"),
						},
					},
				},
				TaskManager: flinkOp.TaskManagerSpec{
					Replicas: 1,
					Resources: corev1.ResourceRequirements{
						Limits: map[corev1.ResourceName]resource.Quantity{
							corev1.ResourceCPU:    resource.MustParse("4"),
							corev1.ResourceMemory: resource.MustParse("4Gi"),
						},
					},
				},
			},
		},
		GeneratedNameMaxLength: &generatedNameMaxLength,
	}

	flinkConfigSection = pluginsConfig.MustRegisterSubSection("flink", defaultConfig)
)
