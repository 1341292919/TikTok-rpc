// Code generated by Kitex v0.12.3. DO NOT EDIT.

package userservice

import (
	user "TikTok-rpc/rpc/user/kitex_gen/user"
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Register": kitex.NewMethodInfo(
		registerHandler,
		newUserServiceRegisterArgs,
		newUserServiceRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newUserServiceLoginArgs,
		newUserServiceLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UploadAvatar": kitex.NewMethodInfo(
		uploadAvatarHandler,
		newUserServiceUploadAvatarArgs,
		newUserServiceUploadAvatarResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetInformation": kitex.NewMethodInfo(
		getInformationHandler,
		newUserServiceGetInformationArgs,
		newUserServiceGetInformationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"SearchImage": kitex.NewMethodInfo(
		searchImageHandler,
		newUserServiceSearchImageArgs,
		newUserServiceSearchImageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetMFA": kitex.NewMethodInfo(
		getMFAHandler,
		newUserServiceGetMFAArgs,
		newUserServiceGetMFAResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"MindBind": kitex.NewMethodInfo(
		mindBindHandler,
		newUserServiceMindBindArgs,
		newUserServiceMindBindResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
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
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
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
		"PackageName": "user",
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

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRegisterArgs)
	realResult := result.(*user.UserServiceRegisterResult)
	success, err := handler.(user.UserService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRegisterArgs() interface{} {
	return user.NewUserServiceRegisterArgs()
}

func newUserServiceRegisterResult() interface{} {
	return user.NewUserServiceRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceLoginArgs)
	realResult := result.(*user.UserServiceLoginResult)
	success, err := handler.(user.UserService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginArgs() interface{} {
	return user.NewUserServiceLoginArgs()
}

func newUserServiceLoginResult() interface{} {
	return user.NewUserServiceLoginResult()
}

func uploadAvatarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUploadAvatarArgs)
	realResult := result.(*user.UserServiceUploadAvatarResult)
	success, err := handler.(user.UserService).UploadAvatar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUploadAvatarArgs() interface{} {
	return user.NewUserServiceUploadAvatarArgs()
}

func newUserServiceUploadAvatarResult() interface{} {
	return user.NewUserServiceUploadAvatarResult()
}

func getInformationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetInformationArgs)
	realResult := result.(*user.UserServiceGetInformationResult)
	success, err := handler.(user.UserService).GetInformation(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetInformationArgs() interface{} {
	return user.NewUserServiceGetInformationArgs()
}

func newUserServiceGetInformationResult() interface{} {
	return user.NewUserServiceGetInformationResult()
}

func searchImageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceSearchImageArgs)
	realResult := result.(*user.UserServiceSearchImageResult)
	success, err := handler.(user.UserService).SearchImage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceSearchImageArgs() interface{} {
	return user.NewUserServiceSearchImageArgs()
}

func newUserServiceSearchImageResult() interface{} {
	return user.NewUserServiceSearchImageResult()
}

func getMFAHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetMFAArgs)
	realResult := result.(*user.UserServiceGetMFAResult)
	success, err := handler.(user.UserService).GetMFA(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetMFAArgs() interface{} {
	return user.NewUserServiceGetMFAArgs()
}

func newUserServiceGetMFAResult() interface{} {
	return user.NewUserServiceGetMFAResult()
}

func mindBindHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceMindBindArgs)
	realResult := result.(*user.UserServiceMindBindResult)
	success, err := handler.(user.UserService).MindBind(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceMindBindArgs() interface{} {
	return user.NewUserServiceMindBindArgs()
}

func newUserServiceMindBindResult() interface{} {
	return user.NewUserServiceMindBindResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, req *user.RegisterRequest) (r *user.RegisterResponse, err error) {
	var _args user.UserServiceRegisterArgs
	_args.Req = req
	var _result user.UserServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *user.LoginRequest) (r *user.LoginResponse, err error) {
	var _args user.UserServiceLoginArgs
	_args.Req = req
	var _result user.UserServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UploadAvatar(ctx context.Context, req *user.UploadAvatarRequest) (r *user.UploadAvatarResponse, err error) {
	var _args user.UserServiceUploadAvatarArgs
	_args.Req = req
	var _result user.UserServiceUploadAvatarResult
	if err = p.c.Call(ctx, "UploadAvatar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetInformation(ctx context.Context, req *user.GetUserInformationRequest) (r *user.GetUserInformationResponse, err error) {
	var _args user.UserServiceGetInformationArgs
	_args.Req = req
	var _result user.UserServiceGetInformationResult
	if err = p.c.Call(ctx, "GetInformation", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SearchImage(ctx context.Context, req *user.SearchImagesRequest) (r *user.SearchImagesResponse, err error) {
	var _args user.UserServiceSearchImageArgs
	_args.Req = req
	var _result user.UserServiceSearchImageResult
	if err = p.c.Call(ctx, "SearchImage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetMFA(ctx context.Context, req *user.GetMFARequest) (r *user.GetMFAResponse, err error) {
	var _args user.UserServiceGetMFAArgs
	_args.Req = req
	var _result user.UserServiceGetMFAResult
	if err = p.c.Call(ctx, "GetMFA", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MindBind(ctx context.Context, req *user.MFABindRequest) (r *user.MFABindResponse, err error) {
	var _args user.UserServiceMindBindArgs
	_args.Req = req
	var _result user.UserServiceMindBindResult
	if err = p.c.Call(ctx, "MindBind", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
