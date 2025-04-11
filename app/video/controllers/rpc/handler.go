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
	id, err := s.useCase.PublishVideo(ctx, v)
	if err != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(err))
	}
	resp.Id = id
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
	videoData, count, err := s.useCase.QueryPublishList(ctx, v)
	if err != nil {
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
	videoData, count, err := s.useCase.SearchVideo(ctx, v)
	if err != nil {
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
	videoData, count, err := s.useCase.PopularVideoList(ctx, v)
	if err != nil {
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
	videoData, count, err := s.useCase.QueryPublishList(ctx, v)
	if err != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(err))
		return
	}
	resp.Data = pack.BuildVideoList(videoData, count)
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}
