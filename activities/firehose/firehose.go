// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package firehose

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/firehose/firehoseiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client firehoseiface.FirehoseAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := firehose.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (firehoseiface.FirehoseAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return firehose.New(sess), nil
}

func (a *Activities) CreateDeliveryStream(ctx context.Context, input *firehose.CreateDeliveryStreamInput) (*firehose.CreateDeliveryStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDeliveryStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDeliveryStream(ctx context.Context, input *firehose.DeleteDeliveryStreamInput) (*firehose.DeleteDeliveryStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDeliveryStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDeliveryStream(ctx context.Context, input *firehose.DescribeDeliveryStreamInput) (*firehose.DescribeDeliveryStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDeliveryStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListDeliveryStreams(ctx context.Context, input *firehose.ListDeliveryStreamsInput) (*firehose.ListDeliveryStreamsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListDeliveryStreamsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForDeliveryStream(ctx context.Context, input *firehose.ListTagsForDeliveryStreamInput) (*firehose.ListTagsForDeliveryStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForDeliveryStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutRecord(ctx context.Context, input *firehose.PutRecordInput) (*firehose.PutRecordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutRecordWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) PutRecordBatch(ctx context.Context, input *firehose.PutRecordBatchInput) (*firehose.PutRecordBatchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.PutRecordBatchWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartDeliveryStreamEncryption(ctx context.Context, input *firehose.StartDeliveryStreamEncryptionInput) (*firehose.StartDeliveryStreamEncryptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartDeliveryStreamEncryptionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopDeliveryStreamEncryption(ctx context.Context, input *firehose.StopDeliveryStreamEncryptionInput) (*firehose.StopDeliveryStreamEncryptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopDeliveryStreamEncryptionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagDeliveryStream(ctx context.Context, input *firehose.TagDeliveryStreamInput) (*firehose.TagDeliveryStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagDeliveryStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagDeliveryStream(ctx context.Context, input *firehose.UntagDeliveryStreamInput) (*firehose.UntagDeliveryStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagDeliveryStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateDestination(ctx context.Context, input *firehose.UpdateDestinationInput) (*firehose.UpdateDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateDestinationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}