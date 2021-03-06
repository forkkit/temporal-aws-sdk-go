// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mediaconnectstub

import (
	"github.com/aws/aws-sdk-go/service/mediaconnect"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type MediaConnectAddFlowOutputsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectAddFlowOutputsFuture) Get(ctx workflow.Context) (*mediaconnect.AddFlowOutputsOutput, error) {
	var output mediaconnect.AddFlowOutputsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectAddFlowSourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectAddFlowSourcesFuture) Get(ctx workflow.Context) (*mediaconnect.AddFlowSourcesOutput, error) {
	var output mediaconnect.AddFlowSourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectAddFlowVpcInterfacesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectAddFlowVpcInterfacesFuture) Get(ctx workflow.Context) (*mediaconnect.AddFlowVpcInterfacesOutput, error) {
	var output mediaconnect.AddFlowVpcInterfacesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectCreateFlowFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectCreateFlowFuture) Get(ctx workflow.Context) (*mediaconnect.CreateFlowOutput, error) {
	var output mediaconnect.CreateFlowOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectDeleteFlowFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectDeleteFlowFuture) Get(ctx workflow.Context) (*mediaconnect.DeleteFlowOutput, error) {
	var output mediaconnect.DeleteFlowOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectDescribeFlowFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectDescribeFlowFuture) Get(ctx workflow.Context) (*mediaconnect.DescribeFlowOutput, error) {
	var output mediaconnect.DescribeFlowOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectDescribeOfferingFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectDescribeOfferingFuture) Get(ctx workflow.Context) (*mediaconnect.DescribeOfferingOutput, error) {
	var output mediaconnect.DescribeOfferingOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectDescribeReservationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectDescribeReservationFuture) Get(ctx workflow.Context) (*mediaconnect.DescribeReservationOutput, error) {
	var output mediaconnect.DescribeReservationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectGrantFlowEntitlementsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectGrantFlowEntitlementsFuture) Get(ctx workflow.Context) (*mediaconnect.GrantFlowEntitlementsOutput, error) {
	var output mediaconnect.GrantFlowEntitlementsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectListEntitlementsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectListEntitlementsFuture) Get(ctx workflow.Context) (*mediaconnect.ListEntitlementsOutput, error) {
	var output mediaconnect.ListEntitlementsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectListFlowsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectListFlowsFuture) Get(ctx workflow.Context) (*mediaconnect.ListFlowsOutput, error) {
	var output mediaconnect.ListFlowsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectListOfferingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectListOfferingsFuture) Get(ctx workflow.Context) (*mediaconnect.ListOfferingsOutput, error) {
	var output mediaconnect.ListOfferingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectListReservationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectListReservationsFuture) Get(ctx workflow.Context) (*mediaconnect.ListReservationsOutput, error) {
	var output mediaconnect.ListReservationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectListTagsForResourceFuture) Get(ctx workflow.Context) (*mediaconnect.ListTagsForResourceOutput, error) {
	var output mediaconnect.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectPurchaseOfferingFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectPurchaseOfferingFuture) Get(ctx workflow.Context) (*mediaconnect.PurchaseOfferingOutput, error) {
	var output mediaconnect.PurchaseOfferingOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectRemoveFlowOutputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectRemoveFlowOutputFuture) Get(ctx workflow.Context) (*mediaconnect.RemoveFlowOutputOutput, error) {
	var output mediaconnect.RemoveFlowOutputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectRemoveFlowSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectRemoveFlowSourceFuture) Get(ctx workflow.Context) (*mediaconnect.RemoveFlowSourceOutput, error) {
	var output mediaconnect.RemoveFlowSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectRemoveFlowVpcInterfaceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectRemoveFlowVpcInterfaceFuture) Get(ctx workflow.Context) (*mediaconnect.RemoveFlowVpcInterfaceOutput, error) {
	var output mediaconnect.RemoveFlowVpcInterfaceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectRevokeFlowEntitlementFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectRevokeFlowEntitlementFuture) Get(ctx workflow.Context) (*mediaconnect.RevokeFlowEntitlementOutput, error) {
	var output mediaconnect.RevokeFlowEntitlementOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectStartFlowFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectStartFlowFuture) Get(ctx workflow.Context) (*mediaconnect.StartFlowOutput, error) {
	var output mediaconnect.StartFlowOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectStopFlowFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectStopFlowFuture) Get(ctx workflow.Context) (*mediaconnect.StopFlowOutput, error) {
	var output mediaconnect.StopFlowOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectTagResourceFuture) Get(ctx workflow.Context) (*mediaconnect.TagResourceOutput, error) {
	var output mediaconnect.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectUntagResourceFuture) Get(ctx workflow.Context) (*mediaconnect.UntagResourceOutput, error) {
	var output mediaconnect.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectUpdateFlowFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectUpdateFlowFuture) Get(ctx workflow.Context) (*mediaconnect.UpdateFlowOutput, error) {
	var output mediaconnect.UpdateFlowOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectUpdateFlowEntitlementFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectUpdateFlowEntitlementFuture) Get(ctx workflow.Context) (*mediaconnect.UpdateFlowEntitlementOutput, error) {
	var output mediaconnect.UpdateFlowEntitlementOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectUpdateFlowOutputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectUpdateFlowOutputFuture) Get(ctx workflow.Context) (*mediaconnect.UpdateFlowOutputOutput, error) {
	var output mediaconnect.UpdateFlowOutputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConnectUpdateFlowSourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConnectUpdateFlowSourceFuture) Get(ctx workflow.Context) (*mediaconnect.UpdateFlowSourceOutput, error) {
	var output mediaconnect.UpdateFlowSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AddFlowOutputs(ctx workflow.Context, input *mediaconnect.AddFlowOutputsInput) (*mediaconnect.AddFlowOutputsOutput, error) {
	var output mediaconnect.AddFlowOutputsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.AddFlowOutputs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddFlowOutputsAsync(ctx workflow.Context, input *mediaconnect.AddFlowOutputsInput) *MediaConnectAddFlowOutputsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.AddFlowOutputs", input)
	return &MediaConnectAddFlowOutputsFuture{Future: future}
}

func (a *stub) AddFlowSources(ctx workflow.Context, input *mediaconnect.AddFlowSourcesInput) (*mediaconnect.AddFlowSourcesOutput, error) {
	var output mediaconnect.AddFlowSourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.AddFlowSources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddFlowSourcesAsync(ctx workflow.Context, input *mediaconnect.AddFlowSourcesInput) *MediaConnectAddFlowSourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.AddFlowSources", input)
	return &MediaConnectAddFlowSourcesFuture{Future: future}
}

func (a *stub) AddFlowVpcInterfaces(ctx workflow.Context, input *mediaconnect.AddFlowVpcInterfacesInput) (*mediaconnect.AddFlowVpcInterfacesOutput, error) {
	var output mediaconnect.AddFlowVpcInterfacesOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.AddFlowVpcInterfaces", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddFlowVpcInterfacesAsync(ctx workflow.Context, input *mediaconnect.AddFlowVpcInterfacesInput) *MediaConnectAddFlowVpcInterfacesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.AddFlowVpcInterfaces", input)
	return &MediaConnectAddFlowVpcInterfacesFuture{Future: future}
}

func (a *stub) CreateFlow(ctx workflow.Context, input *mediaconnect.CreateFlowInput) (*mediaconnect.CreateFlowOutput, error) {
	var output mediaconnect.CreateFlowOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.CreateFlow", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateFlowAsync(ctx workflow.Context, input *mediaconnect.CreateFlowInput) *MediaConnectCreateFlowFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.CreateFlow", input)
	return &MediaConnectCreateFlowFuture{Future: future}
}

func (a *stub) DeleteFlow(ctx workflow.Context, input *mediaconnect.DeleteFlowInput) (*mediaconnect.DeleteFlowOutput, error) {
	var output mediaconnect.DeleteFlowOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.DeleteFlow", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteFlowAsync(ctx workflow.Context, input *mediaconnect.DeleteFlowInput) *MediaConnectDeleteFlowFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.DeleteFlow", input)
	return &MediaConnectDeleteFlowFuture{Future: future}
}

func (a *stub) DescribeFlow(ctx workflow.Context, input *mediaconnect.DescribeFlowInput) (*mediaconnect.DescribeFlowOutput, error) {
	var output mediaconnect.DescribeFlowOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.DescribeFlow", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeFlowAsync(ctx workflow.Context, input *mediaconnect.DescribeFlowInput) *MediaConnectDescribeFlowFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.DescribeFlow", input)
	return &MediaConnectDescribeFlowFuture{Future: future}
}

func (a *stub) DescribeOffering(ctx workflow.Context, input *mediaconnect.DescribeOfferingInput) (*mediaconnect.DescribeOfferingOutput, error) {
	var output mediaconnect.DescribeOfferingOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.DescribeOffering", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeOfferingAsync(ctx workflow.Context, input *mediaconnect.DescribeOfferingInput) *MediaConnectDescribeOfferingFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.DescribeOffering", input)
	return &MediaConnectDescribeOfferingFuture{Future: future}
}

func (a *stub) DescribeReservation(ctx workflow.Context, input *mediaconnect.DescribeReservationInput) (*mediaconnect.DescribeReservationOutput, error) {
	var output mediaconnect.DescribeReservationOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.DescribeReservation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeReservationAsync(ctx workflow.Context, input *mediaconnect.DescribeReservationInput) *MediaConnectDescribeReservationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.DescribeReservation", input)
	return &MediaConnectDescribeReservationFuture{Future: future}
}

func (a *stub) GrantFlowEntitlements(ctx workflow.Context, input *mediaconnect.GrantFlowEntitlementsInput) (*mediaconnect.GrantFlowEntitlementsOutput, error) {
	var output mediaconnect.GrantFlowEntitlementsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.GrantFlowEntitlements", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GrantFlowEntitlementsAsync(ctx workflow.Context, input *mediaconnect.GrantFlowEntitlementsInput) *MediaConnectGrantFlowEntitlementsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.GrantFlowEntitlements", input)
	return &MediaConnectGrantFlowEntitlementsFuture{Future: future}
}

func (a *stub) ListEntitlements(ctx workflow.Context, input *mediaconnect.ListEntitlementsInput) (*mediaconnect.ListEntitlementsOutput, error) {
	var output mediaconnect.ListEntitlementsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.ListEntitlements", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListEntitlementsAsync(ctx workflow.Context, input *mediaconnect.ListEntitlementsInput) *MediaConnectListEntitlementsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.ListEntitlements", input)
	return &MediaConnectListEntitlementsFuture{Future: future}
}

func (a *stub) ListFlows(ctx workflow.Context, input *mediaconnect.ListFlowsInput) (*mediaconnect.ListFlowsOutput, error) {
	var output mediaconnect.ListFlowsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.ListFlows", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListFlowsAsync(ctx workflow.Context, input *mediaconnect.ListFlowsInput) *MediaConnectListFlowsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.ListFlows", input)
	return &MediaConnectListFlowsFuture{Future: future}
}

func (a *stub) ListOfferings(ctx workflow.Context, input *mediaconnect.ListOfferingsInput) (*mediaconnect.ListOfferingsOutput, error) {
	var output mediaconnect.ListOfferingsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.ListOfferings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListOfferingsAsync(ctx workflow.Context, input *mediaconnect.ListOfferingsInput) *MediaConnectListOfferingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.ListOfferings", input)
	return &MediaConnectListOfferingsFuture{Future: future}
}

func (a *stub) ListReservations(ctx workflow.Context, input *mediaconnect.ListReservationsInput) (*mediaconnect.ListReservationsOutput, error) {
	var output mediaconnect.ListReservationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.ListReservations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListReservationsAsync(ctx workflow.Context, input *mediaconnect.ListReservationsInput) *MediaConnectListReservationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.ListReservations", input)
	return &MediaConnectListReservationsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *mediaconnect.ListTagsForResourceInput) (*mediaconnect.ListTagsForResourceOutput, error) {
	var output mediaconnect.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *mediaconnect.ListTagsForResourceInput) *MediaConnectListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.ListTagsForResource", input)
	return &MediaConnectListTagsForResourceFuture{Future: future}
}

func (a *stub) PurchaseOffering(ctx workflow.Context, input *mediaconnect.PurchaseOfferingInput) (*mediaconnect.PurchaseOfferingOutput, error) {
	var output mediaconnect.PurchaseOfferingOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.PurchaseOffering", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PurchaseOfferingAsync(ctx workflow.Context, input *mediaconnect.PurchaseOfferingInput) *MediaConnectPurchaseOfferingFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.PurchaseOffering", input)
	return &MediaConnectPurchaseOfferingFuture{Future: future}
}

func (a *stub) RemoveFlowOutput(ctx workflow.Context, input *mediaconnect.RemoveFlowOutputInput) (*mediaconnect.RemoveFlowOutputOutput, error) {
	var output mediaconnect.RemoveFlowOutputOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.RemoveFlowOutput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveFlowOutputAsync(ctx workflow.Context, input *mediaconnect.RemoveFlowOutputInput) *MediaConnectRemoveFlowOutputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.RemoveFlowOutput", input)
	return &MediaConnectRemoveFlowOutputFuture{Future: future}
}

func (a *stub) RemoveFlowSource(ctx workflow.Context, input *mediaconnect.RemoveFlowSourceInput) (*mediaconnect.RemoveFlowSourceOutput, error) {
	var output mediaconnect.RemoveFlowSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.RemoveFlowSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveFlowSourceAsync(ctx workflow.Context, input *mediaconnect.RemoveFlowSourceInput) *MediaConnectRemoveFlowSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.RemoveFlowSource", input)
	return &MediaConnectRemoveFlowSourceFuture{Future: future}
}

func (a *stub) RemoveFlowVpcInterface(ctx workflow.Context, input *mediaconnect.RemoveFlowVpcInterfaceInput) (*mediaconnect.RemoveFlowVpcInterfaceOutput, error) {
	var output mediaconnect.RemoveFlowVpcInterfaceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.RemoveFlowVpcInterface", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveFlowVpcInterfaceAsync(ctx workflow.Context, input *mediaconnect.RemoveFlowVpcInterfaceInput) *MediaConnectRemoveFlowVpcInterfaceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.RemoveFlowVpcInterface", input)
	return &MediaConnectRemoveFlowVpcInterfaceFuture{Future: future}
}

func (a *stub) RevokeFlowEntitlement(ctx workflow.Context, input *mediaconnect.RevokeFlowEntitlementInput) (*mediaconnect.RevokeFlowEntitlementOutput, error) {
	var output mediaconnect.RevokeFlowEntitlementOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.RevokeFlowEntitlement", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RevokeFlowEntitlementAsync(ctx workflow.Context, input *mediaconnect.RevokeFlowEntitlementInput) *MediaConnectRevokeFlowEntitlementFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.RevokeFlowEntitlement", input)
	return &MediaConnectRevokeFlowEntitlementFuture{Future: future}
}

func (a *stub) StartFlow(ctx workflow.Context, input *mediaconnect.StartFlowInput) (*mediaconnect.StartFlowOutput, error) {
	var output mediaconnect.StartFlowOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.StartFlow", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartFlowAsync(ctx workflow.Context, input *mediaconnect.StartFlowInput) *MediaConnectStartFlowFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.StartFlow", input)
	return &MediaConnectStartFlowFuture{Future: future}
}

func (a *stub) StopFlow(ctx workflow.Context, input *mediaconnect.StopFlowInput) (*mediaconnect.StopFlowOutput, error) {
	var output mediaconnect.StopFlowOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.StopFlow", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopFlowAsync(ctx workflow.Context, input *mediaconnect.StopFlowInput) *MediaConnectStopFlowFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.StopFlow", input)
	return &MediaConnectStopFlowFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *mediaconnect.TagResourceInput) (*mediaconnect.TagResourceOutput, error) {
	var output mediaconnect.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *mediaconnect.TagResourceInput) *MediaConnectTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.TagResource", input)
	return &MediaConnectTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *mediaconnect.UntagResourceInput) (*mediaconnect.UntagResourceOutput, error) {
	var output mediaconnect.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *mediaconnect.UntagResourceInput) *MediaConnectUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.UntagResource", input)
	return &MediaConnectUntagResourceFuture{Future: future}
}

func (a *stub) UpdateFlow(ctx workflow.Context, input *mediaconnect.UpdateFlowInput) (*mediaconnect.UpdateFlowOutput, error) {
	var output mediaconnect.UpdateFlowOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.UpdateFlow", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateFlowAsync(ctx workflow.Context, input *mediaconnect.UpdateFlowInput) *MediaConnectUpdateFlowFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.UpdateFlow", input)
	return &MediaConnectUpdateFlowFuture{Future: future}
}

func (a *stub) UpdateFlowEntitlement(ctx workflow.Context, input *mediaconnect.UpdateFlowEntitlementInput) (*mediaconnect.UpdateFlowEntitlementOutput, error) {
	var output mediaconnect.UpdateFlowEntitlementOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.UpdateFlowEntitlement", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateFlowEntitlementAsync(ctx workflow.Context, input *mediaconnect.UpdateFlowEntitlementInput) *MediaConnectUpdateFlowEntitlementFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.UpdateFlowEntitlement", input)
	return &MediaConnectUpdateFlowEntitlementFuture{Future: future}
}

func (a *stub) UpdateFlowOutput(ctx workflow.Context, input *mediaconnect.UpdateFlowOutputInput) (*mediaconnect.UpdateFlowOutputOutput, error) {
	var output mediaconnect.UpdateFlowOutputOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.UpdateFlowOutput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateFlowOutputAsync(ctx workflow.Context, input *mediaconnect.UpdateFlowOutputInput) *MediaConnectUpdateFlowOutputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.UpdateFlowOutput", input)
	return &MediaConnectUpdateFlowOutputFuture{Future: future}
}

func (a *stub) UpdateFlowSource(ctx workflow.Context, input *mediaconnect.UpdateFlowSourceInput) (*mediaconnect.UpdateFlowSourceOutput, error) {
	var output mediaconnect.UpdateFlowSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconnect.UpdateFlowSource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateFlowSourceAsync(ctx workflow.Context, input *mediaconnect.UpdateFlowSourceInput) *MediaConnectUpdateFlowSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconnect.UpdateFlowSource", input)
	return &MediaConnectUpdateFlowSourceFuture{Future: future}
}
