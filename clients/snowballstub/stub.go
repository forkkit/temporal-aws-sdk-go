// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package snowballstub

import (
	"github.com/aws/aws-sdk-go/service/snowball"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type SnowballCancelClusterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballCancelClusterFuture) Get(ctx workflow.Context) (*snowball.CancelClusterOutput, error) {
	var output snowball.CancelClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballCancelJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballCancelJobFuture) Get(ctx workflow.Context) (*snowball.CancelJobOutput, error) {
	var output snowball.CancelJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballCreateAddressFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballCreateAddressFuture) Get(ctx workflow.Context) (*snowball.CreateAddressOutput, error) {
	var output snowball.CreateAddressOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballCreateClusterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballCreateClusterFuture) Get(ctx workflow.Context) (*snowball.CreateClusterOutput, error) {
	var output snowball.CreateClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballCreateJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballCreateJobFuture) Get(ctx workflow.Context) (*snowball.CreateJobOutput, error) {
	var output snowball.CreateJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballCreateReturnShippingLabelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballCreateReturnShippingLabelFuture) Get(ctx workflow.Context) (*snowball.CreateReturnShippingLabelOutput, error) {
	var output snowball.CreateReturnShippingLabelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballDescribeAddressFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballDescribeAddressFuture) Get(ctx workflow.Context) (*snowball.DescribeAddressOutput, error) {
	var output snowball.DescribeAddressOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballDescribeAddressesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballDescribeAddressesFuture) Get(ctx workflow.Context) (*snowball.DescribeAddressesOutput, error) {
	var output snowball.DescribeAddressesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballDescribeClusterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballDescribeClusterFuture) Get(ctx workflow.Context) (*snowball.DescribeClusterOutput, error) {
	var output snowball.DescribeClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballDescribeJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballDescribeJobFuture) Get(ctx workflow.Context) (*snowball.DescribeJobOutput, error) {
	var output snowball.DescribeJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballDescribeReturnShippingLabelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballDescribeReturnShippingLabelFuture) Get(ctx workflow.Context) (*snowball.DescribeReturnShippingLabelOutput, error) {
	var output snowball.DescribeReturnShippingLabelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballGetJobManifestFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballGetJobManifestFuture) Get(ctx workflow.Context) (*snowball.GetJobManifestOutput, error) {
	var output snowball.GetJobManifestOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballGetJobUnlockCodeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballGetJobUnlockCodeFuture) Get(ctx workflow.Context) (*snowball.GetJobUnlockCodeOutput, error) {
	var output snowball.GetJobUnlockCodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballGetSnowballUsageFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballGetSnowballUsageFuture) Get(ctx workflow.Context) (*snowball.GetSnowballUsageOutput, error) {
	var output snowball.GetSnowballUsageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballGetSoftwareUpdatesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballGetSoftwareUpdatesFuture) Get(ctx workflow.Context) (*snowball.GetSoftwareUpdatesOutput, error) {
	var output snowball.GetSoftwareUpdatesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballListClusterJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballListClusterJobsFuture) Get(ctx workflow.Context) (*snowball.ListClusterJobsOutput, error) {
	var output snowball.ListClusterJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballListClustersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballListClustersFuture) Get(ctx workflow.Context) (*snowball.ListClustersOutput, error) {
	var output snowball.ListClustersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballListCompatibleImagesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballListCompatibleImagesFuture) Get(ctx workflow.Context) (*snowball.ListCompatibleImagesOutput, error) {
	var output snowball.ListCompatibleImagesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballListJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballListJobsFuture) Get(ctx workflow.Context) (*snowball.ListJobsOutput, error) {
	var output snowball.ListJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballUpdateClusterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballUpdateClusterFuture) Get(ctx workflow.Context) (*snowball.UpdateClusterOutput, error) {
	var output snowball.UpdateClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballUpdateJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballUpdateJobFuture) Get(ctx workflow.Context) (*snowball.UpdateJobOutput, error) {
	var output snowball.UpdateJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SnowballUpdateJobShipmentStateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SnowballUpdateJobShipmentStateFuture) Get(ctx workflow.Context) (*snowball.UpdateJobShipmentStateOutput, error) {
	var output snowball.UpdateJobShipmentStateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelCluster(ctx workflow.Context, input *snowball.CancelClusterInput) (*snowball.CancelClusterOutput, error) {
	var output snowball.CancelClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.CancelCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelClusterAsync(ctx workflow.Context, input *snowball.CancelClusterInput) *SnowballCancelClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.CancelCluster", input)
	return &SnowballCancelClusterFuture{Future: future}
}

func (a *stub) CancelJob(ctx workflow.Context, input *snowball.CancelJobInput) (*snowball.CancelJobOutput, error) {
	var output snowball.CancelJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.CancelJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelJobAsync(ctx workflow.Context, input *snowball.CancelJobInput) *SnowballCancelJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.CancelJob", input)
	return &SnowballCancelJobFuture{Future: future}
}

func (a *stub) CreateAddress(ctx workflow.Context, input *snowball.CreateAddressInput) (*snowball.CreateAddressOutput, error) {
	var output snowball.CreateAddressOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.CreateAddress", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAddressAsync(ctx workflow.Context, input *snowball.CreateAddressInput) *SnowballCreateAddressFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.CreateAddress", input)
	return &SnowballCreateAddressFuture{Future: future}
}

func (a *stub) CreateCluster(ctx workflow.Context, input *snowball.CreateClusterInput) (*snowball.CreateClusterOutput, error) {
	var output snowball.CreateClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.CreateCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateClusterAsync(ctx workflow.Context, input *snowball.CreateClusterInput) *SnowballCreateClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.CreateCluster", input)
	return &SnowballCreateClusterFuture{Future: future}
}

func (a *stub) CreateJob(ctx workflow.Context, input *snowball.CreateJobInput) (*snowball.CreateJobOutput, error) {
	var output snowball.CreateJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.CreateJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateJobAsync(ctx workflow.Context, input *snowball.CreateJobInput) *SnowballCreateJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.CreateJob", input)
	return &SnowballCreateJobFuture{Future: future}
}

func (a *stub) CreateReturnShippingLabel(ctx workflow.Context, input *snowball.CreateReturnShippingLabelInput) (*snowball.CreateReturnShippingLabelOutput, error) {
	var output snowball.CreateReturnShippingLabelOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.CreateReturnShippingLabel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateReturnShippingLabelAsync(ctx workflow.Context, input *snowball.CreateReturnShippingLabelInput) *SnowballCreateReturnShippingLabelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.CreateReturnShippingLabel", input)
	return &SnowballCreateReturnShippingLabelFuture{Future: future}
}

func (a *stub) DescribeAddress(ctx workflow.Context, input *snowball.DescribeAddressInput) (*snowball.DescribeAddressOutput, error) {
	var output snowball.DescribeAddressOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.DescribeAddress", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAddressAsync(ctx workflow.Context, input *snowball.DescribeAddressInput) *SnowballDescribeAddressFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.DescribeAddress", input)
	return &SnowballDescribeAddressFuture{Future: future}
}

func (a *stub) DescribeAddresses(ctx workflow.Context, input *snowball.DescribeAddressesInput) (*snowball.DescribeAddressesOutput, error) {
	var output snowball.DescribeAddressesOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.DescribeAddresses", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAddressesAsync(ctx workflow.Context, input *snowball.DescribeAddressesInput) *SnowballDescribeAddressesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.DescribeAddresses", input)
	return &SnowballDescribeAddressesFuture{Future: future}
}

func (a *stub) DescribeCluster(ctx workflow.Context, input *snowball.DescribeClusterInput) (*snowball.DescribeClusterOutput, error) {
	var output snowball.DescribeClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.DescribeCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeClusterAsync(ctx workflow.Context, input *snowball.DescribeClusterInput) *SnowballDescribeClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.DescribeCluster", input)
	return &SnowballDescribeClusterFuture{Future: future}
}

func (a *stub) DescribeJob(ctx workflow.Context, input *snowball.DescribeJobInput) (*snowball.DescribeJobOutput, error) {
	var output snowball.DescribeJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.DescribeJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeJobAsync(ctx workflow.Context, input *snowball.DescribeJobInput) *SnowballDescribeJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.DescribeJob", input)
	return &SnowballDescribeJobFuture{Future: future}
}

func (a *stub) DescribeReturnShippingLabel(ctx workflow.Context, input *snowball.DescribeReturnShippingLabelInput) (*snowball.DescribeReturnShippingLabelOutput, error) {
	var output snowball.DescribeReturnShippingLabelOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.DescribeReturnShippingLabel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeReturnShippingLabelAsync(ctx workflow.Context, input *snowball.DescribeReturnShippingLabelInput) *SnowballDescribeReturnShippingLabelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.DescribeReturnShippingLabel", input)
	return &SnowballDescribeReturnShippingLabelFuture{Future: future}
}

func (a *stub) GetJobManifest(ctx workflow.Context, input *snowball.GetJobManifestInput) (*snowball.GetJobManifestOutput, error) {
	var output snowball.GetJobManifestOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.GetJobManifest", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetJobManifestAsync(ctx workflow.Context, input *snowball.GetJobManifestInput) *SnowballGetJobManifestFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.GetJobManifest", input)
	return &SnowballGetJobManifestFuture{Future: future}
}

func (a *stub) GetJobUnlockCode(ctx workflow.Context, input *snowball.GetJobUnlockCodeInput) (*snowball.GetJobUnlockCodeOutput, error) {
	var output snowball.GetJobUnlockCodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.GetJobUnlockCode", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetJobUnlockCodeAsync(ctx workflow.Context, input *snowball.GetJobUnlockCodeInput) *SnowballGetJobUnlockCodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.GetJobUnlockCode", input)
	return &SnowballGetJobUnlockCodeFuture{Future: future}
}

func (a *stub) GetSnowballUsage(ctx workflow.Context, input *snowball.GetSnowballUsageInput) (*snowball.GetSnowballUsageOutput, error) {
	var output snowball.GetSnowballUsageOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.GetSnowballUsage", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetSnowballUsageAsync(ctx workflow.Context, input *snowball.GetSnowballUsageInput) *SnowballGetSnowballUsageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.GetSnowballUsage", input)
	return &SnowballGetSnowballUsageFuture{Future: future}
}

func (a *stub) GetSoftwareUpdates(ctx workflow.Context, input *snowball.GetSoftwareUpdatesInput) (*snowball.GetSoftwareUpdatesOutput, error) {
	var output snowball.GetSoftwareUpdatesOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.GetSoftwareUpdates", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetSoftwareUpdatesAsync(ctx workflow.Context, input *snowball.GetSoftwareUpdatesInput) *SnowballGetSoftwareUpdatesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.GetSoftwareUpdates", input)
	return &SnowballGetSoftwareUpdatesFuture{Future: future}
}

func (a *stub) ListClusterJobs(ctx workflow.Context, input *snowball.ListClusterJobsInput) (*snowball.ListClusterJobsOutput, error) {
	var output snowball.ListClusterJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.ListClusterJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListClusterJobsAsync(ctx workflow.Context, input *snowball.ListClusterJobsInput) *SnowballListClusterJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.ListClusterJobs", input)
	return &SnowballListClusterJobsFuture{Future: future}
}

func (a *stub) ListClusters(ctx workflow.Context, input *snowball.ListClustersInput) (*snowball.ListClustersOutput, error) {
	var output snowball.ListClustersOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.ListClusters", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListClustersAsync(ctx workflow.Context, input *snowball.ListClustersInput) *SnowballListClustersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.ListClusters", input)
	return &SnowballListClustersFuture{Future: future}
}

func (a *stub) ListCompatibleImages(ctx workflow.Context, input *snowball.ListCompatibleImagesInput) (*snowball.ListCompatibleImagesOutput, error) {
	var output snowball.ListCompatibleImagesOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.ListCompatibleImages", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListCompatibleImagesAsync(ctx workflow.Context, input *snowball.ListCompatibleImagesInput) *SnowballListCompatibleImagesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.ListCompatibleImages", input)
	return &SnowballListCompatibleImagesFuture{Future: future}
}

func (a *stub) ListJobs(ctx workflow.Context, input *snowball.ListJobsInput) (*snowball.ListJobsOutput, error) {
	var output snowball.ListJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.ListJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListJobsAsync(ctx workflow.Context, input *snowball.ListJobsInput) *SnowballListJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.ListJobs", input)
	return &SnowballListJobsFuture{Future: future}
}

func (a *stub) UpdateCluster(ctx workflow.Context, input *snowball.UpdateClusterInput) (*snowball.UpdateClusterOutput, error) {
	var output snowball.UpdateClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.UpdateCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateClusterAsync(ctx workflow.Context, input *snowball.UpdateClusterInput) *SnowballUpdateClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.UpdateCluster", input)
	return &SnowballUpdateClusterFuture{Future: future}
}

func (a *stub) UpdateJob(ctx workflow.Context, input *snowball.UpdateJobInput) (*snowball.UpdateJobOutput, error) {
	var output snowball.UpdateJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.UpdateJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateJobAsync(ctx workflow.Context, input *snowball.UpdateJobInput) *SnowballUpdateJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.UpdateJob", input)
	return &SnowballUpdateJobFuture{Future: future}
}

func (a *stub) UpdateJobShipmentState(ctx workflow.Context, input *snowball.UpdateJobShipmentStateInput) (*snowball.UpdateJobShipmentStateOutput, error) {
	var output snowball.UpdateJobShipmentStateOutput
	err := workflow.ExecuteActivity(ctx, "aws.snowball.UpdateJobShipmentState", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateJobShipmentStateAsync(ctx workflow.Context, input *snowball.UpdateJobShipmentStateInput) *SnowballUpdateJobShipmentStateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.snowball.UpdateJobShipmentState", input)
	return &SnowballUpdateJobShipmentStateFuture{Future: future}
}