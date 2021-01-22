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
