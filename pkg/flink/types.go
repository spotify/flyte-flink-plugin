package flink

import (
	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	pluginsCore "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"
	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/flytek8s/config"
	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/utils"
)

type FlinkProperties map[string]string

func BuildFlinkProperties(config *Config, flinkJob flinkIdl.FlinkJob) FlinkProperties {
	// Start with default config values.
	flinkProperties := make(map[string]string)
	for k, v := range config.DefaultFlinkConfig {
		flinkProperties[k] = v
	}

	for k, v := range flinkJob.GetFlinkProperties() {
		flinkProperties[k] = v
	}

	return flinkProperties
}

type Annotations map[string]string

func GetDefaultAnnotations(taskCtx pluginsCore.TaskExecutionContext) Annotations {
	return utils.UnionMaps(
		config.GetK8sPluginConfig().DefaultAnnotations,
		utils.CopyMap(taskCtx.TaskExecutionMetadata().GetAnnotations()),
	)
}

type Labels map[string]string

func GetDefaultLabels(taskCtx pluginsCore.TaskExecutionContext) Labels {
	return utils.UnionMaps(
		config.GetK8sPluginConfig().DefaultLabels,
		utils.CopyMap(taskCtx.TaskExecutionMetadata().GetLabels()),
	)
}
