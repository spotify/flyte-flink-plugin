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
	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"
	"github.com/lyft/flytepropeller/pkg/controller/executors"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

type kubeClientObj struct {
	client client.Client
	cache  cache.Cache
}

func (k *kubeClientObj) GetClient() client.Client {
	return k.client
}

func (k *kubeClientObj) GetCache() cache.Cache {
	return k.cache
}

func newKubeClientObj(c client.Client, cache cache.Cache) core.KubeClient {
	return &kubeClientObj{client: c, cache: cache}
}

// GetK8sClient ...
func GetK8sClient(config ClusterConfig) (core.KubeClient, error) {
	kubeConf, err := RemoteClusterConfig(config.Endpoint, config.Auth)
	if err != nil {
		return nil, err
	}

	mapper, err := apiutil.NewDynamicRESTMapper(kubeConf)
	if err != nil {
		return nil, err
	}

	cache, err := cache.New(kubeConf, cache.Options{Mapper: mapper})
	if err != nil {
		return nil, err
	}

	c, err := client.New(kubeConf, client.Options{Mapper: mapper})
	if err != nil {
		return nil, err
	}

	fallbackClient := executors.NewFallbackClient(&client.DelegatingClient{
		Reader: &client.DelegatingReader{
			CacheReader:  cache,
			ClientReader: c,
		},
		Writer:       c,
		StatusClient: c,
	}, c)

	return newKubeClientObj(fallbackClient, cache), nil
}
