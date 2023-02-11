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
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/k8s"
	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/utils"

	flinkOp "github.com/spotify/flink-on-k8s-operator/apis/flinkcluster/v1beta1"

	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/io"
	"github.com/flyteorg/flytestdlib/logger"
	structpb "github.com/golang/protobuf/ptypes/struct"
)

type FlinkTaskContext struct {
	ClusterName ClusterName
	Namespace   string
	Annotations map[string]string
	Labels      map[string]string
	Job         flinkIdl.FlinkJob
}

type FlinkTaskExecContext interface {
	TaskReader() pluginsCore.TaskReader
	TaskExecutionMetadata() pluginsCore.TaskExecutionMetadata
	InputReader() io.InputReader
}

func NewFlinkTaskContext(ctx context.Context, taskCtx FlinkTaskExecContext) (*FlinkTaskContext, error) {
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

	err = Validate(&job)
	if err != nil {
		return nil, errors.Wrapf(errors.BadTaskSpecification, err, "invalid FlinkJob [%v], failed to unmarshal", job)
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

	taskMetadata := taskCtx.TaskExecutionMetadata()
	cn, err := NewClusterName(taskMetadata.GetTaskExecutionID().GetGeneratedName())
	if err != nil {
		return nil, errors.Errorf(errors.BadTaskSpecification, "invalid cluster name [%v]", err.Error())
	}

	return &FlinkTaskContext{
		ClusterName: cn,
		Namespace:   taskMetadata.GetNamespace(),
		Annotations: GetDefaultAnnotations(taskMetadata),
		Labels:      GetDefaultLabels(taskMetadata),
		Job:         job,
	}, nil
}

type flinkResourceHandler struct{}

func (flinkResourceHandler) GetProperties() k8s.PluginProperties {
	config := GetFlinkConfig()
	props := k8s.PluginProperties{
		GeneratedNameMaxLength:          config.GeneratedNameMaxLength,
		DisableDeleteResourceOnFinalize: true,
	}

	if config.RemoteClusterConfig.Enabled {
		props.DisableInjectFinalizer = true
		props.DisableInjectOwnerReferences = true
	}

	return props
}

// Creates a new Job that will execute the main container as well as any generated types the result from the execution.
func (flinkResourceHandler) BuildResource(ctx context.Context, taskCtx pluginsCore.TaskExecutionContext) (client.Object, error) {
	// Start with default config values.
	config := GetFlinkConfig()
	flinkTaskCtx, err := NewFlinkTaskContext(ctx, taskCtx)
	if err != nil {
		return nil, errors.Wrapf(errors.BadTaskSpecification, err, "invalid Flink task context")
	}

	cluster, err := NewFlinkCluster(config, *flinkTaskCtx)
	if err != nil {
		return nil, errors.Wrapf(errors.BadTaskSpecification, err, "invalid Flink cluster")
	}

	return cluster, nil
}

func (flinkResourceHandler) BuildIdentityResource(ctx context.Context, taskCtx pluginsCore.TaskExecutionMetadata) (client.Object, error) {
	return &flinkOp.FlinkCluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindFlinkCluster,
			APIVersion: flinkOp.GroupVersion.String(),
		},
	}, nil
}

func (h flinkResourceHandler) OnAbort(ctx context.Context, tCtx pluginsCore.TaskExecutionContext, resource client.Object) (behavior k8s.AbortBehavior, err error) {
	var abortBehavior k8s.AbortBehavior

	annotationPatch, err := NewAnnotationPatch(flinkOp.ControlAnnotation, flinkOp.ControlNameJobCancel)

	if err != nil {
		logger.Error(ctx, "error observed in abort", err)
		return abortBehavior, err
	}

	patchOp := k8s.PatchResourceOperation{Patch: annotationPatch}
	abortBehavior = k8s.AbortBehaviorPatchDefaultResource(patchOp, false)
	return abortBehavior, nil
}

type FlinkClusterIdentifier struct {
	ClusterName string
	Namespace   string
}

func FlinkClusterTaskLogs(ctx context.Context, fci FlinkClusterIdentifier, config *Config) ([]*core.TaskLog, error) {
	var taskLogs []*core.TaskLog

	p, err := logs.InitializeLogPlugins(&config.LogConfig)
	if err != nil {
		return nil, err
	}

	if p == nil {
		return taskLogs, nil
	}

	jobLog, err := p.GetTaskLogs(tasklog.Input{
		PodName:   fci.ClusterName,
		Namespace: fci.Namespace,
		LogName:   "(Job)",
	})
	if err != nil {
		return nil, err
	}
	taskLogs = append(taskLogs, jobLog.TaskLogs...)

	return taskLogs, nil
}

func flinkClusterTaskInfo(ctx context.Context, flinkCluster *flinkOp.FlinkCluster) (*pluginsCore.TaskInfo, error) {
	var taskLogs []*core.TaskLog

	tl, err := FlinkClusterTaskLogs(ctx, FlinkClusterIdentifier{
		ClusterName: flinkCluster.Name,
		Namespace:   flinkCluster.Namespace,
	}, GetFlinkConfig())
	if err != nil {
		return nil, err
	}

	taskLogs = append(taskLogs, tl...)

	info := flinkIdl.FlinkExecutionInfo{}
	components := flinkCluster.Status.Components

	if jmi := components.JobManagerIngress; jmi != nil {
		info.JobManager = &flinkIdl.JobManagerExecutionInfo{
			IngressURLs: jmi.URLs,
		}
	}

	if job := components.Job; job != nil {
		info.Job = &flinkIdl.JobExecutionInfo{Id: job.ID}
	}

	customInfo := &structpb.Struct{}
	err = utils.MarshalStruct(&info, customInfo)
	if err != nil {
		return nil, err
	}

	return &pluginsCore.TaskInfo{
		Logs:       taskLogs,
		CustomInfo: customInfo,
	}, nil
}

func isSubmitterExitCodeRetryable(ctx context.Context, exitCode int32) bool {
	config := GetFlinkConfig()
	for _, ec := range config.NonRetryableExitCodes {
		if exitCode == ec {
			logger.Infof(ctx, "Found non-retryable exit code: %v", ec)
			return false
		}
	}
	return true
}

func flinkClusterJobPhaseInfo(ctx context.Context, jobStatus *flinkOp.JobStatus, occurredAt time.Time, info *pluginsCore.TaskInfo) pluginsCore.PhaseInfo {
	logger.Infof(ctx, "job_state: %s", jobStatus.State)

	msg := fmt.Sprintf("%s %s", jobStatus.ID, jobStatus.State)

	switch jobStatus.State {
	case flinkOp.JobStateCancelled:
		return pluginsCore.PhaseInfoFailure(errors.DownstreamSystemError, msg, info)
	case flinkOp.JobStateFailed, flinkOp.JobStateDeployFailed, flinkOp.JobStateLost:
		if isSubmitterExitCodeRetryable(ctx, jobStatus.SubmitterExitCode) {
			reason := fmt.Sprintf("Flink Job Failed with Error: %v (retryable)", jobStatus.FailureReasons)
			return pluginsCore.PhaseInfoRetryableFailure(errors.DownstreamSystemError, reason, info)
		}
		reason := fmt.Sprintf("Flink Job Failed with Error: %v (non-retryable)", jobStatus.FailureReasons)
		return pluginsCore.PhaseInfoFailure(nonRetryableFlyteCode, reason, info)
	case flinkOp.JobStateRunning:
		return pluginsCore.PhaseInfoRunning(pluginsCore.DefaultPhaseVersion, info)
	case flinkOp.JobStateUpdating, flinkOp.JobStatePending, flinkOp.JobStateDeploying, flinkOp.JobStateRestarting:
		return pluginsCore.PhaseInfoInitializing(occurredAt, pluginsCore.DefaultPhaseVersion, msg, info)
	case flinkOp.JobStateSucceeded:
		if jobStatus.SubmitterExitCode < 0 {
			return pluginsCore.PhaseInfoRunning(pluginsCore.DefaultPhaseVersion, info)
		}
		if jobStatus.SubmitterExitCode == 0 {
			return pluginsCore.PhaseInfoSuccess(info)
		}
		if isSubmitterExitCodeRetryable(ctx, jobStatus.SubmitterExitCode) {
			reason := fmt.Sprintf("Flink jobsubmitter exited with non-zero exit code: %v (retryable)", jobStatus.FailureReasons)
			return pluginsCore.PhaseInfoRetryableFailure(errors.DownstreamSystemError, reason, info)
		}
		reason := fmt.Sprintf("Flink jobsubmitter exited with non-zero exit code: %v (non-retryable)", jobStatus.FailureReasons)
		return pluginsCore.PhaseInfoFailure(nonRetryableFlyteCode, reason, info)
	default:
		msg := fmt.Sprintf("job id: %s with unknown state: %s", jobStatus.ID, jobStatus.State)
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
		return pluginsCore.PhaseInfoWaitingForResourcesInfo(occurredAt, pluginsCore.DefaultPhaseVersion, "cluster starting", info), nil
	case flinkOp.ClusterStateRunning:
		return flinkClusterJobPhaseInfo(ctx, jobStatus, occurredAt, info), nil
	case flinkOp.ClusterStateStopped, flinkOp.ClusterStateStopping, flinkOp.ClusterStatePartiallyStopped:
		return flinkClusterJobPhaseInfo(ctx, jobStatus, occurredAt, info), nil
	}

	return pluginsCore.PhaseInfoRunning(pluginsCore.DefaultPhaseVersion, info), nil
}

func (flinkResourceHandler) GetTaskPhase(ctx context.Context, pluginContext k8s.PluginContext, resource client.Object) (pluginsCore.PhaseInfo, error) {
	app := resource.(*flinkOp.FlinkCluster)
	return flinkClusterPhaseInfo(ctx, app, time.Now())
}
