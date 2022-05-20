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
)

func TestLoadConfig(t *testing.T) {
	flinkConfig := GetFlinkConfig()
	assert.Assert(t, flinkConfig != nil)

	t.Run("remote cluster", func(t *testing.T) {
		config := GetFlinkConfig()
		remoteConfig := ClusterConfig{
			Enabled:  false,
			Endpoint: "127.0.0.1",
			Auth: Auth{
				TokenPath:  "/path/token",
				CaCertPath: "/path/cert",
			},
		}
		assert.DeepEqual(t, config.RemoteClusterConfig, remoteConfig)
	})
}

func init() {
	configAccessor := viper.NewAccessor(config.Options{
		StrictMode:  true,
		SearchPaths: []string{"testdata/config.yaml"},
	})

	configAccessor.UpdateConfig(context.TODO())
}
