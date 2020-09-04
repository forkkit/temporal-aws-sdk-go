package awsclients

import (
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type CodeDeployClient interface {
    AddTagsToOnPremisesInstances(ctx workflow.Context, input *codedeploy.AddTagsToOnPremisesInstancesInput) (*codedeploy.AddTagsToOnPremisesInstancesOutput, error)
    AddTagsToOnPremisesInstancesAsync(ctx workflow.Context, input *codedeploy.AddTagsToOnPremisesInstancesInput) *CodedeployAddTagsToOnPremisesInstancesResult

    BatchGetApplicationRevisions(ctx workflow.Context, input *codedeploy.BatchGetApplicationRevisionsInput) (*codedeploy.BatchGetApplicationRevisionsOutput, error)
    BatchGetApplicationRevisionsAsync(ctx workflow.Context, input *codedeploy.BatchGetApplicationRevisionsInput) *CodedeployBatchGetApplicationRevisionsResult

    BatchGetApplications(ctx workflow.Context, input *codedeploy.BatchGetApplicationsInput) (*codedeploy.BatchGetApplicationsOutput, error)
    BatchGetApplicationsAsync(ctx workflow.Context, input *codedeploy.BatchGetApplicationsInput) *CodedeployBatchGetApplicationsResult

    BatchGetDeploymentGroups(ctx workflow.Context, input *codedeploy.BatchGetDeploymentGroupsInput) (*codedeploy.BatchGetDeploymentGroupsOutput, error)
    BatchGetDeploymentGroupsAsync(ctx workflow.Context, input *codedeploy.BatchGetDeploymentGroupsInput) *CodedeployBatchGetDeploymentGroupsResult

    BatchGetDeploymentInstances(ctx workflow.Context, input *codedeploy.BatchGetDeploymentInstancesInput) (*codedeploy.BatchGetDeploymentInstancesOutput, error)
    BatchGetDeploymentInstancesAsync(ctx workflow.Context, input *codedeploy.BatchGetDeploymentInstancesInput) *CodedeployBatchGetDeploymentInstancesResult

    BatchGetDeploymentTargets(ctx workflow.Context, input *codedeploy.BatchGetDeploymentTargetsInput) (*codedeploy.BatchGetDeploymentTargetsOutput, error)
    BatchGetDeploymentTargetsAsync(ctx workflow.Context, input *codedeploy.BatchGetDeploymentTargetsInput) *CodedeployBatchGetDeploymentTargetsResult

    BatchGetDeployments(ctx workflow.Context, input *codedeploy.BatchGetDeploymentsInput) (*codedeploy.BatchGetDeploymentsOutput, error)
    BatchGetDeploymentsAsync(ctx workflow.Context, input *codedeploy.BatchGetDeploymentsInput) *CodedeployBatchGetDeploymentsResult

    BatchGetOnPremisesInstances(ctx workflow.Context, input *codedeploy.BatchGetOnPremisesInstancesInput) (*codedeploy.BatchGetOnPremisesInstancesOutput, error)
    BatchGetOnPremisesInstancesAsync(ctx workflow.Context, input *codedeploy.BatchGetOnPremisesInstancesInput) *CodedeployBatchGetOnPremisesInstancesResult

    ContinueDeployment(ctx workflow.Context, input *codedeploy.ContinueDeploymentInput) (*codedeploy.ContinueDeploymentOutput, error)
    ContinueDeploymentAsync(ctx workflow.Context, input *codedeploy.ContinueDeploymentInput) *CodedeployContinueDeploymentResult

    CreateApplication(ctx workflow.Context, input *codedeploy.CreateApplicationInput) (*codedeploy.CreateApplicationOutput, error)
    CreateApplicationAsync(ctx workflow.Context, input *codedeploy.CreateApplicationInput) *CodedeployCreateApplicationResult

    CreateDeployment(ctx workflow.Context, input *codedeploy.CreateDeploymentInput) (*codedeploy.CreateDeploymentOutput, error)
    CreateDeploymentAsync(ctx workflow.Context, input *codedeploy.CreateDeploymentInput) *CodedeployCreateDeploymentResult

    CreateDeploymentConfig(ctx workflow.Context, input *codedeploy.CreateDeploymentConfigInput) (*codedeploy.CreateDeploymentConfigOutput, error)
    CreateDeploymentConfigAsync(ctx workflow.Context, input *codedeploy.CreateDeploymentConfigInput) *CodedeployCreateDeploymentConfigResult

    CreateDeploymentGroup(ctx workflow.Context, input *codedeploy.CreateDeploymentGroupInput) (*codedeploy.CreateDeploymentGroupOutput, error)
    CreateDeploymentGroupAsync(ctx workflow.Context, input *codedeploy.CreateDeploymentGroupInput) *CodedeployCreateDeploymentGroupResult

    DeleteApplication(ctx workflow.Context, input *codedeploy.DeleteApplicationInput) (*codedeploy.DeleteApplicationOutput, error)
    DeleteApplicationAsync(ctx workflow.Context, input *codedeploy.DeleteApplicationInput) *CodedeployDeleteApplicationResult

    DeleteDeploymentConfig(ctx workflow.Context, input *codedeploy.DeleteDeploymentConfigInput) (*codedeploy.DeleteDeploymentConfigOutput, error)
    DeleteDeploymentConfigAsync(ctx workflow.Context, input *codedeploy.DeleteDeploymentConfigInput) *CodedeployDeleteDeploymentConfigResult

    DeleteDeploymentGroup(ctx workflow.Context, input *codedeploy.DeleteDeploymentGroupInput) (*codedeploy.DeleteDeploymentGroupOutput, error)
    DeleteDeploymentGroupAsync(ctx workflow.Context, input *codedeploy.DeleteDeploymentGroupInput) *CodedeployDeleteDeploymentGroupResult

    DeleteGitHubAccountToken(ctx workflow.Context, input *codedeploy.DeleteGitHubAccountTokenInput) (*codedeploy.DeleteGitHubAccountTokenOutput, error)
    DeleteGitHubAccountTokenAsync(ctx workflow.Context, input *codedeploy.DeleteGitHubAccountTokenInput) *CodedeployDeleteGitHubAccountTokenResult

    DeleteResourcesByExternalId(ctx workflow.Context, input *codedeploy.DeleteResourcesByExternalIdInput) (*codedeploy.DeleteResourcesByExternalIdOutput, error)
    DeleteResourcesByExternalIdAsync(ctx workflow.Context, input *codedeploy.DeleteResourcesByExternalIdInput) *CodedeployDeleteResourcesByExternalIdResult

    DeregisterOnPremisesInstance(ctx workflow.Context, input *codedeploy.DeregisterOnPremisesInstanceInput) (*codedeploy.DeregisterOnPremisesInstanceOutput, error)
    DeregisterOnPremisesInstanceAsync(ctx workflow.Context, input *codedeploy.DeregisterOnPremisesInstanceInput) *CodedeployDeregisterOnPremisesInstanceResult

    GetApplication(ctx workflow.Context, input *codedeploy.GetApplicationInput) (*codedeploy.GetApplicationOutput, error)
    GetApplicationAsync(ctx workflow.Context, input *codedeploy.GetApplicationInput) *CodedeployGetApplicationResult

    GetApplicationRevision(ctx workflow.Context, input *codedeploy.GetApplicationRevisionInput) (*codedeploy.GetApplicationRevisionOutput, error)
    GetApplicationRevisionAsync(ctx workflow.Context, input *codedeploy.GetApplicationRevisionInput) *CodedeployGetApplicationRevisionResult

    GetDeployment(ctx workflow.Context, input *codedeploy.GetDeploymentInput) (*codedeploy.GetDeploymentOutput, error)
    GetDeploymentAsync(ctx workflow.Context, input *codedeploy.GetDeploymentInput) *CodedeployGetDeploymentResult

    GetDeploymentConfig(ctx workflow.Context, input *codedeploy.GetDeploymentConfigInput) (*codedeploy.GetDeploymentConfigOutput, error)
    GetDeploymentConfigAsync(ctx workflow.Context, input *codedeploy.GetDeploymentConfigInput) *CodedeployGetDeploymentConfigResult

    GetDeploymentGroup(ctx workflow.Context, input *codedeploy.GetDeploymentGroupInput) (*codedeploy.GetDeploymentGroupOutput, error)
    GetDeploymentGroupAsync(ctx workflow.Context, input *codedeploy.GetDeploymentGroupInput) *CodedeployGetDeploymentGroupResult

    GetDeploymentInstance(ctx workflow.Context, input *codedeploy.GetDeploymentInstanceInput) (*codedeploy.GetDeploymentInstanceOutput, error)
    GetDeploymentInstanceAsync(ctx workflow.Context, input *codedeploy.GetDeploymentInstanceInput) *CodedeployGetDeploymentInstanceResult

    GetDeploymentTarget(ctx workflow.Context, input *codedeploy.GetDeploymentTargetInput) (*codedeploy.GetDeploymentTargetOutput, error)
    GetDeploymentTargetAsync(ctx workflow.Context, input *codedeploy.GetDeploymentTargetInput) *CodedeployGetDeploymentTargetResult

    GetOnPremisesInstance(ctx workflow.Context, input *codedeploy.GetOnPremisesInstanceInput) (*codedeploy.GetOnPremisesInstanceOutput, error)
    GetOnPremisesInstanceAsync(ctx workflow.Context, input *codedeploy.GetOnPremisesInstanceInput) *CodedeployGetOnPremisesInstanceResult

    ListApplicationRevisions(ctx workflow.Context, input *codedeploy.ListApplicationRevisionsInput) (*codedeploy.ListApplicationRevisionsOutput, error)
    ListApplicationRevisionsAsync(ctx workflow.Context, input *codedeploy.ListApplicationRevisionsInput) *CodedeployListApplicationRevisionsResult

    ListApplications(ctx workflow.Context, input *codedeploy.ListApplicationsInput) (*codedeploy.ListApplicationsOutput, error)
    ListApplicationsAsync(ctx workflow.Context, input *codedeploy.ListApplicationsInput) *CodedeployListApplicationsResult

    ListDeploymentConfigs(ctx workflow.Context, input *codedeploy.ListDeploymentConfigsInput) (*codedeploy.ListDeploymentConfigsOutput, error)
    ListDeploymentConfigsAsync(ctx workflow.Context, input *codedeploy.ListDeploymentConfigsInput) *CodedeployListDeploymentConfigsResult

    ListDeploymentGroups(ctx workflow.Context, input *codedeploy.ListDeploymentGroupsInput) (*codedeploy.ListDeploymentGroupsOutput, error)
    ListDeploymentGroupsAsync(ctx workflow.Context, input *codedeploy.ListDeploymentGroupsInput) *CodedeployListDeploymentGroupsResult

    ListDeploymentInstances(ctx workflow.Context, input *codedeploy.ListDeploymentInstancesInput) (*codedeploy.ListDeploymentInstancesOutput, error)
    ListDeploymentInstancesAsync(ctx workflow.Context, input *codedeploy.ListDeploymentInstancesInput) *CodedeployListDeploymentInstancesResult

    ListDeploymentTargets(ctx workflow.Context, input *codedeploy.ListDeploymentTargetsInput) (*codedeploy.ListDeploymentTargetsOutput, error)
    ListDeploymentTargetsAsync(ctx workflow.Context, input *codedeploy.ListDeploymentTargetsInput) *CodedeployListDeploymentTargetsResult

    ListDeployments(ctx workflow.Context, input *codedeploy.ListDeploymentsInput) (*codedeploy.ListDeploymentsOutput, error)
    ListDeploymentsAsync(ctx workflow.Context, input *codedeploy.ListDeploymentsInput) *CodedeployListDeploymentsResult

    ListGitHubAccountTokenNames(ctx workflow.Context, input *codedeploy.ListGitHubAccountTokenNamesInput) (*codedeploy.ListGitHubAccountTokenNamesOutput, error)
    ListGitHubAccountTokenNamesAsync(ctx workflow.Context, input *codedeploy.ListGitHubAccountTokenNamesInput) *CodedeployListGitHubAccountTokenNamesResult

    ListOnPremisesInstances(ctx workflow.Context, input *codedeploy.ListOnPremisesInstancesInput) (*codedeploy.ListOnPremisesInstancesOutput, error)
    ListOnPremisesInstancesAsync(ctx workflow.Context, input *codedeploy.ListOnPremisesInstancesInput) *CodedeployListOnPremisesInstancesResult

    ListTagsForResource(ctx workflow.Context, input *codedeploy.ListTagsForResourceInput) (*codedeploy.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *codedeploy.ListTagsForResourceInput) *CodedeployListTagsForResourceResult

    PutLifecycleEventHookExecutionStatus(ctx workflow.Context, input *codedeploy.PutLifecycleEventHookExecutionStatusInput) (*codedeploy.PutLifecycleEventHookExecutionStatusOutput, error)
    PutLifecycleEventHookExecutionStatusAsync(ctx workflow.Context, input *codedeploy.PutLifecycleEventHookExecutionStatusInput) *CodedeployPutLifecycleEventHookExecutionStatusResult

    RegisterApplicationRevision(ctx workflow.Context, input *codedeploy.RegisterApplicationRevisionInput) (*codedeploy.RegisterApplicationRevisionOutput, error)
    RegisterApplicationRevisionAsync(ctx workflow.Context, input *codedeploy.RegisterApplicationRevisionInput) *CodedeployRegisterApplicationRevisionResult

    RegisterOnPremisesInstance(ctx workflow.Context, input *codedeploy.RegisterOnPremisesInstanceInput) (*codedeploy.RegisterOnPremisesInstanceOutput, error)
    RegisterOnPremisesInstanceAsync(ctx workflow.Context, input *codedeploy.RegisterOnPremisesInstanceInput) *CodedeployRegisterOnPremisesInstanceResult

    RemoveTagsFromOnPremisesInstances(ctx workflow.Context, input *codedeploy.RemoveTagsFromOnPremisesInstancesInput) (*codedeploy.RemoveTagsFromOnPremisesInstancesOutput, error)
    RemoveTagsFromOnPremisesInstancesAsync(ctx workflow.Context, input *codedeploy.RemoveTagsFromOnPremisesInstancesInput) *CodedeployRemoveTagsFromOnPremisesInstancesResult

    SkipWaitTimeForInstanceTermination(ctx workflow.Context, input *codedeploy.SkipWaitTimeForInstanceTerminationInput) (*codedeploy.SkipWaitTimeForInstanceTerminationOutput, error)
    SkipWaitTimeForInstanceTerminationAsync(ctx workflow.Context, input *codedeploy.SkipWaitTimeForInstanceTerminationInput) *CodedeploySkipWaitTimeForInstanceTerminationResult

    StopDeployment(ctx workflow.Context, input *codedeploy.StopDeploymentInput) (*codedeploy.StopDeploymentOutput, error)
    StopDeploymentAsync(ctx workflow.Context, input *codedeploy.StopDeploymentInput) *CodedeployStopDeploymentResult

    TagResource(ctx workflow.Context, input *codedeploy.TagResourceInput) (*codedeploy.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *codedeploy.TagResourceInput) *CodedeployTagResourceResult

    UntagResource(ctx workflow.Context, input *codedeploy.UntagResourceInput) (*codedeploy.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *codedeploy.UntagResourceInput) *CodedeployUntagResourceResult

    UpdateApplication(ctx workflow.Context, input *codedeploy.UpdateApplicationInput) (*codedeploy.UpdateApplicationOutput, error)
    UpdateApplicationAsync(ctx workflow.Context, input *codedeploy.UpdateApplicationInput) *CodedeployUpdateApplicationResult

    UpdateDeploymentGroup(ctx workflow.Context, input *codedeploy.UpdateDeploymentGroupInput) (*codedeploy.UpdateDeploymentGroupOutput, error)
    UpdateDeploymentGroupAsync(ctx workflow.Context, input *codedeploy.UpdateDeploymentGroupInput) *CodedeployUpdateDeploymentGroupResult

    WaitUntilDeploymentSuccessful(ctx workflow.Context, input *codedeploy.GetDeploymentInput) error}
type CodedeployAddTagsToOnPremisesInstancesResult struct {
	Result workflow.Future
}

func (r *CodedeployAddTagsToOnPremisesInstancesResult) Get(ctx workflow.Context) (*codedeploy.AddTagsToOnPremisesInstancesOutput, error) {
    var output codedeploy.AddTagsToOnPremisesInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployBatchGetApplicationRevisionsResult struct {
	Result workflow.Future
}

func (r *CodedeployBatchGetApplicationRevisionsResult) Get(ctx workflow.Context) (*codedeploy.BatchGetApplicationRevisionsOutput, error) {
    var output codedeploy.BatchGetApplicationRevisionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployBatchGetApplicationsResult struct {
	Result workflow.Future
}

func (r *CodedeployBatchGetApplicationsResult) Get(ctx workflow.Context) (*codedeploy.BatchGetApplicationsOutput, error) {
    var output codedeploy.BatchGetApplicationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployBatchGetDeploymentGroupsResult struct {
	Result workflow.Future
}

func (r *CodedeployBatchGetDeploymentGroupsResult) Get(ctx workflow.Context) (*codedeploy.BatchGetDeploymentGroupsOutput, error) {
    var output codedeploy.BatchGetDeploymentGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployBatchGetDeploymentInstancesResult struct {
	Result workflow.Future
}

func (r *CodedeployBatchGetDeploymentInstancesResult) Get(ctx workflow.Context) (*codedeploy.BatchGetDeploymentInstancesOutput, error) {
    var output codedeploy.BatchGetDeploymentInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployBatchGetDeploymentTargetsResult struct {
	Result workflow.Future
}

func (r *CodedeployBatchGetDeploymentTargetsResult) Get(ctx workflow.Context) (*codedeploy.BatchGetDeploymentTargetsOutput, error) {
    var output codedeploy.BatchGetDeploymentTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployBatchGetDeploymentsResult struct {
	Result workflow.Future
}

func (r *CodedeployBatchGetDeploymentsResult) Get(ctx workflow.Context) (*codedeploy.BatchGetDeploymentsOutput, error) {
    var output codedeploy.BatchGetDeploymentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployBatchGetOnPremisesInstancesResult struct {
	Result workflow.Future
}

func (r *CodedeployBatchGetOnPremisesInstancesResult) Get(ctx workflow.Context) (*codedeploy.BatchGetOnPremisesInstancesOutput, error) {
    var output codedeploy.BatchGetOnPremisesInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployContinueDeploymentResult struct {
	Result workflow.Future
}

func (r *CodedeployContinueDeploymentResult) Get(ctx workflow.Context) (*codedeploy.ContinueDeploymentOutput, error) {
    var output codedeploy.ContinueDeploymentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployCreateApplicationResult struct {
	Result workflow.Future
}

func (r *CodedeployCreateApplicationResult) Get(ctx workflow.Context) (*codedeploy.CreateApplicationOutput, error) {
    var output codedeploy.CreateApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployCreateDeploymentResult struct {
	Result workflow.Future
}

func (r *CodedeployCreateDeploymentResult) Get(ctx workflow.Context) (*codedeploy.CreateDeploymentOutput, error) {
    var output codedeploy.CreateDeploymentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployCreateDeploymentConfigResult struct {
	Result workflow.Future
}

func (r *CodedeployCreateDeploymentConfigResult) Get(ctx workflow.Context) (*codedeploy.CreateDeploymentConfigOutput, error) {
    var output codedeploy.CreateDeploymentConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployCreateDeploymentGroupResult struct {
	Result workflow.Future
}

func (r *CodedeployCreateDeploymentGroupResult) Get(ctx workflow.Context) (*codedeploy.CreateDeploymentGroupOutput, error) {
    var output codedeploy.CreateDeploymentGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployDeleteApplicationResult struct {
	Result workflow.Future
}

func (r *CodedeployDeleteApplicationResult) Get(ctx workflow.Context) (*codedeploy.DeleteApplicationOutput, error) {
    var output codedeploy.DeleteApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployDeleteDeploymentConfigResult struct {
	Result workflow.Future
}

func (r *CodedeployDeleteDeploymentConfigResult) Get(ctx workflow.Context) (*codedeploy.DeleteDeploymentConfigOutput, error) {
    var output codedeploy.DeleteDeploymentConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployDeleteDeploymentGroupResult struct {
	Result workflow.Future
}

func (r *CodedeployDeleteDeploymentGroupResult) Get(ctx workflow.Context) (*codedeploy.DeleteDeploymentGroupOutput, error) {
    var output codedeploy.DeleteDeploymentGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployDeleteGitHubAccountTokenResult struct {
	Result workflow.Future
}

func (r *CodedeployDeleteGitHubAccountTokenResult) Get(ctx workflow.Context) (*codedeploy.DeleteGitHubAccountTokenOutput, error) {
    var output codedeploy.DeleteGitHubAccountTokenOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployDeleteResourcesByExternalIdResult struct {
	Result workflow.Future
}

func (r *CodedeployDeleteResourcesByExternalIdResult) Get(ctx workflow.Context) (*codedeploy.DeleteResourcesByExternalIdOutput, error) {
    var output codedeploy.DeleteResourcesByExternalIdOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployDeregisterOnPremisesInstanceResult struct {
	Result workflow.Future
}

func (r *CodedeployDeregisterOnPremisesInstanceResult) Get(ctx workflow.Context) (*codedeploy.DeregisterOnPremisesInstanceOutput, error) {
    var output codedeploy.DeregisterOnPremisesInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployGetApplicationResult struct {
	Result workflow.Future
}

func (r *CodedeployGetApplicationResult) Get(ctx workflow.Context) (*codedeploy.GetApplicationOutput, error) {
    var output codedeploy.GetApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployGetApplicationRevisionResult struct {
	Result workflow.Future
}

func (r *CodedeployGetApplicationRevisionResult) Get(ctx workflow.Context) (*codedeploy.GetApplicationRevisionOutput, error) {
    var output codedeploy.GetApplicationRevisionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployGetDeploymentResult struct {
	Result workflow.Future
}

func (r *CodedeployGetDeploymentResult) Get(ctx workflow.Context) (*codedeploy.GetDeploymentOutput, error) {
    var output codedeploy.GetDeploymentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployGetDeploymentConfigResult struct {
	Result workflow.Future
}

func (r *CodedeployGetDeploymentConfigResult) Get(ctx workflow.Context) (*codedeploy.GetDeploymentConfigOutput, error) {
    var output codedeploy.GetDeploymentConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployGetDeploymentGroupResult struct {
	Result workflow.Future
}

func (r *CodedeployGetDeploymentGroupResult) Get(ctx workflow.Context) (*codedeploy.GetDeploymentGroupOutput, error) {
    var output codedeploy.GetDeploymentGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployGetDeploymentInstanceResult struct {
	Result workflow.Future
}

func (r *CodedeployGetDeploymentInstanceResult) Get(ctx workflow.Context) (*codedeploy.GetDeploymentInstanceOutput, error) {
    var output codedeploy.GetDeploymentInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployGetDeploymentTargetResult struct {
	Result workflow.Future
}

func (r *CodedeployGetDeploymentTargetResult) Get(ctx workflow.Context) (*codedeploy.GetDeploymentTargetOutput, error) {
    var output codedeploy.GetDeploymentTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployGetOnPremisesInstanceResult struct {
	Result workflow.Future
}

func (r *CodedeployGetOnPremisesInstanceResult) Get(ctx workflow.Context) (*codedeploy.GetOnPremisesInstanceOutput, error) {
    var output codedeploy.GetOnPremisesInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployListApplicationRevisionsResult struct {
	Result workflow.Future
}

func (r *CodedeployListApplicationRevisionsResult) Get(ctx workflow.Context) (*codedeploy.ListApplicationRevisionsOutput, error) {
    var output codedeploy.ListApplicationRevisionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployListApplicationsResult struct {
	Result workflow.Future
}

func (r *CodedeployListApplicationsResult) Get(ctx workflow.Context) (*codedeploy.ListApplicationsOutput, error) {
    var output codedeploy.ListApplicationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployListDeploymentConfigsResult struct {
	Result workflow.Future
}

func (r *CodedeployListDeploymentConfigsResult) Get(ctx workflow.Context) (*codedeploy.ListDeploymentConfigsOutput, error) {
    var output codedeploy.ListDeploymentConfigsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployListDeploymentGroupsResult struct {
	Result workflow.Future
}

func (r *CodedeployListDeploymentGroupsResult) Get(ctx workflow.Context) (*codedeploy.ListDeploymentGroupsOutput, error) {
    var output codedeploy.ListDeploymentGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployListDeploymentInstancesResult struct {
	Result workflow.Future
}

func (r *CodedeployListDeploymentInstancesResult) Get(ctx workflow.Context) (*codedeploy.ListDeploymentInstancesOutput, error) {
    var output codedeploy.ListDeploymentInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployListDeploymentTargetsResult struct {
	Result workflow.Future
}

func (r *CodedeployListDeploymentTargetsResult) Get(ctx workflow.Context) (*codedeploy.ListDeploymentTargetsOutput, error) {
    var output codedeploy.ListDeploymentTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployListDeploymentsResult struct {
	Result workflow.Future
}

func (r *CodedeployListDeploymentsResult) Get(ctx workflow.Context) (*codedeploy.ListDeploymentsOutput, error) {
    var output codedeploy.ListDeploymentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployListGitHubAccountTokenNamesResult struct {
	Result workflow.Future
}

func (r *CodedeployListGitHubAccountTokenNamesResult) Get(ctx workflow.Context) (*codedeploy.ListGitHubAccountTokenNamesOutput, error) {
    var output codedeploy.ListGitHubAccountTokenNamesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployListOnPremisesInstancesResult struct {
	Result workflow.Future
}

func (r *CodedeployListOnPremisesInstancesResult) Get(ctx workflow.Context) (*codedeploy.ListOnPremisesInstancesOutput, error) {
    var output codedeploy.ListOnPremisesInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *CodedeployListTagsForResourceResult) Get(ctx workflow.Context) (*codedeploy.ListTagsForResourceOutput, error) {
    var output codedeploy.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployPutLifecycleEventHookExecutionStatusResult struct {
	Result workflow.Future
}

func (r *CodedeployPutLifecycleEventHookExecutionStatusResult) Get(ctx workflow.Context) (*codedeploy.PutLifecycleEventHookExecutionStatusOutput, error) {
    var output codedeploy.PutLifecycleEventHookExecutionStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployRegisterApplicationRevisionResult struct {
	Result workflow.Future
}

func (r *CodedeployRegisterApplicationRevisionResult) Get(ctx workflow.Context) (*codedeploy.RegisterApplicationRevisionOutput, error) {
    var output codedeploy.RegisterApplicationRevisionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployRegisterOnPremisesInstanceResult struct {
	Result workflow.Future
}

func (r *CodedeployRegisterOnPremisesInstanceResult) Get(ctx workflow.Context) (*codedeploy.RegisterOnPremisesInstanceOutput, error) {
    var output codedeploy.RegisterOnPremisesInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployRemoveTagsFromOnPremisesInstancesResult struct {
	Result workflow.Future
}

func (r *CodedeployRemoveTagsFromOnPremisesInstancesResult) Get(ctx workflow.Context) (*codedeploy.RemoveTagsFromOnPremisesInstancesOutput, error) {
    var output codedeploy.RemoveTagsFromOnPremisesInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeploySkipWaitTimeForInstanceTerminationResult struct {
	Result workflow.Future
}

func (r *CodedeploySkipWaitTimeForInstanceTerminationResult) Get(ctx workflow.Context) (*codedeploy.SkipWaitTimeForInstanceTerminationOutput, error) {
    var output codedeploy.SkipWaitTimeForInstanceTerminationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployStopDeploymentResult struct {
	Result workflow.Future
}

func (r *CodedeployStopDeploymentResult) Get(ctx workflow.Context) (*codedeploy.StopDeploymentOutput, error) {
    var output codedeploy.StopDeploymentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployTagResourceResult struct {
	Result workflow.Future
}

func (r *CodedeployTagResourceResult) Get(ctx workflow.Context) (*codedeploy.TagResourceOutput, error) {
    var output codedeploy.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployUntagResourceResult struct {
	Result workflow.Future
}

func (r *CodedeployUntagResourceResult) Get(ctx workflow.Context) (*codedeploy.UntagResourceOutput, error) {
    var output codedeploy.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployUpdateApplicationResult struct {
	Result workflow.Future
}

func (r *CodedeployUpdateApplicationResult) Get(ctx workflow.Context) (*codedeploy.UpdateApplicationOutput, error) {
    var output codedeploy.UpdateApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodedeployUpdateDeploymentGroupResult struct {
	Result workflow.Future
}

func (r *CodedeployUpdateDeploymentGroupResult) Get(ctx workflow.Context) (*codedeploy.UpdateDeploymentGroupOutput, error) {
    var output codedeploy.UpdateDeploymentGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type CodeDeployStub struct {
    activities awsactivities.CodeDeployActivities
}

func NewCodeDeployStub() CodeDeployClient {
    return &CodeDeployStub{}
}
func (a *CodeDeployStub) AddTagsToOnPremisesInstances(ctx workflow.Context, input *codedeploy.AddTagsToOnPremisesInstancesInput) (*codedeploy.AddTagsToOnPremisesInstancesOutput, error) {
    var output codedeploy.AddTagsToOnPremisesInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddTagsToOnPremisesInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) AddTagsToOnPremisesInstancesAsync(ctx workflow.Context, input *codedeploy.AddTagsToOnPremisesInstancesInput) *CodedeployAddTagsToOnPremisesInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddTagsToOnPremisesInstances, input)
    return &CodedeployAddTagsToOnPremisesInstancesResult{Result: future}
}
func (a *CodeDeployStub) BatchGetApplicationRevisions(ctx workflow.Context, input *codedeploy.BatchGetApplicationRevisionsInput) (*codedeploy.BatchGetApplicationRevisionsOutput, error) {
    var output codedeploy.BatchGetApplicationRevisionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetApplicationRevisions, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) BatchGetApplicationRevisionsAsync(ctx workflow.Context, input *codedeploy.BatchGetApplicationRevisionsInput) *CodedeployBatchGetApplicationRevisionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetApplicationRevisions, input)
    return &CodedeployBatchGetApplicationRevisionsResult{Result: future}
}
func (a *CodeDeployStub) BatchGetApplications(ctx workflow.Context, input *codedeploy.BatchGetApplicationsInput) (*codedeploy.BatchGetApplicationsOutput, error) {
    var output codedeploy.BatchGetApplicationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetApplications, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) BatchGetApplicationsAsync(ctx workflow.Context, input *codedeploy.BatchGetApplicationsInput) *CodedeployBatchGetApplicationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetApplications, input)
    return &CodedeployBatchGetApplicationsResult{Result: future}
}
func (a *CodeDeployStub) BatchGetDeploymentGroups(ctx workflow.Context, input *codedeploy.BatchGetDeploymentGroupsInput) (*codedeploy.BatchGetDeploymentGroupsOutput, error) {
    var output codedeploy.BatchGetDeploymentGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetDeploymentGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) BatchGetDeploymentGroupsAsync(ctx workflow.Context, input *codedeploy.BatchGetDeploymentGroupsInput) *CodedeployBatchGetDeploymentGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetDeploymentGroups, input)
    return &CodedeployBatchGetDeploymentGroupsResult{Result: future}
}
func (a *CodeDeployStub) BatchGetDeploymentInstances(ctx workflow.Context, input *codedeploy.BatchGetDeploymentInstancesInput) (*codedeploy.BatchGetDeploymentInstancesOutput, error) {
    var output codedeploy.BatchGetDeploymentInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetDeploymentInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) BatchGetDeploymentInstancesAsync(ctx workflow.Context, input *codedeploy.BatchGetDeploymentInstancesInput) *CodedeployBatchGetDeploymentInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetDeploymentInstances, input)
    return &CodedeployBatchGetDeploymentInstancesResult{Result: future}
}
func (a *CodeDeployStub) BatchGetDeploymentTargets(ctx workflow.Context, input *codedeploy.BatchGetDeploymentTargetsInput) (*codedeploy.BatchGetDeploymentTargetsOutput, error) {
    var output codedeploy.BatchGetDeploymentTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetDeploymentTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) BatchGetDeploymentTargetsAsync(ctx workflow.Context, input *codedeploy.BatchGetDeploymentTargetsInput) *CodedeployBatchGetDeploymentTargetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetDeploymentTargets, input)
    return &CodedeployBatchGetDeploymentTargetsResult{Result: future}
}
func (a *CodeDeployStub) BatchGetDeployments(ctx workflow.Context, input *codedeploy.BatchGetDeploymentsInput) (*codedeploy.BatchGetDeploymentsOutput, error) {
    var output codedeploy.BatchGetDeploymentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetDeployments, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) BatchGetDeploymentsAsync(ctx workflow.Context, input *codedeploy.BatchGetDeploymentsInput) *CodedeployBatchGetDeploymentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetDeployments, input)
    return &CodedeployBatchGetDeploymentsResult{Result: future}
}
func (a *CodeDeployStub) BatchGetOnPremisesInstances(ctx workflow.Context, input *codedeploy.BatchGetOnPremisesInstancesInput) (*codedeploy.BatchGetOnPremisesInstancesOutput, error) {
    var output codedeploy.BatchGetOnPremisesInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetOnPremisesInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) BatchGetOnPremisesInstancesAsync(ctx workflow.Context, input *codedeploy.BatchGetOnPremisesInstancesInput) *CodedeployBatchGetOnPremisesInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetOnPremisesInstances, input)
    return &CodedeployBatchGetOnPremisesInstancesResult{Result: future}
}
func (a *CodeDeployStub) ContinueDeployment(ctx workflow.Context, input *codedeploy.ContinueDeploymentInput) (*codedeploy.ContinueDeploymentOutput, error) {
    var output codedeploy.ContinueDeploymentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ContinueDeployment, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) ContinueDeploymentAsync(ctx workflow.Context, input *codedeploy.ContinueDeploymentInput) *CodedeployContinueDeploymentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ContinueDeployment, input)
    return &CodedeployContinueDeploymentResult{Result: future}
}
func (a *CodeDeployStub) CreateApplication(ctx workflow.Context, input *codedeploy.CreateApplicationInput) (*codedeploy.CreateApplicationOutput, error) {
    var output codedeploy.CreateApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) CreateApplicationAsync(ctx workflow.Context, input *codedeploy.CreateApplicationInput) *CodedeployCreateApplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateApplication, input)
    return &CodedeployCreateApplicationResult{Result: future}
}
func (a *CodeDeployStub) CreateDeployment(ctx workflow.Context, input *codedeploy.CreateDeploymentInput) (*codedeploy.CreateDeploymentOutput, error) {
    var output codedeploy.CreateDeploymentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDeployment, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) CreateDeploymentAsync(ctx workflow.Context, input *codedeploy.CreateDeploymentInput) *CodedeployCreateDeploymentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDeployment, input)
    return &CodedeployCreateDeploymentResult{Result: future}
}
func (a *CodeDeployStub) CreateDeploymentConfig(ctx workflow.Context, input *codedeploy.CreateDeploymentConfigInput) (*codedeploy.CreateDeploymentConfigOutput, error) {
    var output codedeploy.CreateDeploymentConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDeploymentConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) CreateDeploymentConfigAsync(ctx workflow.Context, input *codedeploy.CreateDeploymentConfigInput) *CodedeployCreateDeploymentConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDeploymentConfig, input)
    return &CodedeployCreateDeploymentConfigResult{Result: future}
}
func (a *CodeDeployStub) CreateDeploymentGroup(ctx workflow.Context, input *codedeploy.CreateDeploymentGroupInput) (*codedeploy.CreateDeploymentGroupOutput, error) {
    var output codedeploy.CreateDeploymentGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDeploymentGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) CreateDeploymentGroupAsync(ctx workflow.Context, input *codedeploy.CreateDeploymentGroupInput) *CodedeployCreateDeploymentGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDeploymentGroup, input)
    return &CodedeployCreateDeploymentGroupResult{Result: future}
}
func (a *CodeDeployStub) DeleteApplication(ctx workflow.Context, input *codedeploy.DeleteApplicationInput) (*codedeploy.DeleteApplicationOutput, error) {
    var output codedeploy.DeleteApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) DeleteApplicationAsync(ctx workflow.Context, input *codedeploy.DeleteApplicationInput) *CodedeployDeleteApplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApplication, input)
    return &CodedeployDeleteApplicationResult{Result: future}
}
func (a *CodeDeployStub) DeleteDeploymentConfig(ctx workflow.Context, input *codedeploy.DeleteDeploymentConfigInput) (*codedeploy.DeleteDeploymentConfigOutput, error) {
    var output codedeploy.DeleteDeploymentConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDeploymentConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) DeleteDeploymentConfigAsync(ctx workflow.Context, input *codedeploy.DeleteDeploymentConfigInput) *CodedeployDeleteDeploymentConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDeploymentConfig, input)
    return &CodedeployDeleteDeploymentConfigResult{Result: future}
}
func (a *CodeDeployStub) DeleteDeploymentGroup(ctx workflow.Context, input *codedeploy.DeleteDeploymentGroupInput) (*codedeploy.DeleteDeploymentGroupOutput, error) {
    var output codedeploy.DeleteDeploymentGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDeploymentGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) DeleteDeploymentGroupAsync(ctx workflow.Context, input *codedeploy.DeleteDeploymentGroupInput) *CodedeployDeleteDeploymentGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDeploymentGroup, input)
    return &CodedeployDeleteDeploymentGroupResult{Result: future}
}
func (a *CodeDeployStub) DeleteGitHubAccountToken(ctx workflow.Context, input *codedeploy.DeleteGitHubAccountTokenInput) (*codedeploy.DeleteGitHubAccountTokenOutput, error) {
    var output codedeploy.DeleteGitHubAccountTokenOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGitHubAccountToken, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) DeleteGitHubAccountTokenAsync(ctx workflow.Context, input *codedeploy.DeleteGitHubAccountTokenInput) *CodedeployDeleteGitHubAccountTokenResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteGitHubAccountToken, input)
    return &CodedeployDeleteGitHubAccountTokenResult{Result: future}
}
func (a *CodeDeployStub) DeleteResourcesByExternalId(ctx workflow.Context, input *codedeploy.DeleteResourcesByExternalIdInput) (*codedeploy.DeleteResourcesByExternalIdOutput, error) {
    var output codedeploy.DeleteResourcesByExternalIdOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteResourcesByExternalId, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) DeleteResourcesByExternalIdAsync(ctx workflow.Context, input *codedeploy.DeleteResourcesByExternalIdInput) *CodedeployDeleteResourcesByExternalIdResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteResourcesByExternalId, input)
    return &CodedeployDeleteResourcesByExternalIdResult{Result: future}
}
func (a *CodeDeployStub) DeregisterOnPremisesInstance(ctx workflow.Context, input *codedeploy.DeregisterOnPremisesInstanceInput) (*codedeploy.DeregisterOnPremisesInstanceOutput, error) {
    var output codedeploy.DeregisterOnPremisesInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterOnPremisesInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) DeregisterOnPremisesInstanceAsync(ctx workflow.Context, input *codedeploy.DeregisterOnPremisesInstanceInput) *CodedeployDeregisterOnPremisesInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterOnPremisesInstance, input)
    return &CodedeployDeregisterOnPremisesInstanceResult{Result: future}
}
func (a *CodeDeployStub) GetApplication(ctx workflow.Context, input *codedeploy.GetApplicationInput) (*codedeploy.GetApplicationOutput, error) {
    var output codedeploy.GetApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) GetApplicationAsync(ctx workflow.Context, input *codedeploy.GetApplicationInput) *CodedeployGetApplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetApplication, input)
    return &CodedeployGetApplicationResult{Result: future}
}
func (a *CodeDeployStub) GetApplicationRevision(ctx workflow.Context, input *codedeploy.GetApplicationRevisionInput) (*codedeploy.GetApplicationRevisionOutput, error) {
    var output codedeploy.GetApplicationRevisionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetApplicationRevision, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) GetApplicationRevisionAsync(ctx workflow.Context, input *codedeploy.GetApplicationRevisionInput) *CodedeployGetApplicationRevisionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetApplicationRevision, input)
    return &CodedeployGetApplicationRevisionResult{Result: future}
}
func (a *CodeDeployStub) GetDeployment(ctx workflow.Context, input *codedeploy.GetDeploymentInput) (*codedeploy.GetDeploymentOutput, error) {
    var output codedeploy.GetDeploymentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeployment, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) GetDeploymentAsync(ctx workflow.Context, input *codedeploy.GetDeploymentInput) *CodedeployGetDeploymentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDeployment, input)
    return &CodedeployGetDeploymentResult{Result: future}
}
func (a *CodeDeployStub) GetDeploymentConfig(ctx workflow.Context, input *codedeploy.GetDeploymentConfigInput) (*codedeploy.GetDeploymentConfigOutput, error) {
    var output codedeploy.GetDeploymentConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeploymentConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) GetDeploymentConfigAsync(ctx workflow.Context, input *codedeploy.GetDeploymentConfigInput) *CodedeployGetDeploymentConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDeploymentConfig, input)
    return &CodedeployGetDeploymentConfigResult{Result: future}
}
func (a *CodeDeployStub) GetDeploymentGroup(ctx workflow.Context, input *codedeploy.GetDeploymentGroupInput) (*codedeploy.GetDeploymentGroupOutput, error) {
    var output codedeploy.GetDeploymentGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeploymentGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) GetDeploymentGroupAsync(ctx workflow.Context, input *codedeploy.GetDeploymentGroupInput) *CodedeployGetDeploymentGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDeploymentGroup, input)
    return &CodedeployGetDeploymentGroupResult{Result: future}
}
func (a *CodeDeployStub) GetDeploymentInstance(ctx workflow.Context, input *codedeploy.GetDeploymentInstanceInput) (*codedeploy.GetDeploymentInstanceOutput, error) {
    var output codedeploy.GetDeploymentInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeploymentInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) GetDeploymentInstanceAsync(ctx workflow.Context, input *codedeploy.GetDeploymentInstanceInput) *CodedeployGetDeploymentInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDeploymentInstance, input)
    return &CodedeployGetDeploymentInstanceResult{Result: future}
}
func (a *CodeDeployStub) GetDeploymentTarget(ctx workflow.Context, input *codedeploy.GetDeploymentTargetInput) (*codedeploy.GetDeploymentTargetOutput, error) {
    var output codedeploy.GetDeploymentTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeploymentTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) GetDeploymentTargetAsync(ctx workflow.Context, input *codedeploy.GetDeploymentTargetInput) *CodedeployGetDeploymentTargetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDeploymentTarget, input)
    return &CodedeployGetDeploymentTargetResult{Result: future}
}
func (a *CodeDeployStub) GetOnPremisesInstance(ctx workflow.Context, input *codedeploy.GetOnPremisesInstanceInput) (*codedeploy.GetOnPremisesInstanceOutput, error) {
    var output codedeploy.GetOnPremisesInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetOnPremisesInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) GetOnPremisesInstanceAsync(ctx workflow.Context, input *codedeploy.GetOnPremisesInstanceInput) *CodedeployGetOnPremisesInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetOnPremisesInstance, input)
    return &CodedeployGetOnPremisesInstanceResult{Result: future}
}
func (a *CodeDeployStub) ListApplicationRevisions(ctx workflow.Context, input *codedeploy.ListApplicationRevisionsInput) (*codedeploy.ListApplicationRevisionsOutput, error) {
    var output codedeploy.ListApplicationRevisionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListApplicationRevisions, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) ListApplicationRevisionsAsync(ctx workflow.Context, input *codedeploy.ListApplicationRevisionsInput) *CodedeployListApplicationRevisionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListApplicationRevisions, input)
    return &CodedeployListApplicationRevisionsResult{Result: future}
}
func (a *CodeDeployStub) ListApplications(ctx workflow.Context, input *codedeploy.ListApplicationsInput) (*codedeploy.ListApplicationsOutput, error) {
    var output codedeploy.ListApplicationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListApplications, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) ListApplicationsAsync(ctx workflow.Context, input *codedeploy.ListApplicationsInput) *CodedeployListApplicationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListApplications, input)
    return &CodedeployListApplicationsResult{Result: future}
}
func (a *CodeDeployStub) ListDeploymentConfigs(ctx workflow.Context, input *codedeploy.ListDeploymentConfigsInput) (*codedeploy.ListDeploymentConfigsOutput, error) {
    var output codedeploy.ListDeploymentConfigsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeploymentConfigs, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) ListDeploymentConfigsAsync(ctx workflow.Context, input *codedeploy.ListDeploymentConfigsInput) *CodedeployListDeploymentConfigsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDeploymentConfigs, input)
    return &CodedeployListDeploymentConfigsResult{Result: future}
}
func (a *CodeDeployStub) ListDeploymentGroups(ctx workflow.Context, input *codedeploy.ListDeploymentGroupsInput) (*codedeploy.ListDeploymentGroupsOutput, error) {
    var output codedeploy.ListDeploymentGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeploymentGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) ListDeploymentGroupsAsync(ctx workflow.Context, input *codedeploy.ListDeploymentGroupsInput) *CodedeployListDeploymentGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDeploymentGroups, input)
    return &CodedeployListDeploymentGroupsResult{Result: future}
}
func (a *CodeDeployStub) ListDeploymentInstances(ctx workflow.Context, input *codedeploy.ListDeploymentInstancesInput) (*codedeploy.ListDeploymentInstancesOutput, error) {
    var output codedeploy.ListDeploymentInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeploymentInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) ListDeploymentInstancesAsync(ctx workflow.Context, input *codedeploy.ListDeploymentInstancesInput) *CodedeployListDeploymentInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDeploymentInstances, input)
    return &CodedeployListDeploymentInstancesResult{Result: future}
}
func (a *CodeDeployStub) ListDeploymentTargets(ctx workflow.Context, input *codedeploy.ListDeploymentTargetsInput) (*codedeploy.ListDeploymentTargetsOutput, error) {
    var output codedeploy.ListDeploymentTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeploymentTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) ListDeploymentTargetsAsync(ctx workflow.Context, input *codedeploy.ListDeploymentTargetsInput) *CodedeployListDeploymentTargetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDeploymentTargets, input)
    return &CodedeployListDeploymentTargetsResult{Result: future}
}
func (a *CodeDeployStub) ListDeployments(ctx workflow.Context, input *codedeploy.ListDeploymentsInput) (*codedeploy.ListDeploymentsOutput, error) {
    var output codedeploy.ListDeploymentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeployments, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) ListDeploymentsAsync(ctx workflow.Context, input *codedeploy.ListDeploymentsInput) *CodedeployListDeploymentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDeployments, input)
    return &CodedeployListDeploymentsResult{Result: future}
}
func (a *CodeDeployStub) ListGitHubAccountTokenNames(ctx workflow.Context, input *codedeploy.ListGitHubAccountTokenNamesInput) (*codedeploy.ListGitHubAccountTokenNamesOutput, error) {
    var output codedeploy.ListGitHubAccountTokenNamesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGitHubAccountTokenNames, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) ListGitHubAccountTokenNamesAsync(ctx workflow.Context, input *codedeploy.ListGitHubAccountTokenNamesInput) *CodedeployListGitHubAccountTokenNamesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListGitHubAccountTokenNames, input)
    return &CodedeployListGitHubAccountTokenNamesResult{Result: future}
}
func (a *CodeDeployStub) ListOnPremisesInstances(ctx workflow.Context, input *codedeploy.ListOnPremisesInstancesInput) (*codedeploy.ListOnPremisesInstancesOutput, error) {
    var output codedeploy.ListOnPremisesInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOnPremisesInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) ListOnPremisesInstancesAsync(ctx workflow.Context, input *codedeploy.ListOnPremisesInstancesInput) *CodedeployListOnPremisesInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListOnPremisesInstances, input)
    return &CodedeployListOnPremisesInstancesResult{Result: future}
}
func (a *CodeDeployStub) ListTagsForResource(ctx workflow.Context, input *codedeploy.ListTagsForResourceInput) (*codedeploy.ListTagsForResourceOutput, error) {
    var output codedeploy.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) ListTagsForResourceAsync(ctx workflow.Context, input *codedeploy.ListTagsForResourceInput) *CodedeployListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &CodedeployListTagsForResourceResult{Result: future}
}
func (a *CodeDeployStub) PutLifecycleEventHookExecutionStatus(ctx workflow.Context, input *codedeploy.PutLifecycleEventHookExecutionStatusInput) (*codedeploy.PutLifecycleEventHookExecutionStatusOutput, error) {
    var output codedeploy.PutLifecycleEventHookExecutionStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutLifecycleEventHookExecutionStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) PutLifecycleEventHookExecutionStatusAsync(ctx workflow.Context, input *codedeploy.PutLifecycleEventHookExecutionStatusInput) *CodedeployPutLifecycleEventHookExecutionStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutLifecycleEventHookExecutionStatus, input)
    return &CodedeployPutLifecycleEventHookExecutionStatusResult{Result: future}
}
func (a *CodeDeployStub) RegisterApplicationRevision(ctx workflow.Context, input *codedeploy.RegisterApplicationRevisionInput) (*codedeploy.RegisterApplicationRevisionOutput, error) {
    var output codedeploy.RegisterApplicationRevisionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterApplicationRevision, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) RegisterApplicationRevisionAsync(ctx workflow.Context, input *codedeploy.RegisterApplicationRevisionInput) *CodedeployRegisterApplicationRevisionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterApplicationRevision, input)
    return &CodedeployRegisterApplicationRevisionResult{Result: future}
}
func (a *CodeDeployStub) RegisterOnPremisesInstance(ctx workflow.Context, input *codedeploy.RegisterOnPremisesInstanceInput) (*codedeploy.RegisterOnPremisesInstanceOutput, error) {
    var output codedeploy.RegisterOnPremisesInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterOnPremisesInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) RegisterOnPremisesInstanceAsync(ctx workflow.Context, input *codedeploy.RegisterOnPremisesInstanceInput) *CodedeployRegisterOnPremisesInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterOnPremisesInstance, input)
    return &CodedeployRegisterOnPremisesInstanceResult{Result: future}
}
func (a *CodeDeployStub) RemoveTagsFromOnPremisesInstances(ctx workflow.Context, input *codedeploy.RemoveTagsFromOnPremisesInstancesInput) (*codedeploy.RemoveTagsFromOnPremisesInstancesOutput, error) {
    var output codedeploy.RemoveTagsFromOnPremisesInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromOnPremisesInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) RemoveTagsFromOnPremisesInstancesAsync(ctx workflow.Context, input *codedeploy.RemoveTagsFromOnPremisesInstancesInput) *CodedeployRemoveTagsFromOnPremisesInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromOnPremisesInstances, input)
    return &CodedeployRemoveTagsFromOnPremisesInstancesResult{Result: future}
}
func (a *CodeDeployStub) SkipWaitTimeForInstanceTermination(ctx workflow.Context, input *codedeploy.SkipWaitTimeForInstanceTerminationInput) (*codedeploy.SkipWaitTimeForInstanceTerminationOutput, error) {
    var output codedeploy.SkipWaitTimeForInstanceTerminationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SkipWaitTimeForInstanceTermination, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) SkipWaitTimeForInstanceTerminationAsync(ctx workflow.Context, input *codedeploy.SkipWaitTimeForInstanceTerminationInput) *CodedeploySkipWaitTimeForInstanceTerminationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SkipWaitTimeForInstanceTermination, input)
    return &CodedeploySkipWaitTimeForInstanceTerminationResult{Result: future}
}
func (a *CodeDeployStub) StopDeployment(ctx workflow.Context, input *codedeploy.StopDeploymentInput) (*codedeploy.StopDeploymentOutput, error) {
    var output codedeploy.StopDeploymentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopDeployment, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) StopDeploymentAsync(ctx workflow.Context, input *codedeploy.StopDeploymentInput) *CodedeployStopDeploymentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopDeployment, input)
    return &CodedeployStopDeploymentResult{Result: future}
}
func (a *CodeDeployStub) TagResource(ctx workflow.Context, input *codedeploy.TagResourceInput) (*codedeploy.TagResourceOutput, error) {
    var output codedeploy.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) TagResourceAsync(ctx workflow.Context, input *codedeploy.TagResourceInput) *CodedeployTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &CodedeployTagResourceResult{Result: future}
}
func (a *CodeDeployStub) UntagResource(ctx workflow.Context, input *codedeploy.UntagResourceInput) (*codedeploy.UntagResourceOutput, error) {
    var output codedeploy.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) UntagResourceAsync(ctx workflow.Context, input *codedeploy.UntagResourceInput) *CodedeployUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &CodedeployUntagResourceResult{Result: future}
}
func (a *CodeDeployStub) UpdateApplication(ctx workflow.Context, input *codedeploy.UpdateApplicationInput) (*codedeploy.UpdateApplicationOutput, error) {
    var output codedeploy.UpdateApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) UpdateApplicationAsync(ctx workflow.Context, input *codedeploy.UpdateApplicationInput) *CodedeployUpdateApplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateApplication, input)
    return &CodedeployUpdateApplicationResult{Result: future}
}
func (a *CodeDeployStub) UpdateDeploymentGroup(ctx workflow.Context, input *codedeploy.UpdateDeploymentGroupInput) (*codedeploy.UpdateDeploymentGroupOutput, error) {
    var output codedeploy.UpdateDeploymentGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDeploymentGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeDeployStub) UpdateDeploymentGroupAsync(ctx workflow.Context, input *codedeploy.UpdateDeploymentGroupInput) *CodedeployUpdateDeploymentGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDeploymentGroup, input)
    return &CodedeployUpdateDeploymentGroupResult{Result: future}
}

func (a *CodeDeployStub) WaitUntilDeploymentSuccessful(ctx workflow.Context, input *codedeploy.GetDeploymentInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilDeploymentSuccessful, input).Get(ctx, nil)
}

func (a *CodeDeployStub) WaitUntilDeploymentSuccessfulAsync(ctx workflow.Context, input *codedeploy.GetDeploymentInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilDeploymentSuccessful, input)
}