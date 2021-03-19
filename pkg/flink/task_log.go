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

	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/tasklog"
)

type stackdriverLogPlugin struct {
	// the name of the project in GCP that the logs are being published under
	gcpProject string
	// The Log resource name for which the logs are published under
	logResource string
}

func (s *stackdriverLogPlugin) GetTaskLogs(input tasklog.Input) (core.TaskLog, error) {
	return core.TaskLog{
		Uri: fmt.Sprintf(
			"https://console.cloud.google.com/logs/viewer?project=%s&angularJsUrl=%%2Flogs%%2Fviewer%%3Fproject%%3D%s&resource=%s&advancedFilter=resource.labels.pod_name%%3A%s",
			s.gcpProject,
			s.gcpProject,
			s.logResource,
			input.PodName,
		),
		Name:          input.LogName,
		MessageFormat: core.TaskLog_JSON,
	}, nil
}

func NewStackdriverLogPlugin(gcpProject, logResource string) tasklog.Plugin {
	return &stackdriverLogPlugin{
		gcpProject:  gcpProject,
		logResource: logResource,
	}
}
