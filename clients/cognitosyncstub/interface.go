// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package cognitosyncstub

import (
	"github.com/aws/aws-sdk-go/service/cognitosync"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	BulkPublish(ctx workflow.Context, input *cognitosync.BulkPublishInput) (*cognitosync.BulkPublishOutput, error)
	BulkPublishAsync(ctx workflow.Context, input *cognitosync.BulkPublishInput) *CognitoSyncBulkPublishFuture

	DeleteDataset(ctx workflow.Context, input *cognitosync.DeleteDatasetInput) (*cognitosync.DeleteDatasetOutput, error)
	DeleteDatasetAsync(ctx workflow.Context, input *cognitosync.DeleteDatasetInput) *CognitoSyncDeleteDatasetFuture

	DescribeDataset(ctx workflow.Context, input *cognitosync.DescribeDatasetInput) (*cognitosync.DescribeDatasetOutput, error)
	DescribeDatasetAsync(ctx workflow.Context, input *cognitosync.DescribeDatasetInput) *CognitoSyncDescribeDatasetFuture

	DescribeIdentityPoolUsage(ctx workflow.Context, input *cognitosync.DescribeIdentityPoolUsageInput) (*cognitosync.DescribeIdentityPoolUsageOutput, error)
	DescribeIdentityPoolUsageAsync(ctx workflow.Context, input *cognitosync.DescribeIdentityPoolUsageInput) *CognitoSyncDescribeIdentityPoolUsageFuture

	DescribeIdentityUsage(ctx workflow.Context, input *cognitosync.DescribeIdentityUsageInput) (*cognitosync.DescribeIdentityUsageOutput, error)
	DescribeIdentityUsageAsync(ctx workflow.Context, input *cognitosync.DescribeIdentityUsageInput) *CognitoSyncDescribeIdentityUsageFuture

	GetBulkPublishDetails(ctx workflow.Context, input *cognitosync.GetBulkPublishDetailsInput) (*cognitosync.GetBulkPublishDetailsOutput, error)
	GetBulkPublishDetailsAsync(ctx workflow.Context, input *cognitosync.GetBulkPublishDetailsInput) *CognitoSyncGetBulkPublishDetailsFuture

	GetCognitoEvents(ctx workflow.Context, input *cognitosync.GetCognitoEventsInput) (*cognitosync.GetCognitoEventsOutput, error)
	GetCognitoEventsAsync(ctx workflow.Context, input *cognitosync.GetCognitoEventsInput) *CognitoSyncGetCognitoEventsFuture

	GetIdentityPoolConfiguration(ctx workflow.Context, input *cognitosync.GetIdentityPoolConfigurationInput) (*cognitosync.GetIdentityPoolConfigurationOutput, error)
	GetIdentityPoolConfigurationAsync(ctx workflow.Context, input *cognitosync.GetIdentityPoolConfigurationInput) *CognitoSyncGetIdentityPoolConfigurationFuture

	ListDatasets(ctx workflow.Context, input *cognitosync.ListDatasetsInput) (*cognitosync.ListDatasetsOutput, error)
	ListDatasetsAsync(ctx workflow.Context, input *cognitosync.ListDatasetsInput) *CognitoSyncListDatasetsFuture

	ListIdentityPoolUsage(ctx workflow.Context, input *cognitosync.ListIdentityPoolUsageInput) (*cognitosync.ListIdentityPoolUsageOutput, error)
	ListIdentityPoolUsageAsync(ctx workflow.Context, input *cognitosync.ListIdentityPoolUsageInput) *CognitoSyncListIdentityPoolUsageFuture

	ListRecords(ctx workflow.Context, input *cognitosync.ListRecordsInput) (*cognitosync.ListRecordsOutput, error)
	ListRecordsAsync(ctx workflow.Context, input *cognitosync.ListRecordsInput) *CognitoSyncListRecordsFuture

	RegisterDevice(ctx workflow.Context, input *cognitosync.RegisterDeviceInput) (*cognitosync.RegisterDeviceOutput, error)
	RegisterDeviceAsync(ctx workflow.Context, input *cognitosync.RegisterDeviceInput) *CognitoSyncRegisterDeviceFuture

	SetCognitoEvents(ctx workflow.Context, input *cognitosync.SetCognitoEventsInput) (*cognitosync.SetCognitoEventsOutput, error)
	SetCognitoEventsAsync(ctx workflow.Context, input *cognitosync.SetCognitoEventsInput) *CognitoSyncSetCognitoEventsFuture

	SetIdentityPoolConfiguration(ctx workflow.Context, input *cognitosync.SetIdentityPoolConfigurationInput) (*cognitosync.SetIdentityPoolConfigurationOutput, error)
	SetIdentityPoolConfigurationAsync(ctx workflow.Context, input *cognitosync.SetIdentityPoolConfigurationInput) *CognitoSyncSetIdentityPoolConfigurationFuture

	SubscribeToDataset(ctx workflow.Context, input *cognitosync.SubscribeToDatasetInput) (*cognitosync.SubscribeToDatasetOutput, error)
	SubscribeToDatasetAsync(ctx workflow.Context, input *cognitosync.SubscribeToDatasetInput) *CognitoSyncSubscribeToDatasetFuture

	UnsubscribeFromDataset(ctx workflow.Context, input *cognitosync.UnsubscribeFromDatasetInput) (*cognitosync.UnsubscribeFromDatasetOutput, error)
	UnsubscribeFromDatasetAsync(ctx workflow.Context, input *cognitosync.UnsubscribeFromDatasetInput) *CognitoSyncUnsubscribeFromDatasetFuture

	UpdateRecords(ctx workflow.Context, input *cognitosync.UpdateRecordsInput) (*cognitosync.UpdateRecordsOutput, error)
	UpdateRecordsAsync(ctx workflow.Context, input *cognitosync.UpdateRecordsInput) *CognitoSyncUpdateRecordsFuture
}

func NewClient() Client {
	return &stub{}
}
