// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package marketplacecommerceanalyticsstub

import (
	"github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	GenerateDataSet(ctx workflow.Context, input *marketplacecommerceanalytics.GenerateDataSetInput) (*marketplacecommerceanalytics.GenerateDataSetOutput, error)
	GenerateDataSetAsync(ctx workflow.Context, input *marketplacecommerceanalytics.GenerateDataSetInput) *MarketplaceCommerceAnalyticsGenerateDataSetFuture

	StartSupportDataExport(ctx workflow.Context, input *marketplacecommerceanalytics.StartSupportDataExportInput) (*marketplacecommerceanalytics.StartSupportDataExportOutput, error)
	StartSupportDataExportAsync(ctx workflow.Context, input *marketplacecommerceanalytics.StartSupportDataExportInput) *MarketplaceCommerceAnalyticsStartSupportDataExportFuture
}

func NewClient() Client {
	return &stub{}
}
