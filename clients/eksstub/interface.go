// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package eksstub

import (
	"github.com/aws/aws-sdk-go/service/eks"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateCluster(ctx workflow.Context, input *eks.CreateClusterInput) (*eks.CreateClusterOutput, error)
	CreateClusterAsync(ctx workflow.Context, input *eks.CreateClusterInput) *EKSCreateClusterFuture

	CreateFargateProfile(ctx workflow.Context, input *eks.CreateFargateProfileInput) (*eks.CreateFargateProfileOutput, error)
	CreateFargateProfileAsync(ctx workflow.Context, input *eks.CreateFargateProfileInput) *EKSCreateFargateProfileFuture

	CreateNodegroup(ctx workflow.Context, input *eks.CreateNodegroupInput) (*eks.CreateNodegroupOutput, error)
	CreateNodegroupAsync(ctx workflow.Context, input *eks.CreateNodegroupInput) *EKSCreateNodegroupFuture

	DeleteCluster(ctx workflow.Context, input *eks.DeleteClusterInput) (*eks.DeleteClusterOutput, error)
	DeleteClusterAsync(ctx workflow.Context, input *eks.DeleteClusterInput) *EKSDeleteClusterFuture

	DeleteFargateProfile(ctx workflow.Context, input *eks.DeleteFargateProfileInput) (*eks.DeleteFargateProfileOutput, error)
	DeleteFargateProfileAsync(ctx workflow.Context, input *eks.DeleteFargateProfileInput) *EKSDeleteFargateProfileFuture

	DeleteNodegroup(ctx workflow.Context, input *eks.DeleteNodegroupInput) (*eks.DeleteNodegroupOutput, error)
	DeleteNodegroupAsync(ctx workflow.Context, input *eks.DeleteNodegroupInput) *EKSDeleteNodegroupFuture

	DescribeCluster(ctx workflow.Context, input *eks.DescribeClusterInput) (*eks.DescribeClusterOutput, error)
	DescribeClusterAsync(ctx workflow.Context, input *eks.DescribeClusterInput) *EKSDescribeClusterFuture

	DescribeFargateProfile(ctx workflow.Context, input *eks.DescribeFargateProfileInput) (*eks.DescribeFargateProfileOutput, error)
	DescribeFargateProfileAsync(ctx workflow.Context, input *eks.DescribeFargateProfileInput) *EKSDescribeFargateProfileFuture

	DescribeNodegroup(ctx workflow.Context, input *eks.DescribeNodegroupInput) (*eks.DescribeNodegroupOutput, error)
	DescribeNodegroupAsync(ctx workflow.Context, input *eks.DescribeNodegroupInput) *EKSDescribeNodegroupFuture

	DescribeUpdate(ctx workflow.Context, input *eks.DescribeUpdateInput) (*eks.DescribeUpdateOutput, error)
	DescribeUpdateAsync(ctx workflow.Context, input *eks.DescribeUpdateInput) *EKSDescribeUpdateFuture

	ListClusters(ctx workflow.Context, input *eks.ListClustersInput) (*eks.ListClustersOutput, error)
	ListClustersAsync(ctx workflow.Context, input *eks.ListClustersInput) *EKSListClustersFuture

	ListFargateProfiles(ctx workflow.Context, input *eks.ListFargateProfilesInput) (*eks.ListFargateProfilesOutput, error)
	ListFargateProfilesAsync(ctx workflow.Context, input *eks.ListFargateProfilesInput) *EKSListFargateProfilesFuture

	ListNodegroups(ctx workflow.Context, input *eks.ListNodegroupsInput) (*eks.ListNodegroupsOutput, error)
	ListNodegroupsAsync(ctx workflow.Context, input *eks.ListNodegroupsInput) *EKSListNodegroupsFuture

	ListTagsForResource(ctx workflow.Context, input *eks.ListTagsForResourceInput) (*eks.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *eks.ListTagsForResourceInput) *EKSListTagsForResourceFuture

	ListUpdates(ctx workflow.Context, input *eks.ListUpdatesInput) (*eks.ListUpdatesOutput, error)
	ListUpdatesAsync(ctx workflow.Context, input *eks.ListUpdatesInput) *EKSListUpdatesFuture

	TagResource(ctx workflow.Context, input *eks.TagResourceInput) (*eks.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *eks.TagResourceInput) *EKSTagResourceFuture

	UntagResource(ctx workflow.Context, input *eks.UntagResourceInput) (*eks.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *eks.UntagResourceInput) *EKSUntagResourceFuture

	UpdateClusterConfig(ctx workflow.Context, input *eks.UpdateClusterConfigInput) (*eks.UpdateClusterConfigOutput, error)
	UpdateClusterConfigAsync(ctx workflow.Context, input *eks.UpdateClusterConfigInput) *EKSUpdateClusterConfigFuture

	UpdateClusterVersion(ctx workflow.Context, input *eks.UpdateClusterVersionInput) (*eks.UpdateClusterVersionOutput, error)
	UpdateClusterVersionAsync(ctx workflow.Context, input *eks.UpdateClusterVersionInput) *EKSUpdateClusterVersionFuture

	UpdateNodegroupConfig(ctx workflow.Context, input *eks.UpdateNodegroupConfigInput) (*eks.UpdateNodegroupConfigOutput, error)
	UpdateNodegroupConfigAsync(ctx workflow.Context, input *eks.UpdateNodegroupConfigInput) *EKSUpdateNodegroupConfigFuture

	UpdateNodegroupVersion(ctx workflow.Context, input *eks.UpdateNodegroupVersionInput) (*eks.UpdateNodegroupVersionOutput, error)
	UpdateNodegroupVersionAsync(ctx workflow.Context, input *eks.UpdateNodegroupVersionInput) *EKSUpdateNodegroupVersionFuture

	WaitUntilClusterActive(ctx workflow.Context, input *eks.DescribeClusterInput) error
	WaitUntilClusterActiveAsync(ctx workflow.Context, input *eks.DescribeClusterInput) *clients.VoidFuture

	WaitUntilClusterDeleted(ctx workflow.Context, input *eks.DescribeClusterInput) error
	WaitUntilClusterDeletedAsync(ctx workflow.Context, input *eks.DescribeClusterInput) *clients.VoidFuture

	WaitUntilNodegroupActive(ctx workflow.Context, input *eks.DescribeNodegroupInput) error
	WaitUntilNodegroupActiveAsync(ctx workflow.Context, input *eks.DescribeNodegroupInput) *clients.VoidFuture

	WaitUntilNodegroupDeleted(ctx workflow.Context, input *eks.DescribeNodegroupInput) error
	WaitUntilNodegroupDeletedAsync(ctx workflow.Context, input *eks.DescribeNodegroupInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
