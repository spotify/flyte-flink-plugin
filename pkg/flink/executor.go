package flink

import (
	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery"

	pluginsCore "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"

	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/k8s"
	"k8s.io/client-go/kubernetes/scheme"

	flinkOp "github.com/spotify/flink-on-k8s-operator/api/v1beta1"
)

func init() {
	if err := flinkOp.AddToScheme(scheme.Scheme); err != nil {
		panic(err)
	}

	pluginmachinery.PluginRegistry().RegisterK8sPlugin(
		k8s.PluginEntry{
			ID:                  FlinkTaskType,
			RegisteredTaskTypes: []pluginsCore.TaskType{FlinkTaskType},
			ResourceToWatch:     &flinkOp.FlinkCluster{},
			Plugin:              flinkResourceHandler{},
			IsDefault:           false,
			DefaultForTaskTypes: []pluginsCore.TaskType{FlinkTaskType},
		})
}
