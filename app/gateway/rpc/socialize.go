package rpc

import (
	api "TikTok-rpc/app/gateway/model/api/socialize"
	"TikTok-rpc/app/gateway/pack"
	"TikTok-rpc/kitex_gen/socialize"
	"TikTok-rpc/pkg/base/client"
	"TikTok-rpc/pkg/errno"
	"TikTok-rpc/pkg/utils"
	"context"

	"github.com/bytedance/gopkg/util/logger"
)

func InitSocializeRPC() {
	c, err := client.InitSocializeRPC()
	if err != nil {
		logger.Fatalf("api.rpc.socialize InitSocializeRPC failed, err is %v", err)
	}
	socializeClient = *c
}

func FollowRPC(ctx context.Context, req *socialize.FollowRequest) error {
	resp, err := socializeClient.Follow(ctx, req)
	if err != nil {
		logger.Errorf("FollowRPC: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) { // 将其标注为服务错误，那么如果是数据库错误呢
		return errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return nil
}

func QueryFollowList(ctx context.Context, req *socialize.QueryFollowListRequest) (*api.QueryFollowListResponse, error) {
	apiResp := new(api.QueryFollowListResponse)
	resp, err := socializeClient.QueryFollowList(ctx, req)
	if err != nil {
		logger.Errorf("QueryFollowList: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) { // 将其标注为服务错误，那么如果是数据库错误呢
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	apiResp.Data = pack.SimpleUserList(resp.Data)
	return apiResp, nil
}
func QueryFansList(ctx context.Context, req *socialize.QueryFansListRequest) (*api.QueryFansListResponse, error) {
	apiResp := new(api.QueryFansListResponse)
	resp, err := socializeClient.QueryFansList(ctx, req)
	if err != nil {
		logger.Errorf("QueryFansList: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	apiResp.Data = pack.SimpleUserList(resp.Data)
	return apiResp, nil
}
func QueryFriendList(ctx context.Context, req *socialize.QueryFriendListRequest) (*api.QueryFriendListResponse, error) {
	apiResp := new(api.QueryFriendListResponse)
	resp, err := socializeClient.QueryFriendList(ctx, req)
	if err != nil {
		logger.Errorf("QueryFriendList: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	apiResp.Data = pack.SimpleUserList(resp.Data)
	return apiResp, nil
}
