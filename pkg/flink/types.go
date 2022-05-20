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
	pluginsCore "github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/core"
	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/flytek8s/config"
	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/utils"
)

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

type ClusterName string

func (cn ClusterName) Validate() error {
	return ValidateRegEx(string(cn), regexpFlinkClusterName)
}

func (cn ClusterName) String() string {
	return string(cn)
}

func NewClusterName(name string) (ClusterName, error) {
	cn := ClusterName(name)
	if err := cn.Validate(); err != nil {
		return ClusterName(""), nil
	}
	return cn, nil
}
