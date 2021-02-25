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
	"strconv"

	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	pluginsCore "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"
	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/flytek8s/config"
	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/utils"
)

type FlinkProperties map[string]string

func BuildFlinkProperties(config *Config, flinkJob flinkIdl.FlinkJob) FlinkProperties {
	// Start with default config values.
	flinkProperties := make(map[string]string)
	for k, v := range config.FlinkProperties {
		flinkProperties[k] = v
	}

	for k, v := range flinkJob.GetFlinkProperties() {
		flinkProperties[k] = v
	}

	for k, v := range config.FlinkPropertiesOverride {
		flinkProperties[k] = v
	}

	return flinkProperties
}

func (fp FlinkProperties) GetInt(key string) int {
	value, err := strconv.Atoi(fp[key])
	if err != nil {
		panic(fmt.Errorf("cannot parse '%v': %v", fp[key], err))
	}

	return value
}

type Annotations map[string]string

func GetDefaultAnnotations(taskCtx pluginsCore.TaskExecutionMetadata) Annotations {
	return utils.UnionMaps(
		config.GetK8sPluginConfig().DefaultAnnotations,
		utils.CopyMap(taskCtx.GetAnnotations()),
	)
}

type Labels map[string]string

func GetDefaultLabels(taskCtx pluginsCore.TaskExecutionMetadata) Labels {
	return utils.UnionMaps(
		config.GetK8sPluginConfig().DefaultLabels,
		utils.CopyMap(taskCtx.GetLabels()),
	)
}
