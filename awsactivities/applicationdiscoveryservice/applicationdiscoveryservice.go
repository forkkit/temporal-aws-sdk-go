// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package applicationdiscoveryservice

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice"
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice/applicationdiscoveryserviceiface"
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
	client applicationdiscoveryserviceiface.ApplicationDiscoveryServiceAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := applicationdiscoveryservice.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (applicationdiscoveryserviceiface.ApplicationDiscoveryServiceAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return applicationdiscoveryservice.New(sess), nil
}

func (a *Activities) AssociateConfigurationItemsToApplication(ctx context.Context, input *applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput) (*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateConfigurationItemsToApplicationWithContext(ctx, input)
}

func (a *Activities) BatchDeleteImportData(ctx context.Context, input *applicationdiscoveryservice.BatchDeleteImportDataInput) (*applicationdiscoveryservice.BatchDeleteImportDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchDeleteImportDataWithContext(ctx, input)
}

func (a *Activities) CreateApplication(ctx context.Context, input *applicationdiscoveryservice.CreateApplicationInput) (*applicationdiscoveryservice.CreateApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateApplicationWithContext(ctx, input)
}

func (a *Activities) CreateTags(ctx context.Context, input *applicationdiscoveryservice.CreateTagsInput) (*applicationdiscoveryservice.CreateTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateTagsWithContext(ctx, input)
}

func (a *Activities) DeleteApplications(ctx context.Context, input *applicationdiscoveryservice.DeleteApplicationsInput) (*applicationdiscoveryservice.DeleteApplicationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApplicationsWithContext(ctx, input)
}

func (a *Activities) DeleteTags(ctx context.Context, input *applicationdiscoveryservice.DeleteTagsInput) (*applicationdiscoveryservice.DeleteTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteTagsWithContext(ctx, input)
}

func (a *Activities) DescribeAgents(ctx context.Context, input *applicationdiscoveryservice.DescribeAgentsInput) (*applicationdiscoveryservice.DescribeAgentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAgentsWithContext(ctx, input)
}

func (a *Activities) DescribeConfigurations(ctx context.Context, input *applicationdiscoveryservice.DescribeConfigurationsInput) (*applicationdiscoveryservice.DescribeConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeConfigurationsWithContext(ctx, input)
}

func (a *Activities) DescribeContinuousExports(ctx context.Context, input *applicationdiscoveryservice.DescribeContinuousExportsInput) (*applicationdiscoveryservice.DescribeContinuousExportsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeContinuousExportsWithContext(ctx, input)
}

func (a *Activities) DescribeExportConfigurations(ctx context.Context, input *applicationdiscoveryservice.DescribeExportConfigurationsInput) (*applicationdiscoveryservice.DescribeExportConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeExportConfigurationsWithContext(ctx, input)
}

func (a *Activities) DescribeExportTasks(ctx context.Context, input *applicationdiscoveryservice.DescribeExportTasksInput) (*applicationdiscoveryservice.DescribeExportTasksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeExportTasksWithContext(ctx, input)
}

func (a *Activities) DescribeImportTasks(ctx context.Context, input *applicationdiscoveryservice.DescribeImportTasksInput) (*applicationdiscoveryservice.DescribeImportTasksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeImportTasksWithContext(ctx, input)
}

func (a *Activities) DescribeTags(ctx context.Context, input *applicationdiscoveryservice.DescribeTagsInput) (*applicationdiscoveryservice.DescribeTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTagsWithContext(ctx, input)
}

func (a *Activities) DisassociateConfigurationItemsFromApplication(ctx context.Context, input *applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput) (*applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateConfigurationItemsFromApplicationWithContext(ctx, input)
}

func (a *Activities) ExportConfigurations(ctx context.Context, input *applicationdiscoveryservice.ExportConfigurationsInput) (*applicationdiscoveryservice.ExportConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ExportConfigurationsWithContext(ctx, input)
}

func (a *Activities) GetDiscoverySummary(ctx context.Context, input *applicationdiscoveryservice.GetDiscoverySummaryInput) (*applicationdiscoveryservice.GetDiscoverySummaryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDiscoverySummaryWithContext(ctx, input)
}

func (a *Activities) ListConfigurations(ctx context.Context, input *applicationdiscoveryservice.ListConfigurationsInput) (*applicationdiscoveryservice.ListConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListConfigurationsWithContext(ctx, input)
}

func (a *Activities) ListServerNeighbors(ctx context.Context, input *applicationdiscoveryservice.ListServerNeighborsInput) (*applicationdiscoveryservice.ListServerNeighborsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListServerNeighborsWithContext(ctx, input)
}

func (a *Activities) StartContinuousExport(ctx context.Context, input *applicationdiscoveryservice.StartContinuousExportInput) (*applicationdiscoveryservice.StartContinuousExportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartContinuousExportWithContext(ctx, input)
}

func (a *Activities) StartDataCollectionByAgentIds(ctx context.Context, input *applicationdiscoveryservice.StartDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartDataCollectionByAgentIdsWithContext(ctx, input)
}

func (a *Activities) StartExportTask(ctx context.Context, input *applicationdiscoveryservice.StartExportTaskInput) (*applicationdiscoveryservice.StartExportTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartExportTaskWithContext(ctx, input)
}

func (a *Activities) StartImportTask(ctx context.Context, input *applicationdiscoveryservice.StartImportTaskInput) (*applicationdiscoveryservice.StartImportTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartImportTaskWithContext(ctx, input)
}

func (a *Activities) StopContinuousExport(ctx context.Context, input *applicationdiscoveryservice.StopContinuousExportInput) (*applicationdiscoveryservice.StopContinuousExportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopContinuousExportWithContext(ctx, input)
}

func (a *Activities) StopDataCollectionByAgentIds(ctx context.Context, input *applicationdiscoveryservice.StopDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopDataCollectionByAgentIdsWithContext(ctx, input)
}

func (a *Activities) UpdateApplication(ctx context.Context, input *applicationdiscoveryservice.UpdateApplicationInput) (*applicationdiscoveryservice.UpdateApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateApplicationWithContext(ctx, input)
}
