// Code generated by Kitex v0.12.3. DO NOT EDIT.

package socializeservice

import (
	socialize "TikTok-rpc/rpc/socialize/kitex_gen/socialize"
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Follow": kitex.NewMethodInfo(
		followHandler,
		newSocializeServiceFollowArgs,
		newSocializeServiceFollowResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"QueryFollowList": kitex.NewMethodInfo(
		queryFollowListHandler,
		newSocializeServiceQueryFollowListArgs,
		newSocializeServiceQueryFollowListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"QueryFollowerList": kitex.NewMethodInfo(
		queryFollowerListHandler,
		newSocializeServiceQueryFollowerListArgs,
		newSocializeServiceQueryFollowerListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"QueryFriendList": kitex.NewMethodInfo(
		queryFriendListHandler,
		newSocializeServiceQueryFriendListArgs,
		newSocializeServiceQueryFriendListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	socializeServiceServiceInfo                = NewServiceInfo()
	socializeServiceServiceInfoForClient       = NewServiceInfoForClient()
	socializeServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return socializeServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return socializeServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return socializeServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "SocializeService"
	handlerType := (*socialize.SocializeService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "socialize",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.12.3",
		Extra:           extra,
	}
	return svcInfo
}

func followHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*socialize.SocializeServiceFollowArgs)
	realResult := result.(*socialize.SocializeServiceFollowResult)
	success, err := handler.(socialize.SocializeService).Follow(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocializeServiceFollowArgs() interface{} {
	return socialize.NewSocializeServiceFollowArgs()
}

func newSocializeServiceFollowResult() interface{} {
	return socialize.NewSocializeServiceFollowResult()
}

func queryFollowListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*socialize.SocializeServiceQueryFollowListArgs)
	realResult := result.(*socialize.SocializeServiceQueryFollowListResult)
	success, err := handler.(socialize.SocializeService).QueryFollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocializeServiceQueryFollowListArgs() interface{} {
	return socialize.NewSocializeServiceQueryFollowListArgs()
}

func newSocializeServiceQueryFollowListResult() interface{} {
	return socialize.NewSocializeServiceQueryFollowListResult()
}

func queryFollowerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*socialize.SocializeServiceQueryFollowerListArgs)
	realResult := result.(*socialize.SocializeServiceQueryFollowerListResult)
	success, err := handler.(socialize.SocializeService).QueryFollowerList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocializeServiceQueryFollowerListArgs() interface{} {
	return socialize.NewSocializeServiceQueryFollowerListArgs()
}

func newSocializeServiceQueryFollowerListResult() interface{} {
	return socialize.NewSocializeServiceQueryFollowerListResult()
}

func queryFriendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*socialize.SocializeServiceQueryFriendListArgs)
	realResult := result.(*socialize.SocializeServiceQueryFriendListResult)
	success, err := handler.(socialize.SocializeService).QueryFriendList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocializeServiceQueryFriendListArgs() interface{} {
	return socialize.NewSocializeServiceQueryFriendListArgs()
}

func newSocializeServiceQueryFriendListResult() interface{} {
	return socialize.NewSocializeServiceQueryFriendListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Follow(ctx context.Context, req *socialize.FollowRequest) (r *socialize.FollowResponse, err error) {
	var _args socialize.SocializeServiceFollowArgs
	_args.Req = req
	var _result socialize.SocializeServiceFollowResult
	if err = p.c.Call(ctx, "Follow", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryFollowList(ctx context.Context, req *socialize.QueryFollowListRequest) (r *socialize.QueryFollowListResponse, err error) {
	var _args socialize.SocializeServiceQueryFollowListArgs
	_args.Req = req
	var _result socialize.SocializeServiceQueryFollowListResult
	if err = p.c.Call(ctx, "QueryFollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryFollowerList(ctx context.Context, req *socialize.QueryFollowerListRequest) (r *socialize.QueryFollowerListResponse, err error) {
	var _args socialize.SocializeServiceQueryFollowerListArgs
	_args.Req = req
	var _result socialize.SocializeServiceQueryFollowerListResult
	if err = p.c.Call(ctx, "QueryFollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryFriendList(ctx context.Context, req *socialize.QueryFriendListRequest) (r *socialize.QueryFriendListResponse, err error) {
	var _args socialize.SocializeServiceQueryFriendListArgs
	_args.Req = req
	var _result socialize.SocializeServiceQueryFriendListResult
	if err = p.c.Call(ctx, "QueryFriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
