// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package connectstub

import (
	"github.com/aws/aws-sdk-go/service/connect"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateRoutingProfileQueues(ctx workflow.Context, input *connect.AssociateRoutingProfileQueuesInput) (*connect.AssociateRoutingProfileQueuesOutput, error)
	AssociateRoutingProfileQueuesAsync(ctx workflow.Context, input *connect.AssociateRoutingProfileQueuesInput) *ConnectAssociateRoutingProfileQueuesFuture

	CreateContactFlow(ctx workflow.Context, input *connect.CreateContactFlowInput) (*connect.CreateContactFlowOutput, error)
	CreateContactFlowAsync(ctx workflow.Context, input *connect.CreateContactFlowInput) *ConnectCreateContactFlowFuture

	CreateRoutingProfile(ctx workflow.Context, input *connect.CreateRoutingProfileInput) (*connect.CreateRoutingProfileOutput, error)
	CreateRoutingProfileAsync(ctx workflow.Context, input *connect.CreateRoutingProfileInput) *ConnectCreateRoutingProfileFuture

	CreateUser(ctx workflow.Context, input *connect.CreateUserInput) (*connect.CreateUserOutput, error)
	CreateUserAsync(ctx workflow.Context, input *connect.CreateUserInput) *ConnectCreateUserFuture

	DeleteUser(ctx workflow.Context, input *connect.DeleteUserInput) (*connect.DeleteUserOutput, error)
	DeleteUserAsync(ctx workflow.Context, input *connect.DeleteUserInput) *ConnectDeleteUserFuture

	DescribeContactFlow(ctx workflow.Context, input *connect.DescribeContactFlowInput) (*connect.DescribeContactFlowOutput, error)
	DescribeContactFlowAsync(ctx workflow.Context, input *connect.DescribeContactFlowInput) *ConnectDescribeContactFlowFuture

	DescribeRoutingProfile(ctx workflow.Context, input *connect.DescribeRoutingProfileInput) (*connect.DescribeRoutingProfileOutput, error)
	DescribeRoutingProfileAsync(ctx workflow.Context, input *connect.DescribeRoutingProfileInput) *ConnectDescribeRoutingProfileFuture

	DescribeUser(ctx workflow.Context, input *connect.DescribeUserInput) (*connect.DescribeUserOutput, error)
	DescribeUserAsync(ctx workflow.Context, input *connect.DescribeUserInput) *ConnectDescribeUserFuture

	DescribeUserHierarchyGroup(ctx workflow.Context, input *connect.DescribeUserHierarchyGroupInput) (*connect.DescribeUserHierarchyGroupOutput, error)
	DescribeUserHierarchyGroupAsync(ctx workflow.Context, input *connect.DescribeUserHierarchyGroupInput) *ConnectDescribeUserHierarchyGroupFuture

	DescribeUserHierarchyStructure(ctx workflow.Context, input *connect.DescribeUserHierarchyStructureInput) (*connect.DescribeUserHierarchyStructureOutput, error)
	DescribeUserHierarchyStructureAsync(ctx workflow.Context, input *connect.DescribeUserHierarchyStructureInput) *ConnectDescribeUserHierarchyStructureFuture

	DisassociateRoutingProfileQueues(ctx workflow.Context, input *connect.DisassociateRoutingProfileQueuesInput) (*connect.DisassociateRoutingProfileQueuesOutput, error)
	DisassociateRoutingProfileQueuesAsync(ctx workflow.Context, input *connect.DisassociateRoutingProfileQueuesInput) *ConnectDisassociateRoutingProfileQueuesFuture

	GetContactAttributes(ctx workflow.Context, input *connect.GetContactAttributesInput) (*connect.GetContactAttributesOutput, error)
	GetContactAttributesAsync(ctx workflow.Context, input *connect.GetContactAttributesInput) *ConnectGetContactAttributesFuture

	GetCurrentMetricData(ctx workflow.Context, input *connect.GetCurrentMetricDataInput) (*connect.GetCurrentMetricDataOutput, error)
	GetCurrentMetricDataAsync(ctx workflow.Context, input *connect.GetCurrentMetricDataInput) *ConnectGetCurrentMetricDataFuture

	GetFederationToken(ctx workflow.Context, input *connect.GetFederationTokenInput) (*connect.GetFederationTokenOutput, error)
	GetFederationTokenAsync(ctx workflow.Context, input *connect.GetFederationTokenInput) *ConnectGetFederationTokenFuture

	GetMetricData(ctx workflow.Context, input *connect.GetMetricDataInput) (*connect.GetMetricDataOutput, error)
	GetMetricDataAsync(ctx workflow.Context, input *connect.GetMetricDataInput) *ConnectGetMetricDataFuture

	ListContactFlows(ctx workflow.Context, input *connect.ListContactFlowsInput) (*connect.ListContactFlowsOutput, error)
	ListContactFlowsAsync(ctx workflow.Context, input *connect.ListContactFlowsInput) *ConnectListContactFlowsFuture

	ListHoursOfOperations(ctx workflow.Context, input *connect.ListHoursOfOperationsInput) (*connect.ListHoursOfOperationsOutput, error)
	ListHoursOfOperationsAsync(ctx workflow.Context, input *connect.ListHoursOfOperationsInput) *ConnectListHoursOfOperationsFuture

	ListPhoneNumbers(ctx workflow.Context, input *connect.ListPhoneNumbersInput) (*connect.ListPhoneNumbersOutput, error)
	ListPhoneNumbersAsync(ctx workflow.Context, input *connect.ListPhoneNumbersInput) *ConnectListPhoneNumbersFuture

	ListPrompts(ctx workflow.Context, input *connect.ListPromptsInput) (*connect.ListPromptsOutput, error)
	ListPromptsAsync(ctx workflow.Context, input *connect.ListPromptsInput) *ConnectListPromptsFuture

	ListQueues(ctx workflow.Context, input *connect.ListQueuesInput) (*connect.ListQueuesOutput, error)
	ListQueuesAsync(ctx workflow.Context, input *connect.ListQueuesInput) *ConnectListQueuesFuture

	ListRoutingProfileQueues(ctx workflow.Context, input *connect.ListRoutingProfileQueuesInput) (*connect.ListRoutingProfileQueuesOutput, error)
	ListRoutingProfileQueuesAsync(ctx workflow.Context, input *connect.ListRoutingProfileQueuesInput) *ConnectListRoutingProfileQueuesFuture

	ListRoutingProfiles(ctx workflow.Context, input *connect.ListRoutingProfilesInput) (*connect.ListRoutingProfilesOutput, error)
	ListRoutingProfilesAsync(ctx workflow.Context, input *connect.ListRoutingProfilesInput) *ConnectListRoutingProfilesFuture

	ListSecurityProfiles(ctx workflow.Context, input *connect.ListSecurityProfilesInput) (*connect.ListSecurityProfilesOutput, error)
	ListSecurityProfilesAsync(ctx workflow.Context, input *connect.ListSecurityProfilesInput) *ConnectListSecurityProfilesFuture

	ListTagsForResource(ctx workflow.Context, input *connect.ListTagsForResourceInput) (*connect.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *connect.ListTagsForResourceInput) *ConnectListTagsForResourceFuture

	ListUserHierarchyGroups(ctx workflow.Context, input *connect.ListUserHierarchyGroupsInput) (*connect.ListUserHierarchyGroupsOutput, error)
	ListUserHierarchyGroupsAsync(ctx workflow.Context, input *connect.ListUserHierarchyGroupsInput) *ConnectListUserHierarchyGroupsFuture

	ListUsers(ctx workflow.Context, input *connect.ListUsersInput) (*connect.ListUsersOutput, error)
	ListUsersAsync(ctx workflow.Context, input *connect.ListUsersInput) *ConnectListUsersFuture

	ResumeContactRecording(ctx workflow.Context, input *connect.ResumeContactRecordingInput) (*connect.ResumeContactRecordingOutput, error)
	ResumeContactRecordingAsync(ctx workflow.Context, input *connect.ResumeContactRecordingInput) *ConnectResumeContactRecordingFuture

	StartChatContact(ctx workflow.Context, input *connect.StartChatContactInput) (*connect.StartChatContactOutput, error)
	StartChatContactAsync(ctx workflow.Context, input *connect.StartChatContactInput) *ConnectStartChatContactFuture

	StartContactRecording(ctx workflow.Context, input *connect.StartContactRecordingInput) (*connect.StartContactRecordingOutput, error)
	StartContactRecordingAsync(ctx workflow.Context, input *connect.StartContactRecordingInput) *ConnectStartContactRecordingFuture

	StartOutboundVoiceContact(ctx workflow.Context, input *connect.StartOutboundVoiceContactInput) (*connect.StartOutboundVoiceContactOutput, error)
	StartOutboundVoiceContactAsync(ctx workflow.Context, input *connect.StartOutboundVoiceContactInput) *ConnectStartOutboundVoiceContactFuture

	StopContact(ctx workflow.Context, input *connect.StopContactInput) (*connect.StopContactOutput, error)
	StopContactAsync(ctx workflow.Context, input *connect.StopContactInput) *ConnectStopContactFuture

	StopContactRecording(ctx workflow.Context, input *connect.StopContactRecordingInput) (*connect.StopContactRecordingOutput, error)
	StopContactRecordingAsync(ctx workflow.Context, input *connect.StopContactRecordingInput) *ConnectStopContactRecordingFuture

	SuspendContactRecording(ctx workflow.Context, input *connect.SuspendContactRecordingInput) (*connect.SuspendContactRecordingOutput, error)
	SuspendContactRecordingAsync(ctx workflow.Context, input *connect.SuspendContactRecordingInput) *ConnectSuspendContactRecordingFuture

	TagResource(ctx workflow.Context, input *connect.TagResourceInput) (*connect.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *connect.TagResourceInput) *ConnectTagResourceFuture

	UntagResource(ctx workflow.Context, input *connect.UntagResourceInput) (*connect.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *connect.UntagResourceInput) *ConnectUntagResourceFuture

	UpdateContactAttributes(ctx workflow.Context, input *connect.UpdateContactAttributesInput) (*connect.UpdateContactAttributesOutput, error)
	UpdateContactAttributesAsync(ctx workflow.Context, input *connect.UpdateContactAttributesInput) *ConnectUpdateContactAttributesFuture

	UpdateContactFlowContent(ctx workflow.Context, input *connect.UpdateContactFlowContentInput) (*connect.UpdateContactFlowContentOutput, error)
	UpdateContactFlowContentAsync(ctx workflow.Context, input *connect.UpdateContactFlowContentInput) *ConnectUpdateContactFlowContentFuture

	UpdateContactFlowName(ctx workflow.Context, input *connect.UpdateContactFlowNameInput) (*connect.UpdateContactFlowNameOutput, error)
	UpdateContactFlowNameAsync(ctx workflow.Context, input *connect.UpdateContactFlowNameInput) *ConnectUpdateContactFlowNameFuture

	UpdateRoutingProfileConcurrency(ctx workflow.Context, input *connect.UpdateRoutingProfileConcurrencyInput) (*connect.UpdateRoutingProfileConcurrencyOutput, error)
	UpdateRoutingProfileConcurrencyAsync(ctx workflow.Context, input *connect.UpdateRoutingProfileConcurrencyInput) *ConnectUpdateRoutingProfileConcurrencyFuture

	UpdateRoutingProfileDefaultOutboundQueue(ctx workflow.Context, input *connect.UpdateRoutingProfileDefaultOutboundQueueInput) (*connect.UpdateRoutingProfileDefaultOutboundQueueOutput, error)
	UpdateRoutingProfileDefaultOutboundQueueAsync(ctx workflow.Context, input *connect.UpdateRoutingProfileDefaultOutboundQueueInput) *ConnectUpdateRoutingProfileDefaultOutboundQueueFuture

	UpdateRoutingProfileName(ctx workflow.Context, input *connect.UpdateRoutingProfileNameInput) (*connect.UpdateRoutingProfileNameOutput, error)
	UpdateRoutingProfileNameAsync(ctx workflow.Context, input *connect.UpdateRoutingProfileNameInput) *ConnectUpdateRoutingProfileNameFuture

	UpdateRoutingProfileQueues(ctx workflow.Context, input *connect.UpdateRoutingProfileQueuesInput) (*connect.UpdateRoutingProfileQueuesOutput, error)
	UpdateRoutingProfileQueuesAsync(ctx workflow.Context, input *connect.UpdateRoutingProfileQueuesInput) *ConnectUpdateRoutingProfileQueuesFuture

	UpdateUserHierarchy(ctx workflow.Context, input *connect.UpdateUserHierarchyInput) (*connect.UpdateUserHierarchyOutput, error)
	UpdateUserHierarchyAsync(ctx workflow.Context, input *connect.UpdateUserHierarchyInput) *ConnectUpdateUserHierarchyFuture

	UpdateUserIdentityInfo(ctx workflow.Context, input *connect.UpdateUserIdentityInfoInput) (*connect.UpdateUserIdentityInfoOutput, error)
	UpdateUserIdentityInfoAsync(ctx workflow.Context, input *connect.UpdateUserIdentityInfoInput) *ConnectUpdateUserIdentityInfoFuture

	UpdateUserPhoneConfig(ctx workflow.Context, input *connect.UpdateUserPhoneConfigInput) (*connect.UpdateUserPhoneConfigOutput, error)
	UpdateUserPhoneConfigAsync(ctx workflow.Context, input *connect.UpdateUserPhoneConfigInput) *ConnectUpdateUserPhoneConfigFuture

	UpdateUserRoutingProfile(ctx workflow.Context, input *connect.UpdateUserRoutingProfileInput) (*connect.UpdateUserRoutingProfileOutput, error)
	UpdateUserRoutingProfileAsync(ctx workflow.Context, input *connect.UpdateUserRoutingProfileInput) *ConnectUpdateUserRoutingProfileFuture

	UpdateUserSecurityProfiles(ctx workflow.Context, input *connect.UpdateUserSecurityProfilesInput) (*connect.UpdateUserSecurityProfilesOutput, error)
	UpdateUserSecurityProfilesAsync(ctx workflow.Context, input *connect.UpdateUserSecurityProfilesInput) *ConnectUpdateUserSecurityProfilesFuture
}

func NewClient() Client {
	return &stub{}
}
