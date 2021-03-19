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
	"context"
	"fmt"
	"time"

	"github.com/flyteorg/flyteplugins/go/tasks/errors"
	"github.com/flyteorg/flyteplugins/go/tasks/logs"
	pluginsCore "github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/core"
	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/tasklog"

	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/k8s"
	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/utils"

	flinkOp "github.com/spotify/flink-on-k8s-operator/api/v1beta1"

	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/flyteorg/flytestdlib/logger"
)

type flinkResourceHandler struct{}

func (flinkResourceHandler) GetProperties() pluginsCore.PluginProperties {
	return pluginsCore.PluginProperties{}
}

// Creates a new Job that will execute the main container as well as any generated types the result from the execution.
func (flinkResourceHandler) BuildResource(ctx context.Context, taskCtx pluginsCore.TaskExecutionContext) (k8s.Resource, error) {

	taskTemplate, err := taskCtx.TaskReader().Read(ctx)
	if err != nil {
		return nil, errors.Errorf(errors.BadTaskSpecification, "unable to fetch task specification [%v]", err.Error())
	} else if taskTemplate == nil {
		return nil, errors.Errorf(errors.BadTaskSpecification, "nil task specification")
	}

	job := flinkIdl.FlinkJob{}
	err = utils.UnmarshalStruct(taskTemplate.GetCustom(), &job)
	if err != nil {
		return nil, errors.Wrapf(errors.BadTaskSpecification, err, "invalid TaskSpecification [%v], failed to unmarshal", taskTemplate.GetCustom())
	}

	taskInput, err := taskCtx.InputReader().Get(ctx)
	if err != nil {
		return nil, errors.Errorf(errors.BadTaskSpecification, "unable to fetch task inputs [%v]", err.Error())
	}

	// add task input literals to flink job args
	inputs := taskInput.GetLiterals()
	args, err := literalMapToFlinkJobArgs(inputs)
	if err != nil {
		return nil, errors.Errorf(errors.BadTaskSpecification, "not support input arg type [%v]", err.Error())
	}
	job.Args = append(job.Args, args...)

	// Start with default config values.
	config := GetFlinkConfig()

	return BuildFlinkClusterSpec(taskCtx.TaskExecutionMetadata(), job, config)
}

func (flinkResourceHandler) BuildIdentityResource(ctx context.Context, taskCtx pluginsCore.TaskExecutionMetadata) (k8s.Resource, error) {
	return &flinkOp.FlinkCluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindFlinkCluster,
			APIVersion: flinkOp.GroupVersion.String(),
		},
	}, nil
}

func flinkClusterTaskLogs(ctx context.Context, logPlugin logUtils.LogPlugin, flinkCluster *flinkOp.FlinkCluster, via string) ([]*core.TaskLog, error) {
	var taskLogs []*core.TaskLog
	jobManagerStatus := flinkCluster.Status.Components.JobManagerStatefulSet
	taskManagerStatus := flinkCluster.Status.Components.TaskManagerStatefulSet
	jobStatus := flinkCluster.Status.Components.Job

	jobManagerLogName := fmt.Sprintf("JobManager Logs (via %s)", via)
	jobManagerLog, err := logPlugin.GetTaskLog(jobManagerStatus.Name, flinkCluster.Namespace, "", "", jobManagerLogName)
	if err != nil {
		return nil, err
	}
	taskLogs = append(taskLogs, &jobManagerLog)

	taskManagerLogName := fmt.Sprintf("TaskManager Logs (via %s)", via)
	taskManagerLog, err := logPlugin.GetTaskLog(taskManagerStatus.Name, flinkCluster.Namespace, "", "", taskManagerLogName)
	if err != nil {
		return nil, err
	}
	taskLogs = append(taskLogs, &taskManagerLog)

	if jobStatus != nil {
		jobLogName := fmt.Sprintf("Job Logs (via %s)", via)
		jobLog, err := logPlugin.GetTaskLog(jobStatus.Name, flinkCluster.Namespace, "", "", jobLogName)
		if err != nil {
			return nil, err
		}
		taskLogs = append(taskLogs, &jobLog)
	}

	return taskLogs, nil
}

func flinkClusterTaskInfo(ctx context.Context, flinkCluster *flinkOp.FlinkCluster) (*pluginsCore.TaskInfo, error) {
	var taskLogs []*core.TaskLog
	customInfoMap := make(map[string]interface{})

	logConfig := logs.GetLogConfig()

	if logConfig.IsKubernetesEnabled {
		logPlugin := tasklog.NewKubernetesLogPlugin(logConfig.KubernetesURL)
		tl, err := flinkClusterTaskLogs(ctx, logPlugin, flinkCluster, "Kubernetes")
		if err != nil {
			return nil, err
		}

		taskLogs = append(taskLogs, tl...)
	}

	if logConfig.IsStackDriverEnabled {
		logPlugin := NewStackdriverLogPlugin(logConfig.GCPProjectName, logConfig.StackdriverLogResourceName)
		tl, err := flinkClusterTaskLogs(ctx, logPlugin, flinkCluster, "Stackdriver")
		if err != nil {
			return nil, err
		}

		taskLogs = append(taskLogs, tl...)
	}

	if jmi := flinkCluster.Status.Components.JobManagerIngress; jmi != nil {
		customInfoMap["jobmanager-ingress-urls"] = jmi.URLs
	}

	customInfo, err := utils.MarshalObjToStruct(customInfoMap)
	if err != nil {
		return nil, err
	}

	return &pluginsCore.TaskInfo{
		Logs:       taskLogs,
		CustomInfo: customInfo,
	}, nil
}

func flinkClusterJobPhaseInfo(ctx context.Context, jobStatus *flinkOp.JobStatus, occurredAt time.Time, info *pluginsCore.TaskInfo) pluginsCore.PhaseInfo {
	logger.Infof(ctx, "job_state: %s", jobStatus.State)

	msg := fmt.Sprintf("%s %s", jobStatus.Name, jobStatus.State)

	switch jobStatus.State {
	case flinkOp.JobStateCancelled:
		return pluginsCore.PhaseInfoFailure(errors.DownstreamSystemError, msg, info)
	case flinkOp.JobStateFailed:
		return pluginsCore.PhaseInfoRetryableFailure(errors.DownstreamSystemError, msg, info)
	case flinkOp.JobStateRunning:
		return pluginsCore.PhaseInfoRunning(pluginsCore.DefaultPhaseVersion, info)
	case flinkOp.JobStateUpdating, flinkOp.JobStatePending:
		return pluginsCore.PhaseInfoInitializing(occurredAt, pluginsCore.DefaultPhaseVersion, msg, info)
	case flinkOp.JobStateSucceeded:
		return pluginsCore.PhaseInfoSuccess(info)
	default:
		return pluginsCore.PhaseInfoFailure(errors.DownstreamSystemError, msg, info)
	}
}

func flinkClusterPhaseInfo(ctx context.Context, app *flinkOp.FlinkCluster, occurredAt time.Time) (pluginsCore.PhaseInfo, error) {
	info, err := flinkClusterTaskInfo(ctx, app)
	if err != nil {
		return pluginsCore.PhaseInfoUndefined, err
	}

	jobStatus := app.Status.Components.Job

	logger.Infof(ctx, "cluster_state: %s", app.Status.State)

	switch app.Status.State {
	case flinkOp.ClusterStateCreating, flinkOp.ClusterStateReconciling, flinkOp.ClusterStateUpdating:
		return pluginsCore.PhaseInfoWaitingForResources(occurredAt, pluginsCore.DefaultPhaseVersion, "cluster starting"), nil
	case flinkOp.ClusterStateRunning:
		return flinkClusterJobPhaseInfo(ctx, jobStatus, occurredAt, info), nil
	case flinkOp.ClusterStateStopped, flinkOp.ClusterStateStopping, flinkOp.ClusterStatePartiallyStopped:
		return flinkClusterJobPhaseInfo(ctx, jobStatus, occurredAt, info), nil
	}

	return pluginsCore.PhaseInfoRunning(pluginsCore.DefaultPhaseVersion, info), nil
}

func (flinkResourceHandler) GetTaskPhase(ctx context.Context, pluginContext k8s.PluginContext, resource k8s.Resource) (pluginsCore.PhaseInfo, error) {
	app := resource.(*flinkOp.FlinkCluster)
	return flinkClusterPhaseInfo(ctx, app, time.Now())
}
