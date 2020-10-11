// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package backup

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/backup"
	"github.com/aws/aws-sdk-go/service/backup/backupiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client backupiface.BackupAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := backup.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (backupiface.BackupAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return backup.New(sess), nil
}

func (a *Activities) CreateBackupPlan(ctx context.Context, input *backup.CreateBackupPlanInput) (*backup.CreateBackupPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateBackupPlanWithContext(ctx, input)
}

func (a *Activities) CreateBackupSelection(ctx context.Context, input *backup.CreateBackupSelectionInput) (*backup.CreateBackupSelectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateBackupSelectionWithContext(ctx, input)
}

func (a *Activities) CreateBackupVault(ctx context.Context, input *backup.CreateBackupVaultInput) (*backup.CreateBackupVaultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateBackupVaultWithContext(ctx, input)
}

func (a *Activities) DeleteBackupPlan(ctx context.Context, input *backup.DeleteBackupPlanInput) (*backup.DeleteBackupPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBackupPlanWithContext(ctx, input)
}

func (a *Activities) DeleteBackupSelection(ctx context.Context, input *backup.DeleteBackupSelectionInput) (*backup.DeleteBackupSelectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBackupSelectionWithContext(ctx, input)
}

func (a *Activities) DeleteBackupVault(ctx context.Context, input *backup.DeleteBackupVaultInput) (*backup.DeleteBackupVaultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBackupVaultWithContext(ctx, input)
}

func (a *Activities) DeleteBackupVaultAccessPolicy(ctx context.Context, input *backup.DeleteBackupVaultAccessPolicyInput) (*backup.DeleteBackupVaultAccessPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBackupVaultAccessPolicyWithContext(ctx, input)
}

func (a *Activities) DeleteBackupVaultNotifications(ctx context.Context, input *backup.DeleteBackupVaultNotificationsInput) (*backup.DeleteBackupVaultNotificationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBackupVaultNotificationsWithContext(ctx, input)
}

func (a *Activities) DeleteRecoveryPoint(ctx context.Context, input *backup.DeleteRecoveryPointInput) (*backup.DeleteRecoveryPointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRecoveryPointWithContext(ctx, input)
}

func (a *Activities) DescribeBackupJob(ctx context.Context, input *backup.DescribeBackupJobInput) (*backup.DescribeBackupJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeBackupJobWithContext(ctx, input)
}

func (a *Activities) DescribeBackupVault(ctx context.Context, input *backup.DescribeBackupVaultInput) (*backup.DescribeBackupVaultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeBackupVaultWithContext(ctx, input)
}

func (a *Activities) DescribeCopyJob(ctx context.Context, input *backup.DescribeCopyJobInput) (*backup.DescribeCopyJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCopyJobWithContext(ctx, input)
}

func (a *Activities) DescribeProtectedResource(ctx context.Context, input *backup.DescribeProtectedResourceInput) (*backup.DescribeProtectedResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeProtectedResourceWithContext(ctx, input)
}

func (a *Activities) DescribeRecoveryPoint(ctx context.Context, input *backup.DescribeRecoveryPointInput) (*backup.DescribeRecoveryPointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeRecoveryPointWithContext(ctx, input)
}

func (a *Activities) DescribeRegionSettings(ctx context.Context, input *backup.DescribeRegionSettingsInput) (*backup.DescribeRegionSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeRegionSettingsWithContext(ctx, input)
}

func (a *Activities) DescribeRestoreJob(ctx context.Context, input *backup.DescribeRestoreJobInput) (*backup.DescribeRestoreJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeRestoreJobWithContext(ctx, input)
}

func (a *Activities) ExportBackupPlanTemplate(ctx context.Context, input *backup.ExportBackupPlanTemplateInput) (*backup.ExportBackupPlanTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ExportBackupPlanTemplateWithContext(ctx, input)
}

func (a *Activities) GetBackupPlan(ctx context.Context, input *backup.GetBackupPlanInput) (*backup.GetBackupPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBackupPlanWithContext(ctx, input)
}

func (a *Activities) GetBackupPlanFromJSON(ctx context.Context, input *backup.GetBackupPlanFromJSONInput) (*backup.GetBackupPlanFromJSONOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBackupPlanFromJSONWithContext(ctx, input)
}

func (a *Activities) GetBackupPlanFromTemplate(ctx context.Context, input *backup.GetBackupPlanFromTemplateInput) (*backup.GetBackupPlanFromTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBackupPlanFromTemplateWithContext(ctx, input)
}

func (a *Activities) GetBackupSelection(ctx context.Context, input *backup.GetBackupSelectionInput) (*backup.GetBackupSelectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBackupSelectionWithContext(ctx, input)
}

func (a *Activities) GetBackupVaultAccessPolicy(ctx context.Context, input *backup.GetBackupVaultAccessPolicyInput) (*backup.GetBackupVaultAccessPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBackupVaultAccessPolicyWithContext(ctx, input)
}

func (a *Activities) GetBackupVaultNotifications(ctx context.Context, input *backup.GetBackupVaultNotificationsInput) (*backup.GetBackupVaultNotificationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBackupVaultNotificationsWithContext(ctx, input)
}

func (a *Activities) GetRecoveryPointRestoreMetadata(ctx context.Context, input *backup.GetRecoveryPointRestoreMetadataInput) (*backup.GetRecoveryPointRestoreMetadataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRecoveryPointRestoreMetadataWithContext(ctx, input)
}

func (a *Activities) GetSupportedResourceTypes(ctx context.Context, input *backup.GetSupportedResourceTypesInput) (*backup.GetSupportedResourceTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSupportedResourceTypesWithContext(ctx, input)
}

func (a *Activities) ListBackupJobs(ctx context.Context, input *backup.ListBackupJobsInput) (*backup.ListBackupJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBackupJobsWithContext(ctx, input)
}

func (a *Activities) ListBackupPlanTemplates(ctx context.Context, input *backup.ListBackupPlanTemplatesInput) (*backup.ListBackupPlanTemplatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBackupPlanTemplatesWithContext(ctx, input)
}

func (a *Activities) ListBackupPlanVersions(ctx context.Context, input *backup.ListBackupPlanVersionsInput) (*backup.ListBackupPlanVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBackupPlanVersionsWithContext(ctx, input)
}

func (a *Activities) ListBackupPlans(ctx context.Context, input *backup.ListBackupPlansInput) (*backup.ListBackupPlansOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBackupPlansWithContext(ctx, input)
}

func (a *Activities) ListBackupSelections(ctx context.Context, input *backup.ListBackupSelectionsInput) (*backup.ListBackupSelectionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBackupSelectionsWithContext(ctx, input)
}

func (a *Activities) ListBackupVaults(ctx context.Context, input *backup.ListBackupVaultsInput) (*backup.ListBackupVaultsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBackupVaultsWithContext(ctx, input)
}

func (a *Activities) ListCopyJobs(ctx context.Context, input *backup.ListCopyJobsInput) (*backup.ListCopyJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListCopyJobsWithContext(ctx, input)
}

func (a *Activities) ListProtectedResources(ctx context.Context, input *backup.ListProtectedResourcesInput) (*backup.ListProtectedResourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListProtectedResourcesWithContext(ctx, input)
}

func (a *Activities) ListRecoveryPointsByBackupVault(ctx context.Context, input *backup.ListRecoveryPointsByBackupVaultInput) (*backup.ListRecoveryPointsByBackupVaultOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRecoveryPointsByBackupVaultWithContext(ctx, input)
}

func (a *Activities) ListRecoveryPointsByResource(ctx context.Context, input *backup.ListRecoveryPointsByResourceInput) (*backup.ListRecoveryPointsByResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRecoveryPointsByResourceWithContext(ctx, input)
}

func (a *Activities) ListRestoreJobs(ctx context.Context, input *backup.ListRestoreJobsInput) (*backup.ListRestoreJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRestoreJobsWithContext(ctx, input)
}

func (a *Activities) ListTags(ctx context.Context, input *backup.ListTagsInput) (*backup.ListTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsWithContext(ctx, input)
}

func (a *Activities) PutBackupVaultAccessPolicy(ctx context.Context, input *backup.PutBackupVaultAccessPolicyInput) (*backup.PutBackupVaultAccessPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBackupVaultAccessPolicyWithContext(ctx, input)
}

func (a *Activities) PutBackupVaultNotifications(ctx context.Context, input *backup.PutBackupVaultNotificationsInput) (*backup.PutBackupVaultNotificationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutBackupVaultNotificationsWithContext(ctx, input)
}

func (a *Activities) StartBackupJob(ctx context.Context, input *backup.StartBackupJobInput) (*backup.StartBackupJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartBackupJobWithContext(ctx, input)
}

func (a *Activities) StartCopyJob(ctx context.Context, input *backup.StartCopyJobInput) (*backup.StartCopyJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartCopyJobWithContext(ctx, input)
}

func (a *Activities) StartRestoreJob(ctx context.Context, input *backup.StartRestoreJobInput) (*backup.StartRestoreJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartRestoreJobWithContext(ctx, input)
}

func (a *Activities) StopBackupJob(ctx context.Context, input *backup.StopBackupJobInput) (*backup.StopBackupJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopBackupJobWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *backup.TagResourceInput) (*backup.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *backup.UntagResourceInput) (*backup.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *Activities) UpdateBackupPlan(ctx context.Context, input *backup.UpdateBackupPlanInput) (*backup.UpdateBackupPlanOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateBackupPlanWithContext(ctx, input)
}

func (a *Activities) UpdateRecoveryPointLifecycle(ctx context.Context, input *backup.UpdateRecoveryPointLifecycleInput) (*backup.UpdateRecoveryPointLifecycleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRecoveryPointLifecycleWithContext(ctx, input)
}

func (a *Activities) UpdateRegionSettings(ctx context.Context, input *backup.UpdateRegionSettingsInput) (*backup.UpdateRegionSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRegionSettingsWithContext(ctx, input)
}
