package flink

import (
	"context"
	"testing"

	"github.com/lyft/flytestdlib/config"
	"github.com/lyft/flytestdlib/config/viper"
	"gotest.tools/assert"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestLoadConfig(t *testing.T) {
	config := GetFlinkConfig()
	assert.Assert(t, config != nil)

	assert.Equal(t, config.Image, "flink-image")
	assert.Equal(t, config.JobManager.Cpu, resource.MustParse("3.5"))
	assert.Equal(t, config.JobManager.Memory, resource.MustParse("4Gi"))
	assert.Equal(t, config.TaskManager.Cpu, resource.MustParse("4"))
	assert.Equal(t, config.TaskManager.Memory, resource.MustParse("4Gi"))
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
