// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package personalizeruntimestub

import (
	"github.com/aws/aws-sdk-go/service/personalizeruntime"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	GetPersonalizedRanking(ctx workflow.Context, input *personalizeruntime.GetPersonalizedRankingInput) (*personalizeruntime.GetPersonalizedRankingOutput, error)
	GetPersonalizedRankingAsync(ctx workflow.Context, input *personalizeruntime.GetPersonalizedRankingInput) *PersonalizeRuntimeGetPersonalizedRankingFuture

	GetRecommendations(ctx workflow.Context, input *personalizeruntime.GetRecommendationsInput) (*personalizeruntime.GetRecommendationsOutput, error)
	GetRecommendationsAsync(ctx workflow.Context, input *personalizeruntime.GetRecommendationsInput) *PersonalizeRuntimeGetRecommendationsFuture
}

func NewClient() Client {
	return &stub{}
}