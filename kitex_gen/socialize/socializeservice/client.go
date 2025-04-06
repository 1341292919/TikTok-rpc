// Code generated by Kitex v0.12.3. DO NOT EDIT.

package socializeservice

import (
	socialize "TikTok-rpc/kitex_gen/socialize"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Follow(ctx context.Context, req *socialize.FollowRequest, callOptions ...callopt.Option) (r *socialize.FollowResponse, err error)
	QueryFollowList(ctx context.Context, req *socialize.QueryFollowListRequest, callOptions ...callopt.Option) (r *socialize.QueryFollowListResponse, err error)
	QueryFollowerList(ctx context.Context, req *socialize.QueryFollowerListRequest, callOptions ...callopt.Option) (r *socialize.QueryFollowerListResponse, err error)
	QueryFriendList(ctx context.Context, req *socialize.QueryFriendListRequest, callOptions ...callopt.Option) (r *socialize.QueryFriendListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kSocializeServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kSocializeServiceClient struct {
	*kClient
}

func (p *kSocializeServiceClient) Follow(ctx context.Context, req *socialize.FollowRequest, callOptions ...callopt.Option) (r *socialize.FollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Follow(ctx, req)
}

func (p *kSocializeServiceClient) QueryFollowList(ctx context.Context, req *socialize.QueryFollowListRequest, callOptions ...callopt.Option) (r *socialize.QueryFollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryFollowList(ctx, req)
}

func (p *kSocializeServiceClient) QueryFollowerList(ctx context.Context, req *socialize.QueryFollowerListRequest, callOptions ...callopt.Option) (r *socialize.QueryFollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryFollowerList(ctx, req)
}

func (p *kSocializeServiceClient) QueryFriendList(ctx context.Context, req *socialize.QueryFriendListRequest, callOptions ...callopt.Option) (r *socialize.QueryFriendListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryFriendList(ctx, req)
}
