package flink

import (
	"context"
	"testing"

	// pluginsConfig "github.com/lyft/flyteplugins/go/tasks/config"
	"github.com/lyft/flytestdlib/config"
	"github.com/lyft/flytestdlib/config/viper"
	"gotest.tools/assert"
)

func TestLoadConfig(t *testing.T) {
	flinkConfig := GetFlinkConfig()
	assert.Assert(t, flinkConfig != nil)

	assert.Equal(t, flinkConfig.Image, "flink-image")
	assert.Equal(t, flinkConfig.ServiceAccount, "flink-service-account")
}

func init() {
	configAccessor := viper.NewAccessor(config.Options{
		StrictMode:  true,
		SearchPaths: []string{"testdata/config.yaml"},
	})

	configAccessor.UpdateConfig(context.TODO())
}
