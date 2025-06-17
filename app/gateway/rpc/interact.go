package rpc

import (
	api "TikTok-rpc/app/gateway/model/api/interact"
	"TikTok-rpc/app/gateway/pack"
	"TikTok-rpc/kitex_gen/interact"
	"TikTok-rpc/pkg/base/client"
	"TikTok-rpc/pkg/errno"
	"TikTok-rpc/pkg/utils"
	"context"

	"github.com/bytedance/gopkg/util/logger"
)

func InitInteractRPC() {
	c, err := client.InitInteractRPC()
	if err != nil {
		logger.Fatalf("api.rpc.interact InitInteractRPC failed, err is %v", err)
	}
	interactClient = *c
}

func LikeRPC(ctx context.Context, req *interact.LikeRequest) error {
	resp, err := interactClient.Like(ctx, req)
	if err != nil {
		logger.Errorf("LikeRPC:RPC called failed :%v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return nil
}
func QueryLikeListPRC(ctx context.Context, req *interact.QueryLikeListRequest) (*api.QueryLikeListResponse, error) {
	apiResp := new(api.QueryLikeListResponse)
	resp, err := interactClient.QueryLikeList(ctx, req)
	if err != nil {
		logger.Errorf("QueryLikeListPRC:RPC called failed :%v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	apiResp.Data = pack.VideoList(resp.Data)
	return apiResp, nil
}
func CommentRPC(ctx context.Context, req *interact.CommentRequest) error {
	resp, err := interactClient.Comment(ctx, req)
	if err != nil {
		logger.Errorf("CommentRPC:RPC called failed :%v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return nil
}
func DeleteCommentPRC(ctx context.Context, req *interact.DeleteCommentRequest) error {
	resp, err := interactClient.DeleteComment(ctx, req)
	if err != nil {
		logger.Errorf("DeleteCommentPRC:RPC called failed :%v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return nil
}
func QueryCommentListPRC(ctx context.Context, req *interact.QueryCommentListRequest) (*api.QueryCommentListResponse, error) {
	apiResp := new(api.QueryCommentListResponse)
	resp, err := interactClient.QueryCommentList(ctx, req)
	if err != nil {
		logger.Errorf("QueryCommentListPRC:RPC called failed :%v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	apiResp.Data = pack.CommentList(resp.Data)
	return apiResp, nil
}
