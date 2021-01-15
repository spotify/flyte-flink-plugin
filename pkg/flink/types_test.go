package flink

import (
	"testing"

	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	"gotest.tools/assert"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestBuildFlinkProperties(t *testing.T) {
	flinkProperties := BuildFlinkProperties(GetFlinkConfig(), flinkIdl.FlinkJob{})
	assert.Assert(t, len(flinkProperties) > 0)

	assert.Equal(t, flinkProperties.GetResourceQuantity("kubernetes.jobmanager.cores"), resource.MustParse("4"))
	assert.Equal(t, flinkProperties.GetResourceQuantity("kubernetes.jobmanager.memory"), resource.MustParse("4Gi"))

	assert.Equal(t, flinkProperties.GetInt("kubernetes.taskmanager.replicas"), 4)
	assert.Equal(t, flinkProperties.GetResourceQuantity("kubernetes.taskmanager.cores"), resource.MustParse("4"))
	assert.Equal(t, flinkProperties.GetResourceQuantity("kubernetes.taskmanager.memory"), resource.MustParse("4Gi"))
}

func TestBuildFlinkPropertiesOverrides(t *testing.T) {
	flinkJob := flinkIdl.FlinkJob{
		FlinkProperties: map[string]string{
			"kubernetes.taskmanager.replicas": "1",
		},
	}
	flinkProperties := BuildFlinkProperties(GetFlinkConfig(), flinkJob)
	assert.Equal(t, flinkProperties.GetInt("kubernetes.taskmanager.replicas"), 1)
}
