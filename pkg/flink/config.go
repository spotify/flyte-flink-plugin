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
	pluginsConfig "github.com/lyft/flyteplugins/go/tasks/config"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

type IngressConfig struct {
	Enabled bool `json:"enabled" pflag:"Enable JobManager ingress. default: false"`
	UseTLS  bool `json:"useTLS" pflag:"TLS use, default: false."`
}

type JobManagerConfig struct {
	Cpu          resource.Quantity  `json:"cpu" pflag:"number of cores per pod"`
	Memory       resource.Quantity  `json:"memory" pflag:"amount of memory per pod"`
	NodeSelector map[string]string  `json:"nodeSelector" pflag:"Annotates the JobManager resource with desired nodepool type"`
	Sidecars     []corev1.Container `json:"sidecars" pflag:"Sidecar containers running alongside with the JobManager container in the pod"`
	AccessScope  *string            `json:"accessScope" pflag:"Access scope of the JobManager service. Cluster (default), VPC, External, NodePort, Headless."`
	Ingress      IngressConfig      `json:"ingress" pflag:"Provide external access to JobManager UI/API."`
}

type TaskManagerConfig struct {
	Cpu          resource.Quantity  `json:"cpu" pflag:"amout of cpu per pod"`
	Memory       resource.Quantity  `json:"memory" pflag:"amount of memory per pod"`
	Replicas     int                `json:"replicas" pflag:"number of replicas"`
	NodeSelector map[string]string  `json:"nodeSelector" pflag:"Annotates the TasManager resource(s) with desired nodepool type"`
	Sidecars     []corev1.Container `json:"sidecars" pflag:"Sidecar containers running alongside with the TaskManager container in the pod"`
}

// Config ... Flink-specific configs
type Config struct {
	FlinkProperties         map[string]string `json:"flink-properties-default" pflag:",Key value pairs of default flink properties that should be applied to every FlinkJob"`
	FlinkPropertiesOverride map[string]string `json:"flink-properties-override" pflag:",Key value pairs of flink properties to be overridden in every FlinkJob"`
	Image                   string            `json:"image"`
	ServiceAccount          string            `json:"service-account"`
	JobManager              JobManagerConfig  `json:"jobmanager"`
	TaskManager             TaskManagerConfig `json:"taskmanager"`
}

var (
	flinkConfigSection = pluginsConfig.MustRegisterSubSection("flink", &Config{})
)

func GetFlinkConfig() *Config {
	return flinkConfigSection.GetConfig().(*Config)
}

// This method should be used for unit testing only
func setFlinkConfig(cfg *Config) error {
	return flinkConfigSection.SetConfig(cfg)
}
