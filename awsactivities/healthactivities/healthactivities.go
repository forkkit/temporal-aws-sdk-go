// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package healthactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/health"
	"github.com/aws/aws-sdk-go/service/health/healthiface"
	"go.temporal.io/aws-sdk/internal"
	"go.temporal.io/aws-sdk/sessionfactory"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type Activities struct {
	client healthiface.HealthAPI

	sessionFactory sessionfactory.SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := health.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory sessionfactory.SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (healthiface.HealthAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return health.New(sess), nil
}

func (a *Activities) DescribeAffectedAccountsForOrganization(ctx context.Context, input *health.DescribeAffectedAccountsForOrganizationInput) (*health.DescribeAffectedAccountsForOrganizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAffectedAccountsForOrganizationWithContext(ctx, input)
}

func (a *Activities) DescribeAffectedEntities(ctx context.Context, input *health.DescribeAffectedEntitiesInput) (*health.DescribeAffectedEntitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAffectedEntitiesWithContext(ctx, input)
}

func (a *Activities) DescribeAffectedEntitiesForOrganization(ctx context.Context, input *health.DescribeAffectedEntitiesForOrganizationInput) (*health.DescribeAffectedEntitiesForOrganizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAffectedEntitiesForOrganizationWithContext(ctx, input)
}

func (a *Activities) DescribeEntityAggregates(ctx context.Context, input *health.DescribeEntityAggregatesInput) (*health.DescribeEntityAggregatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEntityAggregatesWithContext(ctx, input)
}

func (a *Activities) DescribeEventAggregates(ctx context.Context, input *health.DescribeEventAggregatesInput) (*health.DescribeEventAggregatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventAggregatesWithContext(ctx, input)
}

func (a *Activities) DescribeEventDetails(ctx context.Context, input *health.DescribeEventDetailsInput) (*health.DescribeEventDetailsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventDetailsWithContext(ctx, input)
}

func (a *Activities) DescribeEventDetailsForOrganization(ctx context.Context, input *health.DescribeEventDetailsForOrganizationInput) (*health.DescribeEventDetailsForOrganizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventDetailsForOrganizationWithContext(ctx, input)
}

func (a *Activities) DescribeEventTypes(ctx context.Context, input *health.DescribeEventTypesInput) (*health.DescribeEventTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventTypesWithContext(ctx, input)
}

func (a *Activities) DescribeEvents(ctx context.Context, input *health.DescribeEventsInput) (*health.DescribeEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventsWithContext(ctx, input)
}

func (a *Activities) DescribeEventsForOrganization(ctx context.Context, input *health.DescribeEventsForOrganizationInput) (*health.DescribeEventsForOrganizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventsForOrganizationWithContext(ctx, input)
}

func (a *Activities) DescribeHealthServiceStatusForOrganization(ctx context.Context, input *health.DescribeHealthServiceStatusForOrganizationInput) (*health.DescribeHealthServiceStatusForOrganizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeHealthServiceStatusForOrganizationWithContext(ctx, input)
}

func (a *Activities) DisableHealthServiceAccessForOrganization(ctx context.Context, input *health.DisableHealthServiceAccessForOrganizationInput) (*health.DisableHealthServiceAccessForOrganizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisableHealthServiceAccessForOrganizationWithContext(ctx, input)
}

func (a *Activities) EnableHealthServiceAccessForOrganization(ctx context.Context, input *health.EnableHealthServiceAccessForOrganizationInput) (*health.EnableHealthServiceAccessForOrganizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.EnableHealthServiceAccessForOrganizationWithContext(ctx, input)
}
