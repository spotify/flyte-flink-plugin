package flink

import (
	"testing"

	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	"gotest.tools/assert"
)

func TestBuildFlinkProperties(t *testing.T) {
	flinkProperties := BuildFlinkProperties(GetFlinkConfig(), flinkIdl.FlinkJob{})
	assert.Assert(t, len(flinkProperties) > 0)
}

func TestBuildFlinkPropertiesFullOverride(t *testing.T) {
	flinkJob := flinkIdl.FlinkJob{
		FlinkProperties: map[string]string{
			"akka.ask.timeout": "200s",
		},
	}

	flinkProperties := BuildFlinkProperties(GetFlinkConfig(), flinkJob)
	assert.Equal(t, flinkProperties["akka.ask.timeout"], "200s")
}

func TestBuildFlinkPropertiesFieldLevelOverride(t *testing.T) {
	config := Config{
		FlinkProperties:         map[string]string{"a": "A", "b": "B"},
		FlinkPropertiesOverride: map[string]string{"b": "BOverride", "c": "C"},
		Image:                   "",
		ServiceAccount:          "",
		JobManager:              JobManagerConfig{},
		TaskManager:             TaskManagerConfig{},
	}

	flinkProperties := BuildFlinkProperties(&config, flinkIdl.FlinkJob{})
	assert.Equal(t, flinkProperties["a"], "A")
	assert.Equal(t, flinkProperties["b"], "BOverride")
	assert.Equal(t, flinkProperties["c"], "C")
}
