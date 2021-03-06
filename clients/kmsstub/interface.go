// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package kmsstub

import (
	"github.com/aws/aws-sdk-go/service/kms"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelKeyDeletion(ctx workflow.Context, input *kms.CancelKeyDeletionInput) (*kms.CancelKeyDeletionOutput, error)
	CancelKeyDeletionAsync(ctx workflow.Context, input *kms.CancelKeyDeletionInput) *KMSCancelKeyDeletionFuture

	ConnectCustomKeyStore(ctx workflow.Context, input *kms.ConnectCustomKeyStoreInput) (*kms.ConnectCustomKeyStoreOutput, error)
	ConnectCustomKeyStoreAsync(ctx workflow.Context, input *kms.ConnectCustomKeyStoreInput) *KMSConnectCustomKeyStoreFuture

	CreateAlias(ctx workflow.Context, input *kms.CreateAliasInput) (*kms.CreateAliasOutput, error)
	CreateAliasAsync(ctx workflow.Context, input *kms.CreateAliasInput) *KMSCreateAliasFuture

	CreateCustomKeyStore(ctx workflow.Context, input *kms.CreateCustomKeyStoreInput) (*kms.CreateCustomKeyStoreOutput, error)
	CreateCustomKeyStoreAsync(ctx workflow.Context, input *kms.CreateCustomKeyStoreInput) *KMSCreateCustomKeyStoreFuture

	CreateGrant(ctx workflow.Context, input *kms.CreateGrantInput) (*kms.CreateGrantOutput, error)
	CreateGrantAsync(ctx workflow.Context, input *kms.CreateGrantInput) *KMSCreateGrantFuture

	CreateKey(ctx workflow.Context, input *kms.CreateKeyInput) (*kms.CreateKeyOutput, error)
	CreateKeyAsync(ctx workflow.Context, input *kms.CreateKeyInput) *KMSCreateKeyFuture

	Decrypt(ctx workflow.Context, input *kms.DecryptInput) (*kms.DecryptOutput, error)
	DecryptAsync(ctx workflow.Context, input *kms.DecryptInput) *KMSDecryptFuture

	DeleteAlias(ctx workflow.Context, input *kms.DeleteAliasInput) (*kms.DeleteAliasOutput, error)
	DeleteAliasAsync(ctx workflow.Context, input *kms.DeleteAliasInput) *KMSDeleteAliasFuture

	DeleteCustomKeyStore(ctx workflow.Context, input *kms.DeleteCustomKeyStoreInput) (*kms.DeleteCustomKeyStoreOutput, error)
	DeleteCustomKeyStoreAsync(ctx workflow.Context, input *kms.DeleteCustomKeyStoreInput) *KMSDeleteCustomKeyStoreFuture

	DeleteImportedKeyMaterial(ctx workflow.Context, input *kms.DeleteImportedKeyMaterialInput) (*kms.DeleteImportedKeyMaterialOutput, error)
	DeleteImportedKeyMaterialAsync(ctx workflow.Context, input *kms.DeleteImportedKeyMaterialInput) *KMSDeleteImportedKeyMaterialFuture

	DescribeCustomKeyStores(ctx workflow.Context, input *kms.DescribeCustomKeyStoresInput) (*kms.DescribeCustomKeyStoresOutput, error)
	DescribeCustomKeyStoresAsync(ctx workflow.Context, input *kms.DescribeCustomKeyStoresInput) *KMSDescribeCustomKeyStoresFuture

	DescribeKey(ctx workflow.Context, input *kms.DescribeKeyInput) (*kms.DescribeKeyOutput, error)
	DescribeKeyAsync(ctx workflow.Context, input *kms.DescribeKeyInput) *KMSDescribeKeyFuture

	DisableKey(ctx workflow.Context, input *kms.DisableKeyInput) (*kms.DisableKeyOutput, error)
	DisableKeyAsync(ctx workflow.Context, input *kms.DisableKeyInput) *KMSDisableKeyFuture

	DisableKeyRotation(ctx workflow.Context, input *kms.DisableKeyRotationInput) (*kms.DisableKeyRotationOutput, error)
	DisableKeyRotationAsync(ctx workflow.Context, input *kms.DisableKeyRotationInput) *KMSDisableKeyRotationFuture

	DisconnectCustomKeyStore(ctx workflow.Context, input *kms.DisconnectCustomKeyStoreInput) (*kms.DisconnectCustomKeyStoreOutput, error)
	DisconnectCustomKeyStoreAsync(ctx workflow.Context, input *kms.DisconnectCustomKeyStoreInput) *KMSDisconnectCustomKeyStoreFuture

	EnableKey(ctx workflow.Context, input *kms.EnableKeyInput) (*kms.EnableKeyOutput, error)
	EnableKeyAsync(ctx workflow.Context, input *kms.EnableKeyInput) *KMSEnableKeyFuture

	EnableKeyRotation(ctx workflow.Context, input *kms.EnableKeyRotationInput) (*kms.EnableKeyRotationOutput, error)
	EnableKeyRotationAsync(ctx workflow.Context, input *kms.EnableKeyRotationInput) *KMSEnableKeyRotationFuture

	Encrypt(ctx workflow.Context, input *kms.EncryptInput) (*kms.EncryptOutput, error)
	EncryptAsync(ctx workflow.Context, input *kms.EncryptInput) *KMSEncryptFuture

	GenerateDataKey(ctx workflow.Context, input *kms.GenerateDataKeyInput) (*kms.GenerateDataKeyOutput, error)
	GenerateDataKeyAsync(ctx workflow.Context, input *kms.GenerateDataKeyInput) *KMSGenerateDataKeyFuture

	GenerateDataKeyPair(ctx workflow.Context, input *kms.GenerateDataKeyPairInput) (*kms.GenerateDataKeyPairOutput, error)
	GenerateDataKeyPairAsync(ctx workflow.Context, input *kms.GenerateDataKeyPairInput) *KMSGenerateDataKeyPairFuture

	GenerateDataKeyPairWithoutPlaintext(ctx workflow.Context, input *kms.GenerateDataKeyPairWithoutPlaintextInput) (*kms.GenerateDataKeyPairWithoutPlaintextOutput, error)
	GenerateDataKeyPairWithoutPlaintextAsync(ctx workflow.Context, input *kms.GenerateDataKeyPairWithoutPlaintextInput) *KMSGenerateDataKeyPairWithoutPlaintextFuture

	GenerateDataKeyWithoutPlaintext(ctx workflow.Context, input *kms.GenerateDataKeyWithoutPlaintextInput) (*kms.GenerateDataKeyWithoutPlaintextOutput, error)
	GenerateDataKeyWithoutPlaintextAsync(ctx workflow.Context, input *kms.GenerateDataKeyWithoutPlaintextInput) *KMSGenerateDataKeyWithoutPlaintextFuture

	GenerateRandom(ctx workflow.Context, input *kms.GenerateRandomInput) (*kms.GenerateRandomOutput, error)
	GenerateRandomAsync(ctx workflow.Context, input *kms.GenerateRandomInput) *KMSGenerateRandomFuture

	GetKeyPolicy(ctx workflow.Context, input *kms.GetKeyPolicyInput) (*kms.GetKeyPolicyOutput, error)
	GetKeyPolicyAsync(ctx workflow.Context, input *kms.GetKeyPolicyInput) *KMSGetKeyPolicyFuture

	GetKeyRotationStatus(ctx workflow.Context, input *kms.GetKeyRotationStatusInput) (*kms.GetKeyRotationStatusOutput, error)
	GetKeyRotationStatusAsync(ctx workflow.Context, input *kms.GetKeyRotationStatusInput) *KMSGetKeyRotationStatusFuture

	GetParametersForImport(ctx workflow.Context, input *kms.GetParametersForImportInput) (*kms.GetParametersForImportOutput, error)
	GetParametersForImportAsync(ctx workflow.Context, input *kms.GetParametersForImportInput) *KMSGetParametersForImportFuture

	GetPublicKey(ctx workflow.Context, input *kms.GetPublicKeyInput) (*kms.GetPublicKeyOutput, error)
	GetPublicKeyAsync(ctx workflow.Context, input *kms.GetPublicKeyInput) *KMSGetPublicKeyFuture

	ImportKeyMaterial(ctx workflow.Context, input *kms.ImportKeyMaterialInput) (*kms.ImportKeyMaterialOutput, error)
	ImportKeyMaterialAsync(ctx workflow.Context, input *kms.ImportKeyMaterialInput) *KMSImportKeyMaterialFuture

	ListAliases(ctx workflow.Context, input *kms.ListAliasesInput) (*kms.ListAliasesOutput, error)
	ListAliasesAsync(ctx workflow.Context, input *kms.ListAliasesInput) *KMSListAliasesFuture

	ListGrants(ctx workflow.Context, input *kms.ListGrantsInput) (*kms.ListGrantsResponse, error)
	ListGrantsAsync(ctx workflow.Context, input *kms.ListGrantsInput) *KMSListGrantsFuture

	ListKeyPolicies(ctx workflow.Context, input *kms.ListKeyPoliciesInput) (*kms.ListKeyPoliciesOutput, error)
	ListKeyPoliciesAsync(ctx workflow.Context, input *kms.ListKeyPoliciesInput) *KMSListKeyPoliciesFuture

	ListKeys(ctx workflow.Context, input *kms.ListKeysInput) (*kms.ListKeysOutput, error)
	ListKeysAsync(ctx workflow.Context, input *kms.ListKeysInput) *KMSListKeysFuture

	ListResourceTags(ctx workflow.Context, input *kms.ListResourceTagsInput) (*kms.ListResourceTagsOutput, error)
	ListResourceTagsAsync(ctx workflow.Context, input *kms.ListResourceTagsInput) *KMSListResourceTagsFuture

	ListRetirableGrants(ctx workflow.Context, input *kms.ListRetirableGrantsInput) (*kms.ListGrantsResponse, error)
	ListRetirableGrantsAsync(ctx workflow.Context, input *kms.ListRetirableGrantsInput) *KMSListRetirableGrantsFuture

	PutKeyPolicy(ctx workflow.Context, input *kms.PutKeyPolicyInput) (*kms.PutKeyPolicyOutput, error)
	PutKeyPolicyAsync(ctx workflow.Context, input *kms.PutKeyPolicyInput) *KMSPutKeyPolicyFuture

	ReEncrypt(ctx workflow.Context, input *kms.ReEncryptInput) (*kms.ReEncryptOutput, error)
	ReEncryptAsync(ctx workflow.Context, input *kms.ReEncryptInput) *KMSReEncryptFuture

	RetireGrant(ctx workflow.Context, input *kms.RetireGrantInput) (*kms.RetireGrantOutput, error)
	RetireGrantAsync(ctx workflow.Context, input *kms.RetireGrantInput) *KMSRetireGrantFuture

	RevokeGrant(ctx workflow.Context, input *kms.RevokeGrantInput) (*kms.RevokeGrantOutput, error)
	RevokeGrantAsync(ctx workflow.Context, input *kms.RevokeGrantInput) *KMSRevokeGrantFuture

	ScheduleKeyDeletion(ctx workflow.Context, input *kms.ScheduleKeyDeletionInput) (*kms.ScheduleKeyDeletionOutput, error)
	ScheduleKeyDeletionAsync(ctx workflow.Context, input *kms.ScheduleKeyDeletionInput) *KMSScheduleKeyDeletionFuture

	Sign(ctx workflow.Context, input *kms.SignInput) (*kms.SignOutput, error)
	SignAsync(ctx workflow.Context, input *kms.SignInput) *KMSSignFuture

	TagResource(ctx workflow.Context, input *kms.TagResourceInput) (*kms.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *kms.TagResourceInput) *KMSTagResourceFuture

	UntagResource(ctx workflow.Context, input *kms.UntagResourceInput) (*kms.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *kms.UntagResourceInput) *KMSUntagResourceFuture

	UpdateAlias(ctx workflow.Context, input *kms.UpdateAliasInput) (*kms.UpdateAliasOutput, error)
	UpdateAliasAsync(ctx workflow.Context, input *kms.UpdateAliasInput) *KMSUpdateAliasFuture

	UpdateCustomKeyStore(ctx workflow.Context, input *kms.UpdateCustomKeyStoreInput) (*kms.UpdateCustomKeyStoreOutput, error)
	UpdateCustomKeyStoreAsync(ctx workflow.Context, input *kms.UpdateCustomKeyStoreInput) *KMSUpdateCustomKeyStoreFuture

	UpdateKeyDescription(ctx workflow.Context, input *kms.UpdateKeyDescriptionInput) (*kms.UpdateKeyDescriptionOutput, error)
	UpdateKeyDescriptionAsync(ctx workflow.Context, input *kms.UpdateKeyDescriptionInput) *KMSUpdateKeyDescriptionFuture

	Verify(ctx workflow.Context, input *kms.VerifyInput) (*kms.VerifyOutput, error)
	VerifyAsync(ctx workflow.Context, input *kms.VerifyInput) *KMSVerifyFuture
}

func NewClient() Client {
	return &stub{}
}
