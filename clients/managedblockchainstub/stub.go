// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package managedblockchainstub

import (
	"github.com/aws/aws-sdk-go/service/managedblockchain"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type ManagedBlockchainCreateMemberFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainCreateMemberFuture) Get(ctx workflow.Context) (*managedblockchain.CreateMemberOutput, error) {
	var output managedblockchain.CreateMemberOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainCreateNetworkFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainCreateNetworkFuture) Get(ctx workflow.Context) (*managedblockchain.CreateNetworkOutput, error) {
	var output managedblockchain.CreateNetworkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainCreateNodeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainCreateNodeFuture) Get(ctx workflow.Context) (*managedblockchain.CreateNodeOutput, error) {
	var output managedblockchain.CreateNodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainCreateProposalFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainCreateProposalFuture) Get(ctx workflow.Context) (*managedblockchain.CreateProposalOutput, error) {
	var output managedblockchain.CreateProposalOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainDeleteMemberFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainDeleteMemberFuture) Get(ctx workflow.Context) (*managedblockchain.DeleteMemberOutput, error) {
	var output managedblockchain.DeleteMemberOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainDeleteNodeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainDeleteNodeFuture) Get(ctx workflow.Context) (*managedblockchain.DeleteNodeOutput, error) {
	var output managedblockchain.DeleteNodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainGetMemberFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainGetMemberFuture) Get(ctx workflow.Context) (*managedblockchain.GetMemberOutput, error) {
	var output managedblockchain.GetMemberOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainGetNetworkFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainGetNetworkFuture) Get(ctx workflow.Context) (*managedblockchain.GetNetworkOutput, error) {
	var output managedblockchain.GetNetworkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainGetNodeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainGetNodeFuture) Get(ctx workflow.Context) (*managedblockchain.GetNodeOutput, error) {
	var output managedblockchain.GetNodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainGetProposalFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainGetProposalFuture) Get(ctx workflow.Context) (*managedblockchain.GetProposalOutput, error) {
	var output managedblockchain.GetProposalOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainListInvitationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainListInvitationsFuture) Get(ctx workflow.Context) (*managedblockchain.ListInvitationsOutput, error) {
	var output managedblockchain.ListInvitationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainListMembersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainListMembersFuture) Get(ctx workflow.Context) (*managedblockchain.ListMembersOutput, error) {
	var output managedblockchain.ListMembersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainListNetworksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainListNetworksFuture) Get(ctx workflow.Context) (*managedblockchain.ListNetworksOutput, error) {
	var output managedblockchain.ListNetworksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainListNodesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainListNodesFuture) Get(ctx workflow.Context) (*managedblockchain.ListNodesOutput, error) {
	var output managedblockchain.ListNodesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainListProposalVotesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainListProposalVotesFuture) Get(ctx workflow.Context) (*managedblockchain.ListProposalVotesOutput, error) {
	var output managedblockchain.ListProposalVotesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainListProposalsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainListProposalsFuture) Get(ctx workflow.Context) (*managedblockchain.ListProposalsOutput, error) {
	var output managedblockchain.ListProposalsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainRejectInvitationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainRejectInvitationFuture) Get(ctx workflow.Context) (*managedblockchain.RejectInvitationOutput, error) {
	var output managedblockchain.RejectInvitationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainUpdateMemberFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainUpdateMemberFuture) Get(ctx workflow.Context) (*managedblockchain.UpdateMemberOutput, error) {
	var output managedblockchain.UpdateMemberOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainUpdateNodeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainUpdateNodeFuture) Get(ctx workflow.Context) (*managedblockchain.UpdateNodeOutput, error) {
	var output managedblockchain.UpdateNodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainVoteOnProposalFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ManagedBlockchainVoteOnProposalFuture) Get(ctx workflow.Context) (*managedblockchain.VoteOnProposalOutput, error) {
	var output managedblockchain.VoteOnProposalOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateMember(ctx workflow.Context, input *managedblockchain.CreateMemberInput) (*managedblockchain.CreateMemberOutput, error) {
	var output managedblockchain.CreateMemberOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateMember", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateMemberAsync(ctx workflow.Context, input *managedblockchain.CreateMemberInput) *ManagedBlockchainCreateMemberFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateMember", input)
	return &ManagedBlockchainCreateMemberFuture{Future: future}
}

func (a *stub) CreateNetwork(ctx workflow.Context, input *managedblockchain.CreateNetworkInput) (*managedblockchain.CreateNetworkOutput, error) {
	var output managedblockchain.CreateNetworkOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateNetwork", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateNetworkAsync(ctx workflow.Context, input *managedblockchain.CreateNetworkInput) *ManagedBlockchainCreateNetworkFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateNetwork", input)
	return &ManagedBlockchainCreateNetworkFuture{Future: future}
}

func (a *stub) CreateNode(ctx workflow.Context, input *managedblockchain.CreateNodeInput) (*managedblockchain.CreateNodeOutput, error) {
	var output managedblockchain.CreateNodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateNode", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateNodeAsync(ctx workflow.Context, input *managedblockchain.CreateNodeInput) *ManagedBlockchainCreateNodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateNode", input)
	return &ManagedBlockchainCreateNodeFuture{Future: future}
}

func (a *stub) CreateProposal(ctx workflow.Context, input *managedblockchain.CreateProposalInput) (*managedblockchain.CreateProposalOutput, error) {
	var output managedblockchain.CreateProposalOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateProposal", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateProposalAsync(ctx workflow.Context, input *managedblockchain.CreateProposalInput) *ManagedBlockchainCreateProposalFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateProposal", input)
	return &ManagedBlockchainCreateProposalFuture{Future: future}
}

func (a *stub) DeleteMember(ctx workflow.Context, input *managedblockchain.DeleteMemberInput) (*managedblockchain.DeleteMemberOutput, error) {
	var output managedblockchain.DeleteMemberOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.DeleteMember", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteMemberAsync(ctx workflow.Context, input *managedblockchain.DeleteMemberInput) *ManagedBlockchainDeleteMemberFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.DeleteMember", input)
	return &ManagedBlockchainDeleteMemberFuture{Future: future}
}

func (a *stub) DeleteNode(ctx workflow.Context, input *managedblockchain.DeleteNodeInput) (*managedblockchain.DeleteNodeOutput, error) {
	var output managedblockchain.DeleteNodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.DeleteNode", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteNodeAsync(ctx workflow.Context, input *managedblockchain.DeleteNodeInput) *ManagedBlockchainDeleteNodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.DeleteNode", input)
	return &ManagedBlockchainDeleteNodeFuture{Future: future}
}

func (a *stub) GetMember(ctx workflow.Context, input *managedblockchain.GetMemberInput) (*managedblockchain.GetMemberOutput, error) {
	var output managedblockchain.GetMemberOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetMember", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMemberAsync(ctx workflow.Context, input *managedblockchain.GetMemberInput) *ManagedBlockchainGetMemberFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetMember", input)
	return &ManagedBlockchainGetMemberFuture{Future: future}
}

func (a *stub) GetNetwork(ctx workflow.Context, input *managedblockchain.GetNetworkInput) (*managedblockchain.GetNetworkOutput, error) {
	var output managedblockchain.GetNetworkOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetNetwork", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetNetworkAsync(ctx workflow.Context, input *managedblockchain.GetNetworkInput) *ManagedBlockchainGetNetworkFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetNetwork", input)
	return &ManagedBlockchainGetNetworkFuture{Future: future}
}

func (a *stub) GetNode(ctx workflow.Context, input *managedblockchain.GetNodeInput) (*managedblockchain.GetNodeOutput, error) {
	var output managedblockchain.GetNodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetNode", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetNodeAsync(ctx workflow.Context, input *managedblockchain.GetNodeInput) *ManagedBlockchainGetNodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetNode", input)
	return &ManagedBlockchainGetNodeFuture{Future: future}
}

func (a *stub) GetProposal(ctx workflow.Context, input *managedblockchain.GetProposalInput) (*managedblockchain.GetProposalOutput, error) {
	var output managedblockchain.GetProposalOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetProposal", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetProposalAsync(ctx workflow.Context, input *managedblockchain.GetProposalInput) *ManagedBlockchainGetProposalFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetProposal", input)
	return &ManagedBlockchainGetProposalFuture{Future: future}
}

func (a *stub) ListInvitations(ctx workflow.Context, input *managedblockchain.ListInvitationsInput) (*managedblockchain.ListInvitationsOutput, error) {
	var output managedblockchain.ListInvitationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListInvitations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListInvitationsAsync(ctx workflow.Context, input *managedblockchain.ListInvitationsInput) *ManagedBlockchainListInvitationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListInvitations", input)
	return &ManagedBlockchainListInvitationsFuture{Future: future}
}

func (a *stub) ListMembers(ctx workflow.Context, input *managedblockchain.ListMembersInput) (*managedblockchain.ListMembersOutput, error) {
	var output managedblockchain.ListMembersOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListMembers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListMembersAsync(ctx workflow.Context, input *managedblockchain.ListMembersInput) *ManagedBlockchainListMembersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListMembers", input)
	return &ManagedBlockchainListMembersFuture{Future: future}
}

func (a *stub) ListNetworks(ctx workflow.Context, input *managedblockchain.ListNetworksInput) (*managedblockchain.ListNetworksOutput, error) {
	var output managedblockchain.ListNetworksOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListNetworks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListNetworksAsync(ctx workflow.Context, input *managedblockchain.ListNetworksInput) *ManagedBlockchainListNetworksFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListNetworks", input)
	return &ManagedBlockchainListNetworksFuture{Future: future}
}

func (a *stub) ListNodes(ctx workflow.Context, input *managedblockchain.ListNodesInput) (*managedblockchain.ListNodesOutput, error) {
	var output managedblockchain.ListNodesOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListNodes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListNodesAsync(ctx workflow.Context, input *managedblockchain.ListNodesInput) *ManagedBlockchainListNodesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListNodes", input)
	return &ManagedBlockchainListNodesFuture{Future: future}
}

func (a *stub) ListProposalVotes(ctx workflow.Context, input *managedblockchain.ListProposalVotesInput) (*managedblockchain.ListProposalVotesOutput, error) {
	var output managedblockchain.ListProposalVotesOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListProposalVotes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListProposalVotesAsync(ctx workflow.Context, input *managedblockchain.ListProposalVotesInput) *ManagedBlockchainListProposalVotesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListProposalVotes", input)
	return &ManagedBlockchainListProposalVotesFuture{Future: future}
}

func (a *stub) ListProposals(ctx workflow.Context, input *managedblockchain.ListProposalsInput) (*managedblockchain.ListProposalsOutput, error) {
	var output managedblockchain.ListProposalsOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListProposals", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListProposalsAsync(ctx workflow.Context, input *managedblockchain.ListProposalsInput) *ManagedBlockchainListProposalsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListProposals", input)
	return &ManagedBlockchainListProposalsFuture{Future: future}
}

func (a *stub) RejectInvitation(ctx workflow.Context, input *managedblockchain.RejectInvitationInput) (*managedblockchain.RejectInvitationOutput, error) {
	var output managedblockchain.RejectInvitationOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.RejectInvitation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RejectInvitationAsync(ctx workflow.Context, input *managedblockchain.RejectInvitationInput) *ManagedBlockchainRejectInvitationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.RejectInvitation", input)
	return &ManagedBlockchainRejectInvitationFuture{Future: future}
}

func (a *stub) UpdateMember(ctx workflow.Context, input *managedblockchain.UpdateMemberInput) (*managedblockchain.UpdateMemberOutput, error) {
	var output managedblockchain.UpdateMemberOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.UpdateMember", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateMemberAsync(ctx workflow.Context, input *managedblockchain.UpdateMemberInput) *ManagedBlockchainUpdateMemberFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.UpdateMember", input)
	return &ManagedBlockchainUpdateMemberFuture{Future: future}
}

func (a *stub) UpdateNode(ctx workflow.Context, input *managedblockchain.UpdateNodeInput) (*managedblockchain.UpdateNodeOutput, error) {
	var output managedblockchain.UpdateNodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.UpdateNode", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateNodeAsync(ctx workflow.Context, input *managedblockchain.UpdateNodeInput) *ManagedBlockchainUpdateNodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.UpdateNode", input)
	return &ManagedBlockchainUpdateNodeFuture{Future: future}
}

func (a *stub) VoteOnProposal(ctx workflow.Context, input *managedblockchain.VoteOnProposalInput) (*managedblockchain.VoteOnProposalOutput, error) {
	var output managedblockchain.VoteOnProposalOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.VoteOnProposal", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) VoteOnProposalAsync(ctx workflow.Context, input *managedblockchain.VoteOnProposalInput) *ManagedBlockchainVoteOnProposalFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.VoteOnProposal", input)
	return &ManagedBlockchainVoteOnProposalFuture{Future: future}
}
