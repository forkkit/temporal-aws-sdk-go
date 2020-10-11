// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package iotanalytics

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotanalytics"
	"github.com/aws/aws-sdk-go/service/iotanalytics/iotanalyticsiface"
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
	client iotanalyticsiface.IoTAnalyticsAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := iotanalytics.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (iotanalyticsiface.IoTAnalyticsAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return iotanalytics.New(sess), nil
}

func (a *Activities) BatchPutMessage(ctx context.Context, input *iotanalytics.BatchPutMessageInput) (*iotanalytics.BatchPutMessageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchPutMessageWithContext(ctx, input)
}

func (a *Activities) CancelPipelineReprocessing(ctx context.Context, input *iotanalytics.CancelPipelineReprocessingInput) (*iotanalytics.CancelPipelineReprocessingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CancelPipelineReprocessingWithContext(ctx, input)
}

func (a *Activities) CreateChannel(ctx context.Context, input *iotanalytics.CreateChannelInput) (*iotanalytics.CreateChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateChannelWithContext(ctx, input)
}

func (a *Activities) CreateDataset(ctx context.Context, input *iotanalytics.CreateDatasetInput) (*iotanalytics.CreateDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDatasetWithContext(ctx, input)
}

func (a *Activities) CreateDatasetContent(ctx context.Context, input *iotanalytics.CreateDatasetContentInput) (*iotanalytics.CreateDatasetContentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDatasetContentWithContext(ctx, input)
}

func (a *Activities) CreateDatastore(ctx context.Context, input *iotanalytics.CreateDatastoreInput) (*iotanalytics.CreateDatastoreOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDatastoreWithContext(ctx, input)
}

func (a *Activities) CreatePipeline(ctx context.Context, input *iotanalytics.CreatePipelineInput) (*iotanalytics.CreatePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePipelineWithContext(ctx, input)
}

func (a *Activities) DeleteChannel(ctx context.Context, input *iotanalytics.DeleteChannelInput) (*iotanalytics.DeleteChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteChannelWithContext(ctx, input)
}

func (a *Activities) DeleteDataset(ctx context.Context, input *iotanalytics.DeleteDatasetInput) (*iotanalytics.DeleteDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDatasetWithContext(ctx, input)
}

func (a *Activities) DeleteDatasetContent(ctx context.Context, input *iotanalytics.DeleteDatasetContentInput) (*iotanalytics.DeleteDatasetContentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDatasetContentWithContext(ctx, input)
}

func (a *Activities) DeleteDatastore(ctx context.Context, input *iotanalytics.DeleteDatastoreInput) (*iotanalytics.DeleteDatastoreOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDatastoreWithContext(ctx, input)
}

func (a *Activities) DeletePipeline(ctx context.Context, input *iotanalytics.DeletePipelineInput) (*iotanalytics.DeletePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePipelineWithContext(ctx, input)
}

func (a *Activities) DescribeChannel(ctx context.Context, input *iotanalytics.DescribeChannelInput) (*iotanalytics.DescribeChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeChannelWithContext(ctx, input)
}

func (a *Activities) DescribeDataset(ctx context.Context, input *iotanalytics.DescribeDatasetInput) (*iotanalytics.DescribeDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDatasetWithContext(ctx, input)
}

func (a *Activities) DescribeDatastore(ctx context.Context, input *iotanalytics.DescribeDatastoreInput) (*iotanalytics.DescribeDatastoreOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDatastoreWithContext(ctx, input)
}

func (a *Activities) DescribeLoggingOptions(ctx context.Context, input *iotanalytics.DescribeLoggingOptionsInput) (*iotanalytics.DescribeLoggingOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeLoggingOptionsWithContext(ctx, input)
}

func (a *Activities) DescribePipeline(ctx context.Context, input *iotanalytics.DescribePipelineInput) (*iotanalytics.DescribePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePipelineWithContext(ctx, input)
}

func (a *Activities) GetDatasetContent(ctx context.Context, input *iotanalytics.GetDatasetContentInput) (*iotanalytics.GetDatasetContentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDatasetContentWithContext(ctx, input)
}

func (a *Activities) ListChannels(ctx context.Context, input *iotanalytics.ListChannelsInput) (*iotanalytics.ListChannelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListChannelsWithContext(ctx, input)
}

func (a *Activities) ListDatasetContents(ctx context.Context, input *iotanalytics.ListDatasetContentsInput) (*iotanalytics.ListDatasetContentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDatasetContentsWithContext(ctx, input)
}

func (a *Activities) ListDatasets(ctx context.Context, input *iotanalytics.ListDatasetsInput) (*iotanalytics.ListDatasetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDatasetsWithContext(ctx, input)
}

func (a *Activities) ListDatastores(ctx context.Context, input *iotanalytics.ListDatastoresInput) (*iotanalytics.ListDatastoresOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDatastoresWithContext(ctx, input)
}

func (a *Activities) ListPipelines(ctx context.Context, input *iotanalytics.ListPipelinesInput) (*iotanalytics.ListPipelinesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPipelinesWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *iotanalytics.ListTagsForResourceInput) (*iotanalytics.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) PutLoggingOptions(ctx context.Context, input *iotanalytics.PutLoggingOptionsInput) (*iotanalytics.PutLoggingOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutLoggingOptionsWithContext(ctx, input)
}

func (a *Activities) RunPipelineActivity(ctx context.Context, input *iotanalytics.RunPipelineActivityInput) (*iotanalytics.RunPipelineActivityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RunPipelineActivityWithContext(ctx, input)
}

func (a *Activities) SampleChannelData(ctx context.Context, input *iotanalytics.SampleChannelDataInput) (*iotanalytics.SampleChannelDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SampleChannelDataWithContext(ctx, input)
}

func (a *Activities) StartPipelineReprocessing(ctx context.Context, input *iotanalytics.StartPipelineReprocessingInput) (*iotanalytics.StartPipelineReprocessingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartPipelineReprocessingWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *iotanalytics.TagResourceInput) (*iotanalytics.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *iotanalytics.UntagResourceInput) (*iotanalytics.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *Activities) UpdateChannel(ctx context.Context, input *iotanalytics.UpdateChannelInput) (*iotanalytics.UpdateChannelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateChannelWithContext(ctx, input)
}

func (a *Activities) UpdateDataset(ctx context.Context, input *iotanalytics.UpdateDatasetInput) (*iotanalytics.UpdateDatasetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDatasetWithContext(ctx, input)
}

func (a *Activities) UpdateDatastore(ctx context.Context, input *iotanalytics.UpdateDatastoreInput) (*iotanalytics.UpdateDatastoreOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDatastoreWithContext(ctx, input)
}

func (a *Activities) UpdatePipeline(ctx context.Context, input *iotanalytics.UpdatePipelineInput) (*iotanalytics.UpdatePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdatePipelineWithContext(ctx, input)
}
