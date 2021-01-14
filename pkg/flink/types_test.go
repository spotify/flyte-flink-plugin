package flink

import (
	"testing"

	// pluginsConfig "github.com/lyft/flyteplugins/go/taspks/config"

	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	"gotest.tools/assert"
)

func TestBuildFlinkProperties(t *testing.T) {
	flinkProperties := BuildFlinkProperties(GetFlinkConfig(), flinkIdl.FlinkJob{})
	assert.Assert(t, len(flinkProperties) > 0)
}
