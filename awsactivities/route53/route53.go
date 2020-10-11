// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package route53

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
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
	client route53iface.Route53API

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := route53.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (route53iface.Route53API, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return route53.New(sess), nil
}

func (a *Activities) AssociateVPCWithHostedZone(ctx context.Context, input *route53.AssociateVPCWithHostedZoneInput) (*route53.AssociateVPCWithHostedZoneOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateVPCWithHostedZoneWithContext(ctx, input)
}

func (a *Activities) ChangeResourceRecordSets(ctx context.Context, input *route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ChangeResourceRecordSetsWithContext(ctx, input)
}

func (a *Activities) ChangeTagsForResource(ctx context.Context, input *route53.ChangeTagsForResourceInput) (*route53.ChangeTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ChangeTagsForResourceWithContext(ctx, input)
}

func (a *Activities) CreateHealthCheck(ctx context.Context, input *route53.CreateHealthCheckInput) (*route53.CreateHealthCheckOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateHealthCheckWithContext(ctx, input)
}

func (a *Activities) CreateHostedZone(ctx context.Context, input *route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateHostedZoneWithContext(ctx, input)
}

func (a *Activities) CreateQueryLoggingConfig(ctx context.Context, input *route53.CreateQueryLoggingConfigInput) (*route53.CreateQueryLoggingConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateQueryLoggingConfigWithContext(ctx, input)
}

func (a *Activities) CreateReusableDelegationSet(ctx context.Context, input *route53.CreateReusableDelegationSetInput) (*route53.CreateReusableDelegationSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateReusableDelegationSetWithContext(ctx, input)
}

func (a *Activities) CreateTrafficPolicy(ctx context.Context, input *route53.CreateTrafficPolicyInput) (*route53.CreateTrafficPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateTrafficPolicyWithContext(ctx, input)
}

func (a *Activities) CreateTrafficPolicyInstance(ctx context.Context, input *route53.CreateTrafficPolicyInstanceInput) (*route53.CreateTrafficPolicyInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateTrafficPolicyInstanceWithContext(ctx, input)
}

func (a *Activities) CreateTrafficPolicyVersion(ctx context.Context, input *route53.CreateTrafficPolicyVersionInput) (*route53.CreateTrafficPolicyVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateTrafficPolicyVersionWithContext(ctx, input)
}

func (a *Activities) CreateVPCAssociationAuthorization(ctx context.Context, input *route53.CreateVPCAssociationAuthorizationInput) (*route53.CreateVPCAssociationAuthorizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateVPCAssociationAuthorizationWithContext(ctx, input)
}

func (a *Activities) DeleteHealthCheck(ctx context.Context, input *route53.DeleteHealthCheckInput) (*route53.DeleteHealthCheckOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteHealthCheckWithContext(ctx, input)
}

func (a *Activities) DeleteHostedZone(ctx context.Context, input *route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteHostedZoneWithContext(ctx, input)
}

func (a *Activities) DeleteQueryLoggingConfig(ctx context.Context, input *route53.DeleteQueryLoggingConfigInput) (*route53.DeleteQueryLoggingConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteQueryLoggingConfigWithContext(ctx, input)
}

func (a *Activities) DeleteReusableDelegationSet(ctx context.Context, input *route53.DeleteReusableDelegationSetInput) (*route53.DeleteReusableDelegationSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteReusableDelegationSetWithContext(ctx, input)
}

func (a *Activities) DeleteTrafficPolicy(ctx context.Context, input *route53.DeleteTrafficPolicyInput) (*route53.DeleteTrafficPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteTrafficPolicyWithContext(ctx, input)
}

func (a *Activities) DeleteTrafficPolicyInstance(ctx context.Context, input *route53.DeleteTrafficPolicyInstanceInput) (*route53.DeleteTrafficPolicyInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteTrafficPolicyInstanceWithContext(ctx, input)
}

func (a *Activities) DeleteVPCAssociationAuthorization(ctx context.Context, input *route53.DeleteVPCAssociationAuthorizationInput) (*route53.DeleteVPCAssociationAuthorizationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteVPCAssociationAuthorizationWithContext(ctx, input)
}

func (a *Activities) DisassociateVPCFromHostedZone(ctx context.Context, input *route53.DisassociateVPCFromHostedZoneInput) (*route53.DisassociateVPCFromHostedZoneOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateVPCFromHostedZoneWithContext(ctx, input)
}

func (a *Activities) GetAccountLimit(ctx context.Context, input *route53.GetAccountLimitInput) (*route53.GetAccountLimitOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAccountLimitWithContext(ctx, input)
}

func (a *Activities) GetChange(ctx context.Context, input *route53.GetChangeInput) (*route53.GetChangeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetChangeWithContext(ctx, input)
}

func (a *Activities) GetCheckerIpRanges(ctx context.Context, input *route53.GetCheckerIpRangesInput) (*route53.GetCheckerIpRangesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCheckerIpRangesWithContext(ctx, input)
}

func (a *Activities) GetGeoLocation(ctx context.Context, input *route53.GetGeoLocationInput) (*route53.GetGeoLocationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetGeoLocationWithContext(ctx, input)
}

func (a *Activities) GetHealthCheck(ctx context.Context, input *route53.GetHealthCheckInput) (*route53.GetHealthCheckOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetHealthCheckWithContext(ctx, input)
}

func (a *Activities) GetHealthCheckCount(ctx context.Context, input *route53.GetHealthCheckCountInput) (*route53.GetHealthCheckCountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetHealthCheckCountWithContext(ctx, input)
}

func (a *Activities) GetHealthCheckLastFailureReason(ctx context.Context, input *route53.GetHealthCheckLastFailureReasonInput) (*route53.GetHealthCheckLastFailureReasonOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetHealthCheckLastFailureReasonWithContext(ctx, input)
}

func (a *Activities) GetHealthCheckStatus(ctx context.Context, input *route53.GetHealthCheckStatusInput) (*route53.GetHealthCheckStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetHealthCheckStatusWithContext(ctx, input)
}

func (a *Activities) GetHostedZone(ctx context.Context, input *route53.GetHostedZoneInput) (*route53.GetHostedZoneOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetHostedZoneWithContext(ctx, input)
}

func (a *Activities) GetHostedZoneCount(ctx context.Context, input *route53.GetHostedZoneCountInput) (*route53.GetHostedZoneCountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetHostedZoneCountWithContext(ctx, input)
}

func (a *Activities) GetHostedZoneLimit(ctx context.Context, input *route53.GetHostedZoneLimitInput) (*route53.GetHostedZoneLimitOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetHostedZoneLimitWithContext(ctx, input)
}

func (a *Activities) GetQueryLoggingConfig(ctx context.Context, input *route53.GetQueryLoggingConfigInput) (*route53.GetQueryLoggingConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetQueryLoggingConfigWithContext(ctx, input)
}

func (a *Activities) GetReusableDelegationSet(ctx context.Context, input *route53.GetReusableDelegationSetInput) (*route53.GetReusableDelegationSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetReusableDelegationSetWithContext(ctx, input)
}

func (a *Activities) GetReusableDelegationSetLimit(ctx context.Context, input *route53.GetReusableDelegationSetLimitInput) (*route53.GetReusableDelegationSetLimitOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetReusableDelegationSetLimitWithContext(ctx, input)
}

func (a *Activities) GetTrafficPolicy(ctx context.Context, input *route53.GetTrafficPolicyInput) (*route53.GetTrafficPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTrafficPolicyWithContext(ctx, input)
}

func (a *Activities) GetTrafficPolicyInstance(ctx context.Context, input *route53.GetTrafficPolicyInstanceInput) (*route53.GetTrafficPolicyInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTrafficPolicyInstanceWithContext(ctx, input)
}

func (a *Activities) GetTrafficPolicyInstanceCount(ctx context.Context, input *route53.GetTrafficPolicyInstanceCountInput) (*route53.GetTrafficPolicyInstanceCountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTrafficPolicyInstanceCountWithContext(ctx, input)
}

func (a *Activities) ListGeoLocations(ctx context.Context, input *route53.ListGeoLocationsInput) (*route53.ListGeoLocationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListGeoLocationsWithContext(ctx, input)
}

func (a *Activities) ListHealthChecks(ctx context.Context, input *route53.ListHealthChecksInput) (*route53.ListHealthChecksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListHealthChecksWithContext(ctx, input)
}

func (a *Activities) ListHostedZones(ctx context.Context, input *route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListHostedZonesWithContext(ctx, input)
}

func (a *Activities) ListHostedZonesByName(ctx context.Context, input *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListHostedZonesByNameWithContext(ctx, input)
}

func (a *Activities) ListHostedZonesByVPC(ctx context.Context, input *route53.ListHostedZonesByVPCInput) (*route53.ListHostedZonesByVPCOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListHostedZonesByVPCWithContext(ctx, input)
}

func (a *Activities) ListQueryLoggingConfigs(ctx context.Context, input *route53.ListQueryLoggingConfigsInput) (*route53.ListQueryLoggingConfigsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListQueryLoggingConfigsWithContext(ctx, input)
}

func (a *Activities) ListResourceRecordSets(ctx context.Context, input *route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResourceRecordSetsWithContext(ctx, input)
}

func (a *Activities) ListReusableDelegationSets(ctx context.Context, input *route53.ListReusableDelegationSetsInput) (*route53.ListReusableDelegationSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListReusableDelegationSetsWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) ListTagsForResources(ctx context.Context, input *route53.ListTagsForResourcesInput) (*route53.ListTagsForResourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourcesWithContext(ctx, input)
}

func (a *Activities) ListTrafficPolicies(ctx context.Context, input *route53.ListTrafficPoliciesInput) (*route53.ListTrafficPoliciesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTrafficPoliciesWithContext(ctx, input)
}

func (a *Activities) ListTrafficPolicyInstances(ctx context.Context, input *route53.ListTrafficPolicyInstancesInput) (*route53.ListTrafficPolicyInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTrafficPolicyInstancesWithContext(ctx, input)
}

func (a *Activities) ListTrafficPolicyInstancesByHostedZone(ctx context.Context, input *route53.ListTrafficPolicyInstancesByHostedZoneInput) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTrafficPolicyInstancesByHostedZoneWithContext(ctx, input)
}

func (a *Activities) ListTrafficPolicyInstancesByPolicy(ctx context.Context, input *route53.ListTrafficPolicyInstancesByPolicyInput) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTrafficPolicyInstancesByPolicyWithContext(ctx, input)
}

func (a *Activities) ListTrafficPolicyVersions(ctx context.Context, input *route53.ListTrafficPolicyVersionsInput) (*route53.ListTrafficPolicyVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTrafficPolicyVersionsWithContext(ctx, input)
}

func (a *Activities) ListVPCAssociationAuthorizations(ctx context.Context, input *route53.ListVPCAssociationAuthorizationsInput) (*route53.ListVPCAssociationAuthorizationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListVPCAssociationAuthorizationsWithContext(ctx, input)
}

func (a *Activities) TestDNSAnswer(ctx context.Context, input *route53.TestDNSAnswerInput) (*route53.TestDNSAnswerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TestDNSAnswerWithContext(ctx, input)
}

func (a *Activities) UpdateHealthCheck(ctx context.Context, input *route53.UpdateHealthCheckInput) (*route53.UpdateHealthCheckOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateHealthCheckWithContext(ctx, input)
}

func (a *Activities) UpdateHostedZoneComment(ctx context.Context, input *route53.UpdateHostedZoneCommentInput) (*route53.UpdateHostedZoneCommentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateHostedZoneCommentWithContext(ctx, input)
}

func (a *Activities) UpdateTrafficPolicyComment(ctx context.Context, input *route53.UpdateTrafficPolicyCommentInput) (*route53.UpdateTrafficPolicyCommentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateTrafficPolicyCommentWithContext(ctx, input)
}

func (a *Activities) UpdateTrafficPolicyInstance(ctx context.Context, input *route53.UpdateTrafficPolicyInstanceInput) (*route53.UpdateTrafficPolicyInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateTrafficPolicyInstanceWithContext(ctx, input)
}

func (a *Activities) WaitUntilResourceRecordSetsChanged(ctx context.Context, input *route53.GetChangeInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilResourceRecordSetsChangedWithContext(ctx, input, options...)
	})
}
