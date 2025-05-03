package rpc

import (
	rpcModel "TikTok-rpc/app/interact/domain/model"
	"TikTok-rpc/app/interact/domain/repository"
	"TikTok-rpc/kitex_gen/model"
	"TikTok-rpc/kitex_gen/user"
	"TikTok-rpc/kitex_gen/user/userservice"
	"TikTok-rpc/kitex_gen/video"
	"TikTok-rpc/kitex_gen/video/videoservice"
	"TikTok-rpc/pkg/errno"
	"TikTok-rpc/pkg/utils"
	"context"
	"github.com/bytedance/gopkg/util/logger"
)

type InteractRpcImpl struct {
	video videoservice.Client
	user  userservice.Client
}

func NewInteractRpcImpl(v videoservice.Client, user userservice.Client) repository.RpcPort {
	return &InteractRpcImpl{
		video: v,
		user:  user,
	}
}
func (rpc *InteractRpcImpl) IsVideoExist(ctx context.Context, videoID int64) (bool, error) {
	checkReq := &video.QueryVideoByVIdRequest{
		VideoId: videoID,
	}

	resp, err := rpc.video.QueryVideoById(ctx, checkReq)
	if err != nil {
		logger.Errorf("IsVideoExistRPC: RPC called failed: %v", err.Error())
		return false, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		if resp.Base.Code == errno.ServiceVideoNotExist {
			return false, nil
		}
		return false, errno.NewErrNo(errno.InternalRPCErrorCode, "interact-video rpc failed:"+resp.Base.Msg)
	}
	return true, nil
}

func (rpc *InteractRpcImpl) IsUserExist(ctx context.Context, userId int64) (bool, error) {
	checkReq := &user.GetUserInformationRequest{
		UserId: userId,
	}
	resp, err := rpc.user.GetInformation(ctx, checkReq)
	if err != nil {
		logger.Errorf(" IsUserExistRPC: RPC called failed: %v", err.Error())
		return false, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		if resp.Base.Code == errno.ServiceUserNotExistCode {
			return false, nil
		}
		return false, errno.NewErrNo(errno.InternalRPCErrorCode, "interact-user rpc failed:"+resp.Base.Msg)
	}
	return true, nil
}

// 该函数用于更新视频评论数目
func (rpc *InteractRpcImpl) UpdateVideoCommentCount(ctx context.Context, videoID, count int64) error {
	req := &video.UpdateVideoCommentCountRequest{
		VideoId:     videoID,
		ChangeCount: count,
	}
	resp, err := rpc.video.UpdateCommentCount(ctx, req)
	if err != nil {
		logger.Errorf("UpdateVideoCommentCountRPC: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return errno.NewErrNo(errno.InternalRPCErrorCode, "interact-video rpc failed:"+resp.Base.Msg)
	}
	return nil
}
func (rpc *InteractRpcImpl) UpdateVideoLikeCount(ctx context.Context, videoID, count int64) error {
	req := &video.UpdateVideoLikeCountRequest{
		VideoId:   videoID,
		LikeCount: count,
	}
	resp, err := rpc.video.UpdateLikeCount(ctx, req)
	if err != nil {
		logger.Errorf(" UpdateVideoLikeCountRPC: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return errno.NewErrNo(errno.InternalRPCErrorCode, "interact-video rpc failed:"+resp.Base.Msg)
	}
	return nil
}

func (rpc *InteractRpcImpl) QueryVideoLikeCount(ctx context.Context) ([]*rpcModel.LikeCount, error) {
	req := &video.QueryLikeCountRequest{}
	resp, err := rpc.video.QueryLikeCount(ctx, req)
	if err != nil {
		logger.Errorf(" QueryVideoLikeCountRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !utils.IsRPCSuccess(resp.Base) {
		return nil, errno.NewErrNo(errno.InternalRPCErrorCode, "interact-video rpc failed:"+resp.Base.Msg)
	}
	return buildLikeCountList(resp.Data), nil
}
func (rpc *InteractRpcImpl) AddCount(ctx context.Context, videoID, t int64) error {
	return nil
}

func (rpc *InteractRpcImpl) QueryVideoList(ctx context.Context, videoID []int64) ([]*rpcModel.Video, error) {
	QueryReq := new(video.QueryVideoByVIdRequest)
	videoData := make([]*rpcModel.Video, 0)
	for _, viD := range videoID {
		QueryReq.VideoId = viD
		resp, err := rpc.video.QueryVideoById(ctx, QueryReq)
		if err != nil {
			logger.Errorf(" QueryVideoListRPC: RPC called failed: %v", err.Error())
			return nil, errno.InternalServiceError.WithError(err)
		}
		if !utils.IsRPCSuccess(resp.Base) {
			return nil, errno.NewErrNo(errno.InternalRPCErrorCode, "interact-video rpc failed:"+resp.Base.Msg)
		}
		v := &rpcModel.Video{
			Uid:          resp.Data.UserId,
			Id:           resp.Data.Id,
			Title:        resp.Data.Title,
			Description:  resp.Data.Description,
			VisitCount:   resp.Data.VisitCount,
			LikeCount:    resp.Data.LikeCount,
			VideoUrl:     resp.Data.VideoUrl,
			CoverUrl:     resp.Data.CoverUrl,
			CommentCount: resp.Data.CommentCount,
			CreateAT:     resp.Data.CreatedAt,
			UpdateAT:     resp.Data.UpdatedAt,
		}
		videoData = append(videoData, v)
	}
	return videoData, nil
}

func buildLikeCount(count *model.LikeCount) *rpcModel.LikeCount {
	return &rpcModel.LikeCount{
		Id:    count.VideoId,
		Count: count.Count,
		Type:  0,
	}
}
func buildLikeCountList(data *model.LikeCountList) []*rpcModel.LikeCount {
	likes := make([]*rpcModel.LikeCount, 0)
	for _, v := range data.Items {
		likes = append(likes, buildLikeCount(v))
	}
	return likes
}
