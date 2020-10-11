// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package managedblockchainactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/managedblockchain"
	"github.com/aws/aws-sdk-go/service/managedblockchain/managedblockchainiface"
	"go.temporal.io/aws-sdk/internal"
	"go.temporal.io/aws-sdk/sessionfactory"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type Activities struct {
	client managedblockchainiface.ManagedBlockchainAPI

	sessionFactory sessionfactory.SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := managedblockchain.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory sessionfactory.SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (managedblockchainiface.ManagedBlockchainAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return managedblockchain.New(sess), nil
}

func (a *Activities) CreateMember(ctx context.Context, input *managedblockchain.CreateMemberInput) (*managedblockchain.CreateMemberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateMemberWithContext(ctx, input)
}

func (a *Activities) CreateNetwork(ctx context.Context, input *managedblockchain.CreateNetworkInput) (*managedblockchain.CreateNetworkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateNetworkWithContext(ctx, input)
}

func (a *Activities) CreateNode(ctx context.Context, input *managedblockchain.CreateNodeInput) (*managedblockchain.CreateNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateNodeWithContext(ctx, input)
}

func (a *Activities) CreateProposal(ctx context.Context, input *managedblockchain.CreateProposalInput) (*managedblockchain.CreateProposalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateProposalWithContext(ctx, input)
}

func (a *Activities) DeleteMember(ctx context.Context, input *managedblockchain.DeleteMemberInput) (*managedblockchain.DeleteMemberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteMemberWithContext(ctx, input)
}

func (a *Activities) DeleteNode(ctx context.Context, input *managedblockchain.DeleteNodeInput) (*managedblockchain.DeleteNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteNodeWithContext(ctx, input)
}

func (a *Activities) GetMember(ctx context.Context, input *managedblockchain.GetMemberInput) (*managedblockchain.GetMemberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMemberWithContext(ctx, input)
}

func (a *Activities) GetNetwork(ctx context.Context, input *managedblockchain.GetNetworkInput) (*managedblockchain.GetNetworkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetNetworkWithContext(ctx, input)
}

func (a *Activities) GetNode(ctx context.Context, input *managedblockchain.GetNodeInput) (*managedblockchain.GetNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetNodeWithContext(ctx, input)
}

func (a *Activities) GetProposal(ctx context.Context, input *managedblockchain.GetProposalInput) (*managedblockchain.GetProposalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetProposalWithContext(ctx, input)
}

func (a *Activities) ListInvitations(ctx context.Context, input *managedblockchain.ListInvitationsInput) (*managedblockchain.ListInvitationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListInvitationsWithContext(ctx, input)
}

func (a *Activities) ListMembers(ctx context.Context, input *managedblockchain.ListMembersInput) (*managedblockchain.ListMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListMembersWithContext(ctx, input)
}

func (a *Activities) ListNetworks(ctx context.Context, input *managedblockchain.ListNetworksInput) (*managedblockchain.ListNetworksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListNetworksWithContext(ctx, input)
}

func (a *Activities) ListNodes(ctx context.Context, input *managedblockchain.ListNodesInput) (*managedblockchain.ListNodesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListNodesWithContext(ctx, input)
}

func (a *Activities) ListProposalVotes(ctx context.Context, input *managedblockchain.ListProposalVotesInput) (*managedblockchain.ListProposalVotesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListProposalVotesWithContext(ctx, input)
}

func (a *Activities) ListProposals(ctx context.Context, input *managedblockchain.ListProposalsInput) (*managedblockchain.ListProposalsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListProposalsWithContext(ctx, input)
}

func (a *Activities) RejectInvitation(ctx context.Context, input *managedblockchain.RejectInvitationInput) (*managedblockchain.RejectInvitationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RejectInvitationWithContext(ctx, input)
}

func (a *Activities) UpdateMember(ctx context.Context, input *managedblockchain.UpdateMemberInput) (*managedblockchain.UpdateMemberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateMemberWithContext(ctx, input)
}

func (a *Activities) UpdateNode(ctx context.Context, input *managedblockchain.UpdateNodeInput) (*managedblockchain.UpdateNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateNodeWithContext(ctx, input)
}

func (a *Activities) VoteOnProposal(ctx context.Context, input *managedblockchain.VoteOnProposalInput) (*managedblockchain.VoteOnProposalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.VoteOnProposalWithContext(ctx, input)
}
