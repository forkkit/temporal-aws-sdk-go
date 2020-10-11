// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package dax

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dax"
	"github.com/aws/aws-sdk-go/service/dax/daxiface"
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
	client daxiface.DAXAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := dax.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (daxiface.DAXAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return dax.New(sess), nil
}

func (a *Activities) CreateCluster(ctx context.Context, input *dax.CreateClusterInput) (*dax.CreateClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateClusterWithContext(ctx, input)
}

func (a *Activities) CreateParameterGroup(ctx context.Context, input *dax.CreateParameterGroupInput) (*dax.CreateParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateParameterGroupWithContext(ctx, input)
}

func (a *Activities) CreateSubnetGroup(ctx context.Context, input *dax.CreateSubnetGroupInput) (*dax.CreateSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSubnetGroupWithContext(ctx, input)
}

func (a *Activities) DecreaseReplicationFactor(ctx context.Context, input *dax.DecreaseReplicationFactorInput) (*dax.DecreaseReplicationFactorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DecreaseReplicationFactorWithContext(ctx, input)
}

func (a *Activities) DeleteCluster(ctx context.Context, input *dax.DeleteClusterInput) (*dax.DeleteClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteClusterWithContext(ctx, input)
}

func (a *Activities) DeleteParameterGroup(ctx context.Context, input *dax.DeleteParameterGroupInput) (*dax.DeleteParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteParameterGroupWithContext(ctx, input)
}

func (a *Activities) DeleteSubnetGroup(ctx context.Context, input *dax.DeleteSubnetGroupInput) (*dax.DeleteSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSubnetGroupWithContext(ctx, input)
}

func (a *Activities) DescribeClusters(ctx context.Context, input *dax.DescribeClustersInput) (*dax.DescribeClustersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeClustersWithContext(ctx, input)
}

func (a *Activities) DescribeDefaultParameters(ctx context.Context, input *dax.DescribeDefaultParametersInput) (*dax.DescribeDefaultParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDefaultParametersWithContext(ctx, input)
}

func (a *Activities) DescribeEvents(ctx context.Context, input *dax.DescribeEventsInput) (*dax.DescribeEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventsWithContext(ctx, input)
}

func (a *Activities) DescribeParameterGroups(ctx context.Context, input *dax.DescribeParameterGroupsInput) (*dax.DescribeParameterGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeParameterGroupsWithContext(ctx, input)
}

func (a *Activities) DescribeParameters(ctx context.Context, input *dax.DescribeParametersInput) (*dax.DescribeParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeParametersWithContext(ctx, input)
}

func (a *Activities) DescribeSubnetGroups(ctx context.Context, input *dax.DescribeSubnetGroupsInput) (*dax.DescribeSubnetGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSubnetGroupsWithContext(ctx, input)
}

func (a *Activities) IncreaseReplicationFactor(ctx context.Context, input *dax.IncreaseReplicationFactorInput) (*dax.IncreaseReplicationFactorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.IncreaseReplicationFactorWithContext(ctx, input)
}

func (a *Activities) ListTags(ctx context.Context, input *dax.ListTagsInput) (*dax.ListTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsWithContext(ctx, input)
}

func (a *Activities) RebootNode(ctx context.Context, input *dax.RebootNodeInput) (*dax.RebootNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RebootNodeWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *dax.TagResourceInput) (*dax.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *dax.UntagResourceInput) (*dax.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *Activities) UpdateCluster(ctx context.Context, input *dax.UpdateClusterInput) (*dax.UpdateClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateClusterWithContext(ctx, input)
}

func (a *Activities) UpdateParameterGroup(ctx context.Context, input *dax.UpdateParameterGroupInput) (*dax.UpdateParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateParameterGroupWithContext(ctx, input)
}

func (a *Activities) UpdateSubnetGroup(ctx context.Context, input *dax.UpdateSubnetGroupInput) (*dax.UpdateSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateSubnetGroupWithContext(ctx, input)
}
