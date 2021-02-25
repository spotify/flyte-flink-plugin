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
	logUtils "github.com/lyft/flyteidl/clients/go/coreutils/logs"
	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
)

type stackdriverLogPlugin struct {
	// the name of the project in GCP that the logs are being published under
	gcpProject string
	// The Log resource name for which the logs are published under
	logResource string
}

func (s *stackdriverLogPlugin) GetTaskLog(podName, namespace, containerName, containerID, logName string) (core.TaskLog, error) {
	return core.TaskLog{
		Uri: fmt.Sprintf(
			"https://console.cloud.google.com/logs/viewer?project=%s&angularJsUrl=%%2Flogs%%2Fviewer%%3Fproject%%3D%s&resource=%s&advancedFilter=resource.labels.pod_name%%3A%s",
			s.gcpProject,
			s.gcpProject,
			s.logResource,
			podName,
		),
		Name:          logName,
		MessageFormat: core.TaskLog_JSON,
	}, nil
}

func NewStackdriverLogPlugin(gcpProject, logResource string) logUtils.LogPlugin {
	return &stackdriverLogPlugin{
		gcpProject:  gcpProject,
		logResource: logResource,
	}
}
