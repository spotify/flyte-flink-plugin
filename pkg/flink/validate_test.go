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

	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	"gotest.tools/assert"
)

func TestValidFlinkJob(t *testing.T) {

	t.Run("minimal", func(t *testing.T) {
		job := flinkIdl.FlinkJob{
			JarFiles:  []string{"/foo/bar.jar"},
			MainClass: "com.spotify.job",
		}

		err := job.Validate()
		assert.NilError(t, err)
	})

	t.Run("extended", func(t *testing.T) {
		job := flinkIdl.FlinkJob{
			JarFiles:  []string{"a"},
			MainClass: "com.spotify.job",
			Args:      []string{},
			JobManager: &flinkIdl.JobManager{
				Resource: &flinkIdl.Resource{
					Cpu: &flinkIdl.Resource_Quantity{
						String_: "100",
					},
					Memory: &flinkIdl.Resource_Quantity{
						String_: "100Gi",
					},
					PersistentVolume: &flinkIdl.Resource_PersistentVolume{
						Type: 0,
					},
				},
			},
			TaskManager: &flinkIdl.TaskManager{
				Replicas: 1,
			},
		}

		err := job.Validate()
		assert.NilError(t, err)
	})
}

func TestInvalidFlinkJob(t *testing.T) {
	t.Run("jar", func(t *testing.T) {
		job := flinkIdl.FlinkJob{
			JarFiles:  []string{"gs://foo/bar"},
			MainClass: "com.spotify.job",
			Jflyte: &flinkIdl.JFlyte{
				Artifacts: []*flinkIdl.JFlyte_Artifact{
					{
						Name:     "foo",
						Location: `gs://invalid\jar`,
					},
				},
			},
		}

		err := job.Validate()
		assert.Error(t, err, `invalid FlinkJob.Jflyte: embedded message failed validation | caused by: invalid JFlyte.Artifacts[0]: embedded message failed validation | caused by: invalid JFlyte_Artifact.Location: value must be a valid URI | caused by: parse "gs://invalid\\jar": invalid character "\\" in host name`)
	})

	t.Run("mainClass", func(t *testing.T) {
		job := flinkIdl.FlinkJob{
			JarFiles: []string{"gs://foo/bar"},
		}

		err := job.Validate()
		assert.Error(t, err, `invalid FlinkJob.MainClass: value length must be at least 1 runes`)
	})
}
