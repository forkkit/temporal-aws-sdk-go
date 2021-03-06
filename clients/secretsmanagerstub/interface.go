// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package secretsmanagerstub

import (
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelRotateSecret(ctx workflow.Context, input *secretsmanager.CancelRotateSecretInput) (*secretsmanager.CancelRotateSecretOutput, error)
	CancelRotateSecretAsync(ctx workflow.Context, input *secretsmanager.CancelRotateSecretInput) *SecretsManagerCancelRotateSecretFuture

	CreateSecret(ctx workflow.Context, input *secretsmanager.CreateSecretInput) (*secretsmanager.CreateSecretOutput, error)
	CreateSecretAsync(ctx workflow.Context, input *secretsmanager.CreateSecretInput) *SecretsManagerCreateSecretFuture

	DeleteResourcePolicy(ctx workflow.Context, input *secretsmanager.DeleteResourcePolicyInput) (*secretsmanager.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyAsync(ctx workflow.Context, input *secretsmanager.DeleteResourcePolicyInput) *SecretsManagerDeleteResourcePolicyFuture

	DeleteSecret(ctx workflow.Context, input *secretsmanager.DeleteSecretInput) (*secretsmanager.DeleteSecretOutput, error)
	DeleteSecretAsync(ctx workflow.Context, input *secretsmanager.DeleteSecretInput) *SecretsManagerDeleteSecretFuture

	DescribeSecret(ctx workflow.Context, input *secretsmanager.DescribeSecretInput) (*secretsmanager.DescribeSecretOutput, error)
	DescribeSecretAsync(ctx workflow.Context, input *secretsmanager.DescribeSecretInput) *SecretsManagerDescribeSecretFuture

	GetRandomPassword(ctx workflow.Context, input *secretsmanager.GetRandomPasswordInput) (*secretsmanager.GetRandomPasswordOutput, error)
	GetRandomPasswordAsync(ctx workflow.Context, input *secretsmanager.GetRandomPasswordInput) *SecretsManagerGetRandomPasswordFuture

	GetResourcePolicy(ctx workflow.Context, input *secretsmanager.GetResourcePolicyInput) (*secretsmanager.GetResourcePolicyOutput, error)
	GetResourcePolicyAsync(ctx workflow.Context, input *secretsmanager.GetResourcePolicyInput) *SecretsManagerGetResourcePolicyFuture

	GetSecretValue(ctx workflow.Context, input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error)
	GetSecretValueAsync(ctx workflow.Context, input *secretsmanager.GetSecretValueInput) *SecretsManagerGetSecretValueFuture

	ListSecretVersionIds(ctx workflow.Context, input *secretsmanager.ListSecretVersionIdsInput) (*secretsmanager.ListSecretVersionIdsOutput, error)
	ListSecretVersionIdsAsync(ctx workflow.Context, input *secretsmanager.ListSecretVersionIdsInput) *SecretsManagerListSecretVersionIdsFuture

	ListSecrets(ctx workflow.Context, input *secretsmanager.ListSecretsInput) (*secretsmanager.ListSecretsOutput, error)
	ListSecretsAsync(ctx workflow.Context, input *secretsmanager.ListSecretsInput) *SecretsManagerListSecretsFuture

	PutResourcePolicy(ctx workflow.Context, input *secretsmanager.PutResourcePolicyInput) (*secretsmanager.PutResourcePolicyOutput, error)
	PutResourcePolicyAsync(ctx workflow.Context, input *secretsmanager.PutResourcePolicyInput) *SecretsManagerPutResourcePolicyFuture

	PutSecretValue(ctx workflow.Context, input *secretsmanager.PutSecretValueInput) (*secretsmanager.PutSecretValueOutput, error)
	PutSecretValueAsync(ctx workflow.Context, input *secretsmanager.PutSecretValueInput) *SecretsManagerPutSecretValueFuture

	RestoreSecret(ctx workflow.Context, input *secretsmanager.RestoreSecretInput) (*secretsmanager.RestoreSecretOutput, error)
	RestoreSecretAsync(ctx workflow.Context, input *secretsmanager.RestoreSecretInput) *SecretsManagerRestoreSecretFuture

	RotateSecret(ctx workflow.Context, input *secretsmanager.RotateSecretInput) (*secretsmanager.RotateSecretOutput, error)
	RotateSecretAsync(ctx workflow.Context, input *secretsmanager.RotateSecretInput) *SecretsManagerRotateSecretFuture

	TagResource(ctx workflow.Context, input *secretsmanager.TagResourceInput) (*secretsmanager.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *secretsmanager.TagResourceInput) *SecretsManagerTagResourceFuture

	UntagResource(ctx workflow.Context, input *secretsmanager.UntagResourceInput) (*secretsmanager.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *secretsmanager.UntagResourceInput) *SecretsManagerUntagResourceFuture

	UpdateSecret(ctx workflow.Context, input *secretsmanager.UpdateSecretInput) (*secretsmanager.UpdateSecretOutput, error)
	UpdateSecretAsync(ctx workflow.Context, input *secretsmanager.UpdateSecretInput) *SecretsManagerUpdateSecretFuture

	UpdateSecretVersionStage(ctx workflow.Context, input *secretsmanager.UpdateSecretVersionStageInput) (*secretsmanager.UpdateSecretVersionStageOutput, error)
	UpdateSecretVersionStageAsync(ctx workflow.Context, input *secretsmanager.UpdateSecretVersionStageInput) *SecretsManagerUpdateSecretVersionStageFuture

	ValidateResourcePolicy(ctx workflow.Context, input *secretsmanager.ValidateResourcePolicyInput) (*secretsmanager.ValidateResourcePolicyOutput, error)
	ValidateResourcePolicyAsync(ctx workflow.Context, input *secretsmanager.ValidateResourcePolicyInput) *SecretsManagerValidateResourcePolicyFuture
}

func NewClient() Client {
	return &stub{}
}
