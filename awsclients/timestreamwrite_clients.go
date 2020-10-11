// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/timestreamwrite"
	"go.temporal.io/sdk/workflow"
)

type TimestreamWriteClient interface {
	CreateDatabase(ctx workflow.Context, input *timestreamwrite.CreateDatabaseInput) (*timestreamwrite.CreateDatabaseOutput, error)
	CreateDatabaseAsync(ctx workflow.Context, input *timestreamwrite.CreateDatabaseInput) *TimestreamwriteCreateDatabaseResult

	CreateTable(ctx workflow.Context, input *timestreamwrite.CreateTableInput) (*timestreamwrite.CreateTableOutput, error)
	CreateTableAsync(ctx workflow.Context, input *timestreamwrite.CreateTableInput) *TimestreamwriteCreateTableResult

	DeleteDatabase(ctx workflow.Context, input *timestreamwrite.DeleteDatabaseInput) (*timestreamwrite.DeleteDatabaseOutput, error)
	DeleteDatabaseAsync(ctx workflow.Context, input *timestreamwrite.DeleteDatabaseInput) *TimestreamwriteDeleteDatabaseResult

	DeleteTable(ctx workflow.Context, input *timestreamwrite.DeleteTableInput) (*timestreamwrite.DeleteTableOutput, error)
	DeleteTableAsync(ctx workflow.Context, input *timestreamwrite.DeleteTableInput) *TimestreamwriteDeleteTableResult

	DescribeDatabase(ctx workflow.Context, input *timestreamwrite.DescribeDatabaseInput) (*timestreamwrite.DescribeDatabaseOutput, error)
	DescribeDatabaseAsync(ctx workflow.Context, input *timestreamwrite.DescribeDatabaseInput) *TimestreamwriteDescribeDatabaseResult

	DescribeEndpoints(ctx workflow.Context, input *timestreamwrite.DescribeEndpointsInput) (*timestreamwrite.DescribeEndpointsOutput, error)
	DescribeEndpointsAsync(ctx workflow.Context, input *timestreamwrite.DescribeEndpointsInput) *TimestreamwriteDescribeEndpointsResult

	DescribeTable(ctx workflow.Context, input *timestreamwrite.DescribeTableInput) (*timestreamwrite.DescribeTableOutput, error)
	DescribeTableAsync(ctx workflow.Context, input *timestreamwrite.DescribeTableInput) *TimestreamwriteDescribeTableResult

	ListDatabases(ctx workflow.Context, input *timestreamwrite.ListDatabasesInput) (*timestreamwrite.ListDatabasesOutput, error)
	ListDatabasesAsync(ctx workflow.Context, input *timestreamwrite.ListDatabasesInput) *TimestreamwriteListDatabasesResult

	ListTables(ctx workflow.Context, input *timestreamwrite.ListTablesInput) (*timestreamwrite.ListTablesOutput, error)
	ListTablesAsync(ctx workflow.Context, input *timestreamwrite.ListTablesInput) *TimestreamwriteListTablesResult

	ListTagsForResource(ctx workflow.Context, input *timestreamwrite.ListTagsForResourceInput) (*timestreamwrite.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *timestreamwrite.ListTagsForResourceInput) *TimestreamwriteListTagsForResourceResult

	TagResource(ctx workflow.Context, input *timestreamwrite.TagResourceInput) (*timestreamwrite.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *timestreamwrite.TagResourceInput) *TimestreamwriteTagResourceResult

	UntagResource(ctx workflow.Context, input *timestreamwrite.UntagResourceInput) (*timestreamwrite.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *timestreamwrite.UntagResourceInput) *TimestreamwriteUntagResourceResult

	UpdateDatabase(ctx workflow.Context, input *timestreamwrite.UpdateDatabaseInput) (*timestreamwrite.UpdateDatabaseOutput, error)
	UpdateDatabaseAsync(ctx workflow.Context, input *timestreamwrite.UpdateDatabaseInput) *TimestreamwriteUpdateDatabaseResult

	UpdateTable(ctx workflow.Context, input *timestreamwrite.UpdateTableInput) (*timestreamwrite.UpdateTableOutput, error)
	UpdateTableAsync(ctx workflow.Context, input *timestreamwrite.UpdateTableInput) *TimestreamwriteUpdateTableResult

	WriteRecords(ctx workflow.Context, input *timestreamwrite.WriteRecordsInput) (*timestreamwrite.WriteRecordsOutput, error)
	WriteRecordsAsync(ctx workflow.Context, input *timestreamwrite.WriteRecordsInput) *TimestreamwriteWriteRecordsResult
}

type TimestreamWriteStub struct{}

func NewTimestreamWriteStub() TimestreamWriteClient {
	return &TimestreamWriteStub{}
}

type TimestreamwriteCreateDatabaseResult struct {
	Result workflow.Future
}

func (r *TimestreamwriteCreateDatabaseResult) Get(ctx workflow.Context) (*timestreamwrite.CreateDatabaseOutput, error) {
	var output timestreamwrite.CreateDatabaseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TimestreamwriteCreateTableResult struct {
	Result workflow.Future
}

func (r *TimestreamwriteCreateTableResult) Get(ctx workflow.Context) (*timestreamwrite.CreateTableOutput, error) {
	var output timestreamwrite.CreateTableOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TimestreamwriteDeleteDatabaseResult struct {
	Result workflow.Future
}

func (r *TimestreamwriteDeleteDatabaseResult) Get(ctx workflow.Context) (*timestreamwrite.DeleteDatabaseOutput, error) {
	var output timestreamwrite.DeleteDatabaseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TimestreamwriteDeleteTableResult struct {
	Result workflow.Future
}

func (r *TimestreamwriteDeleteTableResult) Get(ctx workflow.Context) (*timestreamwrite.DeleteTableOutput, error) {
	var output timestreamwrite.DeleteTableOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TimestreamwriteDescribeDatabaseResult struct {
	Result workflow.Future
}

func (r *TimestreamwriteDescribeDatabaseResult) Get(ctx workflow.Context) (*timestreamwrite.DescribeDatabaseOutput, error) {
	var output timestreamwrite.DescribeDatabaseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TimestreamwriteDescribeEndpointsResult struct {
	Result workflow.Future
}

func (r *TimestreamwriteDescribeEndpointsResult) Get(ctx workflow.Context) (*timestreamwrite.DescribeEndpointsOutput, error) {
	var output timestreamwrite.DescribeEndpointsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TimestreamwriteDescribeTableResult struct {
	Result workflow.Future
}

func (r *TimestreamwriteDescribeTableResult) Get(ctx workflow.Context) (*timestreamwrite.DescribeTableOutput, error) {
	var output timestreamwrite.DescribeTableOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TimestreamwriteListDatabasesResult struct {
	Result workflow.Future
}

func (r *TimestreamwriteListDatabasesResult) Get(ctx workflow.Context) (*timestreamwrite.ListDatabasesOutput, error) {
	var output timestreamwrite.ListDatabasesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TimestreamwriteListTablesResult struct {
	Result workflow.Future
}

func (r *TimestreamwriteListTablesResult) Get(ctx workflow.Context) (*timestreamwrite.ListTablesOutput, error) {
	var output timestreamwrite.ListTablesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TimestreamwriteListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *TimestreamwriteListTagsForResourceResult) Get(ctx workflow.Context) (*timestreamwrite.ListTagsForResourceOutput, error) {
	var output timestreamwrite.ListTagsForResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TimestreamwriteTagResourceResult struct {
	Result workflow.Future
}

func (r *TimestreamwriteTagResourceResult) Get(ctx workflow.Context) (*timestreamwrite.TagResourceOutput, error) {
	var output timestreamwrite.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TimestreamwriteUntagResourceResult struct {
	Result workflow.Future
}

func (r *TimestreamwriteUntagResourceResult) Get(ctx workflow.Context) (*timestreamwrite.UntagResourceOutput, error) {
	var output timestreamwrite.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TimestreamwriteUpdateDatabaseResult struct {
	Result workflow.Future
}

func (r *TimestreamwriteUpdateDatabaseResult) Get(ctx workflow.Context) (*timestreamwrite.UpdateDatabaseOutput, error) {
	var output timestreamwrite.UpdateDatabaseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TimestreamwriteUpdateTableResult struct {
	Result workflow.Future
}

func (r *TimestreamwriteUpdateTableResult) Get(ctx workflow.Context) (*timestreamwrite.UpdateTableOutput, error) {
	var output timestreamwrite.UpdateTableOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TimestreamwriteWriteRecordsResult struct {
	Result workflow.Future
}

func (r *TimestreamwriteWriteRecordsResult) Get(ctx workflow.Context) (*timestreamwrite.WriteRecordsOutput, error) {
	var output timestreamwrite.WriteRecordsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) CreateDatabase(ctx workflow.Context, input *timestreamwrite.CreateDatabaseInput) (*timestreamwrite.CreateDatabaseOutput, error) {
	var output timestreamwrite.CreateDatabaseOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.CreateDatabase", input).Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) CreateDatabaseAsync(ctx workflow.Context, input *timestreamwrite.CreateDatabaseInput) *TimestreamwriteCreateDatabaseResult {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.CreateDatabase", input)
	return &TimestreamwriteCreateDatabaseResult{Result: future}
}

func (a *TimestreamWriteStub) CreateTable(ctx workflow.Context, input *timestreamwrite.CreateTableInput) (*timestreamwrite.CreateTableOutput, error) {
	var output timestreamwrite.CreateTableOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.CreateTable", input).Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) CreateTableAsync(ctx workflow.Context, input *timestreamwrite.CreateTableInput) *TimestreamwriteCreateTableResult {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.CreateTable", input)
	return &TimestreamwriteCreateTableResult{Result: future}
}

func (a *TimestreamWriteStub) DeleteDatabase(ctx workflow.Context, input *timestreamwrite.DeleteDatabaseInput) (*timestreamwrite.DeleteDatabaseOutput, error) {
	var output timestreamwrite.DeleteDatabaseOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DeleteDatabase", input).Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) DeleteDatabaseAsync(ctx workflow.Context, input *timestreamwrite.DeleteDatabaseInput) *TimestreamwriteDeleteDatabaseResult {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DeleteDatabase", input)
	return &TimestreamwriteDeleteDatabaseResult{Result: future}
}

func (a *TimestreamWriteStub) DeleteTable(ctx workflow.Context, input *timestreamwrite.DeleteTableInput) (*timestreamwrite.DeleteTableOutput, error) {
	var output timestreamwrite.DeleteTableOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DeleteTable", input).Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) DeleteTableAsync(ctx workflow.Context, input *timestreamwrite.DeleteTableInput) *TimestreamwriteDeleteTableResult {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DeleteTable", input)
	return &TimestreamwriteDeleteTableResult{Result: future}
}

func (a *TimestreamWriteStub) DescribeDatabase(ctx workflow.Context, input *timestreamwrite.DescribeDatabaseInput) (*timestreamwrite.DescribeDatabaseOutput, error) {
	var output timestreamwrite.DescribeDatabaseOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DescribeDatabase", input).Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) DescribeDatabaseAsync(ctx workflow.Context, input *timestreamwrite.DescribeDatabaseInput) *TimestreamwriteDescribeDatabaseResult {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DescribeDatabase", input)
	return &TimestreamwriteDescribeDatabaseResult{Result: future}
}

func (a *TimestreamWriteStub) DescribeEndpoints(ctx workflow.Context, input *timestreamwrite.DescribeEndpointsInput) (*timestreamwrite.DescribeEndpointsOutput, error) {
	var output timestreamwrite.DescribeEndpointsOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DescribeEndpoints", input).Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) DescribeEndpointsAsync(ctx workflow.Context, input *timestreamwrite.DescribeEndpointsInput) *TimestreamwriteDescribeEndpointsResult {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DescribeEndpoints", input)
	return &TimestreamwriteDescribeEndpointsResult{Result: future}
}

func (a *TimestreamWriteStub) DescribeTable(ctx workflow.Context, input *timestreamwrite.DescribeTableInput) (*timestreamwrite.DescribeTableOutput, error) {
	var output timestreamwrite.DescribeTableOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DescribeTable", input).Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) DescribeTableAsync(ctx workflow.Context, input *timestreamwrite.DescribeTableInput) *TimestreamwriteDescribeTableResult {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.DescribeTable", input)
	return &TimestreamwriteDescribeTableResult{Result: future}
}

func (a *TimestreamWriteStub) ListDatabases(ctx workflow.Context, input *timestreamwrite.ListDatabasesInput) (*timestreamwrite.ListDatabasesOutput, error) {
	var output timestreamwrite.ListDatabasesOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.ListDatabases", input).Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) ListDatabasesAsync(ctx workflow.Context, input *timestreamwrite.ListDatabasesInput) *TimestreamwriteListDatabasesResult {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.ListDatabases", input)
	return &TimestreamwriteListDatabasesResult{Result: future}
}

func (a *TimestreamWriteStub) ListTables(ctx workflow.Context, input *timestreamwrite.ListTablesInput) (*timestreamwrite.ListTablesOutput, error) {
	var output timestreamwrite.ListTablesOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.ListTables", input).Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) ListTablesAsync(ctx workflow.Context, input *timestreamwrite.ListTablesInput) *TimestreamwriteListTablesResult {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.ListTables", input)
	return &TimestreamwriteListTablesResult{Result: future}
}

func (a *TimestreamWriteStub) ListTagsForResource(ctx workflow.Context, input *timestreamwrite.ListTagsForResourceInput) (*timestreamwrite.ListTagsForResourceOutput, error) {
	var output timestreamwrite.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) ListTagsForResourceAsync(ctx workflow.Context, input *timestreamwrite.ListTagsForResourceInput) *TimestreamwriteListTagsForResourceResult {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.ListTagsForResource", input)
	return &TimestreamwriteListTagsForResourceResult{Result: future}
}

func (a *TimestreamWriteStub) TagResource(ctx workflow.Context, input *timestreamwrite.TagResourceInput) (*timestreamwrite.TagResourceOutput, error) {
	var output timestreamwrite.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) TagResourceAsync(ctx workflow.Context, input *timestreamwrite.TagResourceInput) *TimestreamwriteTagResourceResult {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.TagResource", input)
	return &TimestreamwriteTagResourceResult{Result: future}
}

func (a *TimestreamWriteStub) UntagResource(ctx workflow.Context, input *timestreamwrite.UntagResourceInput) (*timestreamwrite.UntagResourceOutput, error) {
	var output timestreamwrite.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) UntagResourceAsync(ctx workflow.Context, input *timestreamwrite.UntagResourceInput) *TimestreamwriteUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.UntagResource", input)
	return &TimestreamwriteUntagResourceResult{Result: future}
}

func (a *TimestreamWriteStub) UpdateDatabase(ctx workflow.Context, input *timestreamwrite.UpdateDatabaseInput) (*timestreamwrite.UpdateDatabaseOutput, error) {
	var output timestreamwrite.UpdateDatabaseOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.UpdateDatabase", input).Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) UpdateDatabaseAsync(ctx workflow.Context, input *timestreamwrite.UpdateDatabaseInput) *TimestreamwriteUpdateDatabaseResult {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.UpdateDatabase", input)
	return &TimestreamwriteUpdateDatabaseResult{Result: future}
}

func (a *TimestreamWriteStub) UpdateTable(ctx workflow.Context, input *timestreamwrite.UpdateTableInput) (*timestreamwrite.UpdateTableOutput, error) {
	var output timestreamwrite.UpdateTableOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.UpdateTable", input).Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) UpdateTableAsync(ctx workflow.Context, input *timestreamwrite.UpdateTableInput) *TimestreamwriteUpdateTableResult {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.UpdateTable", input)
	return &TimestreamwriteUpdateTableResult{Result: future}
}

func (a *TimestreamWriteStub) WriteRecords(ctx workflow.Context, input *timestreamwrite.WriteRecordsInput) (*timestreamwrite.WriteRecordsOutput, error) {
	var output timestreamwrite.WriteRecordsOutput
	err := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.WriteRecords", input).Get(ctx, &output)
	return &output, err
}

func (a *TimestreamWriteStub) WriteRecordsAsync(ctx workflow.Context, input *timestreamwrite.WriteRecordsInput) *TimestreamwriteWriteRecordsResult {
	future := workflow.ExecuteActivity(ctx, "aws.timestreamwrite.WriteRecords", input)
	return &TimestreamwriteWriteRecordsResult{Result: future}
}