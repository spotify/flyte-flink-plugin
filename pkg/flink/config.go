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
	"fmt"
	"io/ioutil"
	"time"

	pluginsConfig "github.com/lyft/flyteplugins/go/tasks/config"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/resource"

	restclient "k8s.io/client-go/rest"
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

type ClusterConfig struct {
	Name     string `json:"name" pflag:",Friendly name of the remote cluster"`
	Endpoint string `json:"endpoint" pflag:", Remote K8s cluster endpoint"`
	Auth     Auth   `json:"auth" pflag:"-, Auth setting for the cluster"`
	Enabled  bool   `json:"enabled" pflag:", Boolean flag to enable or disable"`
}

type Auth struct {
	Type      string `json:"type" pflag:", Authentication type"`
	TokenPath string `json:"token-path" pflag:", Token path"`
	CertPath  string `json:"cert-path" pflag:", Certificate path"`
}

func (auth Auth) GetCA() ([]byte, error) {
	cert, err := ioutil.ReadFile(auth.CertPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read k8s CA cert from configured path")
	}
	return cert, nil
}

// GetToken ...
func (auth Auth) GetToken() (string, error) {
	token, err := ioutil.ReadFile(auth.TokenPath)
	if err != nil {
		return "", errors.Wrap(err, "failed to read k8s bearer token from configured path")
	}
	return string(token), nil
}

// Config ... Flink-specific configs
type Config struct {
	RemoteClusterConfig     ClusterConfig     `json:"remote-cluster-config" pflag:"-,Configuration of remote K8s cluster for array jobs"`
	FlinkProperties         map[string]string `json:"flink-properties-default" pflag:",Key value pairs of default flink properties that should be applied to every FlinkJob"`
	FlinkPropertiesOverride map[string]string `json:"flink-properties-override" pflag:",Key value pairs of flink properties to be overridden in every FlinkJob"`
	Image                   string            `json:"image"`
	ServiceAccount          string            `json:"service-account"`
	JobManager              JobManagerConfig  `json:"jobmanager"`
	TaskManager             TaskManagerConfig `json:"taskmanager"`
}

var (
	flinkConfigSection = pluginsConfig.MustRegisterSubSection("flink", &Config{})
	defaultResync      = 120 * time.Second
)

func GetFlinkConfig() *Config {
	return flinkConfigSection.GetConfig().(*Config)
}

// This method should be used for unit testing only
func setFlinkConfig(cfg *Config) error {
	return flinkConfigSection.SetConfig(cfg)
}


// RemoteClusterConfig ...
func RemoteClusterConfig(host string, auth Auth) (*restclient.Config, error) {
	tokenString, err := auth.GetToken()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get auth token: %+v", err))
	}

	caCert, err := auth.GetCA()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get auth CA: %+v", err))
	}

	tlsClientConfig := restclient.TLSClientConfig{}
	tlsClientConfig.CAData = caCert
	return &restclient.Config{
		Host:            host,
		TLSClientConfig: tlsClientConfig,
		BearerToken:     tokenString,
	}, nil
}
