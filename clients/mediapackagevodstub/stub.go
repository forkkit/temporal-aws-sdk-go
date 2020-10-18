// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mediapackagevodstub

import (
	"github.com/aws/aws-sdk-go/service/mediapackagevod"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type MediaPackageVodCreateAssetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodCreateAssetFuture) Get(ctx workflow.Context) (*mediapackagevod.CreateAssetOutput, error) {
	var output mediapackagevod.CreateAssetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageVodCreatePackagingConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodCreatePackagingConfigurationFuture) Get(ctx workflow.Context) (*mediapackagevod.CreatePackagingConfigurationOutput, error) {
	var output mediapackagevod.CreatePackagingConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageVodCreatePackagingGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodCreatePackagingGroupFuture) Get(ctx workflow.Context) (*mediapackagevod.CreatePackagingGroupOutput, error) {
	var output mediapackagevod.CreatePackagingGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageVodDeleteAssetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodDeleteAssetFuture) Get(ctx workflow.Context) (*mediapackagevod.DeleteAssetOutput, error) {
	var output mediapackagevod.DeleteAssetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageVodDeletePackagingConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodDeletePackagingConfigurationFuture) Get(ctx workflow.Context) (*mediapackagevod.DeletePackagingConfigurationOutput, error) {
	var output mediapackagevod.DeletePackagingConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageVodDeletePackagingGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodDeletePackagingGroupFuture) Get(ctx workflow.Context) (*mediapackagevod.DeletePackagingGroupOutput, error) {
	var output mediapackagevod.DeletePackagingGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageVodDescribeAssetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodDescribeAssetFuture) Get(ctx workflow.Context) (*mediapackagevod.DescribeAssetOutput, error) {
	var output mediapackagevod.DescribeAssetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageVodDescribePackagingConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodDescribePackagingConfigurationFuture) Get(ctx workflow.Context) (*mediapackagevod.DescribePackagingConfigurationOutput, error) {
	var output mediapackagevod.DescribePackagingConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageVodDescribePackagingGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodDescribePackagingGroupFuture) Get(ctx workflow.Context) (*mediapackagevod.DescribePackagingGroupOutput, error) {
	var output mediapackagevod.DescribePackagingGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageVodListAssetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodListAssetsFuture) Get(ctx workflow.Context) (*mediapackagevod.ListAssetsOutput, error) {
	var output mediapackagevod.ListAssetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageVodListPackagingConfigurationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodListPackagingConfigurationsFuture) Get(ctx workflow.Context) (*mediapackagevod.ListPackagingConfigurationsOutput, error) {
	var output mediapackagevod.ListPackagingConfigurationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageVodListPackagingGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodListPackagingGroupsFuture) Get(ctx workflow.Context) (*mediapackagevod.ListPackagingGroupsOutput, error) {
	var output mediapackagevod.ListPackagingGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageVodListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodListTagsForResourceFuture) Get(ctx workflow.Context) (*mediapackagevod.ListTagsForResourceOutput, error) {
	var output mediapackagevod.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageVodTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodTagResourceFuture) Get(ctx workflow.Context) (*mediapackagevod.TagResourceOutput, error) {
	var output mediapackagevod.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageVodUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodUntagResourceFuture) Get(ctx workflow.Context) (*mediapackagevod.UntagResourceOutput, error) {
	var output mediapackagevod.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageVodUpdatePackagingGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaPackageVodUpdatePackagingGroupFuture) Get(ctx workflow.Context) (*mediapackagevod.UpdatePackagingGroupOutput, error) {
	var output mediapackagevod.UpdatePackagingGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAsset(ctx workflow.Context, input *mediapackagevod.CreateAssetInput) (*mediapackagevod.CreateAssetOutput, error) {
	var output mediapackagevod.CreateAssetOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.CreateAsset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAssetAsync(ctx workflow.Context, input *mediapackagevod.CreateAssetInput) *MediaPackageVodCreateAssetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.CreateAsset", input)
	return &MediaPackageVodCreateAssetFuture{Future: future}
}

func (a *stub) CreatePackagingConfiguration(ctx workflow.Context, input *mediapackagevod.CreatePackagingConfigurationInput) (*mediapackagevod.CreatePackagingConfigurationOutput, error) {
	var output mediapackagevod.CreatePackagingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.CreatePackagingConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreatePackagingConfigurationAsync(ctx workflow.Context, input *mediapackagevod.CreatePackagingConfigurationInput) *MediaPackageVodCreatePackagingConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.CreatePackagingConfiguration", input)
	return &MediaPackageVodCreatePackagingConfigurationFuture{Future: future}
}

func (a *stub) CreatePackagingGroup(ctx workflow.Context, input *mediapackagevod.CreatePackagingGroupInput) (*mediapackagevod.CreatePackagingGroupOutput, error) {
	var output mediapackagevod.CreatePackagingGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.CreatePackagingGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreatePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.CreatePackagingGroupInput) *MediaPackageVodCreatePackagingGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.CreatePackagingGroup", input)
	return &MediaPackageVodCreatePackagingGroupFuture{Future: future}
}

func (a *stub) DeleteAsset(ctx workflow.Context, input *mediapackagevod.DeleteAssetInput) (*mediapackagevod.DeleteAssetOutput, error) {
	var output mediapackagevod.DeleteAssetOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.DeleteAsset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAssetAsync(ctx workflow.Context, input *mediapackagevod.DeleteAssetInput) *MediaPackageVodDeleteAssetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.DeleteAsset", input)
	return &MediaPackageVodDeleteAssetFuture{Future: future}
}

func (a *stub) DeletePackagingConfiguration(ctx workflow.Context, input *mediapackagevod.DeletePackagingConfigurationInput) (*mediapackagevod.DeletePackagingConfigurationOutput, error) {
	var output mediapackagevod.DeletePackagingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.DeletePackagingConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePackagingConfigurationAsync(ctx workflow.Context, input *mediapackagevod.DeletePackagingConfigurationInput) *MediaPackageVodDeletePackagingConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.DeletePackagingConfiguration", input)
	return &MediaPackageVodDeletePackagingConfigurationFuture{Future: future}
}

func (a *stub) DeletePackagingGroup(ctx workflow.Context, input *mediapackagevod.DeletePackagingGroupInput) (*mediapackagevod.DeletePackagingGroupOutput, error) {
	var output mediapackagevod.DeletePackagingGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.DeletePackagingGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.DeletePackagingGroupInput) *MediaPackageVodDeletePackagingGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.DeletePackagingGroup", input)
	return &MediaPackageVodDeletePackagingGroupFuture{Future: future}
}

func (a *stub) DescribeAsset(ctx workflow.Context, input *mediapackagevod.DescribeAssetInput) (*mediapackagevod.DescribeAssetOutput, error) {
	var output mediapackagevod.DescribeAssetOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.DescribeAsset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAssetAsync(ctx workflow.Context, input *mediapackagevod.DescribeAssetInput) *MediaPackageVodDescribeAssetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.DescribeAsset", input)
	return &MediaPackageVodDescribeAssetFuture{Future: future}
}

func (a *stub) DescribePackagingConfiguration(ctx workflow.Context, input *mediapackagevod.DescribePackagingConfigurationInput) (*mediapackagevod.DescribePackagingConfigurationOutput, error) {
	var output mediapackagevod.DescribePackagingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.DescribePackagingConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribePackagingConfigurationAsync(ctx workflow.Context, input *mediapackagevod.DescribePackagingConfigurationInput) *MediaPackageVodDescribePackagingConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.DescribePackagingConfiguration", input)
	return &MediaPackageVodDescribePackagingConfigurationFuture{Future: future}
}

func (a *stub) DescribePackagingGroup(ctx workflow.Context, input *mediapackagevod.DescribePackagingGroupInput) (*mediapackagevod.DescribePackagingGroupOutput, error) {
	var output mediapackagevod.DescribePackagingGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.DescribePackagingGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.DescribePackagingGroupInput) *MediaPackageVodDescribePackagingGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.DescribePackagingGroup", input)
	return &MediaPackageVodDescribePackagingGroupFuture{Future: future}
}

func (a *stub) ListAssets(ctx workflow.Context, input *mediapackagevod.ListAssetsInput) (*mediapackagevod.ListAssetsOutput, error) {
	var output mediapackagevod.ListAssetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.ListAssets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAssetsAsync(ctx workflow.Context, input *mediapackagevod.ListAssetsInput) *MediaPackageVodListAssetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.ListAssets", input)
	return &MediaPackageVodListAssetsFuture{Future: future}
}

func (a *stub) ListPackagingConfigurations(ctx workflow.Context, input *mediapackagevod.ListPackagingConfigurationsInput) (*mediapackagevod.ListPackagingConfigurationsOutput, error) {
	var output mediapackagevod.ListPackagingConfigurationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.ListPackagingConfigurations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPackagingConfigurationsAsync(ctx workflow.Context, input *mediapackagevod.ListPackagingConfigurationsInput) *MediaPackageVodListPackagingConfigurationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.ListPackagingConfigurations", input)
	return &MediaPackageVodListPackagingConfigurationsFuture{Future: future}
}

func (a *stub) ListPackagingGroups(ctx workflow.Context, input *mediapackagevod.ListPackagingGroupsInput) (*mediapackagevod.ListPackagingGroupsOutput, error) {
	var output mediapackagevod.ListPackagingGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.ListPackagingGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPackagingGroupsAsync(ctx workflow.Context, input *mediapackagevod.ListPackagingGroupsInput) *MediaPackageVodListPackagingGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.ListPackagingGroups", input)
	return &MediaPackageVodListPackagingGroupsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *mediapackagevod.ListTagsForResourceInput) (*mediapackagevod.ListTagsForResourceOutput, error) {
	var output mediapackagevod.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *mediapackagevod.ListTagsForResourceInput) *MediaPackageVodListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.ListTagsForResource", input)
	return &MediaPackageVodListTagsForResourceFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *mediapackagevod.TagResourceInput) (*mediapackagevod.TagResourceOutput, error) {
	var output mediapackagevod.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *mediapackagevod.TagResourceInput) *MediaPackageVodTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.TagResource", input)
	return &MediaPackageVodTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *mediapackagevod.UntagResourceInput) (*mediapackagevod.UntagResourceOutput, error) {
	var output mediapackagevod.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *mediapackagevod.UntagResourceInput) *MediaPackageVodUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.UntagResource", input)
	return &MediaPackageVodUntagResourceFuture{Future: future}
}

func (a *stub) UpdatePackagingGroup(ctx workflow.Context, input *mediapackagevod.UpdatePackagingGroupInput) (*mediapackagevod.UpdatePackagingGroupOutput, error) {
	var output mediapackagevod.UpdatePackagingGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.UpdatePackagingGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdatePackagingGroupAsync(ctx workflow.Context, input *mediapackagevod.UpdatePackagingGroupInput) *MediaPackageVodUpdatePackagingGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackagevod.UpdatePackagingGroup", input)
	return &MediaPackageVodUpdatePackagingGroupFuture{Future: future}
}