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
	config := GetFlinkConfig()
	assert.Assert(t, config != nil)

	assert.Equal(t, config.Image, "flink-image")
	assert.Equal(t, config.ServiceAccount, "flink-service-account")
	assert.Equal(t, config.JobManager.Cpu, resource.MustParse("3.5"))
	assert.Equal(t, config.JobManager.Memory, resource.MustParse("4Gi"))
	assert.DeepEqual(t, config.JobManager.NodeSelector, map[string]string{"gke-nodepool": "nodepool-1"})
	assert.Equal(t, config.TaskManager.Cpu, resource.MustParse("4"))
	assert.Equal(t, config.TaskManager.Memory, resource.MustParse("4Gi"))
	assert.DeepEqual(t, config.TaskManager.NodeSelector, map[string]string{"gke-nodepool": "nodepool-2"})
	assert.Equal(t, config.TaskManager.Replicas, 4)
	assert.Assert(t, len(config.FlinkProperties) > 0)
	assert.Equal(t, config.FlinkPropertiesOverride["jobmanager.archive.fs.dir"], "flink-job-archive-dir")
}

func init() {
	configAccessor := viper.NewAccessor(config.Options{
		StrictMode:  true,
		SearchPaths: []string{"testdata/config.yaml"},
	})

	configAccessor.UpdateConfig(context.TODO())
}
