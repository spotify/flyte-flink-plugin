package flink

import pluginsConfig "github.com/lyft/flyteplugins/go/tasks/config"

// Config ... Flink-specific configs
type Config struct {
	DefaultFlinkConfig map[string]string `json:"flink-config-default" pflag:",Key value pairs of default flink configuration that should be applied to every FlinkJob"`
	Image              string            `json:"image"`
	ServiceAccount     string            `json:"service-account"`
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
