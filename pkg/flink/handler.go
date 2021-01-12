package flink

import (
	"context"
	"fmt"
	"time"

	logUtils "github.com/lyft/flyteidl/clients/go/coreutils/logs"
	"github.com/lyft/flyteplugins/go/tasks/errors"
	"github.com/lyft/flyteplugins/go/tasks/logs"
	pluginsCore "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"

	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/k8s"
	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/utils"

	flinkOp "github.com/regadas/flink-on-k8s-operator/api/v1beta1"

	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/lyft/flytestdlib/logger"
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

	flinkJob := flinkIdl.FlinkJob{}
	err = utils.UnmarshalStruct(taskTemplate.GetCustom(), &flinkJob)
	if err != nil {
		return nil, errors.Wrapf(errors.BadTaskSpecification, err, "invalid TaskSpecification [%v], failed to unmarshal", taskTemplate.GetCustom())
	}

	annotations := GetDefaultAnnotations(taskCtx)
	labels := GetDefaultLabels(taskCtx)

	container := taskTemplate.GetContainer()
	logger.Debugf(ctx, "Container %+v", container)

	// Start with default config values.
	flinkProperties := BuildFlinkProperties(flinkJob)

	jobManager := BuildJobManagerResource(flinkProperties, annotations, labels)
	taskManager := BuildTaskManagerResource(flinkProperties, annotations, labels)
	job := BuildJobResource(taskManager, flinkProperties, flinkJob)
	flinkCluster := BuildFlinkClusterResource(flinkProperties, annotations, labels, jobManager, taskManager, job)

	return &flinkCluster, nil
}

func (flinkResourceHandler) BuildIdentityResource(ctx context.Context, taskCtx pluginsCore.TaskExecutionMetadata) (k8s.Resource, error) {
	return &flinkOp.FlinkCluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindFlinkCluster,
			APIVersion: flinkOp.GroupVersion.String(),
		},
	}, nil
}

func flinkClusterTaskLogs(logPlugin logUtils.LogPlugin, flinkCluster *flinkOp.FlinkCluster, via string) ([]*core.TaskLog, error) {
	var taskLogs []*core.TaskLog
	jobManagerStatus := flinkCluster.Status.Components.JobManagerDeployment
	taskManagerStatus := flinkCluster.Status.Components.TaskManagerDeployment
	jobStatus := flinkCluster.Status.Components.Job

	if jobStatus == nil || jobStatus.Name == "" {
		return taskLogs, nil
	}

	jobLog, err := logPlugin.GetTaskLog(jobStatus.Name, flinkCluster.Namespace, "", "", fmt.Sprintf("Job Logs (via %s)", via))
	if err != nil {
		return nil, err
	}
	jobManagerLog, err := logPlugin.GetTaskLog(jobManagerStatus.Name, flinkCluster.Namespace, "", "", fmt.Sprintf("JobManager Logs (via %s)", via))
	if err != nil {
		return nil, err
	}
	taskManagerLog, err := logPlugin.GetTaskLog(taskManagerStatus.Name, flinkCluster.Namespace, "", "", fmt.Sprintf("TaskManager Logs (via %s)", via))
	if err != nil {
		return nil, err
	}

	return append(taskLogs, &jobLog, &jobManagerLog, &taskManagerLog), nil
}

func flinkClusterTaskInfo(flinkCluster *flinkOp.FlinkCluster) (*pluginsCore.TaskInfo, error) {
	var taskLogs []*core.TaskLog
	customInfoMap := make(map[string]string)

	customInfo, err := utils.MarshalObjToStruct(customInfoMap)
	if err != nil {
		return nil, err
	}

	logConfig := logs.GetLogConfig()

	if logConfig.IsKubernetesEnabled {
		logPlugin := logUtils.NewKubernetesLogPlugin(logConfig.KubernetesURL)
		tl, err := flinkClusterTaskLogs(logPlugin, flinkCluster, "Kubernetes")
		if err != nil {
			return nil, err
		}

		taskLogs = append(taskLogs, tl...)
	}

	if logConfig.IsStackDriverEnabled {
		logPlugin := NewStackdriverLogPlugin(logConfig.GCPProjectName, logConfig.StackdriverLogResourceName)
		tl, err := flinkClusterTaskLogs(logPlugin, flinkCluster, "Stackdriver")
		if err != nil {
			return nil, err
		}

		taskLogs = append(taskLogs, tl...)
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
	info, err := flinkClusterTaskInfo(app)
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
