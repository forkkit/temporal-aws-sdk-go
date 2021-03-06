// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package worklinkstub

import (
	"github.com/aws/aws-sdk-go/service/worklink"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type WorkLinkAssociateDomainFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkAssociateDomainFuture) Get(ctx workflow.Context) (*worklink.AssociateDomainOutput, error) {
	var output worklink.AssociateDomainOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkAssociateWebsiteAuthorizationProviderFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkAssociateWebsiteAuthorizationProviderFuture) Get(ctx workflow.Context) (*worklink.AssociateWebsiteAuthorizationProviderOutput, error) {
	var output worklink.AssociateWebsiteAuthorizationProviderOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkAssociateWebsiteCertificateAuthorityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkAssociateWebsiteCertificateAuthorityFuture) Get(ctx workflow.Context) (*worklink.AssociateWebsiteCertificateAuthorityOutput, error) {
	var output worklink.AssociateWebsiteCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkCreateFleetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkCreateFleetFuture) Get(ctx workflow.Context) (*worklink.CreateFleetOutput, error) {
	var output worklink.CreateFleetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkDeleteFleetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkDeleteFleetFuture) Get(ctx workflow.Context) (*worklink.DeleteFleetOutput, error) {
	var output worklink.DeleteFleetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkDescribeAuditStreamConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkDescribeAuditStreamConfigurationFuture) Get(ctx workflow.Context) (*worklink.DescribeAuditStreamConfigurationOutput, error) {
	var output worklink.DescribeAuditStreamConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkDescribeCompanyNetworkConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkDescribeCompanyNetworkConfigurationFuture) Get(ctx workflow.Context) (*worklink.DescribeCompanyNetworkConfigurationOutput, error) {
	var output worklink.DescribeCompanyNetworkConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkDescribeDeviceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkDescribeDeviceFuture) Get(ctx workflow.Context) (*worklink.DescribeDeviceOutput, error) {
	var output worklink.DescribeDeviceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkDescribeDevicePolicyConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkDescribeDevicePolicyConfigurationFuture) Get(ctx workflow.Context) (*worklink.DescribeDevicePolicyConfigurationOutput, error) {
	var output worklink.DescribeDevicePolicyConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkDescribeDomainFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkDescribeDomainFuture) Get(ctx workflow.Context) (*worklink.DescribeDomainOutput, error) {
	var output worklink.DescribeDomainOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkDescribeFleetMetadataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkDescribeFleetMetadataFuture) Get(ctx workflow.Context) (*worklink.DescribeFleetMetadataOutput, error) {
	var output worklink.DescribeFleetMetadataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkDescribeIdentityProviderConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkDescribeIdentityProviderConfigurationFuture) Get(ctx workflow.Context) (*worklink.DescribeIdentityProviderConfigurationOutput, error) {
	var output worklink.DescribeIdentityProviderConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkDescribeWebsiteCertificateAuthorityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkDescribeWebsiteCertificateAuthorityFuture) Get(ctx workflow.Context) (*worklink.DescribeWebsiteCertificateAuthorityOutput, error) {
	var output worklink.DescribeWebsiteCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkDisassociateDomainFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkDisassociateDomainFuture) Get(ctx workflow.Context) (*worklink.DisassociateDomainOutput, error) {
	var output worklink.DisassociateDomainOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkDisassociateWebsiteAuthorizationProviderFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkDisassociateWebsiteAuthorizationProviderFuture) Get(ctx workflow.Context) (*worklink.DisassociateWebsiteAuthorizationProviderOutput, error) {
	var output worklink.DisassociateWebsiteAuthorizationProviderOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkDisassociateWebsiteCertificateAuthorityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkDisassociateWebsiteCertificateAuthorityFuture) Get(ctx workflow.Context) (*worklink.DisassociateWebsiteCertificateAuthorityOutput, error) {
	var output worklink.DisassociateWebsiteCertificateAuthorityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkListDevicesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkListDevicesFuture) Get(ctx workflow.Context) (*worklink.ListDevicesOutput, error) {
	var output worklink.ListDevicesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkListDomainsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkListDomainsFuture) Get(ctx workflow.Context) (*worklink.ListDomainsOutput, error) {
	var output worklink.ListDomainsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkListFleetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkListFleetsFuture) Get(ctx workflow.Context) (*worklink.ListFleetsOutput, error) {
	var output worklink.ListFleetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkListTagsForResourceFuture) Get(ctx workflow.Context) (*worklink.ListTagsForResourceOutput, error) {
	var output worklink.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkListWebsiteAuthorizationProvidersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkListWebsiteAuthorizationProvidersFuture) Get(ctx workflow.Context) (*worklink.ListWebsiteAuthorizationProvidersOutput, error) {
	var output worklink.ListWebsiteAuthorizationProvidersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkListWebsiteCertificateAuthoritiesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkListWebsiteCertificateAuthoritiesFuture) Get(ctx workflow.Context) (*worklink.ListWebsiteCertificateAuthoritiesOutput, error) {
	var output worklink.ListWebsiteCertificateAuthoritiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkRestoreDomainAccessFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkRestoreDomainAccessFuture) Get(ctx workflow.Context) (*worklink.RestoreDomainAccessOutput, error) {
	var output worklink.RestoreDomainAccessOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkRevokeDomainAccessFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkRevokeDomainAccessFuture) Get(ctx workflow.Context) (*worklink.RevokeDomainAccessOutput, error) {
	var output worklink.RevokeDomainAccessOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkSignOutUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkSignOutUserFuture) Get(ctx workflow.Context) (*worklink.SignOutUserOutput, error) {
	var output worklink.SignOutUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkTagResourceFuture) Get(ctx workflow.Context) (*worklink.TagResourceOutput, error) {
	var output worklink.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkUntagResourceFuture) Get(ctx workflow.Context) (*worklink.UntagResourceOutput, error) {
	var output worklink.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkUpdateAuditStreamConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkUpdateAuditStreamConfigurationFuture) Get(ctx workflow.Context) (*worklink.UpdateAuditStreamConfigurationOutput, error) {
	var output worklink.UpdateAuditStreamConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkUpdateCompanyNetworkConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkUpdateCompanyNetworkConfigurationFuture) Get(ctx workflow.Context) (*worklink.UpdateCompanyNetworkConfigurationOutput, error) {
	var output worklink.UpdateCompanyNetworkConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkUpdateDevicePolicyConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkUpdateDevicePolicyConfigurationFuture) Get(ctx workflow.Context) (*worklink.UpdateDevicePolicyConfigurationOutput, error) {
	var output worklink.UpdateDevicePolicyConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkUpdateDomainMetadataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkUpdateDomainMetadataFuture) Get(ctx workflow.Context) (*worklink.UpdateDomainMetadataOutput, error) {
	var output worklink.UpdateDomainMetadataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkUpdateFleetMetadataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkUpdateFleetMetadataFuture) Get(ctx workflow.Context) (*worklink.UpdateFleetMetadataOutput, error) {
	var output worklink.UpdateFleetMetadataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type WorkLinkUpdateIdentityProviderConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *WorkLinkUpdateIdentityProviderConfigurationFuture) Get(ctx workflow.Context) (*worklink.UpdateIdentityProviderConfigurationOutput, error) {
	var output worklink.UpdateIdentityProviderConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateDomain(ctx workflow.Context, input *worklink.AssociateDomainInput) (*worklink.AssociateDomainOutput, error) {
	var output worklink.AssociateDomainOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.AssociateDomain", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateDomainAsync(ctx workflow.Context, input *worklink.AssociateDomainInput) *WorkLinkAssociateDomainFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.AssociateDomain", input)
	return &WorkLinkAssociateDomainFuture{Future: future}
}

func (a *stub) AssociateWebsiteAuthorizationProvider(ctx workflow.Context, input *worklink.AssociateWebsiteAuthorizationProviderInput) (*worklink.AssociateWebsiteAuthorizationProviderOutput, error) {
	var output worklink.AssociateWebsiteAuthorizationProviderOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.AssociateWebsiteAuthorizationProvider", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateWebsiteAuthorizationProviderAsync(ctx workflow.Context, input *worklink.AssociateWebsiteAuthorizationProviderInput) *WorkLinkAssociateWebsiteAuthorizationProviderFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.AssociateWebsiteAuthorizationProvider", input)
	return &WorkLinkAssociateWebsiteAuthorizationProviderFuture{Future: future}
}

func (a *stub) AssociateWebsiteCertificateAuthority(ctx workflow.Context, input *worklink.AssociateWebsiteCertificateAuthorityInput) (*worklink.AssociateWebsiteCertificateAuthorityOutput, error) {
	var output worklink.AssociateWebsiteCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.AssociateWebsiteCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateWebsiteCertificateAuthorityAsync(ctx workflow.Context, input *worklink.AssociateWebsiteCertificateAuthorityInput) *WorkLinkAssociateWebsiteCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.AssociateWebsiteCertificateAuthority", input)
	return &WorkLinkAssociateWebsiteCertificateAuthorityFuture{Future: future}
}

func (a *stub) CreateFleet(ctx workflow.Context, input *worklink.CreateFleetInput) (*worklink.CreateFleetOutput, error) {
	var output worklink.CreateFleetOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.CreateFleet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateFleetAsync(ctx workflow.Context, input *worklink.CreateFleetInput) *WorkLinkCreateFleetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.CreateFleet", input)
	return &WorkLinkCreateFleetFuture{Future: future}
}

func (a *stub) DeleteFleet(ctx workflow.Context, input *worklink.DeleteFleetInput) (*worklink.DeleteFleetOutput, error) {
	var output worklink.DeleteFleetOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.DeleteFleet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteFleetAsync(ctx workflow.Context, input *worklink.DeleteFleetInput) *WorkLinkDeleteFleetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.DeleteFleet", input)
	return &WorkLinkDeleteFleetFuture{Future: future}
}

func (a *stub) DescribeAuditStreamConfiguration(ctx workflow.Context, input *worklink.DescribeAuditStreamConfigurationInput) (*worklink.DescribeAuditStreamConfigurationOutput, error) {
	var output worklink.DescribeAuditStreamConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeAuditStreamConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAuditStreamConfigurationAsync(ctx workflow.Context, input *worklink.DescribeAuditStreamConfigurationInput) *WorkLinkDescribeAuditStreamConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeAuditStreamConfiguration", input)
	return &WorkLinkDescribeAuditStreamConfigurationFuture{Future: future}
}

func (a *stub) DescribeCompanyNetworkConfiguration(ctx workflow.Context, input *worklink.DescribeCompanyNetworkConfigurationInput) (*worklink.DescribeCompanyNetworkConfigurationOutput, error) {
	var output worklink.DescribeCompanyNetworkConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeCompanyNetworkConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeCompanyNetworkConfigurationAsync(ctx workflow.Context, input *worklink.DescribeCompanyNetworkConfigurationInput) *WorkLinkDescribeCompanyNetworkConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeCompanyNetworkConfiguration", input)
	return &WorkLinkDescribeCompanyNetworkConfigurationFuture{Future: future}
}

func (a *stub) DescribeDevice(ctx workflow.Context, input *worklink.DescribeDeviceInput) (*worklink.DescribeDeviceOutput, error) {
	var output worklink.DescribeDeviceOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeDevice", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDeviceAsync(ctx workflow.Context, input *worklink.DescribeDeviceInput) *WorkLinkDescribeDeviceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeDevice", input)
	return &WorkLinkDescribeDeviceFuture{Future: future}
}

func (a *stub) DescribeDevicePolicyConfiguration(ctx workflow.Context, input *worklink.DescribeDevicePolicyConfigurationInput) (*worklink.DescribeDevicePolicyConfigurationOutput, error) {
	var output worklink.DescribeDevicePolicyConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeDevicePolicyConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDevicePolicyConfigurationAsync(ctx workflow.Context, input *worklink.DescribeDevicePolicyConfigurationInput) *WorkLinkDescribeDevicePolicyConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeDevicePolicyConfiguration", input)
	return &WorkLinkDescribeDevicePolicyConfigurationFuture{Future: future}
}

func (a *stub) DescribeDomain(ctx workflow.Context, input *worklink.DescribeDomainInput) (*worklink.DescribeDomainOutput, error) {
	var output worklink.DescribeDomainOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeDomain", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDomainAsync(ctx workflow.Context, input *worklink.DescribeDomainInput) *WorkLinkDescribeDomainFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeDomain", input)
	return &WorkLinkDescribeDomainFuture{Future: future}
}

func (a *stub) DescribeFleetMetadata(ctx workflow.Context, input *worklink.DescribeFleetMetadataInput) (*worklink.DescribeFleetMetadataOutput, error) {
	var output worklink.DescribeFleetMetadataOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeFleetMetadata", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeFleetMetadataAsync(ctx workflow.Context, input *worklink.DescribeFleetMetadataInput) *WorkLinkDescribeFleetMetadataFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeFleetMetadata", input)
	return &WorkLinkDescribeFleetMetadataFuture{Future: future}
}

func (a *stub) DescribeIdentityProviderConfiguration(ctx workflow.Context, input *worklink.DescribeIdentityProviderConfigurationInput) (*worklink.DescribeIdentityProviderConfigurationOutput, error) {
	var output worklink.DescribeIdentityProviderConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeIdentityProviderConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeIdentityProviderConfigurationAsync(ctx workflow.Context, input *worklink.DescribeIdentityProviderConfigurationInput) *WorkLinkDescribeIdentityProviderConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeIdentityProviderConfiguration", input)
	return &WorkLinkDescribeIdentityProviderConfigurationFuture{Future: future}
}

func (a *stub) DescribeWebsiteCertificateAuthority(ctx workflow.Context, input *worklink.DescribeWebsiteCertificateAuthorityInput) (*worklink.DescribeWebsiteCertificateAuthorityOutput, error) {
	var output worklink.DescribeWebsiteCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeWebsiteCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeWebsiteCertificateAuthorityAsync(ctx workflow.Context, input *worklink.DescribeWebsiteCertificateAuthorityInput) *WorkLinkDescribeWebsiteCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.DescribeWebsiteCertificateAuthority", input)
	return &WorkLinkDescribeWebsiteCertificateAuthorityFuture{Future: future}
}

func (a *stub) DisassociateDomain(ctx workflow.Context, input *worklink.DisassociateDomainInput) (*worklink.DisassociateDomainOutput, error) {
	var output worklink.DisassociateDomainOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.DisassociateDomain", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateDomainAsync(ctx workflow.Context, input *worklink.DisassociateDomainInput) *WorkLinkDisassociateDomainFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.DisassociateDomain", input)
	return &WorkLinkDisassociateDomainFuture{Future: future}
}

func (a *stub) DisassociateWebsiteAuthorizationProvider(ctx workflow.Context, input *worklink.DisassociateWebsiteAuthorizationProviderInput) (*worklink.DisassociateWebsiteAuthorizationProviderOutput, error) {
	var output worklink.DisassociateWebsiteAuthorizationProviderOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.DisassociateWebsiteAuthorizationProvider", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateWebsiteAuthorizationProviderAsync(ctx workflow.Context, input *worklink.DisassociateWebsiteAuthorizationProviderInput) *WorkLinkDisassociateWebsiteAuthorizationProviderFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.DisassociateWebsiteAuthorizationProvider", input)
	return &WorkLinkDisassociateWebsiteAuthorizationProviderFuture{Future: future}
}

func (a *stub) DisassociateWebsiteCertificateAuthority(ctx workflow.Context, input *worklink.DisassociateWebsiteCertificateAuthorityInput) (*worklink.DisassociateWebsiteCertificateAuthorityOutput, error) {
	var output worklink.DisassociateWebsiteCertificateAuthorityOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.DisassociateWebsiteCertificateAuthority", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateWebsiteCertificateAuthorityAsync(ctx workflow.Context, input *worklink.DisassociateWebsiteCertificateAuthorityInput) *WorkLinkDisassociateWebsiteCertificateAuthorityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.DisassociateWebsiteCertificateAuthority", input)
	return &WorkLinkDisassociateWebsiteCertificateAuthorityFuture{Future: future}
}

func (a *stub) ListDevices(ctx workflow.Context, input *worklink.ListDevicesInput) (*worklink.ListDevicesOutput, error) {
	var output worklink.ListDevicesOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.ListDevices", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDevicesAsync(ctx workflow.Context, input *worklink.ListDevicesInput) *WorkLinkListDevicesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.ListDevices", input)
	return &WorkLinkListDevicesFuture{Future: future}
}

func (a *stub) ListDomains(ctx workflow.Context, input *worklink.ListDomainsInput) (*worklink.ListDomainsOutput, error) {
	var output worklink.ListDomainsOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.ListDomains", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDomainsAsync(ctx workflow.Context, input *worklink.ListDomainsInput) *WorkLinkListDomainsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.ListDomains", input)
	return &WorkLinkListDomainsFuture{Future: future}
}

func (a *stub) ListFleets(ctx workflow.Context, input *worklink.ListFleetsInput) (*worklink.ListFleetsOutput, error) {
	var output worklink.ListFleetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.ListFleets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListFleetsAsync(ctx workflow.Context, input *worklink.ListFleetsInput) *WorkLinkListFleetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.ListFleets", input)
	return &WorkLinkListFleetsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *worklink.ListTagsForResourceInput) (*worklink.ListTagsForResourceOutput, error) {
	var output worklink.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *worklink.ListTagsForResourceInput) *WorkLinkListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.ListTagsForResource", input)
	return &WorkLinkListTagsForResourceFuture{Future: future}
}

func (a *stub) ListWebsiteAuthorizationProviders(ctx workflow.Context, input *worklink.ListWebsiteAuthorizationProvidersInput) (*worklink.ListWebsiteAuthorizationProvidersOutput, error) {
	var output worklink.ListWebsiteAuthorizationProvidersOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.ListWebsiteAuthorizationProviders", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListWebsiteAuthorizationProvidersAsync(ctx workflow.Context, input *worklink.ListWebsiteAuthorizationProvidersInput) *WorkLinkListWebsiteAuthorizationProvidersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.ListWebsiteAuthorizationProviders", input)
	return &WorkLinkListWebsiteAuthorizationProvidersFuture{Future: future}
}

func (a *stub) ListWebsiteCertificateAuthorities(ctx workflow.Context, input *worklink.ListWebsiteCertificateAuthoritiesInput) (*worklink.ListWebsiteCertificateAuthoritiesOutput, error) {
	var output worklink.ListWebsiteCertificateAuthoritiesOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.ListWebsiteCertificateAuthorities", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListWebsiteCertificateAuthoritiesAsync(ctx workflow.Context, input *worklink.ListWebsiteCertificateAuthoritiesInput) *WorkLinkListWebsiteCertificateAuthoritiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.ListWebsiteCertificateAuthorities", input)
	return &WorkLinkListWebsiteCertificateAuthoritiesFuture{Future: future}
}

func (a *stub) RestoreDomainAccess(ctx workflow.Context, input *worklink.RestoreDomainAccessInput) (*worklink.RestoreDomainAccessOutput, error) {
	var output worklink.RestoreDomainAccessOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.RestoreDomainAccess", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RestoreDomainAccessAsync(ctx workflow.Context, input *worklink.RestoreDomainAccessInput) *WorkLinkRestoreDomainAccessFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.RestoreDomainAccess", input)
	return &WorkLinkRestoreDomainAccessFuture{Future: future}
}

func (a *stub) RevokeDomainAccess(ctx workflow.Context, input *worklink.RevokeDomainAccessInput) (*worklink.RevokeDomainAccessOutput, error) {
	var output worklink.RevokeDomainAccessOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.RevokeDomainAccess", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RevokeDomainAccessAsync(ctx workflow.Context, input *worklink.RevokeDomainAccessInput) *WorkLinkRevokeDomainAccessFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.RevokeDomainAccess", input)
	return &WorkLinkRevokeDomainAccessFuture{Future: future}
}

func (a *stub) SignOutUser(ctx workflow.Context, input *worklink.SignOutUserInput) (*worklink.SignOutUserOutput, error) {
	var output worklink.SignOutUserOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.SignOutUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SignOutUserAsync(ctx workflow.Context, input *worklink.SignOutUserInput) *WorkLinkSignOutUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.SignOutUser", input)
	return &WorkLinkSignOutUserFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *worklink.TagResourceInput) (*worklink.TagResourceOutput, error) {
	var output worklink.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *worklink.TagResourceInput) *WorkLinkTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.TagResource", input)
	return &WorkLinkTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *worklink.UntagResourceInput) (*worklink.UntagResourceOutput, error) {
	var output worklink.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *worklink.UntagResourceInput) *WorkLinkUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.UntagResource", input)
	return &WorkLinkUntagResourceFuture{Future: future}
}

func (a *stub) UpdateAuditStreamConfiguration(ctx workflow.Context, input *worklink.UpdateAuditStreamConfigurationInput) (*worklink.UpdateAuditStreamConfigurationOutput, error) {
	var output worklink.UpdateAuditStreamConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.UpdateAuditStreamConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateAuditStreamConfigurationAsync(ctx workflow.Context, input *worklink.UpdateAuditStreamConfigurationInput) *WorkLinkUpdateAuditStreamConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.UpdateAuditStreamConfiguration", input)
	return &WorkLinkUpdateAuditStreamConfigurationFuture{Future: future}
}

func (a *stub) UpdateCompanyNetworkConfiguration(ctx workflow.Context, input *worklink.UpdateCompanyNetworkConfigurationInput) (*worklink.UpdateCompanyNetworkConfigurationOutput, error) {
	var output worklink.UpdateCompanyNetworkConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.UpdateCompanyNetworkConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateCompanyNetworkConfigurationAsync(ctx workflow.Context, input *worklink.UpdateCompanyNetworkConfigurationInput) *WorkLinkUpdateCompanyNetworkConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.UpdateCompanyNetworkConfiguration", input)
	return &WorkLinkUpdateCompanyNetworkConfigurationFuture{Future: future}
}

func (a *stub) UpdateDevicePolicyConfiguration(ctx workflow.Context, input *worklink.UpdateDevicePolicyConfigurationInput) (*worklink.UpdateDevicePolicyConfigurationOutput, error) {
	var output worklink.UpdateDevicePolicyConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.UpdateDevicePolicyConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDevicePolicyConfigurationAsync(ctx workflow.Context, input *worklink.UpdateDevicePolicyConfigurationInput) *WorkLinkUpdateDevicePolicyConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.UpdateDevicePolicyConfiguration", input)
	return &WorkLinkUpdateDevicePolicyConfigurationFuture{Future: future}
}

func (a *stub) UpdateDomainMetadata(ctx workflow.Context, input *worklink.UpdateDomainMetadataInput) (*worklink.UpdateDomainMetadataOutput, error) {
	var output worklink.UpdateDomainMetadataOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.UpdateDomainMetadata", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDomainMetadataAsync(ctx workflow.Context, input *worklink.UpdateDomainMetadataInput) *WorkLinkUpdateDomainMetadataFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.UpdateDomainMetadata", input)
	return &WorkLinkUpdateDomainMetadataFuture{Future: future}
}

func (a *stub) UpdateFleetMetadata(ctx workflow.Context, input *worklink.UpdateFleetMetadataInput) (*worklink.UpdateFleetMetadataOutput, error) {
	var output worklink.UpdateFleetMetadataOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.UpdateFleetMetadata", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateFleetMetadataAsync(ctx workflow.Context, input *worklink.UpdateFleetMetadataInput) *WorkLinkUpdateFleetMetadataFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.UpdateFleetMetadata", input)
	return &WorkLinkUpdateFleetMetadataFuture{Future: future}
}

func (a *stub) UpdateIdentityProviderConfiguration(ctx workflow.Context, input *worklink.UpdateIdentityProviderConfigurationInput) (*worklink.UpdateIdentityProviderConfigurationOutput, error) {
	var output worklink.UpdateIdentityProviderConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.worklink.UpdateIdentityProviderConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateIdentityProviderConfigurationAsync(ctx workflow.Context, input *worklink.UpdateIdentityProviderConfigurationInput) *WorkLinkUpdateIdentityProviderConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.worklink.UpdateIdentityProviderConfiguration", input)
	return &WorkLinkUpdateIdentityProviderConfigurationFuture{Future: future}
}
