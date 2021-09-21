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
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	"reflect"
	"testing"

	flinkOp "github.com/spotify/flink-on-k8s-operator/api/v1beta1"
	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	"gotest.tools/assert"
	corev1 "k8s.io/api/core/v1"
)

var artifacts = []string{"gs://bucket/a.jar", "gs://bucket/b.jar", "gs://bucket/c.jar"}

func TestBuildFlinkClusterSpecValid(t *testing.T) {
	parallelism := int32(10)
	jobIdl := flinkIdl.FlinkJob{
		JarFiles:    artifacts,
		Parallelism: parallelism,
		FlinkProperties: map[string]string{
			"taskmanager.numberOfTaskSlots":            "1",
			"metrics.reporter.promgateway.groupingKey": `namespace={{.Namespace}};cluster={{.ClusterName}};execution_id={{index .Labels "execution-id"}}`,
		},
	}
	config := GetFlinkConfig()

	flinkCtx := FlinkTaskContext{
		ClusterName: ClusterName("generated-name"),
		Namespace:   "test-namespace",
		Annotations: make(map[string]string),
		Labels:      map[string]string{"execution-id": "1"},
		Job:         jobIdl,
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

	job := cluster.Spec.Job
	assert.Equal(t, *job.Parallelism, parallelism)
	assert.Equal(t, len(job.Volumes), 1)
	// first one is set through config
	assert.Equal(t, job.Volumes[0], corev1.Volume{Name: "jars"})
	assert.Equal(t, len(job.VolumeMounts), 1)
	// first one is set through config
	assert.Equal(t, job.VolumeMounts[0], corev1.VolumeMount{
		Name:      "jars",
		MountPath: "/jars",
	})

	assert.Equal(t, string(job.CleanupPolicy.AfterJobSucceeds), flinkOp.CleanupActionDeleteCluster)
	assert.Equal(t, cluster.Spec.FlinkProperties["metrics.reporter.promgateway.groupingKey"], "namespace=test-namespace;cluster=generated-name;execution_id=1")
}

func TestWithPersistentVolume(t *testing.T) {
	job := flinkIdl.FlinkJob{
		JarFiles: artifacts,
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
		ClusterName: ClusterName("generated-name"),
		Namespace:   "test-namespace",
		Annotations: make(map[string]string),
		Labels:      make(map[string]string),
		Job:         job,
	}

	cluster, err := NewFlinkCluster(config, flinkCtx)

	assert.NilError(t, err)
	assert.Equal(t, cluster.Spec.Image.Name, "flink-image")
	assert.Equal(t, len(cluster.Spec.TaskManager.VolumeClaimTemplates), 1)
	// it will include cache-colume and the required pv
	assert.Equal(t, len(cluster.Spec.TaskManager.VolumeMounts), 2)
	assert.Equal(t, len(cluster.Spec.TaskManager.Volumes), 1)
	assert.Assert(t, cluster.Spec.FlinkProperties[flinkIoTmpDirsProperty] == volumeClaimMountPath)
}

func TestBuildFlinkClusterSpecInvalid(t *testing.T) {
	job := flinkIdl.FlinkJob{
		JarFiles: artifacts,
		FlinkProperties: map[string]string{
			"taskmanager.numberOfTaskSlots": "1",
		},
	}

	// Use empty config
	config := &Config{}

	flinkCtx := FlinkTaskContext{
		ClusterName: ClusterName("generated-name"),
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
		JarFiles: artifacts,
		FlinkProperties: map[string]string{
			"taskmanager.numberOfTaskSlots": "1",
		},
		ServiceAccount: "flink-user-service-account",
	}
	config := GetFlinkConfig()

	flinkCtx := FlinkTaskContext{
		ClusterName: ClusterName("generated-name"),
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
		JarFiles: artifacts,
		FlinkProperties: map[string]string{
			"taskmanager.numberOfTaskSlots": "1",
		},
		Image: "flink-custom-image",
	}
	config := GetFlinkConfig()

	flinkCtx := FlinkTaskContext{
		ClusterName: ClusterName("generated-name"),
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
		JarFiles: artifacts,
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
		ClusterName: ClusterName("generated-name"),
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

func TestBuildFlinkClusterSpecJobCommand(t *testing.T) {
	job := flinkIdl.FlinkJob{
		JarFiles: []string{},
		FlinkProperties: map[string]string{
			"taskmanager.numberOfTaskSlots": "1",
		},
		Jflyte: &flinkIdl.JFlyte{
			IndexFileLocation: "gs://bucket/index-file.json",
			Artifacts: []*flinkIdl.JFlyte_Artifact{
				{
					Location: "gs://bucket/artifact0.jar",
					Name:     "artifact0.jar",
				},
				{
					Location: "gs://bucket/artifact1.jar",
					Name:     "artifact1.jar",
				},
				{
					Location: "gs://bucket/artifact2.jar",
					Name:     "artifact2.jar",
				},
			},
		},
	}
	config := GetFlinkConfig()

	flinkCtx := FlinkTaskContext{
		ClusterName: ClusterName("generated-name"),
		Namespace:   "test-namespace",
		Annotations: make(map[string]string),
		Labels:      make(map[string]string),
		Job:         job,
	}

	cluster, err := NewFlinkCluster(config, flinkCtx)

	assert.NilError(t, err)
	assert.Equal(t, len(cluster.Spec.Job.InitContainers), 1)

	initCont := cluster.Spec.Job.InitContainers[0]

	assert.Assert(t, reflect.DeepEqual(initCont.Command, []string{"/bin/sh"}))

	args := []string{
		"-c",
		"apt install -y zip && " +
			"mkdir lib && " +
			`if [ -n "${GOOGLE_APPLICATION_CREDENTIALS}" ]; then gcloud auth activate-service-account --key-file $GOOGLE_APPLICATION_CREDENTIALS; fi && ` +
			"gsutil -m cp" +
			" gs://bucket/artifact0.jar" +
			" gs://bucket/artifact1.jar" +
			" gs://bucket/artifact2.jar" +
			" lib && " +
			"zip -r job.jar .",
	}
	assert.Assert(t, reflect.DeepEqual(initCont.Args, args))
}

func TestBuildAnnotationPatch(t *testing.T) {
	patch, err := NewAnnotationPatch("testKey", "testValue")
	assert.Equal(t, patch.Type(), types.MergePatchType)

	bytes, err := patch.Data(nil)

	var jsonData map[string]interface{}
	err = json.Unmarshal(bytes, &jsonData)

	assert.DeepEqual(
		t,
		jsonData["metadata"].(map[string]interface{})["annotations"].(map[string]interface{}),
		map[string]interface{} {"testKey": "testValue"},
	)

	assert.NilError(t, err)
}
