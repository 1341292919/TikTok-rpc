package rpc

import (
	api "TikTok-rpc/app/gateway/model/api/video"
	"TikTok-rpc/app/gateway/pack"
	"TikTok-rpc/kitex_gen/video"
	"TikTok-rpc/pkg/base/client"
	"TikTok-rpc/pkg/errno"
	"TikTok-rpc/pkg/utils"
	"context"

	"github.com/bytedance/gopkg/util/logger"
)

func InitVideoRPC() {
	c, err := client.InitVideoRPC()
	if err != nil {
		logger.Fatalf("api.rpc.video InitVideoRPC failed, err is %v", err)
	}
	videoClient = *c
}
func PublishVideoRPC(ctx context.Context, req *video.PublishRequest) error {
	resp, err := videoClient.PublishVideo(ctx, req)
	if err != nil {
		logger.Errorf("PublishRPC: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return nil
}
func QueryPublishListRPC(ctx context.Context, req *video.QueryPublishListRequest) (*api.QueryPublishListResponse, error) {
	apiResp := &api.QueryPublishListResponse{}
	resp, err := videoClient.QueryList(ctx, req)
	if err != nil {
		logger.Errorf("QueryPublishRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	apiResp.Data = pack.VideoList(resp.Data)
	return apiResp, nil
}
func GetPopularListRPC(ctx context.Context, req *video.GetPopularListRequest) (*api.GetPopularListResponse, error) {
	apiResp := &api.GetPopularListResponse{}
	resp, err := videoClient.GetPopularVideo(ctx, req)
	if err != nil {
		logger.Errorf("PopularRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	apiResp.Data = pack.VideoList(resp.Data)
	return apiResp, nil
}

func SearchVideoRPC(ctx context.Context, req *video.SearchVideoByKeywordRequest) (*api.SearchVideoByKeywordResponse, error) {
	apiResp := &api.SearchVideoByKeywordResponse{}
	resp, err := videoClient.SearchVideo(ctx, req)
	if err != nil {
		logger.Errorf("SearchVideo: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	apiResp.Data = pack.VideoList(resp.Data)
	return apiResp, nil
}
func VideoStreamRPC(ctx context.Context, req *video.VideoStreamRequest) (*api.VideoStreamResponse, error) {
	apiResp := &api.VideoStreamResponse{}
	resp, err := videoClient.GetVideoStream(ctx, req)
	if err != nil {
		logger.Errorf("VideoStream: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	apiResp.Data = pack.VideoList(resp.Data)
	return apiResp, nil
}
