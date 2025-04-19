package usecase

import (
	"TikTok-rpc/app/video/domain/model"
	"TikTok-rpc/pkg/errno"
	"context"
	"fmt"
	"strconv"
)

// 1.参数的检验如：uid的存在与否、2.创建video
func (uc *useCase) PublishVideo(ctx context.Context, video *model.Video) (int64, error) {
	id, err := uc.svc.CreateVideo(ctx, video)
	if err != nil {
		return 0, fmt.Errorf("create video failed: %w", err)
	}
	return id, nil
}

// 1.参数检验 uid存在与否、page参数非负，但uid不存在则返回视频数目为0似乎也不是错误
func (uc *useCase) QueryPublishList(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error) {
	err := uc.svc.Verify(uc.svc.VerifyPageParam(req.PageSize, req.PageNum))
	if err != nil {
		return nil, 0, err
	}
	v, count, err := uc.svc.QueryPublishList(ctx, req)
	if err != nil {
		return nil, -1, fmt.Errorf("query publish list failed: %w", err)
	}
	return v, count, nil
}

// 1.参数检验：日期是否合理，page参数非负
func (uc *useCase) SearchVideo(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error) {
	err := uc.svc.Verify(uc.svc.VerifyDate(req.ToDate, req.FromDate), uc.svc.VerifyPageParam(req.PageSize, req.PageNum))
	if err != nil {
		return nil, -1, err
	}
	v, count, err := uc.svc.SearchVideoByKeyWord(ctx, req)
	if err != nil {
		return nil, -1, fmt.Errorf("search video by key failed: %w", err)
	}
	return v, count, nil
}
func (uc *useCase) PopularVideoList(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error) {
	err := uc.svc.Verify(uc.svc.VerifyDate(req.ToDate, req.FromDate), uc.svc.VerifyPageParam(req.PageSize, req.PageNum))
	if err != nil {
		return nil, -1, err
	}
	v, count, err := uc.svc.QueryPopularVideoList(ctx, req)
	if err != nil {
		return nil, -1, fmt.Errorf("query videos list failed: %w", err)
	}
	return v, count, nil
}

// 按理视频流应该是短时间内多次访问，应该要用redis，同时可能要完成看过的视频不再推送-这一点是怎么做到的
func (uc *useCase) GetVideoStream(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error) {
	err := uc.svc.Verify(uc.svc.VerifyPageParam(req.PageSize, req.PageNum))
	v, count, err := uc.svc.VideoStream(ctx, req)
	if err != nil {
		return nil, -1, fmt.Errorf("video stream failed: %w", err)
	}
	return v, count, nil
}
func (uc *useCase) QueryVideoByID(ctx context.Context, videoid int64) (*model.Video, error) {
	exist, err := uc.db.IsVideoExist(ctx, videoid)
	if err != nil {
		return nil, fmt.Errorf("Check videoId exist failed: %w", err)
	}
	if !exist {
		return nil, errno.NewErrNo(errno.ServiceVideoNotExist, "Video not exist")
	}
	id := strconv.Itoa(int(videoid))
	data, err := uc.db.QueryVideoById(ctx, id)
	return data, err
}
func (uc *useCase) UpdateCommentCount(ctx context.Context, videoid, ccount int64) (err error) {
	return uc.db.UpdateCommentCount(ctx, videoid, ccount)
}
func (uc *useCase) UpdateLikeCount(ctx context.Context, videoid, lcount int64) (err error) {
	return uc.db.UpdateLikeCount(ctx, videoid, lcount)
}
