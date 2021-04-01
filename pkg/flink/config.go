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
	pluginsConfig "github.com/flyteorg/flyteplugins/go/tasks/config"
	"github.com/flyteorg/flyteplugins/go/tasks/logs"
	flinkOp "github.com/spotify/flink-on-k8s-operator/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

type DefaultFlinkCluster = flinkOp.FlinkCluster

// Config ... Flink-specific configs
type Config struct {
	DefaultFlinkCluster     DefaultFlinkCluster `json:"defaultFlinkCluster"`
	FlinkPropertiesOverride map[string]string   `json:"flinkPropertiesOverride" pflag:",Key value pairs of flink properties to be overridden in every FlinkJob"`
	LogConfig               logs.LogConfig      `json:"logs"`
	GeneratedNameMaxLength  *int                `json:"generatedNameMaxLength" pflag:"Specifies the length of TaskExecutionID generated name. default: 50"`
	RemoteClusterConfig     ClusterConfig       `json:"remoteClusterConfig" pflag:"Configuration of remote K8s cluster for array jobs"`
}

var (
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

func GetFlinkConfig() *Config {
	return flinkConfigSection.GetConfig().(*Config)
}

// This method should be used for unit testing only
func setFlinkConfig(cfg *Config) error {
	return flinkConfigSection.SetConfig(cfg)
}
