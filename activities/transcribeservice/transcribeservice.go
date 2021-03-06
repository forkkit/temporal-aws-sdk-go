// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package transcribeservice

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	"github.com/aws/aws-sdk-go/service/transcribeservice/transcribeserviceiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client transcribeserviceiface.TranscribeServiceAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := transcribeservice.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (transcribeserviceiface.TranscribeServiceAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return transcribeservice.New(sess), nil
}

func (a *Activities) CreateLanguageModel(ctx context.Context, input *transcribeservice.CreateLanguageModelInput) (*transcribeservice.CreateLanguageModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateLanguageModelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateMedicalVocabulary(ctx context.Context, input *transcribeservice.CreateMedicalVocabularyInput) (*transcribeservice.CreateMedicalVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateMedicalVocabularyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateVocabulary(ctx context.Context, input *transcribeservice.CreateVocabularyInput) (*transcribeservice.CreateVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateVocabularyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateVocabularyFilter(ctx context.Context, input *transcribeservice.CreateVocabularyFilterInput) (*transcribeservice.CreateVocabularyFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateVocabularyFilterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteLanguageModel(ctx context.Context, input *transcribeservice.DeleteLanguageModelInput) (*transcribeservice.DeleteLanguageModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteLanguageModelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteMedicalTranscriptionJob(ctx context.Context, input *transcribeservice.DeleteMedicalTranscriptionJobInput) (*transcribeservice.DeleteMedicalTranscriptionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteMedicalTranscriptionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteMedicalVocabulary(ctx context.Context, input *transcribeservice.DeleteMedicalVocabularyInput) (*transcribeservice.DeleteMedicalVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteMedicalVocabularyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteTranscriptionJob(ctx context.Context, input *transcribeservice.DeleteTranscriptionJobInput) (*transcribeservice.DeleteTranscriptionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteTranscriptionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteVocabulary(ctx context.Context, input *transcribeservice.DeleteVocabularyInput) (*transcribeservice.DeleteVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteVocabularyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteVocabularyFilter(ctx context.Context, input *transcribeservice.DeleteVocabularyFilterInput) (*transcribeservice.DeleteVocabularyFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteVocabularyFilterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLanguageModel(ctx context.Context, input *transcribeservice.DescribeLanguageModelInput) (*transcribeservice.DescribeLanguageModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLanguageModelWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetMedicalTranscriptionJob(ctx context.Context, input *transcribeservice.GetMedicalTranscriptionJobInput) (*transcribeservice.GetMedicalTranscriptionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetMedicalTranscriptionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetMedicalVocabulary(ctx context.Context, input *transcribeservice.GetMedicalVocabularyInput) (*transcribeservice.GetMedicalVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetMedicalVocabularyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetTranscriptionJob(ctx context.Context, input *transcribeservice.GetTranscriptionJobInput) (*transcribeservice.GetTranscriptionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetTranscriptionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetVocabulary(ctx context.Context, input *transcribeservice.GetVocabularyInput) (*transcribeservice.GetVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetVocabularyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetVocabularyFilter(ctx context.Context, input *transcribeservice.GetVocabularyFilterInput) (*transcribeservice.GetVocabularyFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetVocabularyFilterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListLanguageModels(ctx context.Context, input *transcribeservice.ListLanguageModelsInput) (*transcribeservice.ListLanguageModelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListLanguageModelsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListMedicalTranscriptionJobs(ctx context.Context, input *transcribeservice.ListMedicalTranscriptionJobsInput) (*transcribeservice.ListMedicalTranscriptionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListMedicalTranscriptionJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListMedicalVocabularies(ctx context.Context, input *transcribeservice.ListMedicalVocabulariesInput) (*transcribeservice.ListMedicalVocabulariesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListMedicalVocabulariesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTranscriptionJobs(ctx context.Context, input *transcribeservice.ListTranscriptionJobsInput) (*transcribeservice.ListTranscriptionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTranscriptionJobsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListVocabularies(ctx context.Context, input *transcribeservice.ListVocabulariesInput) (*transcribeservice.ListVocabulariesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListVocabulariesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListVocabularyFilters(ctx context.Context, input *transcribeservice.ListVocabularyFiltersInput) (*transcribeservice.ListVocabularyFiltersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListVocabularyFiltersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartMedicalTranscriptionJob(ctx context.Context, input *transcribeservice.StartMedicalTranscriptionJobInput) (*transcribeservice.StartMedicalTranscriptionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartMedicalTranscriptionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartTranscriptionJob(ctx context.Context, input *transcribeservice.StartTranscriptionJobInput) (*transcribeservice.StartTranscriptionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartTranscriptionJobWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateMedicalVocabulary(ctx context.Context, input *transcribeservice.UpdateMedicalVocabularyInput) (*transcribeservice.UpdateMedicalVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateMedicalVocabularyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateVocabulary(ctx context.Context, input *transcribeservice.UpdateVocabularyInput) (*transcribeservice.UpdateVocabularyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateVocabularyWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateVocabularyFilter(ctx context.Context, input *transcribeservice.UpdateVocabularyFilterInput) (*transcribeservice.UpdateVocabularyFilterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateVocabularyFilterWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
