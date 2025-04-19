package rpc

import (
	"TikTok-rpc/app/video/domain/model"
	"TikTok-rpc/app/video/pack"
	"TikTok-rpc/app/video/usecase"
	"TikTok-rpc/kitex_gen/video"
	"TikTok-rpc/pkg/errno"
	"context"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct {
	useCase usecase.VideoUseCase
}

func NewVideoServiceImpl(useCase usecase.VideoUseCase) *VideoServiceImpl {
	return &VideoServiceImpl{useCase: useCase}
}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishRequest) (resp *video.PublishResponse, err error) {
	resp = new(video.PublishResponse)
	v := &model.Video{
		VideoUrl:    req.VideoUrl,
		CoverUrl:    req.CoverUrl,
		Uid:         req.UserId,
		Title:       req.Title,
		Description: req.Description,
	}
	id, e := s.useCase.PublishVideo(ctx, v)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(err))
		return
	}
	resp.Id = &id
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

// QueryList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) QueryList(ctx context.Context, req *video.QueryPublishListRequest) (resp *video.QueryPublishListResponse, err error) {
	resp = new(video.QueryPublishListResponse)
	v := &model.VideoReq{
		Uid:      req.UserId,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}
	videoData, count, e := s.useCase.QueryPublishList(ctx, v)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(err))
		return
	}
	resp.Data = pack.BuildVideoList(videoData, count)
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

// SearchVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) SearchVideo(ctx context.Context, req *video.SearchVideoByKeywordRequest) (resp *video.SearchVideoByKeywordResponse, err error) {
	resp = new(video.SearchVideoByKeywordResponse)
	v := new(model.VideoReq)
	if req.ToDate == nil || req.FromDate == nil {
		v = &model.VideoReq{
			Keyword:  req.Keyword,
			PageNum:  req.PageNum,
			PageSize: req.PageSize,
		}
	} else {
		v = &model.VideoReq{
			Keyword:  req.Keyword,
			PageNum:  req.PageNum,
			PageSize: req.PageSize,
			ToDate:   *req.ToDate,
			FromDate: *req.FromDate,
		}
	}
	if req.Username != nil {
		v.Username = *req.Username
	}
	videoData, count, e := s.useCase.SearchVideo(ctx, v)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(err))
		return
	}
	resp.Data = pack.BuildVideoList(videoData, count)
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

// GetPopularVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPopularVideo(ctx context.Context, req *video.GetPopularListRequest) (resp *video.GetPopularListResponse, err error) {
	resp = new(video.GetPopularListResponse)
	v := &model.VideoReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}
	videoData, count, e := s.useCase.PopularVideoList(ctx, v)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(err))
		return
	}
	resp.Data = pack.BuildVideoList(videoData, count)
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

// GetVideoStream implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoStream(ctx context.Context, req *video.VideoStreamRequest) (resp *video.VideoStreamResponse, err error) {
	resp = new(video.VideoStreamResponse)
	var v *model.VideoReq
	if req.LatestTime == nil {
		v = &model.VideoReq{
			PageNum:  req.PageNum,
			PageSize: req.PageSize,
		}
	} else {
		v = &model.VideoReq{
			PageNum:  req.PageNum,
			PageSize: req.PageSize,
			ToDate:   *req.LatestTime,
		}
	}
	videoData, count, e := s.useCase.GetVideoStream(ctx, v)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(err))
		return
	}
	resp.Data = pack.BuildVideoList(videoData, count)
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

func (s *VideoServiceImpl) QueryVideoById(ctx context.Context, req *video.QueryVideoByVIdRequest) (resp *video.QueryVideoByVIdResponse, err error) {
	resp = new(video.QueryVideoByVIdResponse)
	videoData, e := s.useCase.QueryVideoByID(ctx, req.VideoId)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Data = pack.BuildVideo(videoData)
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

func (s *VideoServiceImpl) UpdateCommentCount(ctx context.Context, req *video.UpdateVideoCommentCountRequest) (resp *video.UpdateVideoCommentCountResponse, err error) {
	resp = new(video.UpdateVideoCommentCountResponse)
	e := s.useCase.UpdateCommentCount(ctx, req.VideoId, req.ChangeCount)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(err))
		return
	}
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

func (s *VideoServiceImpl) UpdateLikeCount(ctx context.Context, req *video.UpdateVideoLikeCountRequest) (resp *video.UpdateVideoLikeCountResponse, err error) {
	resp = new(video.UpdateVideoLikeCountResponse)
	e := s.useCase.UpdateLikeCount(ctx, req.VideoId, req.LikeCount)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(err))
		return
	}
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}
