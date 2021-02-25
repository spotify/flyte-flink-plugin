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
	"context"
	"strings"

	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery"
	"github.com/lyft/flytestdlib/promutils"

	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"
	pluginsCore "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"

	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/k8s"
	"k8s.io/client-go/kubernetes/scheme"

	"github.com/lyft/flytepropeller/pkg/controller/nodes/task/backoff"
	propeller "github.com/lyft/flytepropeller/pkg/controller/nodes/task/k8s"

	flinkOp "github.com/spotify/flink-on-k8s-operator/api/v1beta1"
)

type flinkSetupContext struct {
	underlying core.SetupContext
	kubeClient core.KubeClient
}

func (f flinkSetupContext) EnqueueOwner() core.EnqueueOwner {
	return f.underlying.EnqueueOwner()
}

// provides a k8s specific owner kind
func (f flinkSetupContext) OwnerKind() string {
	return f.underlying.OwnerKind()
}

// a metrics scope to publish stats under
func (f flinkSetupContext) MetricsScope() promutils.Scope {
	return f.underlying.MetricsScope()
}

// A kubernetes client to the bound cluster
func (f flinkSetupContext) KubeClient() core.KubeClient {
	return f.kubeClient
}

// Returns a secret manager that can retrieve configured secrets for this plugin
func (f flinkSetupContext) SecretManager() core.SecretManager {
	return f.underlying.SecretManager()
}

// Returns a resource negotiator that the plugin can register resource quota against
func (f flinkSetupContext) ResourceRegistrar() core.ResourceRegistrar {
	return f.underlying.ResourceRegistrar()
}

func init() {
	if err := flinkOp.AddToScheme(scheme.Scheme); err != nil {
		panic(err)
	}

	kpe := k8s.PluginEntry{
		ID:                  FlinkTaskType,
		RegisteredTaskTypes: []pluginsCore.TaskType{FlinkTaskType},
		ResourceToWatch:     &flinkOp.FlinkCluster{},
		Plugin:              flinkResourceHandler{},
		IsDefault:           false,
		DefaultForTaskTypes: []pluginsCore.TaskType{FlinkTaskType},
	}

	monitorIndex := propeller.NewResourceMonitorIndex()

	pluginmachinery.PluginRegistry().RegisterCorePlugin(
		core.PluginEntry{
			ID:                  strings.ToLower(kpe.ID),
			RegisteredTaskTypes: kpe.RegisteredTaskTypes,
			LoadPlugin: func(ctx context.Context, iCtx core.SetupContext) (plugin core.Plugin, e error) {
				var kubeClient core.KubeClient

				remoteClusterConfig := GetFlinkConfig().RemoteClusterConfig
				if remoteClusterConfig.Enabled {
					client, err := GetK8sClient(remoteClusterConfig)
					if err != nil {
						return nil, err
					}

					kubeClient = client
				} else {
					kubeClient = iCtx.KubeClient()
				}

				setupCtx := flinkSetupContext{underlying: iCtx, kubeClient: kubeClient}

				backOffController := backoff.NewController(ctx)

				return propeller.NewPluginManagerWithBackOff(ctx, setupCtx, kpe, backOffController, monitorIndex)
			},
			IsDefault:           kpe.IsDefault,
			DefaultForTaskTypes: []pluginsCore.TaskType{FlinkTaskType},
		})
}
