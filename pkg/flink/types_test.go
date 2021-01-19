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

func TestBuildFlinkPropertiesOverrides(t *testing.T) {
	flinkJob := flinkIdl.FlinkJob{
		FlinkProperties: map[string]string{
			"akka.ask.timeout": "200s",
		},
	}

	flinkProperties := BuildFlinkProperties(GetFlinkConfig(), flinkJob)
	assert.Equal(t, flinkProperties["akka.ask.timeout"], "200s")
}
