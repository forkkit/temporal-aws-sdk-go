// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package apigatewaystub

import (
	"github.com/aws/aws-sdk-go/service/apigateway"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateApiKey(ctx workflow.Context, input *apigateway.CreateApiKeyInput) (*apigateway.ApiKey, error)
	CreateApiKeyAsync(ctx workflow.Context, input *apigateway.CreateApiKeyInput) *APIGatewayCreateApiKeyFuture

	CreateAuthorizer(ctx workflow.Context, input *apigateway.CreateAuthorizerInput) (*apigateway.Authorizer, error)
	CreateAuthorizerAsync(ctx workflow.Context, input *apigateway.CreateAuthorizerInput) *APIGatewayCreateAuthorizerFuture

	CreateBasePathMapping(ctx workflow.Context, input *apigateway.CreateBasePathMappingInput) (*apigateway.BasePathMapping, error)
	CreateBasePathMappingAsync(ctx workflow.Context, input *apigateway.CreateBasePathMappingInput) *APIGatewayCreateBasePathMappingFuture

	CreateDeployment(ctx workflow.Context, input *apigateway.CreateDeploymentInput) (*apigateway.Deployment, error)
	CreateDeploymentAsync(ctx workflow.Context, input *apigateway.CreateDeploymentInput) *APIGatewayCreateDeploymentFuture

	CreateDocumentationPart(ctx workflow.Context, input *apigateway.CreateDocumentationPartInput) (*apigateway.DocumentationPart, error)
	CreateDocumentationPartAsync(ctx workflow.Context, input *apigateway.CreateDocumentationPartInput) *APIGatewayCreateDocumentationPartFuture

	CreateDocumentationVersion(ctx workflow.Context, input *apigateway.CreateDocumentationVersionInput) (*apigateway.DocumentationVersion, error)
	CreateDocumentationVersionAsync(ctx workflow.Context, input *apigateway.CreateDocumentationVersionInput) *APIGatewayCreateDocumentationVersionFuture

	CreateDomainName(ctx workflow.Context, input *apigateway.CreateDomainNameInput) (*apigateway.DomainName, error)
	CreateDomainNameAsync(ctx workflow.Context, input *apigateway.CreateDomainNameInput) *APIGatewayCreateDomainNameFuture

	CreateModel(ctx workflow.Context, input *apigateway.CreateModelInput) (*apigateway.Model, error)
	CreateModelAsync(ctx workflow.Context, input *apigateway.CreateModelInput) *APIGatewayCreateModelFuture

	CreateRequestValidator(ctx workflow.Context, input *apigateway.CreateRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error)
	CreateRequestValidatorAsync(ctx workflow.Context, input *apigateway.CreateRequestValidatorInput) *APIGatewayCreateRequestValidatorFuture

	CreateResource(ctx workflow.Context, input *apigateway.CreateResourceInput) (*apigateway.Resource, error)
	CreateResourceAsync(ctx workflow.Context, input *apigateway.CreateResourceInput) *APIGatewayCreateResourceFuture

	CreateRestApi(ctx workflow.Context, input *apigateway.CreateRestApiInput) (*apigateway.RestApi, error)
	CreateRestApiAsync(ctx workflow.Context, input *apigateway.CreateRestApiInput) *APIGatewayCreateRestApiFuture

	CreateStage(ctx workflow.Context, input *apigateway.CreateStageInput) (*apigateway.Stage, error)
	CreateStageAsync(ctx workflow.Context, input *apigateway.CreateStageInput) *APIGatewayCreateStageFuture

	CreateUsagePlan(ctx workflow.Context, input *apigateway.CreateUsagePlanInput) (*apigateway.UsagePlan, error)
	CreateUsagePlanAsync(ctx workflow.Context, input *apigateway.CreateUsagePlanInput) *APIGatewayCreateUsagePlanFuture

	CreateUsagePlanKey(ctx workflow.Context, input *apigateway.CreateUsagePlanKeyInput) (*apigateway.UsagePlanKey, error)
	CreateUsagePlanKeyAsync(ctx workflow.Context, input *apigateway.CreateUsagePlanKeyInput) *APIGatewayCreateUsagePlanKeyFuture

	CreateVpcLink(ctx workflow.Context, input *apigateway.CreateVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error)
	CreateVpcLinkAsync(ctx workflow.Context, input *apigateway.CreateVpcLinkInput) *APIGatewayCreateVpcLinkFuture

	DeleteApiKey(ctx workflow.Context, input *apigateway.DeleteApiKeyInput) (*apigateway.DeleteApiKeyOutput, error)
	DeleteApiKeyAsync(ctx workflow.Context, input *apigateway.DeleteApiKeyInput) *APIGatewayDeleteApiKeyFuture

	DeleteAuthorizer(ctx workflow.Context, input *apigateway.DeleteAuthorizerInput) (*apigateway.DeleteAuthorizerOutput, error)
	DeleteAuthorizerAsync(ctx workflow.Context, input *apigateway.DeleteAuthorizerInput) *APIGatewayDeleteAuthorizerFuture

	DeleteBasePathMapping(ctx workflow.Context, input *apigateway.DeleteBasePathMappingInput) (*apigateway.DeleteBasePathMappingOutput, error)
	DeleteBasePathMappingAsync(ctx workflow.Context, input *apigateway.DeleteBasePathMappingInput) *APIGatewayDeleteBasePathMappingFuture

	DeleteClientCertificate(ctx workflow.Context, input *apigateway.DeleteClientCertificateInput) (*apigateway.DeleteClientCertificateOutput, error)
	DeleteClientCertificateAsync(ctx workflow.Context, input *apigateway.DeleteClientCertificateInput) *APIGatewayDeleteClientCertificateFuture

	DeleteDeployment(ctx workflow.Context, input *apigateway.DeleteDeploymentInput) (*apigateway.DeleteDeploymentOutput, error)
	DeleteDeploymentAsync(ctx workflow.Context, input *apigateway.DeleteDeploymentInput) *APIGatewayDeleteDeploymentFuture

	DeleteDocumentationPart(ctx workflow.Context, input *apigateway.DeleteDocumentationPartInput) (*apigateway.DeleteDocumentationPartOutput, error)
	DeleteDocumentationPartAsync(ctx workflow.Context, input *apigateway.DeleteDocumentationPartInput) *APIGatewayDeleteDocumentationPartFuture

	DeleteDocumentationVersion(ctx workflow.Context, input *apigateway.DeleteDocumentationVersionInput) (*apigateway.DeleteDocumentationVersionOutput, error)
	DeleteDocumentationVersionAsync(ctx workflow.Context, input *apigateway.DeleteDocumentationVersionInput) *APIGatewayDeleteDocumentationVersionFuture

	DeleteDomainName(ctx workflow.Context, input *apigateway.DeleteDomainNameInput) (*apigateway.DeleteDomainNameOutput, error)
	DeleteDomainNameAsync(ctx workflow.Context, input *apigateway.DeleteDomainNameInput) *APIGatewayDeleteDomainNameFuture

	DeleteGatewayResponse(ctx workflow.Context, input *apigateway.DeleteGatewayResponseInput) (*apigateway.DeleteGatewayResponseOutput, error)
	DeleteGatewayResponseAsync(ctx workflow.Context, input *apigateway.DeleteGatewayResponseInput) *APIGatewayDeleteGatewayResponseFuture

	DeleteIntegration(ctx workflow.Context, input *apigateway.DeleteIntegrationInput) (*apigateway.DeleteIntegrationOutput, error)
	DeleteIntegrationAsync(ctx workflow.Context, input *apigateway.DeleteIntegrationInput) *APIGatewayDeleteIntegrationFuture

	DeleteIntegrationResponse(ctx workflow.Context, input *apigateway.DeleteIntegrationResponseInput) (*apigateway.DeleteIntegrationResponseOutput, error)
	DeleteIntegrationResponseAsync(ctx workflow.Context, input *apigateway.DeleteIntegrationResponseInput) *APIGatewayDeleteIntegrationResponseFuture

	DeleteMethod(ctx workflow.Context, input *apigateway.DeleteMethodInput) (*apigateway.DeleteMethodOutput, error)
	DeleteMethodAsync(ctx workflow.Context, input *apigateway.DeleteMethodInput) *APIGatewayDeleteMethodFuture

	DeleteMethodResponse(ctx workflow.Context, input *apigateway.DeleteMethodResponseInput) (*apigateway.DeleteMethodResponseOutput, error)
	DeleteMethodResponseAsync(ctx workflow.Context, input *apigateway.DeleteMethodResponseInput) *APIGatewayDeleteMethodResponseFuture

	DeleteModel(ctx workflow.Context, input *apigateway.DeleteModelInput) (*apigateway.DeleteModelOutput, error)
	DeleteModelAsync(ctx workflow.Context, input *apigateway.DeleteModelInput) *APIGatewayDeleteModelFuture

	DeleteRequestValidator(ctx workflow.Context, input *apigateway.DeleteRequestValidatorInput) (*apigateway.DeleteRequestValidatorOutput, error)
	DeleteRequestValidatorAsync(ctx workflow.Context, input *apigateway.DeleteRequestValidatorInput) *APIGatewayDeleteRequestValidatorFuture

	DeleteResource(ctx workflow.Context, input *apigateway.DeleteResourceInput) (*apigateway.DeleteResourceOutput, error)
	DeleteResourceAsync(ctx workflow.Context, input *apigateway.DeleteResourceInput) *APIGatewayDeleteResourceFuture

	DeleteRestApi(ctx workflow.Context, input *apigateway.DeleteRestApiInput) (*apigateway.DeleteRestApiOutput, error)
	DeleteRestApiAsync(ctx workflow.Context, input *apigateway.DeleteRestApiInput) *APIGatewayDeleteRestApiFuture

	DeleteStage(ctx workflow.Context, input *apigateway.DeleteStageInput) (*apigateway.DeleteStageOutput, error)
	DeleteStageAsync(ctx workflow.Context, input *apigateway.DeleteStageInput) *APIGatewayDeleteStageFuture

	DeleteUsagePlan(ctx workflow.Context, input *apigateway.DeleteUsagePlanInput) (*apigateway.DeleteUsagePlanOutput, error)
	DeleteUsagePlanAsync(ctx workflow.Context, input *apigateway.DeleteUsagePlanInput) *APIGatewayDeleteUsagePlanFuture

	DeleteUsagePlanKey(ctx workflow.Context, input *apigateway.DeleteUsagePlanKeyInput) (*apigateway.DeleteUsagePlanKeyOutput, error)
	DeleteUsagePlanKeyAsync(ctx workflow.Context, input *apigateway.DeleteUsagePlanKeyInput) *APIGatewayDeleteUsagePlanKeyFuture

	DeleteVpcLink(ctx workflow.Context, input *apigateway.DeleteVpcLinkInput) (*apigateway.DeleteVpcLinkOutput, error)
	DeleteVpcLinkAsync(ctx workflow.Context, input *apigateway.DeleteVpcLinkInput) *APIGatewayDeleteVpcLinkFuture

	FlushStageAuthorizersCache(ctx workflow.Context, input *apigateway.FlushStageAuthorizersCacheInput) (*apigateway.FlushStageAuthorizersCacheOutput, error)
	FlushStageAuthorizersCacheAsync(ctx workflow.Context, input *apigateway.FlushStageAuthorizersCacheInput) *APIGatewayFlushStageAuthorizersCacheFuture

	FlushStageCache(ctx workflow.Context, input *apigateway.FlushStageCacheInput) (*apigateway.FlushStageCacheOutput, error)
	FlushStageCacheAsync(ctx workflow.Context, input *apigateway.FlushStageCacheInput) *APIGatewayFlushStageCacheFuture

	GenerateClientCertificate(ctx workflow.Context, input *apigateway.GenerateClientCertificateInput) (*apigateway.ClientCertificate, error)
	GenerateClientCertificateAsync(ctx workflow.Context, input *apigateway.GenerateClientCertificateInput) *APIGatewayGenerateClientCertificateFuture

	GetAccount(ctx workflow.Context, input *apigateway.GetAccountInput) (*apigateway.Account, error)
	GetAccountAsync(ctx workflow.Context, input *apigateway.GetAccountInput) *APIGatewayGetAccountFuture

	GetApiKey(ctx workflow.Context, input *apigateway.GetApiKeyInput) (*apigateway.ApiKey, error)
	GetApiKeyAsync(ctx workflow.Context, input *apigateway.GetApiKeyInput) *APIGatewayGetApiKeyFuture

	GetApiKeys(ctx workflow.Context, input *apigateway.GetApiKeysInput) (*apigateway.GetApiKeysOutput, error)
	GetApiKeysAsync(ctx workflow.Context, input *apigateway.GetApiKeysInput) *APIGatewayGetApiKeysFuture

	GetAuthorizer(ctx workflow.Context, input *apigateway.GetAuthorizerInput) (*apigateway.Authorizer, error)
	GetAuthorizerAsync(ctx workflow.Context, input *apigateway.GetAuthorizerInput) *APIGatewayGetAuthorizerFuture

	GetAuthorizers(ctx workflow.Context, input *apigateway.GetAuthorizersInput) (*apigateway.GetAuthorizersOutput, error)
	GetAuthorizersAsync(ctx workflow.Context, input *apigateway.GetAuthorizersInput) *APIGatewayGetAuthorizersFuture

	GetBasePathMapping(ctx workflow.Context, input *apigateway.GetBasePathMappingInput) (*apigateway.BasePathMapping, error)
	GetBasePathMappingAsync(ctx workflow.Context, input *apigateway.GetBasePathMappingInput) *APIGatewayGetBasePathMappingFuture

	GetBasePathMappings(ctx workflow.Context, input *apigateway.GetBasePathMappingsInput) (*apigateway.GetBasePathMappingsOutput, error)
	GetBasePathMappingsAsync(ctx workflow.Context, input *apigateway.GetBasePathMappingsInput) *APIGatewayGetBasePathMappingsFuture

	GetClientCertificate(ctx workflow.Context, input *apigateway.GetClientCertificateInput) (*apigateway.ClientCertificate, error)
	GetClientCertificateAsync(ctx workflow.Context, input *apigateway.GetClientCertificateInput) *APIGatewayGetClientCertificateFuture

	GetClientCertificates(ctx workflow.Context, input *apigateway.GetClientCertificatesInput) (*apigateway.GetClientCertificatesOutput, error)
	GetClientCertificatesAsync(ctx workflow.Context, input *apigateway.GetClientCertificatesInput) *APIGatewayGetClientCertificatesFuture

	GetDeployment(ctx workflow.Context, input *apigateway.GetDeploymentInput) (*apigateway.Deployment, error)
	GetDeploymentAsync(ctx workflow.Context, input *apigateway.GetDeploymentInput) *APIGatewayGetDeploymentFuture

	GetDeployments(ctx workflow.Context, input *apigateway.GetDeploymentsInput) (*apigateway.GetDeploymentsOutput, error)
	GetDeploymentsAsync(ctx workflow.Context, input *apigateway.GetDeploymentsInput) *APIGatewayGetDeploymentsFuture

	GetDocumentationPart(ctx workflow.Context, input *apigateway.GetDocumentationPartInput) (*apigateway.DocumentationPart, error)
	GetDocumentationPartAsync(ctx workflow.Context, input *apigateway.GetDocumentationPartInput) *APIGatewayGetDocumentationPartFuture

	GetDocumentationParts(ctx workflow.Context, input *apigateway.GetDocumentationPartsInput) (*apigateway.GetDocumentationPartsOutput, error)
	GetDocumentationPartsAsync(ctx workflow.Context, input *apigateway.GetDocumentationPartsInput) *APIGatewayGetDocumentationPartsFuture

	GetDocumentationVersion(ctx workflow.Context, input *apigateway.GetDocumentationVersionInput) (*apigateway.DocumentationVersion, error)
	GetDocumentationVersionAsync(ctx workflow.Context, input *apigateway.GetDocumentationVersionInput) *APIGatewayGetDocumentationVersionFuture

	GetDocumentationVersions(ctx workflow.Context, input *apigateway.GetDocumentationVersionsInput) (*apigateway.GetDocumentationVersionsOutput, error)
	GetDocumentationVersionsAsync(ctx workflow.Context, input *apigateway.GetDocumentationVersionsInput) *APIGatewayGetDocumentationVersionsFuture

	GetDomainName(ctx workflow.Context, input *apigateway.GetDomainNameInput) (*apigateway.DomainName, error)
	GetDomainNameAsync(ctx workflow.Context, input *apigateway.GetDomainNameInput) *APIGatewayGetDomainNameFuture

	GetDomainNames(ctx workflow.Context, input *apigateway.GetDomainNamesInput) (*apigateway.GetDomainNamesOutput, error)
	GetDomainNamesAsync(ctx workflow.Context, input *apigateway.GetDomainNamesInput) *APIGatewayGetDomainNamesFuture

	GetExport(ctx workflow.Context, input *apigateway.GetExportInput) (*apigateway.GetExportOutput, error)
	GetExportAsync(ctx workflow.Context, input *apigateway.GetExportInput) *APIGatewayGetExportFuture

	GetGatewayResponse(ctx workflow.Context, input *apigateway.GetGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error)
	GetGatewayResponseAsync(ctx workflow.Context, input *apigateway.GetGatewayResponseInput) *APIGatewayGetGatewayResponseFuture

	GetGatewayResponses(ctx workflow.Context, input *apigateway.GetGatewayResponsesInput) (*apigateway.GetGatewayResponsesOutput, error)
	GetGatewayResponsesAsync(ctx workflow.Context, input *apigateway.GetGatewayResponsesInput) *APIGatewayGetGatewayResponsesFuture

	GetIntegration(ctx workflow.Context, input *apigateway.GetIntegrationInput) (*apigateway.Integration, error)
	GetIntegrationAsync(ctx workflow.Context, input *apigateway.GetIntegrationInput) *APIGatewayGetIntegrationFuture

	GetIntegrationResponse(ctx workflow.Context, input *apigateway.GetIntegrationResponseInput) (*apigateway.IntegrationResponse, error)
	GetIntegrationResponseAsync(ctx workflow.Context, input *apigateway.GetIntegrationResponseInput) *APIGatewayGetIntegrationResponseFuture

	GetMethod(ctx workflow.Context, input *apigateway.GetMethodInput) (*apigateway.Method, error)
	GetMethodAsync(ctx workflow.Context, input *apigateway.GetMethodInput) *APIGatewayGetMethodFuture

	GetMethodResponse(ctx workflow.Context, input *apigateway.GetMethodResponseInput) (*apigateway.MethodResponse, error)
	GetMethodResponseAsync(ctx workflow.Context, input *apigateway.GetMethodResponseInput) *APIGatewayGetMethodResponseFuture

	GetModel(ctx workflow.Context, input *apigateway.GetModelInput) (*apigateway.Model, error)
	GetModelAsync(ctx workflow.Context, input *apigateway.GetModelInput) *APIGatewayGetModelFuture

	GetModelTemplate(ctx workflow.Context, input *apigateway.GetModelTemplateInput) (*apigateway.GetModelTemplateOutput, error)
	GetModelTemplateAsync(ctx workflow.Context, input *apigateway.GetModelTemplateInput) *APIGatewayGetModelTemplateFuture

	GetModels(ctx workflow.Context, input *apigateway.GetModelsInput) (*apigateway.GetModelsOutput, error)
	GetModelsAsync(ctx workflow.Context, input *apigateway.GetModelsInput) *APIGatewayGetModelsFuture

	GetRequestValidator(ctx workflow.Context, input *apigateway.GetRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error)
	GetRequestValidatorAsync(ctx workflow.Context, input *apigateway.GetRequestValidatorInput) *APIGatewayGetRequestValidatorFuture

	GetRequestValidators(ctx workflow.Context, input *apigateway.GetRequestValidatorsInput) (*apigateway.GetRequestValidatorsOutput, error)
	GetRequestValidatorsAsync(ctx workflow.Context, input *apigateway.GetRequestValidatorsInput) *APIGatewayGetRequestValidatorsFuture

	GetResource(ctx workflow.Context, input *apigateway.GetResourceInput) (*apigateway.Resource, error)
	GetResourceAsync(ctx workflow.Context, input *apigateway.GetResourceInput) *APIGatewayGetResourceFuture

	GetResources(ctx workflow.Context, input *apigateway.GetResourcesInput) (*apigateway.GetResourcesOutput, error)
	GetResourcesAsync(ctx workflow.Context, input *apigateway.GetResourcesInput) *APIGatewayGetResourcesFuture

	GetRestApi(ctx workflow.Context, input *apigateway.GetRestApiInput) (*apigateway.RestApi, error)
	GetRestApiAsync(ctx workflow.Context, input *apigateway.GetRestApiInput) *APIGatewayGetRestApiFuture

	GetRestApis(ctx workflow.Context, input *apigateway.GetRestApisInput) (*apigateway.GetRestApisOutput, error)
	GetRestApisAsync(ctx workflow.Context, input *apigateway.GetRestApisInput) *APIGatewayGetRestApisFuture

	GetSdk(ctx workflow.Context, input *apigateway.GetSdkInput) (*apigateway.GetSdkOutput, error)
	GetSdkAsync(ctx workflow.Context, input *apigateway.GetSdkInput) *APIGatewayGetSdkFuture

	GetSdkType(ctx workflow.Context, input *apigateway.GetSdkTypeInput) (*apigateway.SdkType, error)
	GetSdkTypeAsync(ctx workflow.Context, input *apigateway.GetSdkTypeInput) *APIGatewayGetSdkTypeFuture

	GetSdkTypes(ctx workflow.Context, input *apigateway.GetSdkTypesInput) (*apigateway.GetSdkTypesOutput, error)
	GetSdkTypesAsync(ctx workflow.Context, input *apigateway.GetSdkTypesInput) *APIGatewayGetSdkTypesFuture

	GetStage(ctx workflow.Context, input *apigateway.GetStageInput) (*apigateway.Stage, error)
	GetStageAsync(ctx workflow.Context, input *apigateway.GetStageInput) *APIGatewayGetStageFuture

	GetStages(ctx workflow.Context, input *apigateway.GetStagesInput) (*apigateway.GetStagesOutput, error)
	GetStagesAsync(ctx workflow.Context, input *apigateway.GetStagesInput) *APIGatewayGetStagesFuture

	GetTags(ctx workflow.Context, input *apigateway.GetTagsInput) (*apigateway.GetTagsOutput, error)
	GetTagsAsync(ctx workflow.Context, input *apigateway.GetTagsInput) *APIGatewayGetTagsFuture

	GetUsage(ctx workflow.Context, input *apigateway.GetUsageInput) (*apigateway.Usage, error)
	GetUsageAsync(ctx workflow.Context, input *apigateway.GetUsageInput) *APIGatewayGetUsageFuture

	GetUsagePlan(ctx workflow.Context, input *apigateway.GetUsagePlanInput) (*apigateway.UsagePlan, error)
	GetUsagePlanAsync(ctx workflow.Context, input *apigateway.GetUsagePlanInput) *APIGatewayGetUsagePlanFuture

	GetUsagePlanKey(ctx workflow.Context, input *apigateway.GetUsagePlanKeyInput) (*apigateway.UsagePlanKey, error)
	GetUsagePlanKeyAsync(ctx workflow.Context, input *apigateway.GetUsagePlanKeyInput) *APIGatewayGetUsagePlanKeyFuture

	GetUsagePlanKeys(ctx workflow.Context, input *apigateway.GetUsagePlanKeysInput) (*apigateway.GetUsagePlanKeysOutput, error)
	GetUsagePlanKeysAsync(ctx workflow.Context, input *apigateway.GetUsagePlanKeysInput) *APIGatewayGetUsagePlanKeysFuture

	GetUsagePlans(ctx workflow.Context, input *apigateway.GetUsagePlansInput) (*apigateway.GetUsagePlansOutput, error)
	GetUsagePlansAsync(ctx workflow.Context, input *apigateway.GetUsagePlansInput) *APIGatewayGetUsagePlansFuture

	GetVpcLink(ctx workflow.Context, input *apigateway.GetVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error)
	GetVpcLinkAsync(ctx workflow.Context, input *apigateway.GetVpcLinkInput) *APIGatewayGetVpcLinkFuture

	GetVpcLinks(ctx workflow.Context, input *apigateway.GetVpcLinksInput) (*apigateway.GetVpcLinksOutput, error)
	GetVpcLinksAsync(ctx workflow.Context, input *apigateway.GetVpcLinksInput) *APIGatewayGetVpcLinksFuture

	ImportApiKeys(ctx workflow.Context, input *apigateway.ImportApiKeysInput) (*apigateway.ImportApiKeysOutput, error)
	ImportApiKeysAsync(ctx workflow.Context, input *apigateway.ImportApiKeysInput) *APIGatewayImportApiKeysFuture

	ImportDocumentationParts(ctx workflow.Context, input *apigateway.ImportDocumentationPartsInput) (*apigateway.ImportDocumentationPartsOutput, error)
	ImportDocumentationPartsAsync(ctx workflow.Context, input *apigateway.ImportDocumentationPartsInput) *APIGatewayImportDocumentationPartsFuture

	ImportRestApi(ctx workflow.Context, input *apigateway.ImportRestApiInput) (*apigateway.RestApi, error)
	ImportRestApiAsync(ctx workflow.Context, input *apigateway.ImportRestApiInput) *APIGatewayImportRestApiFuture

	PutGatewayResponse(ctx workflow.Context, input *apigateway.PutGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error)
	PutGatewayResponseAsync(ctx workflow.Context, input *apigateway.PutGatewayResponseInput) *APIGatewayPutGatewayResponseFuture

	PutIntegration(ctx workflow.Context, input *apigateway.PutIntegrationInput) (*apigateway.Integration, error)
	PutIntegrationAsync(ctx workflow.Context, input *apigateway.PutIntegrationInput) *APIGatewayPutIntegrationFuture

	PutIntegrationResponse(ctx workflow.Context, input *apigateway.PutIntegrationResponseInput) (*apigateway.IntegrationResponse, error)
	PutIntegrationResponseAsync(ctx workflow.Context, input *apigateway.PutIntegrationResponseInput) *APIGatewayPutIntegrationResponseFuture

	PutMethod(ctx workflow.Context, input *apigateway.PutMethodInput) (*apigateway.Method, error)
	PutMethodAsync(ctx workflow.Context, input *apigateway.PutMethodInput) *APIGatewayPutMethodFuture

	PutMethodResponse(ctx workflow.Context, input *apigateway.PutMethodResponseInput) (*apigateway.MethodResponse, error)
	PutMethodResponseAsync(ctx workflow.Context, input *apigateway.PutMethodResponseInput) *APIGatewayPutMethodResponseFuture

	PutRestApi(ctx workflow.Context, input *apigateway.PutRestApiInput) (*apigateway.RestApi, error)
	PutRestApiAsync(ctx workflow.Context, input *apigateway.PutRestApiInput) *APIGatewayPutRestApiFuture

	TagResource(ctx workflow.Context, input *apigateway.TagResourceInput) (*apigateway.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *apigateway.TagResourceInput) *APIGatewayTagResourceFuture

	TestInvokeAuthorizer(ctx workflow.Context, input *apigateway.TestInvokeAuthorizerInput) (*apigateway.TestInvokeAuthorizerOutput, error)
	TestInvokeAuthorizerAsync(ctx workflow.Context, input *apigateway.TestInvokeAuthorizerInput) *APIGatewayTestInvokeAuthorizerFuture

	TestInvokeMethod(ctx workflow.Context, input *apigateway.TestInvokeMethodInput) (*apigateway.TestInvokeMethodOutput, error)
	TestInvokeMethodAsync(ctx workflow.Context, input *apigateway.TestInvokeMethodInput) *APIGatewayTestInvokeMethodFuture

	UntagResource(ctx workflow.Context, input *apigateway.UntagResourceInput) (*apigateway.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *apigateway.UntagResourceInput) *APIGatewayUntagResourceFuture

	UpdateAccount(ctx workflow.Context, input *apigateway.UpdateAccountInput) (*apigateway.Account, error)
	UpdateAccountAsync(ctx workflow.Context, input *apigateway.UpdateAccountInput) *APIGatewayUpdateAccountFuture

	UpdateApiKey(ctx workflow.Context, input *apigateway.UpdateApiKeyInput) (*apigateway.ApiKey, error)
	UpdateApiKeyAsync(ctx workflow.Context, input *apigateway.UpdateApiKeyInput) *APIGatewayUpdateApiKeyFuture

	UpdateAuthorizer(ctx workflow.Context, input *apigateway.UpdateAuthorizerInput) (*apigateway.Authorizer, error)
	UpdateAuthorizerAsync(ctx workflow.Context, input *apigateway.UpdateAuthorizerInput) *APIGatewayUpdateAuthorizerFuture

	UpdateBasePathMapping(ctx workflow.Context, input *apigateway.UpdateBasePathMappingInput) (*apigateway.BasePathMapping, error)
	UpdateBasePathMappingAsync(ctx workflow.Context, input *apigateway.UpdateBasePathMappingInput) *APIGatewayUpdateBasePathMappingFuture

	UpdateClientCertificate(ctx workflow.Context, input *apigateway.UpdateClientCertificateInput) (*apigateway.ClientCertificate, error)
	UpdateClientCertificateAsync(ctx workflow.Context, input *apigateway.UpdateClientCertificateInput) *APIGatewayUpdateClientCertificateFuture

	UpdateDeployment(ctx workflow.Context, input *apigateway.UpdateDeploymentInput) (*apigateway.Deployment, error)
	UpdateDeploymentAsync(ctx workflow.Context, input *apigateway.UpdateDeploymentInput) *APIGatewayUpdateDeploymentFuture

	UpdateDocumentationPart(ctx workflow.Context, input *apigateway.UpdateDocumentationPartInput) (*apigateway.DocumentationPart, error)
	UpdateDocumentationPartAsync(ctx workflow.Context, input *apigateway.UpdateDocumentationPartInput) *APIGatewayUpdateDocumentationPartFuture

	UpdateDocumentationVersion(ctx workflow.Context, input *apigateway.UpdateDocumentationVersionInput) (*apigateway.DocumentationVersion, error)
	UpdateDocumentationVersionAsync(ctx workflow.Context, input *apigateway.UpdateDocumentationVersionInput) *APIGatewayUpdateDocumentationVersionFuture

	UpdateDomainName(ctx workflow.Context, input *apigateway.UpdateDomainNameInput) (*apigateway.DomainName, error)
	UpdateDomainNameAsync(ctx workflow.Context, input *apigateway.UpdateDomainNameInput) *APIGatewayUpdateDomainNameFuture

	UpdateGatewayResponse(ctx workflow.Context, input *apigateway.UpdateGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error)
	UpdateGatewayResponseAsync(ctx workflow.Context, input *apigateway.UpdateGatewayResponseInput) *APIGatewayUpdateGatewayResponseFuture

	UpdateIntegration(ctx workflow.Context, input *apigateway.UpdateIntegrationInput) (*apigateway.Integration, error)
	UpdateIntegrationAsync(ctx workflow.Context, input *apigateway.UpdateIntegrationInput) *APIGatewayUpdateIntegrationFuture

	UpdateIntegrationResponse(ctx workflow.Context, input *apigateway.UpdateIntegrationResponseInput) (*apigateway.IntegrationResponse, error)
	UpdateIntegrationResponseAsync(ctx workflow.Context, input *apigateway.UpdateIntegrationResponseInput) *APIGatewayUpdateIntegrationResponseFuture

	UpdateMethod(ctx workflow.Context, input *apigateway.UpdateMethodInput) (*apigateway.Method, error)
	UpdateMethodAsync(ctx workflow.Context, input *apigateway.UpdateMethodInput) *APIGatewayUpdateMethodFuture

	UpdateMethodResponse(ctx workflow.Context, input *apigateway.UpdateMethodResponseInput) (*apigateway.MethodResponse, error)
	UpdateMethodResponseAsync(ctx workflow.Context, input *apigateway.UpdateMethodResponseInput) *APIGatewayUpdateMethodResponseFuture

	UpdateModel(ctx workflow.Context, input *apigateway.UpdateModelInput) (*apigateway.Model, error)
	UpdateModelAsync(ctx workflow.Context, input *apigateway.UpdateModelInput) *APIGatewayUpdateModelFuture

	UpdateRequestValidator(ctx workflow.Context, input *apigateway.UpdateRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error)
	UpdateRequestValidatorAsync(ctx workflow.Context, input *apigateway.UpdateRequestValidatorInput) *APIGatewayUpdateRequestValidatorFuture

	UpdateResource(ctx workflow.Context, input *apigateway.UpdateResourceInput) (*apigateway.Resource, error)
	UpdateResourceAsync(ctx workflow.Context, input *apigateway.UpdateResourceInput) *APIGatewayUpdateResourceFuture

	UpdateRestApi(ctx workflow.Context, input *apigateway.UpdateRestApiInput) (*apigateway.RestApi, error)
	UpdateRestApiAsync(ctx workflow.Context, input *apigateway.UpdateRestApiInput) *APIGatewayUpdateRestApiFuture

	UpdateStage(ctx workflow.Context, input *apigateway.UpdateStageInput) (*apigateway.Stage, error)
	UpdateStageAsync(ctx workflow.Context, input *apigateway.UpdateStageInput) *APIGatewayUpdateStageFuture

	UpdateUsage(ctx workflow.Context, input *apigateway.UpdateUsageInput) (*apigateway.Usage, error)
	UpdateUsageAsync(ctx workflow.Context, input *apigateway.UpdateUsageInput) *APIGatewayUpdateUsageFuture

	UpdateUsagePlan(ctx workflow.Context, input *apigateway.UpdateUsagePlanInput) (*apigateway.UsagePlan, error)
	UpdateUsagePlanAsync(ctx workflow.Context, input *apigateway.UpdateUsagePlanInput) *APIGatewayUpdateUsagePlanFuture

	UpdateVpcLink(ctx workflow.Context, input *apigateway.UpdateVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error)
	UpdateVpcLinkAsync(ctx workflow.Context, input *apigateway.UpdateVpcLinkInput) *APIGatewayUpdateVpcLinkFuture
}

func NewClient() Client {
	return &stub{}
}
