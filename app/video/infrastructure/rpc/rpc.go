package rpc

import (
	"TikTok-rpc/app/video/domain/repository"
	"TikTok-rpc/kitex_gen/user"
	"TikTok-rpc/kitex_gen/user/userservice"
	"TikTok-rpc/pkg/errno"
	"TikTok-rpc/pkg/utils"
	"context"

	"github.com/bytedance/gopkg/util/logger"
)

type VideoRpcImpl struct {
	user userservice.Client
}

func NewVideoRpcImpl(user userservice.Client) repository.VideoRpc {
	return &VideoRpcImpl{user: user}
}
func (rpc *VideoRpcImpl) QueryUserIdByUsername(ctx context.Context, username string) (int64, error) {
	req := &user.QueryUserIdByUsernameRequest{
		Username: username,
	}
	resp, err := rpc.user.QueryUserIdByUsername(ctx, req)
	if err != nil {
		logger.Errorf(" QueryUserIdByUsernameRPC: RPC called failed: %v", err.Error())
		return -1, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return -1, errno.NewErrNo(errno.InternalRPCErrorCode, "interact-video rpc failed:"+resp.Base.Msg)
	}
	return *resp.Id, nil
}
