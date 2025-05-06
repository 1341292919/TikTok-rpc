package rpc

import (
	"TikTok-rpc/kitex_gen/model"
	web "TikTok-rpc/kitex_gen/websocket"
	"TikTok-rpc/pkg/base/client"
	"TikTok-rpc/pkg/errno"
	"TikTok-rpc/pkg/utils"
	"context"

	"github.com/bytedance/gopkg/util/logger"
)

func InitWebsocketRPC() {
	c, err := client.InitWebsocketRPC()
	if err != nil {
		logger.Fatalf("api.rpc.user InitWebsocketRPC failed, err is %v", err)
	}
	websocketClient = *c // UserClinet = c?
}

func AddMessageRPC(ctx context.Context, req *web.AddMessageRequest) error {
	resp, err := websocketClient.AddMessage(ctx, req)
	if err != nil {
		logger.Errorf("AddMessageRPC(: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) { // 将其标注为服务错误，那么如果是数据库错误呢
		return errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return nil
}
func ReadOfflineMessageRPC(ctx context.Context, req *web.QueryOfflineMessageRequest) (*model.ChatMessageList, error) {
	resp, err := websocketClient.QueryOfflineMessage(ctx, req)
	if err != nil {
		logger.Errorf("QueryOfflineMessageRPC(: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp.Data, nil
}
func ReadPrivateHistoryMessageRPC(ctx context.Context, req *web.QueryPrivateHistoryMessageRequest) (*model.ChatMessageList, error) {
	resp, err := websocketClient.QueryPrivateHistoryMessage(ctx, req)
	if err != nil {
		logger.Errorf("QueryPrivateHistoryMessageRPC(: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp.Data, nil
}
func ReadGroupHistoryMessageRPC(ctx context.Context, req *web.QueryGroupHistoryMessageRequest) (*model.ChatMessageList, error) {
	resp, err := websocketClient.QueryGroupHistoryMessage(ctx, req)
	if err != nil {
		logger.Errorf("QueryGroupHistoryMessageRPC(: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp.Data, nil
}
