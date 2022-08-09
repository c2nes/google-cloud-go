// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package accessapproval

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	accessapprovalpb "google.golang.org/genproto/googleapis/cloud/accessapproval/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ListApprovalRequests            []gax.CallOption
	GetApprovalRequest              []gax.CallOption
	ApproveApprovalRequest          []gax.CallOption
	DismissApprovalRequest          []gax.CallOption
	InvalidateApprovalRequest       []gax.CallOption
	GetAccessApprovalSettings       []gax.CallOption
	UpdateAccessApprovalSettings    []gax.CallOption
	DeleteAccessApprovalSettings    []gax.CallOption
	GetAccessApprovalServiceAccount []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("accessapproval.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("accessapproval.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://accessapproval.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		ListApprovalRequests: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetApprovalRequest: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ApproveApprovalRequest:    []gax.CallOption{},
		DismissApprovalRequest:    []gax.CallOption{},
		InvalidateApprovalRequest: []gax.CallOption{},
		GetAccessApprovalSettings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateAccessApprovalSettings:    []gax.CallOption{},
		DeleteAccessApprovalSettings:    []gax.CallOption{},
		GetAccessApprovalServiceAccount: []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods available from Access Approval API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListApprovalRequests(context.Context, *accessapprovalpb.ListApprovalRequestsMessage, ...gax.CallOption) *ApprovalRequestIterator
	GetApprovalRequest(context.Context, *accessapprovalpb.GetApprovalRequestMessage, ...gax.CallOption) (*accessapprovalpb.ApprovalRequest, error)
	ApproveApprovalRequest(context.Context, *accessapprovalpb.ApproveApprovalRequestMessage, ...gax.CallOption) (*accessapprovalpb.ApprovalRequest, error)
	DismissApprovalRequest(context.Context, *accessapprovalpb.DismissApprovalRequestMessage, ...gax.CallOption) (*accessapprovalpb.ApprovalRequest, error)
	InvalidateApprovalRequest(context.Context, *accessapprovalpb.InvalidateApprovalRequestMessage, ...gax.CallOption) (*accessapprovalpb.ApprovalRequest, error)
	GetAccessApprovalSettings(context.Context, *accessapprovalpb.GetAccessApprovalSettingsMessage, ...gax.CallOption) (*accessapprovalpb.AccessApprovalSettings, error)
	UpdateAccessApprovalSettings(context.Context, *accessapprovalpb.UpdateAccessApprovalSettingsMessage, ...gax.CallOption) (*accessapprovalpb.AccessApprovalSettings, error)
	DeleteAccessApprovalSettings(context.Context, *accessapprovalpb.DeleteAccessApprovalSettingsMessage, ...gax.CallOption) error
	GetAccessApprovalServiceAccount(context.Context, *accessapprovalpb.GetAccessApprovalServiceAccountMessage, ...gax.CallOption) (*accessapprovalpb.AccessApprovalServiceAccount, error)
}

// Client is a client for interacting with Access Approval API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// This API allows a customer to manage accesses to cloud resources by
// Google personnel. It defines the following resource model:
//
//	The API has a collection of
//	ApprovalRequest
//	resources, named approvalRequests/{approval_request}
//
//	The API has top-level settings per Project/Folder/Organization, named
//	accessApprovalSettings
//
// The service also periodically emails a list of recipients, defined at the
// Project/Folder/Organization level in the accessApprovalSettings, when there
// is a pending ApprovalRequest for them to act on. The ApprovalRequests can
// also optionally be published to a Pub/Sub topic owned by the customer
// (contact support if you would like to enable Pub/Sub notifications).
//
// ApprovalRequests can be approved or dismissed. Google personnel can only
// access the indicated resource or resources if the request is approved
// (subject to some exclusions:
// https://cloud.google.com/access-approval/docs/overview#exclusions (at https://cloud.google.com/access-approval/docs/overview#exclusions)).
//
// Note: Using Access Approval functionality will mean that Google may not be
// able to meet the SLAs for your chosen products, as any support response times
// may be dramatically increased. As such the SLAs do not apply to any service
// disruption to the extent impacted by Customer’s use of Access Approval. Do
// not enable Access Approval for projects where you may require high service
// availability and rapid response by Google Cloud Support.
//
// After a request is approved or dismissed, no further action may be taken on
// it. Requests with the requested_expiration in the past or with no activity
// for 14 days are considered dismissed. When an approval expires, the request
// is considered dismissed.
//
// If a request is not approved or dismissed, we call it pending.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListApprovalRequests lists approval requests associated with a project, folder, or organization.
// Approval requests can be filtered by state (pending, active, dismissed).
// The order is reverse chronological.
func (c *Client) ListApprovalRequests(ctx context.Context, req *accessapprovalpb.ListApprovalRequestsMessage, opts ...gax.CallOption) *ApprovalRequestIterator {
	return c.internalClient.ListApprovalRequests(ctx, req, opts...)
}

// GetApprovalRequest gets an approval request. Returns NOT_FOUND if the request does not exist.
func (c *Client) GetApprovalRequest(ctx context.Context, req *accessapprovalpb.GetApprovalRequestMessage, opts ...gax.CallOption) (*accessapprovalpb.ApprovalRequest, error) {
	return c.internalClient.GetApprovalRequest(ctx, req, opts...)
}

// ApproveApprovalRequest approves a request and returns the updated ApprovalRequest.
//
// Returns NOT_FOUND if the request does not exist. Returns
// FAILED_PRECONDITION if the request exists but is not in a pending state.
func (c *Client) ApproveApprovalRequest(ctx context.Context, req *accessapprovalpb.ApproveApprovalRequestMessage, opts ...gax.CallOption) (*accessapprovalpb.ApprovalRequest, error) {
	return c.internalClient.ApproveApprovalRequest(ctx, req, opts...)
}

// DismissApprovalRequest dismisses a request. Returns the updated ApprovalRequest.
//
// NOTE: This does not deny access to the resource if another request has been
// made and approved. It is equivalent in effect to ignoring the request
// altogether.
//
// Returns NOT_FOUND if the request does not exist.
//
// Returns FAILED_PRECONDITION if the request exists but is not in a pending
// state.
func (c *Client) DismissApprovalRequest(ctx context.Context, req *accessapprovalpb.DismissApprovalRequestMessage, opts ...gax.CallOption) (*accessapprovalpb.ApprovalRequest, error) {
	return c.internalClient.DismissApprovalRequest(ctx, req, opts...)
}

// InvalidateApprovalRequest invalidates an existing ApprovalRequest. Returns the updated
// ApprovalRequest.
//
// NOTE: This does not deny access to the resource if another request has been
// made and approved. It only invalidates a single approval.
//
// Returns FAILED_PRECONDITION if the request exists but is not in an approved
// state.
func (c *Client) InvalidateApprovalRequest(ctx context.Context, req *accessapprovalpb.InvalidateApprovalRequestMessage, opts ...gax.CallOption) (*accessapprovalpb.ApprovalRequest, error) {
	return c.internalClient.InvalidateApprovalRequest(ctx, req, opts...)
}

// GetAccessApprovalSettings gets the settings associated with a project, folder, or organization.
func (c *Client) GetAccessApprovalSettings(ctx context.Context, req *accessapprovalpb.GetAccessApprovalSettingsMessage, opts ...gax.CallOption) (*accessapprovalpb.AccessApprovalSettings, error) {
	return c.internalClient.GetAccessApprovalSettings(ctx, req, opts...)
}

// UpdateAccessApprovalSettings updates the settings associated with a project, folder, or organization.
// Settings to update are determined by the value of field_mask.
func (c *Client) UpdateAccessApprovalSettings(ctx context.Context, req *accessapprovalpb.UpdateAccessApprovalSettingsMessage, opts ...gax.CallOption) (*accessapprovalpb.AccessApprovalSettings, error) {
	return c.internalClient.UpdateAccessApprovalSettings(ctx, req, opts...)
}

// DeleteAccessApprovalSettings deletes the settings associated with a project, folder, or organization.
// This will have the effect of disabling Access Approval for the project,
// folder, or organization, but only if all ancestors also have Access
// Approval disabled. If Access Approval is enabled at a higher level of the
// hierarchy, then Access Approval will still be enabled at this level as
// the settings are inherited.
func (c *Client) DeleteAccessApprovalSettings(ctx context.Context, req *accessapprovalpb.DeleteAccessApprovalSettingsMessage, opts ...gax.CallOption) error {
	return c.internalClient.DeleteAccessApprovalSettings(ctx, req, opts...)
}

// GetAccessApprovalServiceAccount retrieves the service account that is used by Access Approval to access KMS
// keys for signing approved approval requests.
func (c *Client) GetAccessApprovalServiceAccount(ctx context.Context, req *accessapprovalpb.GetAccessApprovalServiceAccountMessage, opts ...gax.CallOption) (*accessapprovalpb.AccessApprovalServiceAccount, error) {
	return c.internalClient.GetAccessApprovalServiceAccount(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Access Approval API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client accessapprovalpb.AccessApprovalClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new access approval client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// This API allows a customer to manage accesses to cloud resources by
// Google personnel. It defines the following resource model:
//
//	The API has a collection of
//	ApprovalRequest
//	resources, named approvalRequests/{approval_request}
//
//	The API has top-level settings per Project/Folder/Organization, named
//	accessApprovalSettings
//
// The service also periodically emails a list of recipients, defined at the
// Project/Folder/Organization level in the accessApprovalSettings, when there
// is a pending ApprovalRequest for them to act on. The ApprovalRequests can
// also optionally be published to a Pub/Sub topic owned by the customer
// (contact support if you would like to enable Pub/Sub notifications).
//
// ApprovalRequests can be approved or dismissed. Google personnel can only
// access the indicated resource or resources if the request is approved
// (subject to some exclusions:
// https://cloud.google.com/access-approval/docs/overview#exclusions (at https://cloud.google.com/access-approval/docs/overview#exclusions)).
//
// Note: Using Access Approval functionality will mean that Google may not be
// able to meet the SLAs for your chosen products, as any support response times
// may be dramatically increased. As such the SLAs do not apply to any service
// disruption to the extent impacted by Customer’s use of Access Approval. Do
// not enable Access Approval for projects where you may require high service
// availability and rapid response by Google Cloud Support.
//
// After a request is approved or dismissed, no further action may be taken on
// it. Requests with the requested_expiration in the past or with no activity
// for 14 days are considered dismissed. When an approval expires, the request
// is considered dismissed.
//
// If a request is not approved or dismissed, we call it pending.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		client:           accessapprovalpb.NewAccessApprovalClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) ListApprovalRequests(ctx context.Context, req *accessapprovalpb.ListApprovalRequestsMessage, opts ...gax.CallOption) *ApprovalRequestIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListApprovalRequests[0:len((*c.CallOptions).ListApprovalRequests):len((*c.CallOptions).ListApprovalRequests)], opts...)
	it := &ApprovalRequestIterator{}
	req = proto.Clone(req).(*accessapprovalpb.ListApprovalRequestsMessage)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*accessapprovalpb.ApprovalRequest, string, error) {
		resp := &accessapprovalpb.ListApprovalRequestsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListApprovalRequests(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetApprovalRequests(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) GetApprovalRequest(ctx context.Context, req *accessapprovalpb.GetApprovalRequestMessage, opts ...gax.CallOption) (*accessapprovalpb.ApprovalRequest, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetApprovalRequest[0:len((*c.CallOptions).GetApprovalRequest):len((*c.CallOptions).GetApprovalRequest)], opts...)
	var resp *accessapprovalpb.ApprovalRequest
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetApprovalRequest(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ApproveApprovalRequest(ctx context.Context, req *accessapprovalpb.ApproveApprovalRequestMessage, opts ...gax.CallOption) (*accessapprovalpb.ApprovalRequest, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ApproveApprovalRequest[0:len((*c.CallOptions).ApproveApprovalRequest):len((*c.CallOptions).ApproveApprovalRequest)], opts...)
	var resp *accessapprovalpb.ApprovalRequest
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ApproveApprovalRequest(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DismissApprovalRequest(ctx context.Context, req *accessapprovalpb.DismissApprovalRequestMessage, opts ...gax.CallOption) (*accessapprovalpb.ApprovalRequest, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DismissApprovalRequest[0:len((*c.CallOptions).DismissApprovalRequest):len((*c.CallOptions).DismissApprovalRequest)], opts...)
	var resp *accessapprovalpb.ApprovalRequest
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.DismissApprovalRequest(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) InvalidateApprovalRequest(ctx context.Context, req *accessapprovalpb.InvalidateApprovalRequestMessage, opts ...gax.CallOption) (*accessapprovalpb.ApprovalRequest, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).InvalidateApprovalRequest[0:len((*c.CallOptions).InvalidateApprovalRequest):len((*c.CallOptions).InvalidateApprovalRequest)], opts...)
	var resp *accessapprovalpb.ApprovalRequest
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.InvalidateApprovalRequest(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetAccessApprovalSettings(ctx context.Context, req *accessapprovalpb.GetAccessApprovalSettingsMessage, opts ...gax.CallOption) (*accessapprovalpb.AccessApprovalSettings, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAccessApprovalSettings[0:len((*c.CallOptions).GetAccessApprovalSettings):len((*c.CallOptions).GetAccessApprovalSettings)], opts...)
	var resp *accessapprovalpb.AccessApprovalSettings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetAccessApprovalSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) UpdateAccessApprovalSettings(ctx context.Context, req *accessapprovalpb.UpdateAccessApprovalSettingsMessage, opts ...gax.CallOption) (*accessapprovalpb.AccessApprovalSettings, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "settings.name", url.QueryEscape(req.GetSettings().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateAccessApprovalSettings[0:len((*c.CallOptions).UpdateAccessApprovalSettings):len((*c.CallOptions).UpdateAccessApprovalSettings)], opts...)
	var resp *accessapprovalpb.AccessApprovalSettings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateAccessApprovalSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeleteAccessApprovalSettings(ctx context.Context, req *accessapprovalpb.DeleteAccessApprovalSettingsMessage, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteAccessApprovalSettings[0:len((*c.CallOptions).DeleteAccessApprovalSettings):len((*c.CallOptions).DeleteAccessApprovalSettings)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteAccessApprovalSettings(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) GetAccessApprovalServiceAccount(ctx context.Context, req *accessapprovalpb.GetAccessApprovalServiceAccountMessage, opts ...gax.CallOption) (*accessapprovalpb.AccessApprovalServiceAccount, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAccessApprovalServiceAccount[0:len((*c.CallOptions).GetAccessApprovalServiceAccount):len((*c.CallOptions).GetAccessApprovalServiceAccount)], opts...)
	var resp *accessapprovalpb.AccessApprovalServiceAccount
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetAccessApprovalServiceAccount(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ApprovalRequestIterator manages a stream of *accessapprovalpb.ApprovalRequest.
type ApprovalRequestIterator struct {
	items    []*accessapprovalpb.ApprovalRequest
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*accessapprovalpb.ApprovalRequest, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ApprovalRequestIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ApprovalRequestIterator) Next() (*accessapprovalpb.ApprovalRequest, error) {
	var item *accessapprovalpb.ApprovalRequest
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ApprovalRequestIterator) bufLen() int {
	return len(it.items)
}

func (it *ApprovalRequestIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
