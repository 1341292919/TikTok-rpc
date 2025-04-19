package usecase

import (
	"TikTok-rpc/app/video/domain/model"
	"TikTok-rpc/app/video/domain/repository"
	"TikTok-rpc/app/video/domain/service"
	"context"
)

type VideoUseCase interface {
	PublishVideo(ctx context.Context, video *model.Video) (id int64, err error)
	QueryPublishList(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error)
	SearchVideo(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error)
	PopularVideoList(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error)
	GetVideoStream(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error)
	QueryVideoByID(ctx context.Context, videoid int64) (*model.Video, error)
	UpdateCommentCount(ctx context.Context, videoid, ccount int64) (err error)
	UpdateLikeCount(ctx context.Context, videoid, lcount int64) (err error)
}

type useCase struct {
	db    repository.VideoDB
	svc   *service.VideoService
	cache repository.VideoCache
	Rpc   repository.VideoRpc
}

func NewVideoUseCase(db repository.VideoDB, svc *service.VideoService, cache repository.VideoCache, rpc repository.VideoRpc) *useCase {
	return &useCase{
		db:    db,
		svc:   svc,
		cache: cache,
		Rpc:   rpc,
	}
}
