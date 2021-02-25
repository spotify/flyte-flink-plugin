package flink

import (
	pluginsConfig "github.com/lyft/flyteplugins/go/tasks/config"
	"k8s.io/apimachinery/pkg/api/resource"
)

type JobManagerConfig struct {
	Cpu    resource.Quantity `json:"cpu" pflag:"number of cores per pod"`
	Memory resource.Quantity `json:"memory" pflag:"amount of memory per pod"`
}

type TaskManagerConfig struct {
	Cpu      resource.Quantity `json:"cpu" pflag:"amout of cpu per pod"`
	Memory   resource.Quantity `json:"memory" pflag:"amount of memory per pod"`
	Replicas int               `json:"replicas" pflag:"number of replicas"`
}

// Config ... Flink-specific configs
type Config struct {
	FlinkProperties         map[string]string `json:"flink-properties-default" pflag:",Key value pairs of default flink properties that should be applied to every FlinkJob"`
	FlinkPropertiesOverride map[string]string `json:"flink-properties-override" pflag:",Key value pairs of flink properties to be overridden in every FlinkJob"`
	Image                   string            `json:"image"`
	ServiceAccount          string            `json:"service-account"`
	JobManager              JobManagerConfig  `json:"jobmanager"`
	TaskManager             TaskManagerConfig `json:"taskmanager"`
}

var (
	flinkConfigSection = pluginsConfig.MustRegisterSubSection("flink", &Config{})
)

func GetFlinkConfig() *Config {
	return flinkConfigSection.GetConfig().(*Config)
}

// This method should be used for unit testing only
func setFlinkConfig(cfg *Config) error {
	return flinkConfigSection.SetConfig(cfg)
}
