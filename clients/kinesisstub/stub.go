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

type stub struct{}

type KinesisAddTagsToStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisAddTagsToStreamFuture) Get(ctx workflow.Context) (*kinesis.AddTagsToStreamOutput, error) {
	var output kinesis.AddTagsToStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisCreateStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisCreateStreamFuture) Get(ctx workflow.Context) (*kinesis.CreateStreamOutput, error) {
	var output kinesis.CreateStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisDecreaseStreamRetentionPeriodFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisDecreaseStreamRetentionPeriodFuture) Get(ctx workflow.Context) (*kinesis.DecreaseStreamRetentionPeriodOutput, error) {
	var output kinesis.DecreaseStreamRetentionPeriodOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisDeleteStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisDeleteStreamFuture) Get(ctx workflow.Context) (*kinesis.DeleteStreamOutput, error) {
	var output kinesis.DeleteStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisDeregisterStreamConsumerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisDeregisterStreamConsumerFuture) Get(ctx workflow.Context) (*kinesis.DeregisterStreamConsumerOutput, error) {
	var output kinesis.DeregisterStreamConsumerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisDescribeLimitsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisDescribeLimitsFuture) Get(ctx workflow.Context) (*kinesis.DescribeLimitsOutput, error) {
	var output kinesis.DescribeLimitsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisDescribeStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisDescribeStreamFuture) Get(ctx workflow.Context) (*kinesis.DescribeStreamOutput, error) {
	var output kinesis.DescribeStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisDescribeStreamConsumerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisDescribeStreamConsumerFuture) Get(ctx workflow.Context) (*kinesis.DescribeStreamConsumerOutput, error) {
	var output kinesis.DescribeStreamConsumerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisDescribeStreamSummaryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisDescribeStreamSummaryFuture) Get(ctx workflow.Context) (*kinesis.DescribeStreamSummaryOutput, error) {
	var output kinesis.DescribeStreamSummaryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisDisableEnhancedMonitoringFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisDisableEnhancedMonitoringFuture) Get(ctx workflow.Context) (*kinesis.EnhancedMonitoringOutput, error) {
	var output kinesis.EnhancedMonitoringOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisEnableEnhancedMonitoringFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisEnableEnhancedMonitoringFuture) Get(ctx workflow.Context) (*kinesis.EnhancedMonitoringOutput, error) {
	var output kinesis.EnhancedMonitoringOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisGetRecordsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisGetRecordsFuture) Get(ctx workflow.Context) (*kinesis.GetRecordsOutput, error) {
	var output kinesis.GetRecordsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisGetShardIteratorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisGetShardIteratorFuture) Get(ctx workflow.Context) (*kinesis.GetShardIteratorOutput, error) {
	var output kinesis.GetShardIteratorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisIncreaseStreamRetentionPeriodFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisIncreaseStreamRetentionPeriodFuture) Get(ctx workflow.Context) (*kinesis.IncreaseStreamRetentionPeriodOutput, error) {
	var output kinesis.IncreaseStreamRetentionPeriodOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisListShardsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisListShardsFuture) Get(ctx workflow.Context) (*kinesis.ListShardsOutput, error) {
	var output kinesis.ListShardsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisListStreamConsumersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisListStreamConsumersFuture) Get(ctx workflow.Context) (*kinesis.ListStreamConsumersOutput, error) {
	var output kinesis.ListStreamConsumersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisListStreamsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisListStreamsFuture) Get(ctx workflow.Context) (*kinesis.ListStreamsOutput, error) {
	var output kinesis.ListStreamsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisListTagsForStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisListTagsForStreamFuture) Get(ctx workflow.Context) (*kinesis.ListTagsForStreamOutput, error) {
	var output kinesis.ListTagsForStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisMergeShardsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisMergeShardsFuture) Get(ctx workflow.Context) (*kinesis.MergeShardsOutput, error) {
	var output kinesis.MergeShardsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisPutRecordFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisPutRecordFuture) Get(ctx workflow.Context) (*kinesis.PutRecordOutput, error) {
	var output kinesis.PutRecordOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisPutRecordsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisPutRecordsFuture) Get(ctx workflow.Context) (*kinesis.PutRecordsOutput, error) {
	var output kinesis.PutRecordsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisRegisterStreamConsumerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisRegisterStreamConsumerFuture) Get(ctx workflow.Context) (*kinesis.RegisterStreamConsumerOutput, error) {
	var output kinesis.RegisterStreamConsumerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisRemoveTagsFromStreamFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisRemoveTagsFromStreamFuture) Get(ctx workflow.Context) (*kinesis.RemoveTagsFromStreamOutput, error) {
	var output kinesis.RemoveTagsFromStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisSplitShardFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisSplitShardFuture) Get(ctx workflow.Context) (*kinesis.SplitShardOutput, error) {
	var output kinesis.SplitShardOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisStartStreamEncryptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisStartStreamEncryptionFuture) Get(ctx workflow.Context) (*kinesis.StartStreamEncryptionOutput, error) {
	var output kinesis.StartStreamEncryptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisStopStreamEncryptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisStopStreamEncryptionFuture) Get(ctx workflow.Context) (*kinesis.StopStreamEncryptionOutput, error) {
	var output kinesis.StopStreamEncryptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisSubscribeToShardFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisSubscribeToShardFuture) Get(ctx workflow.Context) (*kinesis.SubscribeToShardOutput, error) {
	var output kinesis.SubscribeToShardOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisUpdateShardCountFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *KinesisUpdateShardCountFuture) Get(ctx workflow.Context) (*kinesis.UpdateShardCountOutput, error) {
	var output kinesis.UpdateShardCountOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTagsToStream(ctx workflow.Context, input *kinesis.AddTagsToStreamInput) (*kinesis.AddTagsToStreamOutput, error) {
	var output kinesis.AddTagsToStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.AddTagsToStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTagsToStreamAsync(ctx workflow.Context, input *kinesis.AddTagsToStreamInput) *KinesisAddTagsToStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.AddTagsToStream", input)
	return &KinesisAddTagsToStreamFuture{Future: future}
}

func (a *stub) CreateStream(ctx workflow.Context, input *kinesis.CreateStreamInput) (*kinesis.CreateStreamOutput, error) {
	var output kinesis.CreateStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.CreateStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateStreamAsync(ctx workflow.Context, input *kinesis.CreateStreamInput) *KinesisCreateStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.CreateStream", input)
	return &KinesisCreateStreamFuture{Future: future}
}

func (a *stub) DecreaseStreamRetentionPeriod(ctx workflow.Context, input *kinesis.DecreaseStreamRetentionPeriodInput) (*kinesis.DecreaseStreamRetentionPeriodOutput, error) {
	var output kinesis.DecreaseStreamRetentionPeriodOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DecreaseStreamRetentionPeriod", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DecreaseStreamRetentionPeriodAsync(ctx workflow.Context, input *kinesis.DecreaseStreamRetentionPeriodInput) *KinesisDecreaseStreamRetentionPeriodFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DecreaseStreamRetentionPeriod", input)
	return &KinesisDecreaseStreamRetentionPeriodFuture{Future: future}
}

func (a *stub) DeleteStream(ctx workflow.Context, input *kinesis.DeleteStreamInput) (*kinesis.DeleteStreamOutput, error) {
	var output kinesis.DeleteStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DeleteStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteStreamAsync(ctx workflow.Context, input *kinesis.DeleteStreamInput) *KinesisDeleteStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DeleteStream", input)
	return &KinesisDeleteStreamFuture{Future: future}
}

func (a *stub) DeregisterStreamConsumer(ctx workflow.Context, input *kinesis.DeregisterStreamConsumerInput) (*kinesis.DeregisterStreamConsumerOutput, error) {
	var output kinesis.DeregisterStreamConsumerOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DeregisterStreamConsumer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeregisterStreamConsumerAsync(ctx workflow.Context, input *kinesis.DeregisterStreamConsumerInput) *KinesisDeregisterStreamConsumerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DeregisterStreamConsumer", input)
	return &KinesisDeregisterStreamConsumerFuture{Future: future}
}

func (a *stub) DescribeLimits(ctx workflow.Context, input *kinesis.DescribeLimitsInput) (*kinesis.DescribeLimitsOutput, error) {
	var output kinesis.DescribeLimitsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeLimits", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLimitsAsync(ctx workflow.Context, input *kinesis.DescribeLimitsInput) *KinesisDescribeLimitsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeLimits", input)
	return &KinesisDescribeLimitsFuture{Future: future}
}

func (a *stub) DescribeStream(ctx workflow.Context, input *kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error) {
	var output kinesis.DescribeStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeStreamAsync(ctx workflow.Context, input *kinesis.DescribeStreamInput) *KinesisDescribeStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeStream", input)
	return &KinesisDescribeStreamFuture{Future: future}
}

func (a *stub) DescribeStreamConsumer(ctx workflow.Context, input *kinesis.DescribeStreamConsumerInput) (*kinesis.DescribeStreamConsumerOutput, error) {
	var output kinesis.DescribeStreamConsumerOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeStreamConsumer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeStreamConsumerAsync(ctx workflow.Context, input *kinesis.DescribeStreamConsumerInput) *KinesisDescribeStreamConsumerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeStreamConsumer", input)
	return &KinesisDescribeStreamConsumerFuture{Future: future}
}

func (a *stub) DescribeStreamSummary(ctx workflow.Context, input *kinesis.DescribeStreamSummaryInput) (*kinesis.DescribeStreamSummaryOutput, error) {
	var output kinesis.DescribeStreamSummaryOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeStreamSummary", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeStreamSummaryAsync(ctx workflow.Context, input *kinesis.DescribeStreamSummaryInput) *KinesisDescribeStreamSummaryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DescribeStreamSummary", input)
	return &KinesisDescribeStreamSummaryFuture{Future: future}
}

func (a *stub) DisableEnhancedMonitoring(ctx workflow.Context, input *kinesis.DisableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error) {
	var output kinesis.EnhancedMonitoringOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.DisableEnhancedMonitoring", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisableEnhancedMonitoringAsync(ctx workflow.Context, input *kinesis.DisableEnhancedMonitoringInput) *KinesisDisableEnhancedMonitoringFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.DisableEnhancedMonitoring", input)
	return &KinesisDisableEnhancedMonitoringFuture{Future: future}
}

func (a *stub) EnableEnhancedMonitoring(ctx workflow.Context, input *kinesis.EnableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error) {
	var output kinesis.EnhancedMonitoringOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.EnableEnhancedMonitoring", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EnableEnhancedMonitoringAsync(ctx workflow.Context, input *kinesis.EnableEnhancedMonitoringInput) *KinesisEnableEnhancedMonitoringFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.EnableEnhancedMonitoring", input)
	return &KinesisEnableEnhancedMonitoringFuture{Future: future}
}

func (a *stub) GetRecords(ctx workflow.Context, input *kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error) {
	var output kinesis.GetRecordsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.GetRecords", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRecordsAsync(ctx workflow.Context, input *kinesis.GetRecordsInput) *KinesisGetRecordsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.GetRecords", input)
	return &KinesisGetRecordsFuture{Future: future}
}

func (a *stub) GetShardIterator(ctx workflow.Context, input *kinesis.GetShardIteratorInput) (*kinesis.GetShardIteratorOutput, error) {
	var output kinesis.GetShardIteratorOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.GetShardIterator", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetShardIteratorAsync(ctx workflow.Context, input *kinesis.GetShardIteratorInput) *KinesisGetShardIteratorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.GetShardIterator", input)
	return &KinesisGetShardIteratorFuture{Future: future}
}

func (a *stub) IncreaseStreamRetentionPeriod(ctx workflow.Context, input *kinesis.IncreaseStreamRetentionPeriodInput) (*kinesis.IncreaseStreamRetentionPeriodOutput, error) {
	var output kinesis.IncreaseStreamRetentionPeriodOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.IncreaseStreamRetentionPeriod", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) IncreaseStreamRetentionPeriodAsync(ctx workflow.Context, input *kinesis.IncreaseStreamRetentionPeriodInput) *KinesisIncreaseStreamRetentionPeriodFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.IncreaseStreamRetentionPeriod", input)
	return &KinesisIncreaseStreamRetentionPeriodFuture{Future: future}
}

func (a *stub) ListShards(ctx workflow.Context, input *kinesis.ListShardsInput) (*kinesis.ListShardsOutput, error) {
	var output kinesis.ListShardsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.ListShards", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListShardsAsync(ctx workflow.Context, input *kinesis.ListShardsInput) *KinesisListShardsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.ListShards", input)
	return &KinesisListShardsFuture{Future: future}
}

func (a *stub) ListStreamConsumers(ctx workflow.Context, input *kinesis.ListStreamConsumersInput) (*kinesis.ListStreamConsumersOutput, error) {
	var output kinesis.ListStreamConsumersOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.ListStreamConsumers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListStreamConsumersAsync(ctx workflow.Context, input *kinesis.ListStreamConsumersInput) *KinesisListStreamConsumersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.ListStreamConsumers", input)
	return &KinesisListStreamConsumersFuture{Future: future}
}

func (a *stub) ListStreams(ctx workflow.Context, input *kinesis.ListStreamsInput) (*kinesis.ListStreamsOutput, error) {
	var output kinesis.ListStreamsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.ListStreams", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListStreamsAsync(ctx workflow.Context, input *kinesis.ListStreamsInput) *KinesisListStreamsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.ListStreams", input)
	return &KinesisListStreamsFuture{Future: future}
}

func (a *stub) ListTagsForStream(ctx workflow.Context, input *kinesis.ListTagsForStreamInput) (*kinesis.ListTagsForStreamOutput, error) {
	var output kinesis.ListTagsForStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.ListTagsForStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForStreamAsync(ctx workflow.Context, input *kinesis.ListTagsForStreamInput) *KinesisListTagsForStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.ListTagsForStream", input)
	return &KinesisListTagsForStreamFuture{Future: future}
}

func (a *stub) MergeShards(ctx workflow.Context, input *kinesis.MergeShardsInput) (*kinesis.MergeShardsOutput, error) {
	var output kinesis.MergeShardsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.MergeShards", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) MergeShardsAsync(ctx workflow.Context, input *kinesis.MergeShardsInput) *KinesisMergeShardsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.MergeShards", input)
	return &KinesisMergeShardsFuture{Future: future}
}

func (a *stub) PutRecord(ctx workflow.Context, input *kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error) {
	var output kinesis.PutRecordOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.PutRecord", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutRecordAsync(ctx workflow.Context, input *kinesis.PutRecordInput) *KinesisPutRecordFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.PutRecord", input)
	return &KinesisPutRecordFuture{Future: future}
}

func (a *stub) PutRecords(ctx workflow.Context, input *kinesis.PutRecordsInput) (*kinesis.PutRecordsOutput, error) {
	var output kinesis.PutRecordsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.PutRecords", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutRecordsAsync(ctx workflow.Context, input *kinesis.PutRecordsInput) *KinesisPutRecordsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.PutRecords", input)
	return &KinesisPutRecordsFuture{Future: future}
}

func (a *stub) RegisterStreamConsumer(ctx workflow.Context, input *kinesis.RegisterStreamConsumerInput) (*kinesis.RegisterStreamConsumerOutput, error) {
	var output kinesis.RegisterStreamConsumerOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.RegisterStreamConsumer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RegisterStreamConsumerAsync(ctx workflow.Context, input *kinesis.RegisterStreamConsumerInput) *KinesisRegisterStreamConsumerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.RegisterStreamConsumer", input)
	return &KinesisRegisterStreamConsumerFuture{Future: future}
}

func (a *stub) RemoveTagsFromStream(ctx workflow.Context, input *kinesis.RemoveTagsFromStreamInput) (*kinesis.RemoveTagsFromStreamOutput, error) {
	var output kinesis.RemoveTagsFromStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.RemoveTagsFromStream", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveTagsFromStreamAsync(ctx workflow.Context, input *kinesis.RemoveTagsFromStreamInput) *KinesisRemoveTagsFromStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.RemoveTagsFromStream", input)
	return &KinesisRemoveTagsFromStreamFuture{Future: future}
}

func (a *stub) SplitShard(ctx workflow.Context, input *kinesis.SplitShardInput) (*kinesis.SplitShardOutput, error) {
	var output kinesis.SplitShardOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.SplitShard", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SplitShardAsync(ctx workflow.Context, input *kinesis.SplitShardInput) *KinesisSplitShardFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.SplitShard", input)
	return &KinesisSplitShardFuture{Future: future}
}

func (a *stub) StartStreamEncryption(ctx workflow.Context, input *kinesis.StartStreamEncryptionInput) (*kinesis.StartStreamEncryptionOutput, error) {
	var output kinesis.StartStreamEncryptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.StartStreamEncryption", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartStreamEncryptionAsync(ctx workflow.Context, input *kinesis.StartStreamEncryptionInput) *KinesisStartStreamEncryptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.StartStreamEncryption", input)
	return &KinesisStartStreamEncryptionFuture{Future: future}
}

func (a *stub) StopStreamEncryption(ctx workflow.Context, input *kinesis.StopStreamEncryptionInput) (*kinesis.StopStreamEncryptionOutput, error) {
	var output kinesis.StopStreamEncryptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.StopStreamEncryption", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StopStreamEncryptionAsync(ctx workflow.Context, input *kinesis.StopStreamEncryptionInput) *KinesisStopStreamEncryptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.StopStreamEncryption", input)
	return &KinesisStopStreamEncryptionFuture{Future: future}
}

func (a *stub) SubscribeToShard(ctx workflow.Context, input *kinesis.SubscribeToShardInput) (*kinesis.SubscribeToShardOutput, error) {
	var output kinesis.SubscribeToShardOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.SubscribeToShard", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SubscribeToShardAsync(ctx workflow.Context, input *kinesis.SubscribeToShardInput) *KinesisSubscribeToShardFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.SubscribeToShard", input)
	return &KinesisSubscribeToShardFuture{Future: future}
}

func (a *stub) UpdateShardCount(ctx workflow.Context, input *kinesis.UpdateShardCountInput) (*kinesis.UpdateShardCountOutput, error) {
	var output kinesis.UpdateShardCountOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesis.UpdateShardCount", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateShardCountAsync(ctx workflow.Context, input *kinesis.UpdateShardCountInput) *KinesisUpdateShardCountFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.UpdateShardCount", input)
	return &KinesisUpdateShardCountFuture{Future: future}
}

func (a *stub) WaitUntilStreamExists(ctx workflow.Context, input *kinesis.DescribeStreamInput) error {
	return workflow.ExecuteActivity(ctx, "aws.kinesis.WaitUntilStreamExists", input).Get(ctx, nil)
}

func (a *stub) WaitUntilStreamExistsAsync(ctx workflow.Context, input *kinesis.DescribeStreamInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.WaitUntilStreamExists", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilStreamNotExists(ctx workflow.Context, input *kinesis.DescribeStreamInput) error {
	return workflow.ExecuteActivity(ctx, "aws.kinesis.WaitUntilStreamNotExists", input).Get(ctx, nil)
}

func (a *stub) WaitUntilStreamNotExistsAsync(ctx workflow.Context, input *kinesis.DescribeStreamInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesis.WaitUntilStreamNotExists", input)
	return clients.NewVoidFuture(future)
}
