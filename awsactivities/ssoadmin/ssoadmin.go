// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ssoadmin

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssoadmin"
	"github.com/aws/aws-sdk-go/service/ssoadmin/ssoadminiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client ssoadminiface.SSOAdminAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := ssoadmin.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (ssoadminiface.SSOAdminAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return ssoadmin.New(sess), nil
}

func (a *Activities) AttachManagedPolicyToPermissionSet(ctx context.Context, input *ssoadmin.AttachManagedPolicyToPermissionSetInput) (*ssoadmin.AttachManagedPolicyToPermissionSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AttachManagedPolicyToPermissionSetWithContext(ctx, input)
}

func (a *Activities) CreateAccountAssignment(ctx context.Context, input *ssoadmin.CreateAccountAssignmentInput) (*ssoadmin.CreateAccountAssignmentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateAccountAssignmentWithContext(ctx, input)
}

func (a *Activities) CreatePermissionSet(ctx context.Context, input *ssoadmin.CreatePermissionSetInput) (*ssoadmin.CreatePermissionSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePermissionSetWithContext(ctx, input)
}

func (a *Activities) DeleteAccountAssignment(ctx context.Context, input *ssoadmin.DeleteAccountAssignmentInput) (*ssoadmin.DeleteAccountAssignmentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAccountAssignmentWithContext(ctx, input)
}

func (a *Activities) DeleteInlinePolicyFromPermissionSet(ctx context.Context, input *ssoadmin.DeleteInlinePolicyFromPermissionSetInput) (*ssoadmin.DeleteInlinePolicyFromPermissionSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteInlinePolicyFromPermissionSetWithContext(ctx, input)
}

func (a *Activities) DeletePermissionSet(ctx context.Context, input *ssoadmin.DeletePermissionSetInput) (*ssoadmin.DeletePermissionSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePermissionSetWithContext(ctx, input)
}

func (a *Activities) DescribeAccountAssignmentCreationStatus(ctx context.Context, input *ssoadmin.DescribeAccountAssignmentCreationStatusInput) (*ssoadmin.DescribeAccountAssignmentCreationStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAccountAssignmentCreationStatusWithContext(ctx, input)
}

func (a *Activities) DescribeAccountAssignmentDeletionStatus(ctx context.Context, input *ssoadmin.DescribeAccountAssignmentDeletionStatusInput) (*ssoadmin.DescribeAccountAssignmentDeletionStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAccountAssignmentDeletionStatusWithContext(ctx, input)
}

func (a *Activities) DescribePermissionSet(ctx context.Context, input *ssoadmin.DescribePermissionSetInput) (*ssoadmin.DescribePermissionSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePermissionSetWithContext(ctx, input)
}

func (a *Activities) DescribePermissionSetProvisioningStatus(ctx context.Context, input *ssoadmin.DescribePermissionSetProvisioningStatusInput) (*ssoadmin.DescribePermissionSetProvisioningStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePermissionSetProvisioningStatusWithContext(ctx, input)
}

func (a *Activities) DetachManagedPolicyFromPermissionSet(ctx context.Context, input *ssoadmin.DetachManagedPolicyFromPermissionSetInput) (*ssoadmin.DetachManagedPolicyFromPermissionSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetachManagedPolicyFromPermissionSetWithContext(ctx, input)
}

func (a *Activities) GetInlinePolicyForPermissionSet(ctx context.Context, input *ssoadmin.GetInlinePolicyForPermissionSetInput) (*ssoadmin.GetInlinePolicyForPermissionSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInlinePolicyForPermissionSetWithContext(ctx, input)
}

func (a *Activities) ListAccountAssignmentCreationStatus(ctx context.Context, input *ssoadmin.ListAccountAssignmentCreationStatusInput) (*ssoadmin.ListAccountAssignmentCreationStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAccountAssignmentCreationStatusWithContext(ctx, input)
}

func (a *Activities) ListAccountAssignmentDeletionStatus(ctx context.Context, input *ssoadmin.ListAccountAssignmentDeletionStatusInput) (*ssoadmin.ListAccountAssignmentDeletionStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAccountAssignmentDeletionStatusWithContext(ctx, input)
}

func (a *Activities) ListAccountAssignments(ctx context.Context, input *ssoadmin.ListAccountAssignmentsInput) (*ssoadmin.ListAccountAssignmentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAccountAssignmentsWithContext(ctx, input)
}

func (a *Activities) ListAccountsForProvisionedPermissionSet(ctx context.Context, input *ssoadmin.ListAccountsForProvisionedPermissionSetInput) (*ssoadmin.ListAccountsForProvisionedPermissionSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAccountsForProvisionedPermissionSetWithContext(ctx, input)
}

func (a *Activities) ListInstances(ctx context.Context, input *ssoadmin.ListInstancesInput) (*ssoadmin.ListInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListInstancesWithContext(ctx, input)
}

func (a *Activities) ListManagedPoliciesInPermissionSet(ctx context.Context, input *ssoadmin.ListManagedPoliciesInPermissionSetInput) (*ssoadmin.ListManagedPoliciesInPermissionSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListManagedPoliciesInPermissionSetWithContext(ctx, input)
}

func (a *Activities) ListPermissionSetProvisioningStatus(ctx context.Context, input *ssoadmin.ListPermissionSetProvisioningStatusInput) (*ssoadmin.ListPermissionSetProvisioningStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPermissionSetProvisioningStatusWithContext(ctx, input)
}

func (a *Activities) ListPermissionSets(ctx context.Context, input *ssoadmin.ListPermissionSetsInput) (*ssoadmin.ListPermissionSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPermissionSetsWithContext(ctx, input)
}

func (a *Activities) ListPermissionSetsProvisionedToAccount(ctx context.Context, input *ssoadmin.ListPermissionSetsProvisionedToAccountInput) (*ssoadmin.ListPermissionSetsProvisionedToAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPermissionSetsProvisionedToAccountWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *ssoadmin.ListTagsForResourceInput) (*ssoadmin.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) ProvisionPermissionSet(ctx context.Context, input *ssoadmin.ProvisionPermissionSetInput) (*ssoadmin.ProvisionPermissionSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ProvisionPermissionSetWithContext(ctx, input)
}

func (a *Activities) PutInlinePolicyToPermissionSet(ctx context.Context, input *ssoadmin.PutInlinePolicyToPermissionSetInput) (*ssoadmin.PutInlinePolicyToPermissionSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutInlinePolicyToPermissionSetWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *ssoadmin.TagResourceInput) (*ssoadmin.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *ssoadmin.UntagResourceInput) (*ssoadmin.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *Activities) UpdatePermissionSet(ctx context.Context, input *ssoadmin.UpdatePermissionSetInput) (*ssoadmin.UpdatePermissionSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdatePermissionSetWithContext(ctx, input)
}
