package repository

import (
	"TikTok-rpc/app/video/domain/model"
	"context"
)

type VideoDB interface {
	CreateVideo(ctx context.Context, video *model.Video) (int64, error)
	QueryVideoByKeyWord(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error)
	QueryVideoByUid(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error)
	QueryVideoListById(ctx context.Context, id []string) ([]*model.Video, error)
	QueryPopularVideo(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error)
	IsVideoExist(ctx context.Context, uid int64) (bool, error)
	QueryVideoById(ctx context.Context, id string) (*model.Video, error)
	UpdateCommentCount(ctx context.Context, videoid, changecount int64) error
	UpdateLikeCount(ctx context.Context, videoid, likecount int64) error
	QueryVideoDuringTime(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error)
	QueryLikeCount(ctx context.Context) ([]*model.LikeCount, error)
}
type VideoCache interface {
	NewIdToRank(ctx context.Context, vid int64) error
	GetVideoIdByRank(ctx context.Context, count int64) ([]string, error)
	GetVideoByRank(ctx context.Context, count int64) ([]*model.Video, error)
	AddVideoToRank(ctx context.Context, video []*model.Video) error
	DeleteVideoRank(ctx context.Context) error
	DeleteVideoIdRank(ctx context.Context) error
}

type VideoRpc interface {
	QueryUserIdByUsername(ctx context.Context, username string) (int64, error)
}
