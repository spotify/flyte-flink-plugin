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
	"github.com/flyteorg/flyteplugins/go/tasks/logs"
	flinkOp "github.com/spotify/flink-on-k8s-operator/apis/flinkcluster/v1beta1"
)

// Config ... Flink-specific configs
type Config struct {
	DefaultFlinkCluster     flinkOp.FlinkCluster `json:"defaultFlinkCluster"`
	FlinkPropertiesOverride map[string]string    `json:"flinkPropertiesOverride" pflag:",Key value pairs of flink properties to be overridden in every FlinkJob"`
	LogConfig               logs.LogConfig       `json:"logs"`
	GeneratedNameMaxLength  *int                 `json:"generatedNameMaxLength" pflag:"Specifies the length of TaskExecutionID generated name. default: 50"`
	RemoteClusterConfig     ClusterConfig        `json:"remoteClusterConfig" pflag:"Configuration of remote K8s cluster for array jobs"`
	NonRetryableExitCodes   []int32              `json:"nonRetryableExitCodes" pfFlag:"Defines which job submitter exit codes should not be retried"`
	NonRetryableFlyteCode   *string              `json:"nonRetryableFlyteCode,omitempty" pfFlag:"Defines which code should be returned in case of nonRetryable exit codes"`
}

func GetFlinkConfig() *Config {
	return flinkConfigSection.GetConfig().(*Config)
}

// This method should be used for unit testing only
func setFlinkConfig(cfg *Config) error {
	return flinkConfigSection.SetConfig(cfg)
}
