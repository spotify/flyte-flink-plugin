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
	"testing"

	flinkOp "github.com/spotify/flink-on-k8s-operator/api/v1beta2"
	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	"gotest.tools/assert"
)

func TestMergeProperties(t *testing.T) {
	properties := MergeProperties(
		GetFlinkConfig().DefaultFlinkCluster.Spec.FlinkProperties,
		flinkIdl.FlinkJob{}.FlinkProperties,
		GetFlinkConfig().FlinkPropertiesOverride,
	)
	assert.Assert(t, len(properties) > 0)
}

func TestMergePropertiesFullOverride(t *testing.T) {
	flinkJob := flinkIdl.FlinkJob{
		FlinkProperties: map[string]string{
			"akka.ask.timeout": "200s",
		},
	}

	properties := MergeProperties(
		GetFlinkConfig().DefaultFlinkCluster.Spec.FlinkProperties,
		flinkJob.FlinkProperties,
		GetFlinkConfig().FlinkPropertiesOverride,
	)
	assert.Equal(t, properties["akka.ask.timeout"], "200s")
}

func TestMergePropertiesFieldLevelOverride(t *testing.T) {
	config := Config{
		DefaultFlinkCluster: flinkOp.FlinkCluster{
			Spec: flinkOp.FlinkClusterSpec{
				FlinkProperties: map[string]string{"a": "A", "b": "B"},
			},
		},
		FlinkPropertiesOverride: map[string]string{"b": "BOverride", "c": "C"},
	}

	properties := MergeProperties(
		config.DefaultFlinkCluster.Spec.FlinkProperties,
		flinkIdl.FlinkJob{}.FlinkProperties,
		config.FlinkPropertiesOverride,
	)
	assert.Equal(t, properties["a"], "A")
	assert.Equal(t, properties["b"], "BOverride")
	assert.Equal(t, properties["c"], "C")
}
