// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mobilestub

import (
	"github.com/aws/aws-sdk-go/service/mobile"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type MobileCreateProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MobileCreateProjectFuture) Get(ctx workflow.Context) (*mobile.CreateProjectOutput, error) {
	var output mobile.CreateProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MobileDeleteProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MobileDeleteProjectFuture) Get(ctx workflow.Context) (*mobile.DeleteProjectOutput, error) {
	var output mobile.DeleteProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MobileDescribeBundleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MobileDescribeBundleFuture) Get(ctx workflow.Context) (*mobile.DescribeBundleOutput, error) {
	var output mobile.DescribeBundleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MobileDescribeProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MobileDescribeProjectFuture) Get(ctx workflow.Context) (*mobile.DescribeProjectOutput, error) {
	var output mobile.DescribeProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MobileExportBundleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MobileExportBundleFuture) Get(ctx workflow.Context) (*mobile.ExportBundleOutput, error) {
	var output mobile.ExportBundleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MobileExportProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MobileExportProjectFuture) Get(ctx workflow.Context) (*mobile.ExportProjectOutput, error) {
	var output mobile.ExportProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MobileListBundlesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MobileListBundlesFuture) Get(ctx workflow.Context) (*mobile.ListBundlesOutput, error) {
	var output mobile.ListBundlesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MobileListProjectsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MobileListProjectsFuture) Get(ctx workflow.Context) (*mobile.ListProjectsOutput, error) {
	var output mobile.ListProjectsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MobileUpdateProjectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MobileUpdateProjectFuture) Get(ctx workflow.Context) (*mobile.UpdateProjectOutput, error) {
	var output mobile.UpdateProjectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateProject(ctx workflow.Context, input *mobile.CreateProjectInput) (*mobile.CreateProjectOutput, error) {
	var output mobile.CreateProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws.mobile.CreateProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateProjectAsync(ctx workflow.Context, input *mobile.CreateProjectInput) *MobileCreateProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mobile.CreateProject", input)
	return &MobileCreateProjectFuture{Future: future}
}

func (a *stub) DeleteProject(ctx workflow.Context, input *mobile.DeleteProjectInput) (*mobile.DeleteProjectOutput, error) {
	var output mobile.DeleteProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws.mobile.DeleteProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteProjectAsync(ctx workflow.Context, input *mobile.DeleteProjectInput) *MobileDeleteProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mobile.DeleteProject", input)
	return &MobileDeleteProjectFuture{Future: future}
}

func (a *stub) DescribeBundle(ctx workflow.Context, input *mobile.DescribeBundleInput) (*mobile.DescribeBundleOutput, error) {
	var output mobile.DescribeBundleOutput
	err := workflow.ExecuteActivity(ctx, "aws.mobile.DescribeBundle", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeBundleAsync(ctx workflow.Context, input *mobile.DescribeBundleInput) *MobileDescribeBundleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mobile.DescribeBundle", input)
	return &MobileDescribeBundleFuture{Future: future}
}

func (a *stub) DescribeProject(ctx workflow.Context, input *mobile.DescribeProjectInput) (*mobile.DescribeProjectOutput, error) {
	var output mobile.DescribeProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws.mobile.DescribeProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeProjectAsync(ctx workflow.Context, input *mobile.DescribeProjectInput) *MobileDescribeProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mobile.DescribeProject", input)
	return &MobileDescribeProjectFuture{Future: future}
}

func (a *stub) ExportBundle(ctx workflow.Context, input *mobile.ExportBundleInput) (*mobile.ExportBundleOutput, error) {
	var output mobile.ExportBundleOutput
	err := workflow.ExecuteActivity(ctx, "aws.mobile.ExportBundle", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ExportBundleAsync(ctx workflow.Context, input *mobile.ExportBundleInput) *MobileExportBundleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mobile.ExportBundle", input)
	return &MobileExportBundleFuture{Future: future}
}

func (a *stub) ExportProject(ctx workflow.Context, input *mobile.ExportProjectInput) (*mobile.ExportProjectOutput, error) {
	var output mobile.ExportProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws.mobile.ExportProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ExportProjectAsync(ctx workflow.Context, input *mobile.ExportProjectInput) *MobileExportProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mobile.ExportProject", input)
	return &MobileExportProjectFuture{Future: future}
}

func (a *stub) ListBundles(ctx workflow.Context, input *mobile.ListBundlesInput) (*mobile.ListBundlesOutput, error) {
	var output mobile.ListBundlesOutput
	err := workflow.ExecuteActivity(ctx, "aws.mobile.ListBundles", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListBundlesAsync(ctx workflow.Context, input *mobile.ListBundlesInput) *MobileListBundlesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mobile.ListBundles", input)
	return &MobileListBundlesFuture{Future: future}
}

func (a *stub) ListProjects(ctx workflow.Context, input *mobile.ListProjectsInput) (*mobile.ListProjectsOutput, error) {
	var output mobile.ListProjectsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mobile.ListProjects", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListProjectsAsync(ctx workflow.Context, input *mobile.ListProjectsInput) *MobileListProjectsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mobile.ListProjects", input)
	return &MobileListProjectsFuture{Future: future}
}

func (a *stub) UpdateProject(ctx workflow.Context, input *mobile.UpdateProjectInput) (*mobile.UpdateProjectOutput, error) {
	var output mobile.UpdateProjectOutput
	err := workflow.ExecuteActivity(ctx, "aws.mobile.UpdateProject", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateProjectAsync(ctx workflow.Context, input *mobile.UpdateProjectInput) *MobileUpdateProjectFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mobile.UpdateProject", input)
	return &MobileUpdateProjectFuture{Future: future}
}
