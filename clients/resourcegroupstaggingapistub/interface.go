// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package resourcegroupstaggingapistub

import (
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DescribeReportCreation(ctx workflow.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error)
	DescribeReportCreationAsync(ctx workflow.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput) *ResourceGroupsTaggingAPIDescribeReportCreationFuture

	GetComplianceSummary(ctx workflow.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error)
	GetComplianceSummaryAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput) *ResourceGroupsTaggingAPIGetComplianceSummaryFuture

	GetResources(ctx workflow.Context, input *resourcegroupstaggingapi.GetResourcesInput) (*resourcegroupstaggingapi.GetResourcesOutput, error)
	GetResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetResourcesInput) *ResourceGroupsTaggingAPIGetResourcesFuture

	GetTagKeys(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagKeysInput) (*resourcegroupstaggingapi.GetTagKeysOutput, error)
	GetTagKeysAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagKeysInput) *ResourceGroupsTaggingAPIGetTagKeysFuture

	GetTagValues(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagValuesInput) (*resourcegroupstaggingapi.GetTagValuesOutput, error)
	GetTagValuesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagValuesInput) *ResourceGroupsTaggingAPIGetTagValuesFuture

	StartReportCreation(ctx workflow.Context, input *resourcegroupstaggingapi.StartReportCreationInput) (*resourcegroupstaggingapi.StartReportCreationOutput, error)
	StartReportCreationAsync(ctx workflow.Context, input *resourcegroupstaggingapi.StartReportCreationInput) *ResourceGroupsTaggingAPIStartReportCreationFuture

	TagResources(ctx workflow.Context, input *resourcegroupstaggingapi.TagResourcesInput) (*resourcegroupstaggingapi.TagResourcesOutput, error)
	TagResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.TagResourcesInput) *ResourceGroupsTaggingAPITagResourcesFuture

	UntagResources(ctx workflow.Context, input *resourcegroupstaggingapi.UntagResourcesInput) (*resourcegroupstaggingapi.UntagResourcesOutput, error)
	UntagResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.UntagResourcesInput) *ResourceGroupsTaggingAPIUntagResourcesFuture
}

func NewClient() Client {
	return &stub{}
}
