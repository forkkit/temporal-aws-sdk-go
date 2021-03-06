// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package kinesisstub

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddTagsToStream(ctx workflow.Context, input *kinesis.AddTagsToStreamInput) (*kinesis.AddTagsToStreamOutput, error)
	AddTagsToStreamAsync(ctx workflow.Context, input *kinesis.AddTagsToStreamInput) *KinesisAddTagsToStreamFuture

	CreateStream(ctx workflow.Context, input *kinesis.CreateStreamInput) (*kinesis.CreateStreamOutput, error)
	CreateStreamAsync(ctx workflow.Context, input *kinesis.CreateStreamInput) *KinesisCreateStreamFuture

	DecreaseStreamRetentionPeriod(ctx workflow.Context, input *kinesis.DecreaseStreamRetentionPeriodInput) (*kinesis.DecreaseStreamRetentionPeriodOutput, error)
	DecreaseStreamRetentionPeriodAsync(ctx workflow.Context, input *kinesis.DecreaseStreamRetentionPeriodInput) *KinesisDecreaseStreamRetentionPeriodFuture

	DeleteStream(ctx workflow.Context, input *kinesis.DeleteStreamInput) (*kinesis.DeleteStreamOutput, error)
	DeleteStreamAsync(ctx workflow.Context, input *kinesis.DeleteStreamInput) *KinesisDeleteStreamFuture

	DeregisterStreamConsumer(ctx workflow.Context, input *kinesis.DeregisterStreamConsumerInput) (*kinesis.DeregisterStreamConsumerOutput, error)
	DeregisterStreamConsumerAsync(ctx workflow.Context, input *kinesis.DeregisterStreamConsumerInput) *KinesisDeregisterStreamConsumerFuture

	DescribeLimits(ctx workflow.Context, input *kinesis.DescribeLimitsInput) (*kinesis.DescribeLimitsOutput, error)
	DescribeLimitsAsync(ctx workflow.Context, input *kinesis.DescribeLimitsInput) *KinesisDescribeLimitsFuture

	DescribeStream(ctx workflow.Context, input *kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error)
	DescribeStreamAsync(ctx workflow.Context, input *kinesis.DescribeStreamInput) *KinesisDescribeStreamFuture

	DescribeStreamConsumer(ctx workflow.Context, input *kinesis.DescribeStreamConsumerInput) (*kinesis.DescribeStreamConsumerOutput, error)
	DescribeStreamConsumerAsync(ctx workflow.Context, input *kinesis.DescribeStreamConsumerInput) *KinesisDescribeStreamConsumerFuture

	DescribeStreamSummary(ctx workflow.Context, input *kinesis.DescribeStreamSummaryInput) (*kinesis.DescribeStreamSummaryOutput, error)
	DescribeStreamSummaryAsync(ctx workflow.Context, input *kinesis.DescribeStreamSummaryInput) *KinesisDescribeStreamSummaryFuture

	DisableEnhancedMonitoring(ctx workflow.Context, input *kinesis.DisableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error)
	DisableEnhancedMonitoringAsync(ctx workflow.Context, input *kinesis.DisableEnhancedMonitoringInput) *KinesisDisableEnhancedMonitoringFuture

	EnableEnhancedMonitoring(ctx workflow.Context, input *kinesis.EnableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error)
	EnableEnhancedMonitoringAsync(ctx workflow.Context, input *kinesis.EnableEnhancedMonitoringInput) *KinesisEnableEnhancedMonitoringFuture

	GetRecords(ctx workflow.Context, input *kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error)
	GetRecordsAsync(ctx workflow.Context, input *kinesis.GetRecordsInput) *KinesisGetRecordsFuture

	GetShardIterator(ctx workflow.Context, input *kinesis.GetShardIteratorInput) (*kinesis.GetShardIteratorOutput, error)
	GetShardIteratorAsync(ctx workflow.Context, input *kinesis.GetShardIteratorInput) *KinesisGetShardIteratorFuture

	IncreaseStreamRetentionPeriod(ctx workflow.Context, input *kinesis.IncreaseStreamRetentionPeriodInput) (*kinesis.IncreaseStreamRetentionPeriodOutput, error)
	IncreaseStreamRetentionPeriodAsync(ctx workflow.Context, input *kinesis.IncreaseStreamRetentionPeriodInput) *KinesisIncreaseStreamRetentionPeriodFuture

	ListShards(ctx workflow.Context, input *kinesis.ListShardsInput) (*kinesis.ListShardsOutput, error)
	ListShardsAsync(ctx workflow.Context, input *kinesis.ListShardsInput) *KinesisListShardsFuture

	ListStreamConsumers(ctx workflow.Context, input *kinesis.ListStreamConsumersInput) (*kinesis.ListStreamConsumersOutput, error)
	ListStreamConsumersAsync(ctx workflow.Context, input *kinesis.ListStreamConsumersInput) *KinesisListStreamConsumersFuture

	ListStreams(ctx workflow.Context, input *kinesis.ListStreamsInput) (*kinesis.ListStreamsOutput, error)
	ListStreamsAsync(ctx workflow.Context, input *kinesis.ListStreamsInput) *KinesisListStreamsFuture

	ListTagsForStream(ctx workflow.Context, input *kinesis.ListTagsForStreamInput) (*kinesis.ListTagsForStreamOutput, error)
	ListTagsForStreamAsync(ctx workflow.Context, input *kinesis.ListTagsForStreamInput) *KinesisListTagsForStreamFuture

	MergeShards(ctx workflow.Context, input *kinesis.MergeShardsInput) (*kinesis.MergeShardsOutput, error)
	MergeShardsAsync(ctx workflow.Context, input *kinesis.MergeShardsInput) *KinesisMergeShardsFuture

	PutRecord(ctx workflow.Context, input *kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error)
	PutRecordAsync(ctx workflow.Context, input *kinesis.PutRecordInput) *KinesisPutRecordFuture

	PutRecords(ctx workflow.Context, input *kinesis.PutRecordsInput) (*kinesis.PutRecordsOutput, error)
	PutRecordsAsync(ctx workflow.Context, input *kinesis.PutRecordsInput) *KinesisPutRecordsFuture

	RegisterStreamConsumer(ctx workflow.Context, input *kinesis.RegisterStreamConsumerInput) (*kinesis.RegisterStreamConsumerOutput, error)
	RegisterStreamConsumerAsync(ctx workflow.Context, input *kinesis.RegisterStreamConsumerInput) *KinesisRegisterStreamConsumerFuture

	RemoveTagsFromStream(ctx workflow.Context, input *kinesis.RemoveTagsFromStreamInput) (*kinesis.RemoveTagsFromStreamOutput, error)
	RemoveTagsFromStreamAsync(ctx workflow.Context, input *kinesis.RemoveTagsFromStreamInput) *KinesisRemoveTagsFromStreamFuture

	SplitShard(ctx workflow.Context, input *kinesis.SplitShardInput) (*kinesis.SplitShardOutput, error)
	SplitShardAsync(ctx workflow.Context, input *kinesis.SplitShardInput) *KinesisSplitShardFuture

	StartStreamEncryption(ctx workflow.Context, input *kinesis.StartStreamEncryptionInput) (*kinesis.StartStreamEncryptionOutput, error)
	StartStreamEncryptionAsync(ctx workflow.Context, input *kinesis.StartStreamEncryptionInput) *KinesisStartStreamEncryptionFuture

	StopStreamEncryption(ctx workflow.Context, input *kinesis.StopStreamEncryptionInput) (*kinesis.StopStreamEncryptionOutput, error)
	StopStreamEncryptionAsync(ctx workflow.Context, input *kinesis.StopStreamEncryptionInput) *KinesisStopStreamEncryptionFuture

	SubscribeToShard(ctx workflow.Context, input *kinesis.SubscribeToShardInput) (*kinesis.SubscribeToShardOutput, error)
	SubscribeToShardAsync(ctx workflow.Context, input *kinesis.SubscribeToShardInput) *KinesisSubscribeToShardFuture

	UpdateShardCount(ctx workflow.Context, input *kinesis.UpdateShardCountInput) (*kinesis.UpdateShardCountOutput, error)
	UpdateShardCountAsync(ctx workflow.Context, input *kinesis.UpdateShardCountInput) *KinesisUpdateShardCountFuture

	WaitUntilStreamExists(ctx workflow.Context, input *kinesis.DescribeStreamInput) error
	WaitUntilStreamExistsAsync(ctx workflow.Context, input *kinesis.DescribeStreamInput) *clients.VoidFuture

	WaitUntilStreamNotExists(ctx workflow.Context, input *kinesis.DescribeStreamInput) error
	WaitUntilStreamNotExistsAsync(ctx workflow.Context, input *kinesis.DescribeStreamInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
