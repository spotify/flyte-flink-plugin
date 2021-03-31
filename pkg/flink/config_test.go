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
	"testing"

	"github.com/flyteorg/flytestdlib/config"
	"github.com/flyteorg/flytestdlib/config/viper"
	"gotest.tools/assert"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestLoadConfig(t *testing.T) {
	flinkConfig := GetFlinkConfig()
	assert.Assert(t, flinkConfig != nil)

	t.Run("uses defaults", func(t *testing.T) {
		assert.Equal(t, flinkConfig.JobManager.Memory, resource.MustParse("4Gi"))
		assert.Equal(t, flinkConfig.TaskManager.Cpu, resource.MustParse("4"))
		assert.Equal(t, flinkConfig.TaskManager.Memory, resource.MustParse("4Gi"))
		assert.Equal(t, flinkConfig.RemoteClusterConfig.Enabled, false)
	})

	t.Run("overrides defaults", func(t *testing.T) {
		assert.Equal(t, flinkConfig.TaskManager.Replicas, 4)
		assert.Equal(t, flinkConfig.JobManager.Cpu, resource.MustParse("3.5"))
		assert.Equal(t, flinkConfig.ServiceAccount, "flink-service-account")
		assert.Equal(t, *flinkConfig.GeneratedNameMaxLength, 50)
	})

	t.Run("sets properties with no defaults", func(t *testing.T) {
		assert.DeepEqual(t, flinkConfig.JobManager.NodeSelector, map[string]string{"gke-nodepool": "nodepool-1"})
		assert.DeepEqual(t, flinkConfig.TaskManager.NodeSelector, map[string]string{"gke-nodepool": "nodepool-2"})
		assert.Equal(t, flinkConfig.Image, "flink-image")
		assert.Assert(t, len(flinkConfig.FlinkProperties) > 0)
		assert.Equal(t, flinkConfig.FlinkPropertiesOverride["jobmanager.archive.fs.dir"], "flink-job-archive-dir")
	})

	t.Run("flink log configs", func(t *testing.T) {
		assert.Assert(t, len(flinkConfig.FlinkLogConfig) == 9)
		assert.Equal(t, flinkConfig.FlinkLogConfig["log4j.logger.org.apache.flink"], "INFO")
		assert.Equal(t, flinkConfig.FlinkLogConfig["log4j.appender.console"], "org.apache.log4j.ConsoleAppender")
	})

	t.Run("remote cluster", func(t *testing.T) {
		configAccessor := viper.NewAccessor(config.Options{
			StrictMode:  true,
			SearchPaths: []string{"testdata/config_remote_cluster.yaml"},
		})

		configAccessor.UpdateConfig(context.TODO())
		remoteFlinkConfig := GetFlinkConfig()

		remoteConfig := ClusterConfig{
			Enabled:  true,
			Endpoint: "127.0.0.1",
			Auth: Auth{
				TokenPath:  "/path/token",
				CaCertPath: "/path/cert",
			},
		}
		assert.DeepEqual(t, remoteFlinkConfig.RemoteClusterConfig, remoteConfig)
	})
}

func init() {
	configAccessor := viper.NewAccessor(config.Options{
		StrictMode:  true,
		SearchPaths: []string{"testdata/config.yaml"},
	})

	configAccessor.UpdateConfig(context.TODO())
}
