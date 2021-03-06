// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package fsxstub

import (
	"github.com/aws/aws-sdk-go/service/fsx"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type FSxCancelDataRepositoryTaskFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FSxCancelDataRepositoryTaskFuture) Get(ctx workflow.Context) (*fsx.CancelDataRepositoryTaskOutput, error) {
	var output fsx.CancelDataRepositoryTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FSxCreateBackupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FSxCreateBackupFuture) Get(ctx workflow.Context) (*fsx.CreateBackupOutput, error) {
	var output fsx.CreateBackupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FSxCreateDataRepositoryTaskFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FSxCreateDataRepositoryTaskFuture) Get(ctx workflow.Context) (*fsx.CreateDataRepositoryTaskOutput, error) {
	var output fsx.CreateDataRepositoryTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FSxCreateFileSystemFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FSxCreateFileSystemFuture) Get(ctx workflow.Context) (*fsx.CreateFileSystemOutput, error) {
	var output fsx.CreateFileSystemOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FSxCreateFileSystemFromBackupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FSxCreateFileSystemFromBackupFuture) Get(ctx workflow.Context) (*fsx.CreateFileSystemFromBackupOutput, error) {
	var output fsx.CreateFileSystemFromBackupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FSxDeleteBackupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FSxDeleteBackupFuture) Get(ctx workflow.Context) (*fsx.DeleteBackupOutput, error) {
	var output fsx.DeleteBackupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FSxDeleteFileSystemFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FSxDeleteFileSystemFuture) Get(ctx workflow.Context) (*fsx.DeleteFileSystemOutput, error) {
	var output fsx.DeleteFileSystemOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FSxDescribeBackupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FSxDescribeBackupsFuture) Get(ctx workflow.Context) (*fsx.DescribeBackupsOutput, error) {
	var output fsx.DescribeBackupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FSxDescribeDataRepositoryTasksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FSxDescribeDataRepositoryTasksFuture) Get(ctx workflow.Context) (*fsx.DescribeDataRepositoryTasksOutput, error) {
	var output fsx.DescribeDataRepositoryTasksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FSxDescribeFileSystemsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FSxDescribeFileSystemsFuture) Get(ctx workflow.Context) (*fsx.DescribeFileSystemsOutput, error) {
	var output fsx.DescribeFileSystemsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FSxListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FSxListTagsForResourceFuture) Get(ctx workflow.Context) (*fsx.ListTagsForResourceOutput, error) {
	var output fsx.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FSxTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FSxTagResourceFuture) Get(ctx workflow.Context) (*fsx.TagResourceOutput, error) {
	var output fsx.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FSxUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FSxUntagResourceFuture) Get(ctx workflow.Context) (*fsx.UntagResourceOutput, error) {
	var output fsx.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type FSxUpdateFileSystemFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *FSxUpdateFileSystemFuture) Get(ctx workflow.Context) (*fsx.UpdateFileSystemOutput, error) {
	var output fsx.UpdateFileSystemOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelDataRepositoryTask(ctx workflow.Context, input *fsx.CancelDataRepositoryTaskInput) (*fsx.CancelDataRepositoryTaskOutput, error) {
	var output fsx.CancelDataRepositoryTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws.fsx.CancelDataRepositoryTask", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelDataRepositoryTaskAsync(ctx workflow.Context, input *fsx.CancelDataRepositoryTaskInput) *FSxCancelDataRepositoryTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fsx.CancelDataRepositoryTask", input)
	return &FSxCancelDataRepositoryTaskFuture{Future: future}
}

func (a *stub) CreateBackup(ctx workflow.Context, input *fsx.CreateBackupInput) (*fsx.CreateBackupOutput, error) {
	var output fsx.CreateBackupOutput
	err := workflow.ExecuteActivity(ctx, "aws.fsx.CreateBackup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBackupAsync(ctx workflow.Context, input *fsx.CreateBackupInput) *FSxCreateBackupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fsx.CreateBackup", input)
	return &FSxCreateBackupFuture{Future: future}
}

func (a *stub) CreateDataRepositoryTask(ctx workflow.Context, input *fsx.CreateDataRepositoryTaskInput) (*fsx.CreateDataRepositoryTaskOutput, error) {
	var output fsx.CreateDataRepositoryTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws.fsx.CreateDataRepositoryTask", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDataRepositoryTaskAsync(ctx workflow.Context, input *fsx.CreateDataRepositoryTaskInput) *FSxCreateDataRepositoryTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fsx.CreateDataRepositoryTask", input)
	return &FSxCreateDataRepositoryTaskFuture{Future: future}
}

func (a *stub) CreateFileSystem(ctx workflow.Context, input *fsx.CreateFileSystemInput) (*fsx.CreateFileSystemOutput, error) {
	var output fsx.CreateFileSystemOutput
	err := workflow.ExecuteActivity(ctx, "aws.fsx.CreateFileSystem", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateFileSystemAsync(ctx workflow.Context, input *fsx.CreateFileSystemInput) *FSxCreateFileSystemFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fsx.CreateFileSystem", input)
	return &FSxCreateFileSystemFuture{Future: future}
}

func (a *stub) CreateFileSystemFromBackup(ctx workflow.Context, input *fsx.CreateFileSystemFromBackupInput) (*fsx.CreateFileSystemFromBackupOutput, error) {
	var output fsx.CreateFileSystemFromBackupOutput
	err := workflow.ExecuteActivity(ctx, "aws.fsx.CreateFileSystemFromBackup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateFileSystemFromBackupAsync(ctx workflow.Context, input *fsx.CreateFileSystemFromBackupInput) *FSxCreateFileSystemFromBackupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fsx.CreateFileSystemFromBackup", input)
	return &FSxCreateFileSystemFromBackupFuture{Future: future}
}

func (a *stub) DeleteBackup(ctx workflow.Context, input *fsx.DeleteBackupInput) (*fsx.DeleteBackupOutput, error) {
	var output fsx.DeleteBackupOutput
	err := workflow.ExecuteActivity(ctx, "aws.fsx.DeleteBackup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBackupAsync(ctx workflow.Context, input *fsx.DeleteBackupInput) *FSxDeleteBackupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fsx.DeleteBackup", input)
	return &FSxDeleteBackupFuture{Future: future}
}

func (a *stub) DeleteFileSystem(ctx workflow.Context, input *fsx.DeleteFileSystemInput) (*fsx.DeleteFileSystemOutput, error) {
	var output fsx.DeleteFileSystemOutput
	err := workflow.ExecuteActivity(ctx, "aws.fsx.DeleteFileSystem", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteFileSystemAsync(ctx workflow.Context, input *fsx.DeleteFileSystemInput) *FSxDeleteFileSystemFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fsx.DeleteFileSystem", input)
	return &FSxDeleteFileSystemFuture{Future: future}
}

func (a *stub) DescribeBackups(ctx workflow.Context, input *fsx.DescribeBackupsInput) (*fsx.DescribeBackupsOutput, error) {
	var output fsx.DescribeBackupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.fsx.DescribeBackups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeBackupsAsync(ctx workflow.Context, input *fsx.DescribeBackupsInput) *FSxDescribeBackupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fsx.DescribeBackups", input)
	return &FSxDescribeBackupsFuture{Future: future}
}

func (a *stub) DescribeDataRepositoryTasks(ctx workflow.Context, input *fsx.DescribeDataRepositoryTasksInput) (*fsx.DescribeDataRepositoryTasksOutput, error) {
	var output fsx.DescribeDataRepositoryTasksOutput
	err := workflow.ExecuteActivity(ctx, "aws.fsx.DescribeDataRepositoryTasks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDataRepositoryTasksAsync(ctx workflow.Context, input *fsx.DescribeDataRepositoryTasksInput) *FSxDescribeDataRepositoryTasksFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fsx.DescribeDataRepositoryTasks", input)
	return &FSxDescribeDataRepositoryTasksFuture{Future: future}
}

func (a *stub) DescribeFileSystems(ctx workflow.Context, input *fsx.DescribeFileSystemsInput) (*fsx.DescribeFileSystemsOutput, error) {
	var output fsx.DescribeFileSystemsOutput
	err := workflow.ExecuteActivity(ctx, "aws.fsx.DescribeFileSystems", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeFileSystemsAsync(ctx workflow.Context, input *fsx.DescribeFileSystemsInput) *FSxDescribeFileSystemsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fsx.DescribeFileSystems", input)
	return &FSxDescribeFileSystemsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *fsx.ListTagsForResourceInput) (*fsx.ListTagsForResourceOutput, error) {
	var output fsx.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.fsx.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *fsx.ListTagsForResourceInput) *FSxListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fsx.ListTagsForResource", input)
	return &FSxListTagsForResourceFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *fsx.TagResourceInput) (*fsx.TagResourceOutput, error) {
	var output fsx.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.fsx.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *fsx.TagResourceInput) *FSxTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fsx.TagResource", input)
	return &FSxTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *fsx.UntagResourceInput) (*fsx.UntagResourceOutput, error) {
	var output fsx.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.fsx.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *fsx.UntagResourceInput) *FSxUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fsx.UntagResource", input)
	return &FSxUntagResourceFuture{Future: future}
}

func (a *stub) UpdateFileSystem(ctx workflow.Context, input *fsx.UpdateFileSystemInput) (*fsx.UpdateFileSystemOutput, error) {
	var output fsx.UpdateFileSystemOutput
	err := workflow.ExecuteActivity(ctx, "aws.fsx.UpdateFileSystem", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateFileSystemAsync(ctx workflow.Context, input *fsx.UpdateFileSystemInput) *FSxUpdateFileSystemFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fsx.UpdateFileSystem", input)
	return &FSxUpdateFileSystemFuture{Future: future}
}
