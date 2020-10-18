// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package servicecatalogstub

import (
	"github.com/aws/aws-sdk-go/service/servicecatalog"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AcceptPortfolioShare(ctx workflow.Context, input *servicecatalog.AcceptPortfolioShareInput) (*servicecatalog.AcceptPortfolioShareOutput, error)
	AcceptPortfolioShareAsync(ctx workflow.Context, input *servicecatalog.AcceptPortfolioShareInput) *ServiceCatalogAcceptPortfolioShareFuture

	AssociateBudgetWithResource(ctx workflow.Context, input *servicecatalog.AssociateBudgetWithResourceInput) (*servicecatalog.AssociateBudgetWithResourceOutput, error)
	AssociateBudgetWithResourceAsync(ctx workflow.Context, input *servicecatalog.AssociateBudgetWithResourceInput) *ServiceCatalogAssociateBudgetWithResourceFuture

	AssociatePrincipalWithPortfolio(ctx workflow.Context, input *servicecatalog.AssociatePrincipalWithPortfolioInput) (*servicecatalog.AssociatePrincipalWithPortfolioOutput, error)
	AssociatePrincipalWithPortfolioAsync(ctx workflow.Context, input *servicecatalog.AssociatePrincipalWithPortfolioInput) *ServiceCatalogAssociatePrincipalWithPortfolioFuture

	AssociateProductWithPortfolio(ctx workflow.Context, input *servicecatalog.AssociateProductWithPortfolioInput) (*servicecatalog.AssociateProductWithPortfolioOutput, error)
	AssociateProductWithPortfolioAsync(ctx workflow.Context, input *servicecatalog.AssociateProductWithPortfolioInput) *ServiceCatalogAssociateProductWithPortfolioFuture

	AssociateServiceActionWithProvisioningArtifact(ctx workflow.Context, input *servicecatalog.AssociateServiceActionWithProvisioningArtifactInput) (*servicecatalog.AssociateServiceActionWithProvisioningArtifactOutput, error)
	AssociateServiceActionWithProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.AssociateServiceActionWithProvisioningArtifactInput) *ServiceCatalogAssociateServiceActionWithProvisioningArtifactFuture

	AssociateTagOptionWithResource(ctx workflow.Context, input *servicecatalog.AssociateTagOptionWithResourceInput) (*servicecatalog.AssociateTagOptionWithResourceOutput, error)
	AssociateTagOptionWithResourceAsync(ctx workflow.Context, input *servicecatalog.AssociateTagOptionWithResourceInput) *ServiceCatalogAssociateTagOptionWithResourceFuture

	BatchAssociateServiceActionWithProvisioningArtifact(ctx workflow.Context, input *servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput) (*servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactOutput, error)
	BatchAssociateServiceActionWithProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput) *ServiceCatalogBatchAssociateServiceActionWithProvisioningArtifactFuture

	BatchDisassociateServiceActionFromProvisioningArtifact(ctx workflow.Context, input *servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput) (*servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactOutput, error)
	BatchDisassociateServiceActionFromProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput) *ServiceCatalogBatchDisassociateServiceActionFromProvisioningArtifactFuture

	CopyProduct(ctx workflow.Context, input *servicecatalog.CopyProductInput) (*servicecatalog.CopyProductOutput, error)
	CopyProductAsync(ctx workflow.Context, input *servicecatalog.CopyProductInput) *ServiceCatalogCopyProductFuture

	CreateConstraint(ctx workflow.Context, input *servicecatalog.CreateConstraintInput) (*servicecatalog.CreateConstraintOutput, error)
	CreateConstraintAsync(ctx workflow.Context, input *servicecatalog.CreateConstraintInput) *ServiceCatalogCreateConstraintFuture

	CreatePortfolio(ctx workflow.Context, input *servicecatalog.CreatePortfolioInput) (*servicecatalog.CreatePortfolioOutput, error)
	CreatePortfolioAsync(ctx workflow.Context, input *servicecatalog.CreatePortfolioInput) *ServiceCatalogCreatePortfolioFuture

	CreatePortfolioShare(ctx workflow.Context, input *servicecatalog.CreatePortfolioShareInput) (*servicecatalog.CreatePortfolioShareOutput, error)
	CreatePortfolioShareAsync(ctx workflow.Context, input *servicecatalog.CreatePortfolioShareInput) *ServiceCatalogCreatePortfolioShareFuture

	CreateProduct(ctx workflow.Context, input *servicecatalog.CreateProductInput) (*servicecatalog.CreateProductOutput, error)
	CreateProductAsync(ctx workflow.Context, input *servicecatalog.CreateProductInput) *ServiceCatalogCreateProductFuture

	CreateProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.CreateProvisionedProductPlanInput) (*servicecatalog.CreateProvisionedProductPlanOutput, error)
	CreateProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.CreateProvisionedProductPlanInput) *ServiceCatalogCreateProvisionedProductPlanFuture

	CreateProvisioningArtifact(ctx workflow.Context, input *servicecatalog.CreateProvisioningArtifactInput) (*servicecatalog.CreateProvisioningArtifactOutput, error)
	CreateProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.CreateProvisioningArtifactInput) *ServiceCatalogCreateProvisioningArtifactFuture

	CreateServiceAction(ctx workflow.Context, input *servicecatalog.CreateServiceActionInput) (*servicecatalog.CreateServiceActionOutput, error)
	CreateServiceActionAsync(ctx workflow.Context, input *servicecatalog.CreateServiceActionInput) *ServiceCatalogCreateServiceActionFuture

	CreateTagOption(ctx workflow.Context, input *servicecatalog.CreateTagOptionInput) (*servicecatalog.CreateTagOptionOutput, error)
	CreateTagOptionAsync(ctx workflow.Context, input *servicecatalog.CreateTagOptionInput) *ServiceCatalogCreateTagOptionFuture

	DeleteConstraint(ctx workflow.Context, input *servicecatalog.DeleteConstraintInput) (*servicecatalog.DeleteConstraintOutput, error)
	DeleteConstraintAsync(ctx workflow.Context, input *servicecatalog.DeleteConstraintInput) *ServiceCatalogDeleteConstraintFuture

	DeletePortfolio(ctx workflow.Context, input *servicecatalog.DeletePortfolioInput) (*servicecatalog.DeletePortfolioOutput, error)
	DeletePortfolioAsync(ctx workflow.Context, input *servicecatalog.DeletePortfolioInput) *ServiceCatalogDeletePortfolioFuture

	DeletePortfolioShare(ctx workflow.Context, input *servicecatalog.DeletePortfolioShareInput) (*servicecatalog.DeletePortfolioShareOutput, error)
	DeletePortfolioShareAsync(ctx workflow.Context, input *servicecatalog.DeletePortfolioShareInput) *ServiceCatalogDeletePortfolioShareFuture

	DeleteProduct(ctx workflow.Context, input *servicecatalog.DeleteProductInput) (*servicecatalog.DeleteProductOutput, error)
	DeleteProductAsync(ctx workflow.Context, input *servicecatalog.DeleteProductInput) *ServiceCatalogDeleteProductFuture

	DeleteProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.DeleteProvisionedProductPlanInput) (*servicecatalog.DeleteProvisionedProductPlanOutput, error)
	DeleteProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.DeleteProvisionedProductPlanInput) *ServiceCatalogDeleteProvisionedProductPlanFuture

	DeleteProvisioningArtifact(ctx workflow.Context, input *servicecatalog.DeleteProvisioningArtifactInput) (*servicecatalog.DeleteProvisioningArtifactOutput, error)
	DeleteProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.DeleteProvisioningArtifactInput) *ServiceCatalogDeleteProvisioningArtifactFuture

	DeleteServiceAction(ctx workflow.Context, input *servicecatalog.DeleteServiceActionInput) (*servicecatalog.DeleteServiceActionOutput, error)
	DeleteServiceActionAsync(ctx workflow.Context, input *servicecatalog.DeleteServiceActionInput) *ServiceCatalogDeleteServiceActionFuture

	DeleteTagOption(ctx workflow.Context, input *servicecatalog.DeleteTagOptionInput) (*servicecatalog.DeleteTagOptionOutput, error)
	DeleteTagOptionAsync(ctx workflow.Context, input *servicecatalog.DeleteTagOptionInput) *ServiceCatalogDeleteTagOptionFuture

	DescribeConstraint(ctx workflow.Context, input *servicecatalog.DescribeConstraintInput) (*servicecatalog.DescribeConstraintOutput, error)
	DescribeConstraintAsync(ctx workflow.Context, input *servicecatalog.DescribeConstraintInput) *ServiceCatalogDescribeConstraintFuture

	DescribeCopyProductStatus(ctx workflow.Context, input *servicecatalog.DescribeCopyProductStatusInput) (*servicecatalog.DescribeCopyProductStatusOutput, error)
	DescribeCopyProductStatusAsync(ctx workflow.Context, input *servicecatalog.DescribeCopyProductStatusInput) *ServiceCatalogDescribeCopyProductStatusFuture

	DescribePortfolio(ctx workflow.Context, input *servicecatalog.DescribePortfolioInput) (*servicecatalog.DescribePortfolioOutput, error)
	DescribePortfolioAsync(ctx workflow.Context, input *servicecatalog.DescribePortfolioInput) *ServiceCatalogDescribePortfolioFuture

	DescribePortfolioShareStatus(ctx workflow.Context, input *servicecatalog.DescribePortfolioShareStatusInput) (*servicecatalog.DescribePortfolioShareStatusOutput, error)
	DescribePortfolioShareStatusAsync(ctx workflow.Context, input *servicecatalog.DescribePortfolioShareStatusInput) *ServiceCatalogDescribePortfolioShareStatusFuture

	DescribeProduct(ctx workflow.Context, input *servicecatalog.DescribeProductInput) (*servicecatalog.DescribeProductOutput, error)
	DescribeProductAsync(ctx workflow.Context, input *servicecatalog.DescribeProductInput) *ServiceCatalogDescribeProductFuture

	DescribeProductAsAdmin(ctx workflow.Context, input *servicecatalog.DescribeProductAsAdminInput) (*servicecatalog.DescribeProductAsAdminOutput, error)
	DescribeProductAsAdminAsync(ctx workflow.Context, input *servicecatalog.DescribeProductAsAdminInput) *ServiceCatalogDescribeProductAsAdminFuture

	DescribeProductView(ctx workflow.Context, input *servicecatalog.DescribeProductViewInput) (*servicecatalog.DescribeProductViewOutput, error)
	DescribeProductViewAsync(ctx workflow.Context, input *servicecatalog.DescribeProductViewInput) *ServiceCatalogDescribeProductViewFuture

	DescribeProvisionedProduct(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductInput) (*servicecatalog.DescribeProvisionedProductOutput, error)
	DescribeProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductInput) *ServiceCatalogDescribeProvisionedProductFuture

	DescribeProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductPlanInput) (*servicecatalog.DescribeProvisionedProductPlanOutput, error)
	DescribeProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductPlanInput) *ServiceCatalogDescribeProvisionedProductPlanFuture

	DescribeProvisioningArtifact(ctx workflow.Context, input *servicecatalog.DescribeProvisioningArtifactInput) (*servicecatalog.DescribeProvisioningArtifactOutput, error)
	DescribeProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisioningArtifactInput) *ServiceCatalogDescribeProvisioningArtifactFuture

	DescribeProvisioningParameters(ctx workflow.Context, input *servicecatalog.DescribeProvisioningParametersInput) (*servicecatalog.DescribeProvisioningParametersOutput, error)
	DescribeProvisioningParametersAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisioningParametersInput) *ServiceCatalogDescribeProvisioningParametersFuture

	DescribeRecord(ctx workflow.Context, input *servicecatalog.DescribeRecordInput) (*servicecatalog.DescribeRecordOutput, error)
	DescribeRecordAsync(ctx workflow.Context, input *servicecatalog.DescribeRecordInput) *ServiceCatalogDescribeRecordFuture

	DescribeServiceAction(ctx workflow.Context, input *servicecatalog.DescribeServiceActionInput) (*servicecatalog.DescribeServiceActionOutput, error)
	DescribeServiceActionAsync(ctx workflow.Context, input *servicecatalog.DescribeServiceActionInput) *ServiceCatalogDescribeServiceActionFuture

	DescribeServiceActionExecutionParameters(ctx workflow.Context, input *servicecatalog.DescribeServiceActionExecutionParametersInput) (*servicecatalog.DescribeServiceActionExecutionParametersOutput, error)
	DescribeServiceActionExecutionParametersAsync(ctx workflow.Context, input *servicecatalog.DescribeServiceActionExecutionParametersInput) *ServiceCatalogDescribeServiceActionExecutionParametersFuture

	DescribeTagOption(ctx workflow.Context, input *servicecatalog.DescribeTagOptionInput) (*servicecatalog.DescribeTagOptionOutput, error)
	DescribeTagOptionAsync(ctx workflow.Context, input *servicecatalog.DescribeTagOptionInput) *ServiceCatalogDescribeTagOptionFuture

	DisableAWSOrganizationsAccess(ctx workflow.Context, input *servicecatalog.DisableAWSOrganizationsAccessInput) (*servicecatalog.DisableAWSOrganizationsAccessOutput, error)
	DisableAWSOrganizationsAccessAsync(ctx workflow.Context, input *servicecatalog.DisableAWSOrganizationsAccessInput) *ServiceCatalogDisableAWSOrganizationsAccessFuture

	DisassociateBudgetFromResource(ctx workflow.Context, input *servicecatalog.DisassociateBudgetFromResourceInput) (*servicecatalog.DisassociateBudgetFromResourceOutput, error)
	DisassociateBudgetFromResourceAsync(ctx workflow.Context, input *servicecatalog.DisassociateBudgetFromResourceInput) *ServiceCatalogDisassociateBudgetFromResourceFuture

	DisassociatePrincipalFromPortfolio(ctx workflow.Context, input *servicecatalog.DisassociatePrincipalFromPortfolioInput) (*servicecatalog.DisassociatePrincipalFromPortfolioOutput, error)
	DisassociatePrincipalFromPortfolioAsync(ctx workflow.Context, input *servicecatalog.DisassociatePrincipalFromPortfolioInput) *ServiceCatalogDisassociatePrincipalFromPortfolioFuture

	DisassociateProductFromPortfolio(ctx workflow.Context, input *servicecatalog.DisassociateProductFromPortfolioInput) (*servicecatalog.DisassociateProductFromPortfolioOutput, error)
	DisassociateProductFromPortfolioAsync(ctx workflow.Context, input *servicecatalog.DisassociateProductFromPortfolioInput) *ServiceCatalogDisassociateProductFromPortfolioFuture

	DisassociateServiceActionFromProvisioningArtifact(ctx workflow.Context, input *servicecatalog.DisassociateServiceActionFromProvisioningArtifactInput) (*servicecatalog.DisassociateServiceActionFromProvisioningArtifactOutput, error)
	DisassociateServiceActionFromProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.DisassociateServiceActionFromProvisioningArtifactInput) *ServiceCatalogDisassociateServiceActionFromProvisioningArtifactFuture

	DisassociateTagOptionFromResource(ctx workflow.Context, input *servicecatalog.DisassociateTagOptionFromResourceInput) (*servicecatalog.DisassociateTagOptionFromResourceOutput, error)
	DisassociateTagOptionFromResourceAsync(ctx workflow.Context, input *servicecatalog.DisassociateTagOptionFromResourceInput) *ServiceCatalogDisassociateTagOptionFromResourceFuture

	EnableAWSOrganizationsAccess(ctx workflow.Context, input *servicecatalog.EnableAWSOrganizationsAccessInput) (*servicecatalog.EnableAWSOrganizationsAccessOutput, error)
	EnableAWSOrganizationsAccessAsync(ctx workflow.Context, input *servicecatalog.EnableAWSOrganizationsAccessInput) *ServiceCatalogEnableAWSOrganizationsAccessFuture

	ExecuteProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductPlanInput) (*servicecatalog.ExecuteProvisionedProductPlanOutput, error)
	ExecuteProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductPlanInput) *ServiceCatalogExecuteProvisionedProductPlanFuture

	ExecuteProvisionedProductServiceAction(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductServiceActionInput) (*servicecatalog.ExecuteProvisionedProductServiceActionOutput, error)
	ExecuteProvisionedProductServiceActionAsync(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductServiceActionInput) *ServiceCatalogExecuteProvisionedProductServiceActionFuture

	GetAWSOrganizationsAccessStatus(ctx workflow.Context, input *servicecatalog.GetAWSOrganizationsAccessStatusInput) (*servicecatalog.GetAWSOrganizationsAccessStatusOutput, error)
	GetAWSOrganizationsAccessStatusAsync(ctx workflow.Context, input *servicecatalog.GetAWSOrganizationsAccessStatusInput) *ServiceCatalogGetAWSOrganizationsAccessStatusFuture

	GetProvisionedProductOutputs(ctx workflow.Context, input *servicecatalog.GetProvisionedProductOutputsInput) (*servicecatalog.GetProvisionedProductOutputsOutput, error)
	GetProvisionedProductOutputsAsync(ctx workflow.Context, input *servicecatalog.GetProvisionedProductOutputsInput) *ServiceCatalogGetProvisionedProductOutputsFuture

	ListAcceptedPortfolioShares(ctx workflow.Context, input *servicecatalog.ListAcceptedPortfolioSharesInput) (*servicecatalog.ListAcceptedPortfolioSharesOutput, error)
	ListAcceptedPortfolioSharesAsync(ctx workflow.Context, input *servicecatalog.ListAcceptedPortfolioSharesInput) *ServiceCatalogListAcceptedPortfolioSharesFuture

	ListBudgetsForResource(ctx workflow.Context, input *servicecatalog.ListBudgetsForResourceInput) (*servicecatalog.ListBudgetsForResourceOutput, error)
	ListBudgetsForResourceAsync(ctx workflow.Context, input *servicecatalog.ListBudgetsForResourceInput) *ServiceCatalogListBudgetsForResourceFuture

	ListConstraintsForPortfolio(ctx workflow.Context, input *servicecatalog.ListConstraintsForPortfolioInput) (*servicecatalog.ListConstraintsForPortfolioOutput, error)
	ListConstraintsForPortfolioAsync(ctx workflow.Context, input *servicecatalog.ListConstraintsForPortfolioInput) *ServiceCatalogListConstraintsForPortfolioFuture

	ListLaunchPaths(ctx workflow.Context, input *servicecatalog.ListLaunchPathsInput) (*servicecatalog.ListLaunchPathsOutput, error)
	ListLaunchPathsAsync(ctx workflow.Context, input *servicecatalog.ListLaunchPathsInput) *ServiceCatalogListLaunchPathsFuture

	ListOrganizationPortfolioAccess(ctx workflow.Context, input *servicecatalog.ListOrganizationPortfolioAccessInput) (*servicecatalog.ListOrganizationPortfolioAccessOutput, error)
	ListOrganizationPortfolioAccessAsync(ctx workflow.Context, input *servicecatalog.ListOrganizationPortfolioAccessInput) *ServiceCatalogListOrganizationPortfolioAccessFuture

	ListPortfolioAccess(ctx workflow.Context, input *servicecatalog.ListPortfolioAccessInput) (*servicecatalog.ListPortfolioAccessOutput, error)
	ListPortfolioAccessAsync(ctx workflow.Context, input *servicecatalog.ListPortfolioAccessInput) *ServiceCatalogListPortfolioAccessFuture

	ListPortfolios(ctx workflow.Context, input *servicecatalog.ListPortfoliosInput) (*servicecatalog.ListPortfoliosOutput, error)
	ListPortfoliosAsync(ctx workflow.Context, input *servicecatalog.ListPortfoliosInput) *ServiceCatalogListPortfoliosFuture

	ListPortfoliosForProduct(ctx workflow.Context, input *servicecatalog.ListPortfoliosForProductInput) (*servicecatalog.ListPortfoliosForProductOutput, error)
	ListPortfoliosForProductAsync(ctx workflow.Context, input *servicecatalog.ListPortfoliosForProductInput) *ServiceCatalogListPortfoliosForProductFuture

	ListPrincipalsForPortfolio(ctx workflow.Context, input *servicecatalog.ListPrincipalsForPortfolioInput) (*servicecatalog.ListPrincipalsForPortfolioOutput, error)
	ListPrincipalsForPortfolioAsync(ctx workflow.Context, input *servicecatalog.ListPrincipalsForPortfolioInput) *ServiceCatalogListPrincipalsForPortfolioFuture

	ListProvisionedProductPlans(ctx workflow.Context, input *servicecatalog.ListProvisionedProductPlansInput) (*servicecatalog.ListProvisionedProductPlansOutput, error)
	ListProvisionedProductPlansAsync(ctx workflow.Context, input *servicecatalog.ListProvisionedProductPlansInput) *ServiceCatalogListProvisionedProductPlansFuture

	ListProvisioningArtifacts(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsInput) (*servicecatalog.ListProvisioningArtifactsOutput, error)
	ListProvisioningArtifactsAsync(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsInput) *ServiceCatalogListProvisioningArtifactsFuture

	ListProvisioningArtifactsForServiceAction(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsForServiceActionInput) (*servicecatalog.ListProvisioningArtifactsForServiceActionOutput, error)
	ListProvisioningArtifactsForServiceActionAsync(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsForServiceActionInput) *ServiceCatalogListProvisioningArtifactsForServiceActionFuture

	ListRecordHistory(ctx workflow.Context, input *servicecatalog.ListRecordHistoryInput) (*servicecatalog.ListRecordHistoryOutput, error)
	ListRecordHistoryAsync(ctx workflow.Context, input *servicecatalog.ListRecordHistoryInput) *ServiceCatalogListRecordHistoryFuture

	ListResourcesForTagOption(ctx workflow.Context, input *servicecatalog.ListResourcesForTagOptionInput) (*servicecatalog.ListResourcesForTagOptionOutput, error)
	ListResourcesForTagOptionAsync(ctx workflow.Context, input *servicecatalog.ListResourcesForTagOptionInput) *ServiceCatalogListResourcesForTagOptionFuture

	ListServiceActions(ctx workflow.Context, input *servicecatalog.ListServiceActionsInput) (*servicecatalog.ListServiceActionsOutput, error)
	ListServiceActionsAsync(ctx workflow.Context, input *servicecatalog.ListServiceActionsInput) *ServiceCatalogListServiceActionsFuture

	ListServiceActionsForProvisioningArtifact(ctx workflow.Context, input *servicecatalog.ListServiceActionsForProvisioningArtifactInput) (*servicecatalog.ListServiceActionsForProvisioningArtifactOutput, error)
	ListServiceActionsForProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.ListServiceActionsForProvisioningArtifactInput) *ServiceCatalogListServiceActionsForProvisioningArtifactFuture

	ListStackInstancesForProvisionedProduct(ctx workflow.Context, input *servicecatalog.ListStackInstancesForProvisionedProductInput) (*servicecatalog.ListStackInstancesForProvisionedProductOutput, error)
	ListStackInstancesForProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.ListStackInstancesForProvisionedProductInput) *ServiceCatalogListStackInstancesForProvisionedProductFuture

	ListTagOptions(ctx workflow.Context, input *servicecatalog.ListTagOptionsInput) (*servicecatalog.ListTagOptionsOutput, error)
	ListTagOptionsAsync(ctx workflow.Context, input *servicecatalog.ListTagOptionsInput) *ServiceCatalogListTagOptionsFuture

	ProvisionProduct(ctx workflow.Context, input *servicecatalog.ProvisionProductInput) (*servicecatalog.ProvisionProductOutput, error)
	ProvisionProductAsync(ctx workflow.Context, input *servicecatalog.ProvisionProductInput) *ServiceCatalogProvisionProductFuture

	RejectPortfolioShare(ctx workflow.Context, input *servicecatalog.RejectPortfolioShareInput) (*servicecatalog.RejectPortfolioShareOutput, error)
	RejectPortfolioShareAsync(ctx workflow.Context, input *servicecatalog.RejectPortfolioShareInput) *ServiceCatalogRejectPortfolioShareFuture

	ScanProvisionedProducts(ctx workflow.Context, input *servicecatalog.ScanProvisionedProductsInput) (*servicecatalog.ScanProvisionedProductsOutput, error)
	ScanProvisionedProductsAsync(ctx workflow.Context, input *servicecatalog.ScanProvisionedProductsInput) *ServiceCatalogScanProvisionedProductsFuture

	SearchProducts(ctx workflow.Context, input *servicecatalog.SearchProductsInput) (*servicecatalog.SearchProductsOutput, error)
	SearchProductsAsync(ctx workflow.Context, input *servicecatalog.SearchProductsInput) *ServiceCatalogSearchProductsFuture

	SearchProductsAsAdmin(ctx workflow.Context, input *servicecatalog.SearchProductsAsAdminInput) (*servicecatalog.SearchProductsAsAdminOutput, error)
	SearchProductsAsAdminAsync(ctx workflow.Context, input *servicecatalog.SearchProductsAsAdminInput) *ServiceCatalogSearchProductsAsAdminFuture

	SearchProvisionedProducts(ctx workflow.Context, input *servicecatalog.SearchProvisionedProductsInput) (*servicecatalog.SearchProvisionedProductsOutput, error)
	SearchProvisionedProductsAsync(ctx workflow.Context, input *servicecatalog.SearchProvisionedProductsInput) *ServiceCatalogSearchProvisionedProductsFuture

	TerminateProvisionedProduct(ctx workflow.Context, input *servicecatalog.TerminateProvisionedProductInput) (*servicecatalog.TerminateProvisionedProductOutput, error)
	TerminateProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.TerminateProvisionedProductInput) *ServiceCatalogTerminateProvisionedProductFuture

	UpdateConstraint(ctx workflow.Context, input *servicecatalog.UpdateConstraintInput) (*servicecatalog.UpdateConstraintOutput, error)
	UpdateConstraintAsync(ctx workflow.Context, input *servicecatalog.UpdateConstraintInput) *ServiceCatalogUpdateConstraintFuture

	UpdatePortfolio(ctx workflow.Context, input *servicecatalog.UpdatePortfolioInput) (*servicecatalog.UpdatePortfolioOutput, error)
	UpdatePortfolioAsync(ctx workflow.Context, input *servicecatalog.UpdatePortfolioInput) *ServiceCatalogUpdatePortfolioFuture

	UpdateProduct(ctx workflow.Context, input *servicecatalog.UpdateProductInput) (*servicecatalog.UpdateProductOutput, error)
	UpdateProductAsync(ctx workflow.Context, input *servicecatalog.UpdateProductInput) *ServiceCatalogUpdateProductFuture

	UpdateProvisionedProduct(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductInput) (*servicecatalog.UpdateProvisionedProductOutput, error)
	UpdateProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductInput) *ServiceCatalogUpdateProvisionedProductFuture

	UpdateProvisionedProductProperties(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductPropertiesInput) (*servicecatalog.UpdateProvisionedProductPropertiesOutput, error)
	UpdateProvisionedProductPropertiesAsync(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductPropertiesInput) *ServiceCatalogUpdateProvisionedProductPropertiesFuture

	UpdateProvisioningArtifact(ctx workflow.Context, input *servicecatalog.UpdateProvisioningArtifactInput) (*servicecatalog.UpdateProvisioningArtifactOutput, error)
	UpdateProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.UpdateProvisioningArtifactInput) *ServiceCatalogUpdateProvisioningArtifactFuture

	UpdateServiceAction(ctx workflow.Context, input *servicecatalog.UpdateServiceActionInput) (*servicecatalog.UpdateServiceActionOutput, error)
	UpdateServiceActionAsync(ctx workflow.Context, input *servicecatalog.UpdateServiceActionInput) *ServiceCatalogUpdateServiceActionFuture

	UpdateTagOption(ctx workflow.Context, input *servicecatalog.UpdateTagOptionInput) (*servicecatalog.UpdateTagOptionOutput, error)
	UpdateTagOptionAsync(ctx workflow.Context, input *servicecatalog.UpdateTagOptionInput) *ServiceCatalogUpdateTagOptionFuture
}

func NewClient() Client {
	return &stub{}
}