package rpc

import (
	api "TikTok-rpc/app/gateway/model/api/user"
	apiModel "TikTok-rpc/app/gateway/model/model"
	"TikTok-rpc/app/gateway/pack"
	"TikTok-rpc/kitex_gen/user"
	"TikTok-rpc/pkg/base/client"
	"TikTok-rpc/pkg/errno"
	"TikTok-rpc/pkg/utils"
	"context"
	"github.com/bytedance/gopkg/util/logger"
)

func InitUserRPC() {
	c, err := client.InitUserRPC()
	if err != nil {
		logger.Fatalf("api.rpc.user InitUserRPC failed, err is %v", err)
	}
	userClient = *c // UserClinet = c?
}

// 注册
// 传入的是rpc的请求,应该返回hz的Resp
// 由于只要求返回Base
func RegisterRPC(ctx context.Context, req *user.RegisterRequest) error {
	resp, err := userClient.Register(ctx, req)
	if err != nil {
		logger.Errorf("LoginRPC: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) { // 将其标注为服务错误，那么如果是数据库错误呢
		return errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return nil
}

// 登录
func LoginRPC(ctx context.Context, req *user.LoginRequest) (*api.LoginResponse, error) {
	apiResp := new(api.LoginResponse)
	resp, err := userClient.Login(ctx, req)
	if err != nil {
		logger.Errorf("LoginRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) { // 将其标注为服务错误，那么如果是数据库错误呢
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	// 对数据进行封装
	apiResp.Data = pack.User(resp.Data)
	return apiResp, nil
}

// 以图搜图
func SearchImageRPC(ctx context.Context, req *user.SearchImagesRequest) (*api.SearchImagesResponse, error) {
	apiResp := new(api.SearchImagesResponse)
	resp, err := userClient.SearchImage(ctx, req)
	if err != nil {
		logger.Errorf("SearchImageRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return apiResp, nil
}

// 上传头像
func UploadAvatarRPC(ctx context.Context, req *user.UploadAvatarRequest) (*api.UploadAvatarResponse, error) {
	apiResp := new(api.UploadAvatarResponse)
	resp, err := userClient.UploadAvatar(ctx, req)
	if err != nil {
		logger.Errorf("UploadAvatarRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	apiResp.Data = pack.User(resp.Data)
	return apiResp, nil
}

// 获取用户信息
func GetUserMessagesRPC(ctx context.Context, req *user.GetUserInformationRequest) (*api.GetUserInformationResponse, error) {
	apiResp := new(api.GetUserInformationResponse)
	resp, err := userClient.GetInformation(ctx, req)
	if err != nil {
		logger.Errorf("GetUserMessagesRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	apiResp.Data = pack.User(resp.Data)
	return apiResp, nil
}

// 获取qrcode
func GetQrcodeRPC(ctx context.Context, req *user.GetMFARequest) (*api.GetMFAResponse, error) {
	apiResp := new(api.GetMFAResponse)
	resp, err := userClient.GetMFA(ctx, req)
	if err != nil {
		logger.Errorf("GetQrcodeRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	apiResp = &api.GetMFAResponse{
		Data: &apiModel.MFAMessage{
			Qrcode: resp.Data.Qrcode,
			Secret: resp.Data.Secret,
		},
	}
	return apiResp, nil
}

// 绑定多因素身份认证(MFA)
func MFABindRPC(ctx context.Context, req *user.MFABindRequest) (*api.MFABindResponse, error) {
	apiResp := new(api.MFABindResponse)
	resp, err := userClient.MindBind(ctx, req)
	if err != nil {
		logger.Errorf("MFABindRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return apiResp, nil
}
