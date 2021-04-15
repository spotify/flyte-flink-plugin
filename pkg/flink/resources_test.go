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
	"testing"

	flinkOp "github.com/spotify/flink-on-k8s-operator/api/v1beta1"
	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	"gotest.tools/assert"
)

func TestBuildFlinkClusterSpecValid(t *testing.T) {
	job := flinkIdl.FlinkJob{
		JarFile: "job.jar",
		FlinkProperties: map[string]string{
			"taskmanager.numberOfTaskSlots": "1",
		},
	}
	config := GetFlinkConfig()

	flinkCtx := FlinkTaskContext{
		Name:        "generated-name",
		Namespace:   "test-namespace",
		Annotations: make(map[string]string),
		Labels:      make(map[string]string),
		Job:         job,
	}

	cluster, err := NewFlinkCluster(config, flinkCtx)

	assert.NilError(t, err)
	assert.Equal(t, cluster.Spec.Image.Name, "flink-image")
	assert.DeepEqual(t, cluster.Spec.JobManager.NodeSelector, map[string]string{"gke-nodepool": "nodepool-1"})
	assert.DeepEqual(t, cluster.Spec.TaskManager.NodeSelector, map[string]string{"gke-nodepool": "nodepool-2"})

	sidecars := cluster.Spec.JobManager.Sidecars
	assert.Assert(t, len(sidecars) == 1)
	assert.Equal(t, sidecars[0].Name, "sidecar")
	assert.Equal(t, sidecars[0].Image, "sidecar-image")

	assert.Equal(t, cluster.Spec.JobManager.AccessScope, "External")

	assert.Assert(t, cluster.Spec.JobManager.Ingress != nil)
	assert.Equal(t, *cluster.Spec.JobManager.Ingress.UseTLS, true)

	assert.Equal(t, cluster.Spec.Job.JarFile, job.JarFile)
	assert.Equal(t, len(cluster.Spec.Job.Volumes), 1)
	assert.Equal(t, len(cluster.Spec.Job.VolumeMounts), 1)
}

func TestWithPersistentVolume(t *testing.T) {
	job := flinkIdl.FlinkJob{
		JarFile: "job.jar",
		FlinkProperties: map[string]string{
			"taskmanager.numberOfTaskSlots": "1",
		},
		TaskManager: &flinkIdl.TaskManager{
			Resource: &flinkIdl.Resource{
				PersistentVolume: &flinkIdl.Resource_PersistentVolume{
					Type: flinkIdl.Resource_PersistentVolume_PD_SSD,
					Size: &flinkIdl.Resource_Quantity{String_: "100Gi"},
				},
			},
			Replicas: 1,
		},
	}
	config := GetFlinkConfig()

	flinkCtx := FlinkTaskContext{
		Name:        "generated-name",
		Namespace:   "test-namespace",
		Annotations: make(map[string]string),
		Labels:      make(map[string]string),
		Job:         job,
	}

	cluster, err := NewFlinkCluster(config, flinkCtx)

	assert.NilError(t, err)
	assert.Equal(t, cluster.Spec.Image.Name, "flink-image")
	assert.Assert(t, len(cluster.Spec.TaskManager.VolumeClaimTemplates) == 1)
	// it will include cache-colume and the required pv
	assert.Assert(t, len(cluster.Spec.TaskManager.VolumeMounts) == 2)
	assert.Assert(t, len(cluster.Spec.TaskManager.Volumes) == 1)
	assert.Assert(t, cluster.Spec.FlinkProperties[flinkIoTmpDirsProperty] == volumeClaimMountPath)
}

func TestBuildFlinkClusterSpecInvalid(t *testing.T) {
	job := flinkIdl.FlinkJob{
		FlinkProperties: map[string]string{
			"taskmanager.numberOfTaskSlots": "1",
		},
	}

	// Use empty config
	config := &Config{}

	flinkCtx := FlinkTaskContext{
		Name:        "generated-name",
		Namespace:   "test-namespace",
		Annotations: make(map[string]string),
		Labels:      make(map[string]string),
		Job:         job,
	}

	_, err := NewFlinkCluster(config, flinkCtx)
	assert.Error(t, err, "image name is unspecified")
}

func TestBuildFlinkClusterSpecServiceAccount(t *testing.T) {
	job := flinkIdl.FlinkJob{
		JarFile: "job.jar",
		FlinkProperties: map[string]string{
			"taskmanager.numberOfTaskSlots": "1",
		},
		ServiceAccount: "flink-user-service-account",
	}
	config := GetFlinkConfig()

	flinkCtx := FlinkTaskContext{
		Name:        "generated-name",
		Namespace:   "test-namespace",
		Annotations: make(map[string]string),
		Labels:      make(map[string]string),
		Job:         job,
	}

	cluster, err := NewFlinkCluster(config, flinkCtx)

	assert.NilError(t, err)
	assert.Equal(t, *cluster.Spec.ServiceAccountName, "flink-user-service-account")
}

func TestBuildFlinkClusterSpecImage(t *testing.T) {
	job := flinkIdl.FlinkJob{
		JarFile: "job.jar",
		FlinkProperties: map[string]string{
			"taskmanager.numberOfTaskSlots": "1",
		},
		Image: "flink-custom-image",
	}
	config := GetFlinkConfig()

	flinkCtx := FlinkTaskContext{
		Name:        "generated-name",
		Namespace:   "test-namespace",
		Annotations: make(map[string]string),
		Labels:      make(map[string]string),
		Job:         job,
	}

	cluster, err := NewFlinkCluster(config, flinkCtx)

	assert.NilError(t, err)
	assert.Equal(t, cluster.Spec.Image.Name, "flink-custom-image")
}

func TestBuildFlinkClusterWithIngress(t *testing.T) {
	job := flinkIdl.FlinkJob{
		JarFile: "job.jar",
		FlinkProperties: map[string]string{
			"taskmanager.numberOfTaskSlots": "1",
		},
		Image: "flink-custom-image",
	}

	config := GetFlinkConfig()
	config.DefaultFlinkCluster.Spec.JobManager.Ingress = &flinkOp.JobManagerIngressSpec{
		Annotations: map[string]string{
			"cluster-autoscaler.kubernetes.io/safe-to-evict": "false",
			"kubernetes.io/ingress.class":                    "gce-internal",
		},
	}

	flinkCtx := FlinkTaskContext{
		Name:        "generated-name",
		Namespace:   "test-namespace",
		Annotations: make(map[string]string),
		Labels:      make(map[string]string),
		Job:         job,
	}

	cluster, err := NewFlinkCluster(config, flinkCtx)
	assert.NilError(t, err)

	assert.Assert(t, cluster.Spec.JobManager.Ingress != nil)
	assert.DeepEqual(t, cluster.Spec.JobManager.Ingress.Annotations, map[string]string{
		"cluster-autoscaler.kubernetes.io/safe-to-evict": "false",
		"kubernetes.io/ingress.class":                    "gce-internal",
	})
}

func TestBuildFlinkClusterSpecInvalidClusterName(t *testing.T) {
	job := flinkIdl.FlinkJob{
		JarFile: "job.jar",
		FlinkProperties: map[string]string{
			"taskmanager.numberOfTaskSlots": "1",
		},
	}
	config := GetFlinkConfig()

	flinkCtx := FlinkTaskContext{
		Name:        ".generated-name",
		Namespace:   "test-namespace",
		Annotations: make(map[string]string),
		Labels:      make(map[string]string),
		Job:         job,
	}

	_, err := NewFlinkCluster(config, flinkCtx)
	assert.ErrorContains(t, err, "Validation error: ")
}
