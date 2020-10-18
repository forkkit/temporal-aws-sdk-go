// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mediaconvertstub

import (
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type MediaConvertAssociateCertificateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertAssociateCertificateFuture) Get(ctx workflow.Context) (*mediaconvert.AssociateCertificateOutput, error) {
	var output mediaconvert.AssociateCertificateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertCancelJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertCancelJobFuture) Get(ctx workflow.Context) (*mediaconvert.CancelJobOutput, error) {
	var output mediaconvert.CancelJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertCreateJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertCreateJobFuture) Get(ctx workflow.Context) (*mediaconvert.CreateJobOutput, error) {
	var output mediaconvert.CreateJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertCreateJobTemplateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertCreateJobTemplateFuture) Get(ctx workflow.Context) (*mediaconvert.CreateJobTemplateOutput, error) {
	var output mediaconvert.CreateJobTemplateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertCreatePresetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertCreatePresetFuture) Get(ctx workflow.Context) (*mediaconvert.CreatePresetOutput, error) {
	var output mediaconvert.CreatePresetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertCreateQueueFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertCreateQueueFuture) Get(ctx workflow.Context) (*mediaconvert.CreateQueueOutput, error) {
	var output mediaconvert.CreateQueueOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertDeleteJobTemplateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertDeleteJobTemplateFuture) Get(ctx workflow.Context) (*mediaconvert.DeleteJobTemplateOutput, error) {
	var output mediaconvert.DeleteJobTemplateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertDeletePresetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertDeletePresetFuture) Get(ctx workflow.Context) (*mediaconvert.DeletePresetOutput, error) {
	var output mediaconvert.DeletePresetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertDeleteQueueFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertDeleteQueueFuture) Get(ctx workflow.Context) (*mediaconvert.DeleteQueueOutput, error) {
	var output mediaconvert.DeleteQueueOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertDescribeEndpointsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertDescribeEndpointsFuture) Get(ctx workflow.Context) (*mediaconvert.DescribeEndpointsOutput, error) {
	var output mediaconvert.DescribeEndpointsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertDisassociateCertificateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertDisassociateCertificateFuture) Get(ctx workflow.Context) (*mediaconvert.DisassociateCertificateOutput, error) {
	var output mediaconvert.DisassociateCertificateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertGetJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertGetJobFuture) Get(ctx workflow.Context) (*mediaconvert.GetJobOutput, error) {
	var output mediaconvert.GetJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertGetJobTemplateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertGetJobTemplateFuture) Get(ctx workflow.Context) (*mediaconvert.GetJobTemplateOutput, error) {
	var output mediaconvert.GetJobTemplateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertGetPresetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertGetPresetFuture) Get(ctx workflow.Context) (*mediaconvert.GetPresetOutput, error) {
	var output mediaconvert.GetPresetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertGetQueueFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertGetQueueFuture) Get(ctx workflow.Context) (*mediaconvert.GetQueueOutput, error) {
	var output mediaconvert.GetQueueOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertListJobTemplatesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertListJobTemplatesFuture) Get(ctx workflow.Context) (*mediaconvert.ListJobTemplatesOutput, error) {
	var output mediaconvert.ListJobTemplatesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertListJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertListJobsFuture) Get(ctx workflow.Context) (*mediaconvert.ListJobsOutput, error) {
	var output mediaconvert.ListJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertListPresetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertListPresetsFuture) Get(ctx workflow.Context) (*mediaconvert.ListPresetsOutput, error) {
	var output mediaconvert.ListPresetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertListQueuesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertListQueuesFuture) Get(ctx workflow.Context) (*mediaconvert.ListQueuesOutput, error) {
	var output mediaconvert.ListQueuesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertListTagsForResourceFuture) Get(ctx workflow.Context) (*mediaconvert.ListTagsForResourceOutput, error) {
	var output mediaconvert.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertTagResourceFuture) Get(ctx workflow.Context) (*mediaconvert.TagResourceOutput, error) {
	var output mediaconvert.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertUntagResourceFuture) Get(ctx workflow.Context) (*mediaconvert.UntagResourceOutput, error) {
	var output mediaconvert.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertUpdateJobTemplateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertUpdateJobTemplateFuture) Get(ctx workflow.Context) (*mediaconvert.UpdateJobTemplateOutput, error) {
	var output mediaconvert.UpdateJobTemplateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertUpdatePresetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertUpdatePresetFuture) Get(ctx workflow.Context) (*mediaconvert.UpdatePresetOutput, error) {
	var output mediaconvert.UpdatePresetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaConvertUpdateQueueFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MediaConvertUpdateQueueFuture) Get(ctx workflow.Context) (*mediaconvert.UpdateQueueOutput, error) {
	var output mediaconvert.UpdateQueueOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateCertificate(ctx workflow.Context, input *mediaconvert.AssociateCertificateInput) (*mediaconvert.AssociateCertificateOutput, error) {
	var output mediaconvert.AssociateCertificateOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.AssociateCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateCertificateAsync(ctx workflow.Context, input *mediaconvert.AssociateCertificateInput) *MediaConvertAssociateCertificateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.AssociateCertificate", input)
	return &MediaConvertAssociateCertificateFuture{Future: future}
}

func (a *stub) CancelJob(ctx workflow.Context, input *mediaconvert.CancelJobInput) (*mediaconvert.CancelJobOutput, error) {
	var output mediaconvert.CancelJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.CancelJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CancelJobAsync(ctx workflow.Context, input *mediaconvert.CancelJobInput) *MediaConvertCancelJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.CancelJob", input)
	return &MediaConvertCancelJobFuture{Future: future}
}

func (a *stub) CreateJob(ctx workflow.Context, input *mediaconvert.CreateJobInput) (*mediaconvert.CreateJobOutput, error) {
	var output mediaconvert.CreateJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.CreateJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateJobAsync(ctx workflow.Context, input *mediaconvert.CreateJobInput) *MediaConvertCreateJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.CreateJob", input)
	return &MediaConvertCreateJobFuture{Future: future}
}

func (a *stub) CreateJobTemplate(ctx workflow.Context, input *mediaconvert.CreateJobTemplateInput) (*mediaconvert.CreateJobTemplateOutput, error) {
	var output mediaconvert.CreateJobTemplateOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.CreateJobTemplate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateJobTemplateAsync(ctx workflow.Context, input *mediaconvert.CreateJobTemplateInput) *MediaConvertCreateJobTemplateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.CreateJobTemplate", input)
	return &MediaConvertCreateJobTemplateFuture{Future: future}
}

func (a *stub) CreatePreset(ctx workflow.Context, input *mediaconvert.CreatePresetInput) (*mediaconvert.CreatePresetOutput, error) {
	var output mediaconvert.CreatePresetOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.CreatePreset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreatePresetAsync(ctx workflow.Context, input *mediaconvert.CreatePresetInput) *MediaConvertCreatePresetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.CreatePreset", input)
	return &MediaConvertCreatePresetFuture{Future: future}
}

func (a *stub) CreateQueue(ctx workflow.Context, input *mediaconvert.CreateQueueInput) (*mediaconvert.CreateQueueOutput, error) {
	var output mediaconvert.CreateQueueOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.CreateQueue", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateQueueAsync(ctx workflow.Context, input *mediaconvert.CreateQueueInput) *MediaConvertCreateQueueFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.CreateQueue", input)
	return &MediaConvertCreateQueueFuture{Future: future}
}

func (a *stub) DeleteJobTemplate(ctx workflow.Context, input *mediaconvert.DeleteJobTemplateInput) (*mediaconvert.DeleteJobTemplateOutput, error) {
	var output mediaconvert.DeleteJobTemplateOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.DeleteJobTemplate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteJobTemplateAsync(ctx workflow.Context, input *mediaconvert.DeleteJobTemplateInput) *MediaConvertDeleteJobTemplateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.DeleteJobTemplate", input)
	return &MediaConvertDeleteJobTemplateFuture{Future: future}
}

func (a *stub) DeletePreset(ctx workflow.Context, input *mediaconvert.DeletePresetInput) (*mediaconvert.DeletePresetOutput, error) {
	var output mediaconvert.DeletePresetOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.DeletePreset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePresetAsync(ctx workflow.Context, input *mediaconvert.DeletePresetInput) *MediaConvertDeletePresetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.DeletePreset", input)
	return &MediaConvertDeletePresetFuture{Future: future}
}

func (a *stub) DeleteQueue(ctx workflow.Context, input *mediaconvert.DeleteQueueInput) (*mediaconvert.DeleteQueueOutput, error) {
	var output mediaconvert.DeleteQueueOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.DeleteQueue", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteQueueAsync(ctx workflow.Context, input *mediaconvert.DeleteQueueInput) *MediaConvertDeleteQueueFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.DeleteQueue", input)
	return &MediaConvertDeleteQueueFuture{Future: future}
}

func (a *stub) DescribeEndpoints(ctx workflow.Context, input *mediaconvert.DescribeEndpointsInput) (*mediaconvert.DescribeEndpointsOutput, error) {
	var output mediaconvert.DescribeEndpointsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.DescribeEndpoints", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEndpointsAsync(ctx workflow.Context, input *mediaconvert.DescribeEndpointsInput) *MediaConvertDescribeEndpointsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.DescribeEndpoints", input)
	return &MediaConvertDescribeEndpointsFuture{Future: future}
}

func (a *stub) DisassociateCertificate(ctx workflow.Context, input *mediaconvert.DisassociateCertificateInput) (*mediaconvert.DisassociateCertificateOutput, error) {
	var output mediaconvert.DisassociateCertificateOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.DisassociateCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateCertificateAsync(ctx workflow.Context, input *mediaconvert.DisassociateCertificateInput) *MediaConvertDisassociateCertificateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.DisassociateCertificate", input)
	return &MediaConvertDisassociateCertificateFuture{Future: future}
}

func (a *stub) GetJob(ctx workflow.Context, input *mediaconvert.GetJobInput) (*mediaconvert.GetJobOutput, error) {
	var output mediaconvert.GetJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.GetJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetJobAsync(ctx workflow.Context, input *mediaconvert.GetJobInput) *MediaConvertGetJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.GetJob", input)
	return &MediaConvertGetJobFuture{Future: future}
}

func (a *stub) GetJobTemplate(ctx workflow.Context, input *mediaconvert.GetJobTemplateInput) (*mediaconvert.GetJobTemplateOutput, error) {
	var output mediaconvert.GetJobTemplateOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.GetJobTemplate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetJobTemplateAsync(ctx workflow.Context, input *mediaconvert.GetJobTemplateInput) *MediaConvertGetJobTemplateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.GetJobTemplate", input)
	return &MediaConvertGetJobTemplateFuture{Future: future}
}

func (a *stub) GetPreset(ctx workflow.Context, input *mediaconvert.GetPresetInput) (*mediaconvert.GetPresetOutput, error) {
	var output mediaconvert.GetPresetOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.GetPreset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPresetAsync(ctx workflow.Context, input *mediaconvert.GetPresetInput) *MediaConvertGetPresetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.GetPreset", input)
	return &MediaConvertGetPresetFuture{Future: future}
}

func (a *stub) GetQueue(ctx workflow.Context, input *mediaconvert.GetQueueInput) (*mediaconvert.GetQueueOutput, error) {
	var output mediaconvert.GetQueueOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.GetQueue", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetQueueAsync(ctx workflow.Context, input *mediaconvert.GetQueueInput) *MediaConvertGetQueueFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.GetQueue", input)
	return &MediaConvertGetQueueFuture{Future: future}
}

func (a *stub) ListJobTemplates(ctx workflow.Context, input *mediaconvert.ListJobTemplatesInput) (*mediaconvert.ListJobTemplatesOutput, error) {
	var output mediaconvert.ListJobTemplatesOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.ListJobTemplates", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListJobTemplatesAsync(ctx workflow.Context, input *mediaconvert.ListJobTemplatesInput) *MediaConvertListJobTemplatesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.ListJobTemplates", input)
	return &MediaConvertListJobTemplatesFuture{Future: future}
}

func (a *stub) ListJobs(ctx workflow.Context, input *mediaconvert.ListJobsInput) (*mediaconvert.ListJobsOutput, error) {
	var output mediaconvert.ListJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.ListJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListJobsAsync(ctx workflow.Context, input *mediaconvert.ListJobsInput) *MediaConvertListJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.ListJobs", input)
	return &MediaConvertListJobsFuture{Future: future}
}

func (a *stub) ListPresets(ctx workflow.Context, input *mediaconvert.ListPresetsInput) (*mediaconvert.ListPresetsOutput, error) {
	var output mediaconvert.ListPresetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.ListPresets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPresetsAsync(ctx workflow.Context, input *mediaconvert.ListPresetsInput) *MediaConvertListPresetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.ListPresets", input)
	return &MediaConvertListPresetsFuture{Future: future}
}

func (a *stub) ListQueues(ctx workflow.Context, input *mediaconvert.ListQueuesInput) (*mediaconvert.ListQueuesOutput, error) {
	var output mediaconvert.ListQueuesOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.ListQueues", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListQueuesAsync(ctx workflow.Context, input *mediaconvert.ListQueuesInput) *MediaConvertListQueuesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.ListQueues", input)
	return &MediaConvertListQueuesFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *mediaconvert.ListTagsForResourceInput) (*mediaconvert.ListTagsForResourceOutput, error) {
	var output mediaconvert.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *mediaconvert.ListTagsForResourceInput) *MediaConvertListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.ListTagsForResource", input)
	return &MediaConvertListTagsForResourceFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *mediaconvert.TagResourceInput) (*mediaconvert.TagResourceOutput, error) {
	var output mediaconvert.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *mediaconvert.TagResourceInput) *MediaConvertTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.TagResource", input)
	return &MediaConvertTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *mediaconvert.UntagResourceInput) (*mediaconvert.UntagResourceOutput, error) {
	var output mediaconvert.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *mediaconvert.UntagResourceInput) *MediaConvertUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.UntagResource", input)
	return &MediaConvertUntagResourceFuture{Future: future}
}

func (a *stub) UpdateJobTemplate(ctx workflow.Context, input *mediaconvert.UpdateJobTemplateInput) (*mediaconvert.UpdateJobTemplateOutput, error) {
	var output mediaconvert.UpdateJobTemplateOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.UpdateJobTemplate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateJobTemplateAsync(ctx workflow.Context, input *mediaconvert.UpdateJobTemplateInput) *MediaConvertUpdateJobTemplateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.UpdateJobTemplate", input)
	return &MediaConvertUpdateJobTemplateFuture{Future: future}
}

func (a *stub) UpdatePreset(ctx workflow.Context, input *mediaconvert.UpdatePresetInput) (*mediaconvert.UpdatePresetOutput, error) {
	var output mediaconvert.UpdatePresetOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.UpdatePreset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdatePresetAsync(ctx workflow.Context, input *mediaconvert.UpdatePresetInput) *MediaConvertUpdatePresetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.UpdatePreset", input)
	return &MediaConvertUpdatePresetFuture{Future: future}
}

func (a *stub) UpdateQueue(ctx workflow.Context, input *mediaconvert.UpdateQueueInput) (*mediaconvert.UpdateQueueOutput, error) {
	var output mediaconvert.UpdateQueueOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediaconvert.UpdateQueue", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateQueueAsync(ctx workflow.Context, input *mediaconvert.UpdateQueueInput) *MediaConvertUpdateQueueFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediaconvert.UpdateQueue", input)
	return &MediaConvertUpdateQueueFuture{Future: future}
}