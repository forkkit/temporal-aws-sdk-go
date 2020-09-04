package awsclients

import (
	"github.com/aws/aws-sdk-go/service/lambda"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type LambdaClient interface {
    AddLayerVersionPermission(ctx workflow.Context, input *lambda.AddLayerVersionPermissionInput) (*lambda.AddLayerVersionPermissionOutput, error)
    AddLayerVersionPermissionAsync(ctx workflow.Context, input *lambda.AddLayerVersionPermissionInput) *LambdaAddLayerVersionPermissionResult

    AddPermission(ctx workflow.Context, input *lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error)
    AddPermissionAsync(ctx workflow.Context, input *lambda.AddPermissionInput) *LambdaAddPermissionResult

    CreateAlias(ctx workflow.Context, input *lambda.CreateAliasInput) (*lambda.AliasConfiguration, error)
    CreateAliasAsync(ctx workflow.Context, input *lambda.CreateAliasInput) *LambdaCreateAliasResult

    CreateEventSourceMapping(ctx workflow.Context, input *lambda.CreateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
    CreateEventSourceMappingAsync(ctx workflow.Context, input *lambda.CreateEventSourceMappingInput) *LambdaCreateEventSourceMappingResult

    CreateFunction(ctx workflow.Context, input *lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, error)
    CreateFunctionAsync(ctx workflow.Context, input *lambda.CreateFunctionInput) *LambdaCreateFunctionResult

    DeleteAlias(ctx workflow.Context, input *lambda.DeleteAliasInput) (*lambda.DeleteAliasOutput, error)
    DeleteAliasAsync(ctx workflow.Context, input *lambda.DeleteAliasInput) *LambdaDeleteAliasResult

    DeleteEventSourceMapping(ctx workflow.Context, input *lambda.DeleteEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
    DeleteEventSourceMappingAsync(ctx workflow.Context, input *lambda.DeleteEventSourceMappingInput) *LambdaDeleteEventSourceMappingResult

    DeleteFunction(ctx workflow.Context, input *lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error)
    DeleteFunctionAsync(ctx workflow.Context, input *lambda.DeleteFunctionInput) *LambdaDeleteFunctionResult

    DeleteFunctionConcurrency(ctx workflow.Context, input *lambda.DeleteFunctionConcurrencyInput) (*lambda.DeleteFunctionConcurrencyOutput, error)
    DeleteFunctionConcurrencyAsync(ctx workflow.Context, input *lambda.DeleteFunctionConcurrencyInput) *LambdaDeleteFunctionConcurrencyResult

    DeleteFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.DeleteFunctionEventInvokeConfigInput) (*lambda.DeleteFunctionEventInvokeConfigOutput, error)
    DeleteFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.DeleteFunctionEventInvokeConfigInput) *LambdaDeleteFunctionEventInvokeConfigResult

    DeleteLayerVersion(ctx workflow.Context, input *lambda.DeleteLayerVersionInput) (*lambda.DeleteLayerVersionOutput, error)
    DeleteLayerVersionAsync(ctx workflow.Context, input *lambda.DeleteLayerVersionInput) *LambdaDeleteLayerVersionResult

    DeleteProvisionedConcurrencyConfig(ctx workflow.Context, input *lambda.DeleteProvisionedConcurrencyConfigInput) (*lambda.DeleteProvisionedConcurrencyConfigOutput, error)
    DeleteProvisionedConcurrencyConfigAsync(ctx workflow.Context, input *lambda.DeleteProvisionedConcurrencyConfigInput) *LambdaDeleteProvisionedConcurrencyConfigResult

    GetAccountSettings(ctx workflow.Context, input *lambda.GetAccountSettingsInput) (*lambda.GetAccountSettingsOutput, error)
    GetAccountSettingsAsync(ctx workflow.Context, input *lambda.GetAccountSettingsInput) *LambdaGetAccountSettingsResult

    GetAlias(ctx workflow.Context, input *lambda.GetAliasInput) (*lambda.AliasConfiguration, error)
    GetAliasAsync(ctx workflow.Context, input *lambda.GetAliasInput) *LambdaGetAliasResult

    GetEventSourceMapping(ctx workflow.Context, input *lambda.GetEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
    GetEventSourceMappingAsync(ctx workflow.Context, input *lambda.GetEventSourceMappingInput) *LambdaGetEventSourceMappingResult

    GetFunction(ctx workflow.Context, input *lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error)
    GetFunctionAsync(ctx workflow.Context, input *lambda.GetFunctionInput) *LambdaGetFunctionResult

    GetFunctionConcurrency(ctx workflow.Context, input *lambda.GetFunctionConcurrencyInput) (*lambda.GetFunctionConcurrencyOutput, error)
    GetFunctionConcurrencyAsync(ctx workflow.Context, input *lambda.GetFunctionConcurrencyInput) *LambdaGetFunctionConcurrencyResult

    GetFunctionConfiguration(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)
    GetFunctionConfigurationAsync(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) *LambdaGetFunctionConfigurationResult

    GetFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.GetFunctionEventInvokeConfigInput) (*lambda.GetFunctionEventInvokeConfigOutput, error)
    GetFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.GetFunctionEventInvokeConfigInput) *LambdaGetFunctionEventInvokeConfigResult

    GetLayerVersion(ctx workflow.Context, input *lambda.GetLayerVersionInput) (*lambda.GetLayerVersionOutput, error)
    GetLayerVersionAsync(ctx workflow.Context, input *lambda.GetLayerVersionInput) *LambdaGetLayerVersionResult

    GetLayerVersionByArn(ctx workflow.Context, input *lambda.GetLayerVersionByArnInput) (*lambda.GetLayerVersionByArnOutput, error)
    GetLayerVersionByArnAsync(ctx workflow.Context, input *lambda.GetLayerVersionByArnInput) *LambdaGetLayerVersionByArnResult

    GetLayerVersionPolicy(ctx workflow.Context, input *lambda.GetLayerVersionPolicyInput) (*lambda.GetLayerVersionPolicyOutput, error)
    GetLayerVersionPolicyAsync(ctx workflow.Context, input *lambda.GetLayerVersionPolicyInput) *LambdaGetLayerVersionPolicyResult

    GetPolicy(ctx workflow.Context, input *lambda.GetPolicyInput) (*lambda.GetPolicyOutput, error)
    GetPolicyAsync(ctx workflow.Context, input *lambda.GetPolicyInput) *LambdaGetPolicyResult

    GetProvisionedConcurrencyConfig(ctx workflow.Context, input *lambda.GetProvisionedConcurrencyConfigInput) (*lambda.GetProvisionedConcurrencyConfigOutput, error)
    GetProvisionedConcurrencyConfigAsync(ctx workflow.Context, input *lambda.GetProvisionedConcurrencyConfigInput) *LambdaGetProvisionedConcurrencyConfigResult

    Invoke(ctx workflow.Context, input *lambda.InvokeInput) (*lambda.InvokeOutput, error)
    InvokeAsync(ctx workflow.Context, input *lambda.InvokeInput) *LambdaInvokeResult

    ListAliases(ctx workflow.Context, input *lambda.ListAliasesInput) (*lambda.ListAliasesOutput, error)
    ListAliasesAsync(ctx workflow.Context, input *lambda.ListAliasesInput) *LambdaListAliasesResult

    ListEventSourceMappings(ctx workflow.Context, input *lambda.ListEventSourceMappingsInput) (*lambda.ListEventSourceMappingsOutput, error)
    ListEventSourceMappingsAsync(ctx workflow.Context, input *lambda.ListEventSourceMappingsInput) *LambdaListEventSourceMappingsResult

    ListFunctionEventInvokeConfigs(ctx workflow.Context, input *lambda.ListFunctionEventInvokeConfigsInput) (*lambda.ListFunctionEventInvokeConfigsOutput, error)
    ListFunctionEventInvokeConfigsAsync(ctx workflow.Context, input *lambda.ListFunctionEventInvokeConfigsInput) *LambdaListFunctionEventInvokeConfigsResult

    ListFunctions(ctx workflow.Context, input *lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error)
    ListFunctionsAsync(ctx workflow.Context, input *lambda.ListFunctionsInput) *LambdaListFunctionsResult

    ListLayerVersions(ctx workflow.Context, input *lambda.ListLayerVersionsInput) (*lambda.ListLayerVersionsOutput, error)
    ListLayerVersionsAsync(ctx workflow.Context, input *lambda.ListLayerVersionsInput) *LambdaListLayerVersionsResult

    ListLayers(ctx workflow.Context, input *lambda.ListLayersInput) (*lambda.ListLayersOutput, error)
    ListLayersAsync(ctx workflow.Context, input *lambda.ListLayersInput) *LambdaListLayersResult

    ListProvisionedConcurrencyConfigs(ctx workflow.Context, input *lambda.ListProvisionedConcurrencyConfigsInput) (*lambda.ListProvisionedConcurrencyConfigsOutput, error)
    ListProvisionedConcurrencyConfigsAsync(ctx workflow.Context, input *lambda.ListProvisionedConcurrencyConfigsInput) *LambdaListProvisionedConcurrencyConfigsResult

    ListTags(ctx workflow.Context, input *lambda.ListTagsInput) (*lambda.ListTagsOutput, error)
    ListTagsAsync(ctx workflow.Context, input *lambda.ListTagsInput) *LambdaListTagsResult

    ListVersionsByFunction(ctx workflow.Context, input *lambda.ListVersionsByFunctionInput) (*lambda.ListVersionsByFunctionOutput, error)
    ListVersionsByFunctionAsync(ctx workflow.Context, input *lambda.ListVersionsByFunctionInput) *LambdaListVersionsByFunctionResult

    PublishLayerVersion(ctx workflow.Context, input *lambda.PublishLayerVersionInput) (*lambda.PublishLayerVersionOutput, error)
    PublishLayerVersionAsync(ctx workflow.Context, input *lambda.PublishLayerVersionInput) *LambdaPublishLayerVersionResult

    PublishVersion(ctx workflow.Context, input *lambda.PublishVersionInput) (*lambda.FunctionConfiguration, error)
    PublishVersionAsync(ctx workflow.Context, input *lambda.PublishVersionInput) *LambdaPublishVersionResult

    PutFunctionConcurrency(ctx workflow.Context, input *lambda.PutFunctionConcurrencyInput) (*lambda.PutFunctionConcurrencyOutput, error)
    PutFunctionConcurrencyAsync(ctx workflow.Context, input *lambda.PutFunctionConcurrencyInput) *LambdaPutFunctionConcurrencyResult

    PutFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.PutFunctionEventInvokeConfigInput) (*lambda.PutFunctionEventInvokeConfigOutput, error)
    PutFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.PutFunctionEventInvokeConfigInput) *LambdaPutFunctionEventInvokeConfigResult

    PutProvisionedConcurrencyConfig(ctx workflow.Context, input *lambda.PutProvisionedConcurrencyConfigInput) (*lambda.PutProvisionedConcurrencyConfigOutput, error)
    PutProvisionedConcurrencyConfigAsync(ctx workflow.Context, input *lambda.PutProvisionedConcurrencyConfigInput) *LambdaPutProvisionedConcurrencyConfigResult

    RemoveLayerVersionPermission(ctx workflow.Context, input *lambda.RemoveLayerVersionPermissionInput) (*lambda.RemoveLayerVersionPermissionOutput, error)
    RemoveLayerVersionPermissionAsync(ctx workflow.Context, input *lambda.RemoveLayerVersionPermissionInput) *LambdaRemoveLayerVersionPermissionResult

    RemovePermission(ctx workflow.Context, input *lambda.RemovePermissionInput) (*lambda.RemovePermissionOutput, error)
    RemovePermissionAsync(ctx workflow.Context, input *lambda.RemovePermissionInput) *LambdaRemovePermissionResult

    TagResource(ctx workflow.Context, input *lambda.TagResourceInput) (*lambda.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *lambda.TagResourceInput) *LambdaTagResourceResult

    UntagResource(ctx workflow.Context, input *lambda.UntagResourceInput) (*lambda.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *lambda.UntagResourceInput) *LambdaUntagResourceResult

    UpdateAlias(ctx workflow.Context, input *lambda.UpdateAliasInput) (*lambda.AliasConfiguration, error)
    UpdateAliasAsync(ctx workflow.Context, input *lambda.UpdateAliasInput) *LambdaUpdateAliasResult

    UpdateEventSourceMapping(ctx workflow.Context, input *lambda.UpdateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
    UpdateEventSourceMappingAsync(ctx workflow.Context, input *lambda.UpdateEventSourceMappingInput) *LambdaUpdateEventSourceMappingResult

    UpdateFunctionCode(ctx workflow.Context, input *lambda.UpdateFunctionCodeInput) (*lambda.FunctionConfiguration, error)
    UpdateFunctionCodeAsync(ctx workflow.Context, input *lambda.UpdateFunctionCodeInput) *LambdaUpdateFunctionCodeResult

    UpdateFunctionConfiguration(ctx workflow.Context, input *lambda.UpdateFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)
    UpdateFunctionConfigurationAsync(ctx workflow.Context, input *lambda.UpdateFunctionConfigurationInput) *LambdaUpdateFunctionConfigurationResult

    UpdateFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.UpdateFunctionEventInvokeConfigInput) (*lambda.UpdateFunctionEventInvokeConfigOutput, error)
    UpdateFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.UpdateFunctionEventInvokeConfigInput) *LambdaUpdateFunctionEventInvokeConfigResult

    WaitUntilFunctionActive(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) error
    WaitUntilFunctionExists(ctx workflow.Context, input *lambda.GetFunctionInput) error
    WaitUntilFunctionUpdated(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) error}
type LambdaAddLayerVersionPermissionResult struct {
	Result workflow.Future
}

func (r *LambdaAddLayerVersionPermissionResult) Get(ctx workflow.Context) (*lambda.AddLayerVersionPermissionOutput, error) {
    var output lambda.AddLayerVersionPermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaAddPermissionResult struct {
	Result workflow.Future
}

func (r *LambdaAddPermissionResult) Get(ctx workflow.Context) (*lambda.AddPermissionOutput, error) {
    var output lambda.AddPermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaCreateAliasResult struct {
	Result workflow.Future
}

func (r *LambdaCreateAliasResult) Get(ctx workflow.Context) (*lambda.AliasConfiguration, error) {
    var output lambda.AliasConfiguration
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaCreateEventSourceMappingResult struct {
	Result workflow.Future
}

func (r *LambdaCreateEventSourceMappingResult) Get(ctx workflow.Context) (*lambda.EventSourceMappingConfiguration, error) {
    var output lambda.EventSourceMappingConfiguration
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaCreateFunctionResult struct {
	Result workflow.Future
}

func (r *LambdaCreateFunctionResult) Get(ctx workflow.Context) (*lambda.FunctionConfiguration, error) {
    var output lambda.FunctionConfiguration
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaDeleteAliasResult struct {
	Result workflow.Future
}

func (r *LambdaDeleteAliasResult) Get(ctx workflow.Context) (*lambda.DeleteAliasOutput, error) {
    var output lambda.DeleteAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaDeleteEventSourceMappingResult struct {
	Result workflow.Future
}

func (r *LambdaDeleteEventSourceMappingResult) Get(ctx workflow.Context) (*lambda.EventSourceMappingConfiguration, error) {
    var output lambda.EventSourceMappingConfiguration
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaDeleteFunctionResult struct {
	Result workflow.Future
}

func (r *LambdaDeleteFunctionResult) Get(ctx workflow.Context) (*lambda.DeleteFunctionOutput, error) {
    var output lambda.DeleteFunctionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaDeleteFunctionConcurrencyResult struct {
	Result workflow.Future
}

func (r *LambdaDeleteFunctionConcurrencyResult) Get(ctx workflow.Context) (*lambda.DeleteFunctionConcurrencyOutput, error) {
    var output lambda.DeleteFunctionConcurrencyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaDeleteFunctionEventInvokeConfigResult struct {
	Result workflow.Future
}

func (r *LambdaDeleteFunctionEventInvokeConfigResult) Get(ctx workflow.Context) (*lambda.DeleteFunctionEventInvokeConfigOutput, error) {
    var output lambda.DeleteFunctionEventInvokeConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaDeleteLayerVersionResult struct {
	Result workflow.Future
}

func (r *LambdaDeleteLayerVersionResult) Get(ctx workflow.Context) (*lambda.DeleteLayerVersionOutput, error) {
    var output lambda.DeleteLayerVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaDeleteProvisionedConcurrencyConfigResult struct {
	Result workflow.Future
}

func (r *LambdaDeleteProvisionedConcurrencyConfigResult) Get(ctx workflow.Context) (*lambda.DeleteProvisionedConcurrencyConfigOutput, error) {
    var output lambda.DeleteProvisionedConcurrencyConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaGetAccountSettingsResult struct {
	Result workflow.Future
}

func (r *LambdaGetAccountSettingsResult) Get(ctx workflow.Context) (*lambda.GetAccountSettingsOutput, error) {
    var output lambda.GetAccountSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaGetAliasResult struct {
	Result workflow.Future
}

func (r *LambdaGetAliasResult) Get(ctx workflow.Context) (*lambda.AliasConfiguration, error) {
    var output lambda.AliasConfiguration
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaGetEventSourceMappingResult struct {
	Result workflow.Future
}

func (r *LambdaGetEventSourceMappingResult) Get(ctx workflow.Context) (*lambda.EventSourceMappingConfiguration, error) {
    var output lambda.EventSourceMappingConfiguration
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaGetFunctionResult struct {
	Result workflow.Future
}

func (r *LambdaGetFunctionResult) Get(ctx workflow.Context) (*lambda.GetFunctionOutput, error) {
    var output lambda.GetFunctionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaGetFunctionConcurrencyResult struct {
	Result workflow.Future
}

func (r *LambdaGetFunctionConcurrencyResult) Get(ctx workflow.Context) (*lambda.GetFunctionConcurrencyOutput, error) {
    var output lambda.GetFunctionConcurrencyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaGetFunctionConfigurationResult struct {
	Result workflow.Future
}

func (r *LambdaGetFunctionConfigurationResult) Get(ctx workflow.Context) (*lambda.FunctionConfiguration, error) {
    var output lambda.FunctionConfiguration
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaGetFunctionEventInvokeConfigResult struct {
	Result workflow.Future
}

func (r *LambdaGetFunctionEventInvokeConfigResult) Get(ctx workflow.Context) (*lambda.GetFunctionEventInvokeConfigOutput, error) {
    var output lambda.GetFunctionEventInvokeConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaGetLayerVersionResult struct {
	Result workflow.Future
}

func (r *LambdaGetLayerVersionResult) Get(ctx workflow.Context) (*lambda.GetLayerVersionOutput, error) {
    var output lambda.GetLayerVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaGetLayerVersionByArnResult struct {
	Result workflow.Future
}

func (r *LambdaGetLayerVersionByArnResult) Get(ctx workflow.Context) (*lambda.GetLayerVersionByArnOutput, error) {
    var output lambda.GetLayerVersionByArnOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaGetLayerVersionPolicyResult struct {
	Result workflow.Future
}

func (r *LambdaGetLayerVersionPolicyResult) Get(ctx workflow.Context) (*lambda.GetLayerVersionPolicyOutput, error) {
    var output lambda.GetLayerVersionPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaGetPolicyResult struct {
	Result workflow.Future
}

func (r *LambdaGetPolicyResult) Get(ctx workflow.Context) (*lambda.GetPolicyOutput, error) {
    var output lambda.GetPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaGetProvisionedConcurrencyConfigResult struct {
	Result workflow.Future
}

func (r *LambdaGetProvisionedConcurrencyConfigResult) Get(ctx workflow.Context) (*lambda.GetProvisionedConcurrencyConfigOutput, error) {
    var output lambda.GetProvisionedConcurrencyConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaInvokeResult struct {
	Result workflow.Future
}

func (r *LambdaInvokeResult) Get(ctx workflow.Context) (*lambda.InvokeOutput, error) {
    var output lambda.InvokeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaListAliasesResult struct {
	Result workflow.Future
}

func (r *LambdaListAliasesResult) Get(ctx workflow.Context) (*lambda.ListAliasesOutput, error) {
    var output lambda.ListAliasesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaListEventSourceMappingsResult struct {
	Result workflow.Future
}

func (r *LambdaListEventSourceMappingsResult) Get(ctx workflow.Context) (*lambda.ListEventSourceMappingsOutput, error) {
    var output lambda.ListEventSourceMappingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaListFunctionEventInvokeConfigsResult struct {
	Result workflow.Future
}

func (r *LambdaListFunctionEventInvokeConfigsResult) Get(ctx workflow.Context) (*lambda.ListFunctionEventInvokeConfigsOutput, error) {
    var output lambda.ListFunctionEventInvokeConfigsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaListFunctionsResult struct {
	Result workflow.Future
}

func (r *LambdaListFunctionsResult) Get(ctx workflow.Context) (*lambda.ListFunctionsOutput, error) {
    var output lambda.ListFunctionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaListLayerVersionsResult struct {
	Result workflow.Future
}

func (r *LambdaListLayerVersionsResult) Get(ctx workflow.Context) (*lambda.ListLayerVersionsOutput, error) {
    var output lambda.ListLayerVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaListLayersResult struct {
	Result workflow.Future
}

func (r *LambdaListLayersResult) Get(ctx workflow.Context) (*lambda.ListLayersOutput, error) {
    var output lambda.ListLayersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaListProvisionedConcurrencyConfigsResult struct {
	Result workflow.Future
}

func (r *LambdaListProvisionedConcurrencyConfigsResult) Get(ctx workflow.Context) (*lambda.ListProvisionedConcurrencyConfigsOutput, error) {
    var output lambda.ListProvisionedConcurrencyConfigsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaListTagsResult struct {
	Result workflow.Future
}

func (r *LambdaListTagsResult) Get(ctx workflow.Context) (*lambda.ListTagsOutput, error) {
    var output lambda.ListTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaListVersionsByFunctionResult struct {
	Result workflow.Future
}

func (r *LambdaListVersionsByFunctionResult) Get(ctx workflow.Context) (*lambda.ListVersionsByFunctionOutput, error) {
    var output lambda.ListVersionsByFunctionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaPublishLayerVersionResult struct {
	Result workflow.Future
}

func (r *LambdaPublishLayerVersionResult) Get(ctx workflow.Context) (*lambda.PublishLayerVersionOutput, error) {
    var output lambda.PublishLayerVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaPublishVersionResult struct {
	Result workflow.Future
}

func (r *LambdaPublishVersionResult) Get(ctx workflow.Context) (*lambda.FunctionConfiguration, error) {
    var output lambda.FunctionConfiguration
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaPutFunctionConcurrencyResult struct {
	Result workflow.Future
}

func (r *LambdaPutFunctionConcurrencyResult) Get(ctx workflow.Context) (*lambda.PutFunctionConcurrencyOutput, error) {
    var output lambda.PutFunctionConcurrencyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaPutFunctionEventInvokeConfigResult struct {
	Result workflow.Future
}

func (r *LambdaPutFunctionEventInvokeConfigResult) Get(ctx workflow.Context) (*lambda.PutFunctionEventInvokeConfigOutput, error) {
    var output lambda.PutFunctionEventInvokeConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaPutProvisionedConcurrencyConfigResult struct {
	Result workflow.Future
}

func (r *LambdaPutProvisionedConcurrencyConfigResult) Get(ctx workflow.Context) (*lambda.PutProvisionedConcurrencyConfigOutput, error) {
    var output lambda.PutProvisionedConcurrencyConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaRemoveLayerVersionPermissionResult struct {
	Result workflow.Future
}

func (r *LambdaRemoveLayerVersionPermissionResult) Get(ctx workflow.Context) (*lambda.RemoveLayerVersionPermissionOutput, error) {
    var output lambda.RemoveLayerVersionPermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaRemovePermissionResult struct {
	Result workflow.Future
}

func (r *LambdaRemovePermissionResult) Get(ctx workflow.Context) (*lambda.RemovePermissionOutput, error) {
    var output lambda.RemovePermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaTagResourceResult struct {
	Result workflow.Future
}

func (r *LambdaTagResourceResult) Get(ctx workflow.Context) (*lambda.TagResourceOutput, error) {
    var output lambda.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaUntagResourceResult struct {
	Result workflow.Future
}

func (r *LambdaUntagResourceResult) Get(ctx workflow.Context) (*lambda.UntagResourceOutput, error) {
    var output lambda.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaUpdateAliasResult struct {
	Result workflow.Future
}

func (r *LambdaUpdateAliasResult) Get(ctx workflow.Context) (*lambda.AliasConfiguration, error) {
    var output lambda.AliasConfiguration
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaUpdateEventSourceMappingResult struct {
	Result workflow.Future
}

func (r *LambdaUpdateEventSourceMappingResult) Get(ctx workflow.Context) (*lambda.EventSourceMappingConfiguration, error) {
    var output lambda.EventSourceMappingConfiguration
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaUpdateFunctionCodeResult struct {
	Result workflow.Future
}

func (r *LambdaUpdateFunctionCodeResult) Get(ctx workflow.Context) (*lambda.FunctionConfiguration, error) {
    var output lambda.FunctionConfiguration
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaUpdateFunctionConfigurationResult struct {
	Result workflow.Future
}

func (r *LambdaUpdateFunctionConfigurationResult) Get(ctx workflow.Context) (*lambda.FunctionConfiguration, error) {
    var output lambda.FunctionConfiguration
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LambdaUpdateFunctionEventInvokeConfigResult struct {
	Result workflow.Future
}

func (r *LambdaUpdateFunctionEventInvokeConfigResult) Get(ctx workflow.Context) (*lambda.UpdateFunctionEventInvokeConfigOutput, error) {
    var output lambda.UpdateFunctionEventInvokeConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type LambdaStub struct {
    activities awsactivities.LambdaActivities
}

func NewLambdaStub() LambdaClient {
    return &LambdaStub{}
}
func (a *LambdaStub) AddLayerVersionPermission(ctx workflow.Context, input *lambda.AddLayerVersionPermissionInput) (*lambda.AddLayerVersionPermissionOutput, error) {
    var output lambda.AddLayerVersionPermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddLayerVersionPermission, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) AddLayerVersionPermissionAsync(ctx workflow.Context, input *lambda.AddLayerVersionPermissionInput) *LambdaAddLayerVersionPermissionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddLayerVersionPermission, input)
    return &LambdaAddLayerVersionPermissionResult{Result: future}
}
func (a *LambdaStub) AddPermission(ctx workflow.Context, input *lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error) {
    var output lambda.AddPermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddPermission, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) AddPermissionAsync(ctx workflow.Context, input *lambda.AddPermissionInput) *LambdaAddPermissionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddPermission, input)
    return &LambdaAddPermissionResult{Result: future}
}
func (a *LambdaStub) CreateAlias(ctx workflow.Context, input *lambda.CreateAliasInput) (*lambda.AliasConfiguration, error) {
    var output lambda.AliasConfiguration
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) CreateAliasAsync(ctx workflow.Context, input *lambda.CreateAliasInput) *LambdaCreateAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateAlias, input)
    return &LambdaCreateAliasResult{Result: future}
}
func (a *LambdaStub) CreateEventSourceMapping(ctx workflow.Context, input *lambda.CreateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
    var output lambda.EventSourceMappingConfiguration
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEventSourceMapping, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) CreateEventSourceMappingAsync(ctx workflow.Context, input *lambda.CreateEventSourceMappingInput) *LambdaCreateEventSourceMappingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateEventSourceMapping, input)
    return &LambdaCreateEventSourceMappingResult{Result: future}
}
func (a *LambdaStub) CreateFunction(ctx workflow.Context, input *lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, error) {
    var output lambda.FunctionConfiguration
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFunction, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) CreateFunctionAsync(ctx workflow.Context, input *lambda.CreateFunctionInput) *LambdaCreateFunctionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateFunction, input)
    return &LambdaCreateFunctionResult{Result: future}
}
func (a *LambdaStub) DeleteAlias(ctx workflow.Context, input *lambda.DeleteAliasInput) (*lambda.DeleteAliasOutput, error) {
    var output lambda.DeleteAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) DeleteAliasAsync(ctx workflow.Context, input *lambda.DeleteAliasInput) *LambdaDeleteAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAlias, input)
    return &LambdaDeleteAliasResult{Result: future}
}
func (a *LambdaStub) DeleteEventSourceMapping(ctx workflow.Context, input *lambda.DeleteEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
    var output lambda.EventSourceMappingConfiguration
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEventSourceMapping, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) DeleteEventSourceMappingAsync(ctx workflow.Context, input *lambda.DeleteEventSourceMappingInput) *LambdaDeleteEventSourceMappingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEventSourceMapping, input)
    return &LambdaDeleteEventSourceMappingResult{Result: future}
}
func (a *LambdaStub) DeleteFunction(ctx workflow.Context, input *lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error) {
    var output lambda.DeleteFunctionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFunction, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) DeleteFunctionAsync(ctx workflow.Context, input *lambda.DeleteFunctionInput) *LambdaDeleteFunctionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFunction, input)
    return &LambdaDeleteFunctionResult{Result: future}
}
func (a *LambdaStub) DeleteFunctionConcurrency(ctx workflow.Context, input *lambda.DeleteFunctionConcurrencyInput) (*lambda.DeleteFunctionConcurrencyOutput, error) {
    var output lambda.DeleteFunctionConcurrencyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFunctionConcurrency, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) DeleteFunctionConcurrencyAsync(ctx workflow.Context, input *lambda.DeleteFunctionConcurrencyInput) *LambdaDeleteFunctionConcurrencyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFunctionConcurrency, input)
    return &LambdaDeleteFunctionConcurrencyResult{Result: future}
}
func (a *LambdaStub) DeleteFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.DeleteFunctionEventInvokeConfigInput) (*lambda.DeleteFunctionEventInvokeConfigOutput, error) {
    var output lambda.DeleteFunctionEventInvokeConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFunctionEventInvokeConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) DeleteFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.DeleteFunctionEventInvokeConfigInput) *LambdaDeleteFunctionEventInvokeConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFunctionEventInvokeConfig, input)
    return &LambdaDeleteFunctionEventInvokeConfigResult{Result: future}
}
func (a *LambdaStub) DeleteLayerVersion(ctx workflow.Context, input *lambda.DeleteLayerVersionInput) (*lambda.DeleteLayerVersionOutput, error) {
    var output lambda.DeleteLayerVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLayerVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) DeleteLayerVersionAsync(ctx workflow.Context, input *lambda.DeleteLayerVersionInput) *LambdaDeleteLayerVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLayerVersion, input)
    return &LambdaDeleteLayerVersionResult{Result: future}
}
func (a *LambdaStub) DeleteProvisionedConcurrencyConfig(ctx workflow.Context, input *lambda.DeleteProvisionedConcurrencyConfigInput) (*lambda.DeleteProvisionedConcurrencyConfigOutput, error) {
    var output lambda.DeleteProvisionedConcurrencyConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProvisionedConcurrencyConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) DeleteProvisionedConcurrencyConfigAsync(ctx workflow.Context, input *lambda.DeleteProvisionedConcurrencyConfigInput) *LambdaDeleteProvisionedConcurrencyConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteProvisionedConcurrencyConfig, input)
    return &LambdaDeleteProvisionedConcurrencyConfigResult{Result: future}
}
func (a *LambdaStub) GetAccountSettings(ctx workflow.Context, input *lambda.GetAccountSettingsInput) (*lambda.GetAccountSettingsOutput, error) {
    var output lambda.GetAccountSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccountSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) GetAccountSettingsAsync(ctx workflow.Context, input *lambda.GetAccountSettingsInput) *LambdaGetAccountSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAccountSettings, input)
    return &LambdaGetAccountSettingsResult{Result: future}
}
func (a *LambdaStub) GetAlias(ctx workflow.Context, input *lambda.GetAliasInput) (*lambda.AliasConfiguration, error) {
    var output lambda.AliasConfiguration
    err := workflow.ExecuteActivity(ctx, a.activities.GetAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) GetAliasAsync(ctx workflow.Context, input *lambda.GetAliasInput) *LambdaGetAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAlias, input)
    return &LambdaGetAliasResult{Result: future}
}
func (a *LambdaStub) GetEventSourceMapping(ctx workflow.Context, input *lambda.GetEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
    var output lambda.EventSourceMappingConfiguration
    err := workflow.ExecuteActivity(ctx, a.activities.GetEventSourceMapping, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) GetEventSourceMappingAsync(ctx workflow.Context, input *lambda.GetEventSourceMappingInput) *LambdaGetEventSourceMappingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetEventSourceMapping, input)
    return &LambdaGetEventSourceMappingResult{Result: future}
}
func (a *LambdaStub) GetFunction(ctx workflow.Context, input *lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error) {
    var output lambda.GetFunctionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFunction, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) GetFunctionAsync(ctx workflow.Context, input *lambda.GetFunctionInput) *LambdaGetFunctionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetFunction, input)
    return &LambdaGetFunctionResult{Result: future}
}
func (a *LambdaStub) GetFunctionConcurrency(ctx workflow.Context, input *lambda.GetFunctionConcurrencyInput) (*lambda.GetFunctionConcurrencyOutput, error) {
    var output lambda.GetFunctionConcurrencyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFunctionConcurrency, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) GetFunctionConcurrencyAsync(ctx workflow.Context, input *lambda.GetFunctionConcurrencyInput) *LambdaGetFunctionConcurrencyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetFunctionConcurrency, input)
    return &LambdaGetFunctionConcurrencyResult{Result: future}
}
func (a *LambdaStub) GetFunctionConfiguration(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) (*lambda.FunctionConfiguration, error) {
    var output lambda.FunctionConfiguration
    err := workflow.ExecuteActivity(ctx, a.activities.GetFunctionConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) GetFunctionConfigurationAsync(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) *LambdaGetFunctionConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetFunctionConfiguration, input)
    return &LambdaGetFunctionConfigurationResult{Result: future}
}
func (a *LambdaStub) GetFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.GetFunctionEventInvokeConfigInput) (*lambda.GetFunctionEventInvokeConfigOutput, error) {
    var output lambda.GetFunctionEventInvokeConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFunctionEventInvokeConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) GetFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.GetFunctionEventInvokeConfigInput) *LambdaGetFunctionEventInvokeConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetFunctionEventInvokeConfig, input)
    return &LambdaGetFunctionEventInvokeConfigResult{Result: future}
}
func (a *LambdaStub) GetLayerVersion(ctx workflow.Context, input *lambda.GetLayerVersionInput) (*lambda.GetLayerVersionOutput, error) {
    var output lambda.GetLayerVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLayerVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) GetLayerVersionAsync(ctx workflow.Context, input *lambda.GetLayerVersionInput) *LambdaGetLayerVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetLayerVersion, input)
    return &LambdaGetLayerVersionResult{Result: future}
}
func (a *LambdaStub) GetLayerVersionByArn(ctx workflow.Context, input *lambda.GetLayerVersionByArnInput) (*lambda.GetLayerVersionByArnOutput, error) {
    var output lambda.GetLayerVersionByArnOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLayerVersionByArn, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) GetLayerVersionByArnAsync(ctx workflow.Context, input *lambda.GetLayerVersionByArnInput) *LambdaGetLayerVersionByArnResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetLayerVersionByArn, input)
    return &LambdaGetLayerVersionByArnResult{Result: future}
}
func (a *LambdaStub) GetLayerVersionPolicy(ctx workflow.Context, input *lambda.GetLayerVersionPolicyInput) (*lambda.GetLayerVersionPolicyOutput, error) {
    var output lambda.GetLayerVersionPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLayerVersionPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) GetLayerVersionPolicyAsync(ctx workflow.Context, input *lambda.GetLayerVersionPolicyInput) *LambdaGetLayerVersionPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetLayerVersionPolicy, input)
    return &LambdaGetLayerVersionPolicyResult{Result: future}
}
func (a *LambdaStub) GetPolicy(ctx workflow.Context, input *lambda.GetPolicyInput) (*lambda.GetPolicyOutput, error) {
    var output lambda.GetPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) GetPolicyAsync(ctx workflow.Context, input *lambda.GetPolicyInput) *LambdaGetPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPolicy, input)
    return &LambdaGetPolicyResult{Result: future}
}
func (a *LambdaStub) GetProvisionedConcurrencyConfig(ctx workflow.Context, input *lambda.GetProvisionedConcurrencyConfigInput) (*lambda.GetProvisionedConcurrencyConfigOutput, error) {
    var output lambda.GetProvisionedConcurrencyConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetProvisionedConcurrencyConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) GetProvisionedConcurrencyConfigAsync(ctx workflow.Context, input *lambda.GetProvisionedConcurrencyConfigInput) *LambdaGetProvisionedConcurrencyConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetProvisionedConcurrencyConfig, input)
    return &LambdaGetProvisionedConcurrencyConfigResult{Result: future}
}
func (a *LambdaStub) Invoke(ctx workflow.Context, input *lambda.InvokeInput) (*lambda.InvokeOutput, error) {
    var output lambda.InvokeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.Invoke, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) InvokeAsync(ctx workflow.Context, input *lambda.InvokeInput) *LambdaInvokeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.Invoke, input)
    return &LambdaInvokeResult{Result: future}
}
func (a *LambdaStub) ListAliases(ctx workflow.Context, input *lambda.ListAliasesInput) (*lambda.ListAliasesOutput, error) {
    var output lambda.ListAliasesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAliases, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) ListAliasesAsync(ctx workflow.Context, input *lambda.ListAliasesInput) *LambdaListAliasesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAliases, input)
    return &LambdaListAliasesResult{Result: future}
}
func (a *LambdaStub) ListEventSourceMappings(ctx workflow.Context, input *lambda.ListEventSourceMappingsInput) (*lambda.ListEventSourceMappingsOutput, error) {
    var output lambda.ListEventSourceMappingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEventSourceMappings, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) ListEventSourceMappingsAsync(ctx workflow.Context, input *lambda.ListEventSourceMappingsInput) *LambdaListEventSourceMappingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListEventSourceMappings, input)
    return &LambdaListEventSourceMappingsResult{Result: future}
}
func (a *LambdaStub) ListFunctionEventInvokeConfigs(ctx workflow.Context, input *lambda.ListFunctionEventInvokeConfigsInput) (*lambda.ListFunctionEventInvokeConfigsOutput, error) {
    var output lambda.ListFunctionEventInvokeConfigsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFunctionEventInvokeConfigs, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) ListFunctionEventInvokeConfigsAsync(ctx workflow.Context, input *lambda.ListFunctionEventInvokeConfigsInput) *LambdaListFunctionEventInvokeConfigsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFunctionEventInvokeConfigs, input)
    return &LambdaListFunctionEventInvokeConfigsResult{Result: future}
}
func (a *LambdaStub) ListFunctions(ctx workflow.Context, input *lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error) {
    var output lambda.ListFunctionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFunctions, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) ListFunctionsAsync(ctx workflow.Context, input *lambda.ListFunctionsInput) *LambdaListFunctionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFunctions, input)
    return &LambdaListFunctionsResult{Result: future}
}
func (a *LambdaStub) ListLayerVersions(ctx workflow.Context, input *lambda.ListLayerVersionsInput) (*lambda.ListLayerVersionsOutput, error) {
    var output lambda.ListLayerVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLayerVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) ListLayerVersionsAsync(ctx workflow.Context, input *lambda.ListLayerVersionsInput) *LambdaListLayerVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListLayerVersions, input)
    return &LambdaListLayerVersionsResult{Result: future}
}
func (a *LambdaStub) ListLayers(ctx workflow.Context, input *lambda.ListLayersInput) (*lambda.ListLayersOutput, error) {
    var output lambda.ListLayersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLayers, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) ListLayersAsync(ctx workflow.Context, input *lambda.ListLayersInput) *LambdaListLayersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListLayers, input)
    return &LambdaListLayersResult{Result: future}
}
func (a *LambdaStub) ListProvisionedConcurrencyConfigs(ctx workflow.Context, input *lambda.ListProvisionedConcurrencyConfigsInput) (*lambda.ListProvisionedConcurrencyConfigsOutput, error) {
    var output lambda.ListProvisionedConcurrencyConfigsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProvisionedConcurrencyConfigs, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) ListProvisionedConcurrencyConfigsAsync(ctx workflow.Context, input *lambda.ListProvisionedConcurrencyConfigsInput) *LambdaListProvisionedConcurrencyConfigsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListProvisionedConcurrencyConfigs, input)
    return &LambdaListProvisionedConcurrencyConfigsResult{Result: future}
}
func (a *LambdaStub) ListTags(ctx workflow.Context, input *lambda.ListTagsInput) (*lambda.ListTagsOutput, error) {
    var output lambda.ListTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTags, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) ListTagsAsync(ctx workflow.Context, input *lambda.ListTagsInput) *LambdaListTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTags, input)
    return &LambdaListTagsResult{Result: future}
}
func (a *LambdaStub) ListVersionsByFunction(ctx workflow.Context, input *lambda.ListVersionsByFunctionInput) (*lambda.ListVersionsByFunctionOutput, error) {
    var output lambda.ListVersionsByFunctionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListVersionsByFunction, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) ListVersionsByFunctionAsync(ctx workflow.Context, input *lambda.ListVersionsByFunctionInput) *LambdaListVersionsByFunctionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListVersionsByFunction, input)
    return &LambdaListVersionsByFunctionResult{Result: future}
}
func (a *LambdaStub) PublishLayerVersion(ctx workflow.Context, input *lambda.PublishLayerVersionInput) (*lambda.PublishLayerVersionOutput, error) {
    var output lambda.PublishLayerVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PublishLayerVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) PublishLayerVersionAsync(ctx workflow.Context, input *lambda.PublishLayerVersionInput) *LambdaPublishLayerVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PublishLayerVersion, input)
    return &LambdaPublishLayerVersionResult{Result: future}
}
func (a *LambdaStub) PublishVersion(ctx workflow.Context, input *lambda.PublishVersionInput) (*lambda.FunctionConfiguration, error) {
    var output lambda.FunctionConfiguration
    err := workflow.ExecuteActivity(ctx, a.activities.PublishVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) PublishVersionAsync(ctx workflow.Context, input *lambda.PublishVersionInput) *LambdaPublishVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PublishVersion, input)
    return &LambdaPublishVersionResult{Result: future}
}
func (a *LambdaStub) PutFunctionConcurrency(ctx workflow.Context, input *lambda.PutFunctionConcurrencyInput) (*lambda.PutFunctionConcurrencyOutput, error) {
    var output lambda.PutFunctionConcurrencyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutFunctionConcurrency, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) PutFunctionConcurrencyAsync(ctx workflow.Context, input *lambda.PutFunctionConcurrencyInput) *LambdaPutFunctionConcurrencyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutFunctionConcurrency, input)
    return &LambdaPutFunctionConcurrencyResult{Result: future}
}
func (a *LambdaStub) PutFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.PutFunctionEventInvokeConfigInput) (*lambda.PutFunctionEventInvokeConfigOutput, error) {
    var output lambda.PutFunctionEventInvokeConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutFunctionEventInvokeConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) PutFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.PutFunctionEventInvokeConfigInput) *LambdaPutFunctionEventInvokeConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutFunctionEventInvokeConfig, input)
    return &LambdaPutFunctionEventInvokeConfigResult{Result: future}
}
func (a *LambdaStub) PutProvisionedConcurrencyConfig(ctx workflow.Context, input *lambda.PutProvisionedConcurrencyConfigInput) (*lambda.PutProvisionedConcurrencyConfigOutput, error) {
    var output lambda.PutProvisionedConcurrencyConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutProvisionedConcurrencyConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) PutProvisionedConcurrencyConfigAsync(ctx workflow.Context, input *lambda.PutProvisionedConcurrencyConfigInput) *LambdaPutProvisionedConcurrencyConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutProvisionedConcurrencyConfig, input)
    return &LambdaPutProvisionedConcurrencyConfigResult{Result: future}
}
func (a *LambdaStub) RemoveLayerVersionPermission(ctx workflow.Context, input *lambda.RemoveLayerVersionPermissionInput) (*lambda.RemoveLayerVersionPermissionOutput, error) {
    var output lambda.RemoveLayerVersionPermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveLayerVersionPermission, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) RemoveLayerVersionPermissionAsync(ctx workflow.Context, input *lambda.RemoveLayerVersionPermissionInput) *LambdaRemoveLayerVersionPermissionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveLayerVersionPermission, input)
    return &LambdaRemoveLayerVersionPermissionResult{Result: future}
}
func (a *LambdaStub) RemovePermission(ctx workflow.Context, input *lambda.RemovePermissionInput) (*lambda.RemovePermissionOutput, error) {
    var output lambda.RemovePermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemovePermission, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) RemovePermissionAsync(ctx workflow.Context, input *lambda.RemovePermissionInput) *LambdaRemovePermissionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemovePermission, input)
    return &LambdaRemovePermissionResult{Result: future}
}
func (a *LambdaStub) TagResource(ctx workflow.Context, input *lambda.TagResourceInput) (*lambda.TagResourceOutput, error) {
    var output lambda.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) TagResourceAsync(ctx workflow.Context, input *lambda.TagResourceInput) *LambdaTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &LambdaTagResourceResult{Result: future}
}
func (a *LambdaStub) UntagResource(ctx workflow.Context, input *lambda.UntagResourceInput) (*lambda.UntagResourceOutput, error) {
    var output lambda.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) UntagResourceAsync(ctx workflow.Context, input *lambda.UntagResourceInput) *LambdaUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &LambdaUntagResourceResult{Result: future}
}
func (a *LambdaStub) UpdateAlias(ctx workflow.Context, input *lambda.UpdateAliasInput) (*lambda.AliasConfiguration, error) {
    var output lambda.AliasConfiguration
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) UpdateAliasAsync(ctx workflow.Context, input *lambda.UpdateAliasInput) *LambdaUpdateAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateAlias, input)
    return &LambdaUpdateAliasResult{Result: future}
}
func (a *LambdaStub) UpdateEventSourceMapping(ctx workflow.Context, input *lambda.UpdateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
    var output lambda.EventSourceMappingConfiguration
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateEventSourceMapping, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) UpdateEventSourceMappingAsync(ctx workflow.Context, input *lambda.UpdateEventSourceMappingInput) *LambdaUpdateEventSourceMappingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateEventSourceMapping, input)
    return &LambdaUpdateEventSourceMappingResult{Result: future}
}
func (a *LambdaStub) UpdateFunctionCode(ctx workflow.Context, input *lambda.UpdateFunctionCodeInput) (*lambda.FunctionConfiguration, error) {
    var output lambda.FunctionConfiguration
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFunctionCode, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) UpdateFunctionCodeAsync(ctx workflow.Context, input *lambda.UpdateFunctionCodeInput) *LambdaUpdateFunctionCodeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateFunctionCode, input)
    return &LambdaUpdateFunctionCodeResult{Result: future}
}
func (a *LambdaStub) UpdateFunctionConfiguration(ctx workflow.Context, input *lambda.UpdateFunctionConfigurationInput) (*lambda.FunctionConfiguration, error) {
    var output lambda.FunctionConfiguration
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFunctionConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) UpdateFunctionConfigurationAsync(ctx workflow.Context, input *lambda.UpdateFunctionConfigurationInput) *LambdaUpdateFunctionConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateFunctionConfiguration, input)
    return &LambdaUpdateFunctionConfigurationResult{Result: future}
}
func (a *LambdaStub) UpdateFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.UpdateFunctionEventInvokeConfigInput) (*lambda.UpdateFunctionEventInvokeConfigOutput, error) {
    var output lambda.UpdateFunctionEventInvokeConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFunctionEventInvokeConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *LambdaStub) UpdateFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.UpdateFunctionEventInvokeConfigInput) *LambdaUpdateFunctionEventInvokeConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateFunctionEventInvokeConfig, input)
    return &LambdaUpdateFunctionEventInvokeConfigResult{Result: future}
}

func (a *LambdaStub) WaitUntilFunctionActive(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilFunctionActive, input).Get(ctx, nil)
}

func (a *LambdaStub) WaitUntilFunctionActiveAsync(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilFunctionActive, input)
}

func (a *LambdaStub) WaitUntilFunctionExists(ctx workflow.Context, input *lambda.GetFunctionInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilFunctionExists, input).Get(ctx, nil)
}

func (a *LambdaStub) WaitUntilFunctionExistsAsync(ctx workflow.Context, input *lambda.GetFunctionInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilFunctionExists, input)
}

func (a *LambdaStub) WaitUntilFunctionUpdated(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilFunctionUpdated, input).Get(ctx, nil)
}

func (a *LambdaStub) WaitUntilFunctionUpdatedAsync(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilFunctionUpdated, input)
}