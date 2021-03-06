// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package iotsecuretunnelingstub

import (
	"github.com/aws/aws-sdk-go/service/iotsecuretunneling"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CloseTunnel(ctx workflow.Context, input *iotsecuretunneling.CloseTunnelInput) (*iotsecuretunneling.CloseTunnelOutput, error)
	CloseTunnelAsync(ctx workflow.Context, input *iotsecuretunneling.CloseTunnelInput) *IoTSecureTunnelingCloseTunnelFuture

	DescribeTunnel(ctx workflow.Context, input *iotsecuretunneling.DescribeTunnelInput) (*iotsecuretunneling.DescribeTunnelOutput, error)
	DescribeTunnelAsync(ctx workflow.Context, input *iotsecuretunneling.DescribeTunnelInput) *IoTSecureTunnelingDescribeTunnelFuture

	ListTagsForResource(ctx workflow.Context, input *iotsecuretunneling.ListTagsForResourceInput) (*iotsecuretunneling.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *iotsecuretunneling.ListTagsForResourceInput) *IoTSecureTunnelingListTagsForResourceFuture

	ListTunnels(ctx workflow.Context, input *iotsecuretunneling.ListTunnelsInput) (*iotsecuretunneling.ListTunnelsOutput, error)
	ListTunnelsAsync(ctx workflow.Context, input *iotsecuretunneling.ListTunnelsInput) *IoTSecureTunnelingListTunnelsFuture

	OpenTunnel(ctx workflow.Context, input *iotsecuretunneling.OpenTunnelInput) (*iotsecuretunneling.OpenTunnelOutput, error)
	OpenTunnelAsync(ctx workflow.Context, input *iotsecuretunneling.OpenTunnelInput) *IoTSecureTunnelingOpenTunnelFuture

	TagResource(ctx workflow.Context, input *iotsecuretunneling.TagResourceInput) (*iotsecuretunneling.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *iotsecuretunneling.TagResourceInput) *IoTSecureTunnelingTagResourceFuture

	UntagResource(ctx workflow.Context, input *iotsecuretunneling.UntagResourceInput) (*iotsecuretunneling.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *iotsecuretunneling.UntagResourceInput) *IoTSecureTunnelingUntagResourceFuture
}

func NewClient() Client {
	return &stub{}
}
