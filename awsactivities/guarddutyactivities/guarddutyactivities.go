// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package guarddutyactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/guardduty"
	"github.com/aws/aws-sdk-go/service/guardduty/guarddutyiface"
	"go.temporal.io/aws-sdk/internal"
	"go.temporal.io/aws-sdk/sessionfactory"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type Activities struct {
	client guarddutyiface.GuardDutyAPI

	sessionFactory sessionfactory.SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := guardduty.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory sessionfactory.SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (guarddutyiface.GuardDutyAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return guardduty.New(sess), nil
}

func (a *Activities) AcceptInvitation(ctx context.Context, input *guardduty.AcceptInvitationInput) (*guardduty.AcceptInvitationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AcceptInvitationWithContext(ctx, input)
}

func (a *Activities) ArchiveFindings(ctx context.Context, input *guardduty.ArchiveFindingsInput) (*guardduty.ArchiveFindingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ArchiveFindingsWithContext(ctx, input)
}

func (a *Activities) CreateDetector(ctx context.Context, input *guardduty.CreateDetectorInput) (*guardduty.CreateDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateDetectorWithContext(ctx, input)
}

func (a *Activities) CreateFilter(ctx context.Context, input *guardduty.CreateFilterInput) (*guardduty.CreateFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateFilterWithContext(ctx, input)
}

func (a *Activities) CreateIPSet(ctx context.Context, input *guardduty.CreateIPSetInput) (*guardduty.CreateIPSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateIPSetWithContext(ctx, input)
}

func (a *Activities) CreateMembers(ctx context.Context, input *guardduty.CreateMembersInput) (*guardduty.CreateMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateMembersWithContext(ctx, input)
}

func (a *Activities) CreatePublishingDestination(ctx context.Context, input *guardduty.CreatePublishingDestinationInput) (*guardduty.CreatePublishingDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreatePublishingDestinationWithContext(ctx, input)
}

func (a *Activities) CreateSampleFindings(ctx context.Context, input *guardduty.CreateSampleFindingsInput) (*guardduty.CreateSampleFindingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSampleFindingsWithContext(ctx, input)
}

func (a *Activities) CreateThreatIntelSet(ctx context.Context, input *guardduty.CreateThreatIntelSetInput) (*guardduty.CreateThreatIntelSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateThreatIntelSetWithContext(ctx, input)
}

func (a *Activities) DeclineInvitations(ctx context.Context, input *guardduty.DeclineInvitationsInput) (*guardduty.DeclineInvitationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeclineInvitationsWithContext(ctx, input)
}

func (a *Activities) DeleteDetector(ctx context.Context, input *guardduty.DeleteDetectorInput) (*guardduty.DeleteDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDetectorWithContext(ctx, input)
}

func (a *Activities) DeleteFilter(ctx context.Context, input *guardduty.DeleteFilterInput) (*guardduty.DeleteFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteFilterWithContext(ctx, input)
}

func (a *Activities) DeleteIPSet(ctx context.Context, input *guardduty.DeleteIPSetInput) (*guardduty.DeleteIPSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteIPSetWithContext(ctx, input)
}

func (a *Activities) DeleteInvitations(ctx context.Context, input *guardduty.DeleteInvitationsInput) (*guardduty.DeleteInvitationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteInvitationsWithContext(ctx, input)
}

func (a *Activities) DeleteMembers(ctx context.Context, input *guardduty.DeleteMembersInput) (*guardduty.DeleteMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteMembersWithContext(ctx, input)
}

func (a *Activities) DeletePublishingDestination(ctx context.Context, input *guardduty.DeletePublishingDestinationInput) (*guardduty.DeletePublishingDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePublishingDestinationWithContext(ctx, input)
}

func (a *Activities) DeleteThreatIntelSet(ctx context.Context, input *guardduty.DeleteThreatIntelSetInput) (*guardduty.DeleteThreatIntelSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteThreatIntelSetWithContext(ctx, input)
}

func (a *Activities) DescribeOrganizationConfiguration(ctx context.Context, input *guardduty.DescribeOrganizationConfigurationInput) (*guardduty.DescribeOrganizationConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeOrganizationConfigurationWithContext(ctx, input)
}

func (a *Activities) DescribePublishingDestination(ctx context.Context, input *guardduty.DescribePublishingDestinationInput) (*guardduty.DescribePublishingDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePublishingDestinationWithContext(ctx, input)
}

func (a *Activities) DisableOrganizationAdminAccount(ctx context.Context, input *guardduty.DisableOrganizationAdminAccountInput) (*guardduty.DisableOrganizationAdminAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisableOrganizationAdminAccountWithContext(ctx, input)
}

func (a *Activities) DisassociateFromMasterAccount(ctx context.Context, input *guardduty.DisassociateFromMasterAccountInput) (*guardduty.DisassociateFromMasterAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateFromMasterAccountWithContext(ctx, input)
}

func (a *Activities) DisassociateMembers(ctx context.Context, input *guardduty.DisassociateMembersInput) (*guardduty.DisassociateMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateMembersWithContext(ctx, input)
}

func (a *Activities) EnableOrganizationAdminAccount(ctx context.Context, input *guardduty.EnableOrganizationAdminAccountInput) (*guardduty.EnableOrganizationAdminAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.EnableOrganizationAdminAccountWithContext(ctx, input)
}

func (a *Activities) GetDetector(ctx context.Context, input *guardduty.GetDetectorInput) (*guardduty.GetDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDetectorWithContext(ctx, input)
}

func (a *Activities) GetFilter(ctx context.Context, input *guardduty.GetFilterInput) (*guardduty.GetFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFilterWithContext(ctx, input)
}

func (a *Activities) GetFindings(ctx context.Context, input *guardduty.GetFindingsInput) (*guardduty.GetFindingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFindingsWithContext(ctx, input)
}

func (a *Activities) GetFindingsStatistics(ctx context.Context, input *guardduty.GetFindingsStatisticsInput) (*guardduty.GetFindingsStatisticsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFindingsStatisticsWithContext(ctx, input)
}

func (a *Activities) GetIPSet(ctx context.Context, input *guardduty.GetIPSetInput) (*guardduty.GetIPSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetIPSetWithContext(ctx, input)
}

func (a *Activities) GetInvitationsCount(ctx context.Context, input *guardduty.GetInvitationsCountInput) (*guardduty.GetInvitationsCountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInvitationsCountWithContext(ctx, input)
}

func (a *Activities) GetMasterAccount(ctx context.Context, input *guardduty.GetMasterAccountInput) (*guardduty.GetMasterAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMasterAccountWithContext(ctx, input)
}

func (a *Activities) GetMemberDetectors(ctx context.Context, input *guardduty.GetMemberDetectorsInput) (*guardduty.GetMemberDetectorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMemberDetectorsWithContext(ctx, input)
}

func (a *Activities) GetMembers(ctx context.Context, input *guardduty.GetMembersInput) (*guardduty.GetMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMembersWithContext(ctx, input)
}

func (a *Activities) GetThreatIntelSet(ctx context.Context, input *guardduty.GetThreatIntelSetInput) (*guardduty.GetThreatIntelSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetThreatIntelSetWithContext(ctx, input)
}

func (a *Activities) GetUsageStatistics(ctx context.Context, input *guardduty.GetUsageStatisticsInput) (*guardduty.GetUsageStatisticsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetUsageStatisticsWithContext(ctx, input)
}

func (a *Activities) InviteMembers(ctx context.Context, input *guardduty.InviteMembersInput) (*guardduty.InviteMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.InviteMembersWithContext(ctx, input)
}

func (a *Activities) ListDetectors(ctx context.Context, input *guardduty.ListDetectorsInput) (*guardduty.ListDetectorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDetectorsWithContext(ctx, input)
}

func (a *Activities) ListFilters(ctx context.Context, input *guardduty.ListFiltersInput) (*guardduty.ListFiltersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFiltersWithContext(ctx, input)
}

func (a *Activities) ListFindings(ctx context.Context, input *guardduty.ListFindingsInput) (*guardduty.ListFindingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFindingsWithContext(ctx, input)
}

func (a *Activities) ListIPSets(ctx context.Context, input *guardduty.ListIPSetsInput) (*guardduty.ListIPSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListIPSetsWithContext(ctx, input)
}

func (a *Activities) ListInvitations(ctx context.Context, input *guardduty.ListInvitationsInput) (*guardduty.ListInvitationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListInvitationsWithContext(ctx, input)
}

func (a *Activities) ListMembers(ctx context.Context, input *guardduty.ListMembersInput) (*guardduty.ListMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListMembersWithContext(ctx, input)
}

func (a *Activities) ListOrganizationAdminAccounts(ctx context.Context, input *guardduty.ListOrganizationAdminAccountsInput) (*guardduty.ListOrganizationAdminAccountsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListOrganizationAdminAccountsWithContext(ctx, input)
}

func (a *Activities) ListPublishingDestinations(ctx context.Context, input *guardduty.ListPublishingDestinationsInput) (*guardduty.ListPublishingDestinationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPublishingDestinationsWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *guardduty.ListTagsForResourceInput) (*guardduty.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) ListThreatIntelSets(ctx context.Context, input *guardduty.ListThreatIntelSetsInput) (*guardduty.ListThreatIntelSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListThreatIntelSetsWithContext(ctx, input)
}

func (a *Activities) StartMonitoringMembers(ctx context.Context, input *guardduty.StartMonitoringMembersInput) (*guardduty.StartMonitoringMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartMonitoringMembersWithContext(ctx, input)
}

func (a *Activities) StopMonitoringMembers(ctx context.Context, input *guardduty.StopMonitoringMembersInput) (*guardduty.StopMonitoringMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopMonitoringMembersWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *guardduty.TagResourceInput) (*guardduty.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UnarchiveFindings(ctx context.Context, input *guardduty.UnarchiveFindingsInput) (*guardduty.UnarchiveFindingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UnarchiveFindingsWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *guardduty.UntagResourceInput) (*guardduty.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *Activities) UpdateDetector(ctx context.Context, input *guardduty.UpdateDetectorInput) (*guardduty.UpdateDetectorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDetectorWithContext(ctx, input)
}

func (a *Activities) UpdateFilter(ctx context.Context, input *guardduty.UpdateFilterInput) (*guardduty.UpdateFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateFilterWithContext(ctx, input)
}

func (a *Activities) UpdateFindingsFeedback(ctx context.Context, input *guardduty.UpdateFindingsFeedbackInput) (*guardduty.UpdateFindingsFeedbackOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateFindingsFeedbackWithContext(ctx, input)
}

func (a *Activities) UpdateIPSet(ctx context.Context, input *guardduty.UpdateIPSetInput) (*guardduty.UpdateIPSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateIPSetWithContext(ctx, input)
}

func (a *Activities) UpdateMemberDetectors(ctx context.Context, input *guardduty.UpdateMemberDetectorsInput) (*guardduty.UpdateMemberDetectorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateMemberDetectorsWithContext(ctx, input)
}

func (a *Activities) UpdateOrganizationConfiguration(ctx context.Context, input *guardduty.UpdateOrganizationConfigurationInput) (*guardduty.UpdateOrganizationConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateOrganizationConfigurationWithContext(ctx, input)
}

func (a *Activities) UpdatePublishingDestination(ctx context.Context, input *guardduty.UpdatePublishingDestinationInput) (*guardduty.UpdatePublishingDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdatePublishingDestinationWithContext(ctx, input)
}

func (a *Activities) UpdateThreatIntelSet(ctx context.Context, input *guardduty.UpdateThreatIntelSetInput) (*guardduty.UpdateThreatIntelSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateThreatIntelSetWithContext(ctx, input)
}
