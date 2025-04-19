package repository

import (
	"TikTok-rpc/app/interact/domain/model"
	"context"
)

type InteractDB interface {
	IsCommentExist(ctx context.Context, id int64) (bool, error)
	IsVideoLikeExist(ctx context.Context, id, uid int64) (bool, error)
	IsCommentLikeExist(ctx context.Context, id, uid int64) (bool, error)
	UpdateCommentLikeCount(ctx context.Context, cid, newcount int64) error
	CreateNewUserLike(ctx context.Context, cid, uid, t int64) error
	DeleteUserLike(ctx context.Context, targetid, uid, t int64) error
	QueryAllUserLike(ctx context.Context) ([]*model.UserLike, error)
	QueryUserLikeByUid(ctx context.Context, uid int64) ([]*model.UserLike, error)
	CreateNewComment(ctx context.Context, req *model.InteractReq) (int64, error)
	DeleteComment(ctx context.Context, req *model.InteractReq) (*model.Comment, error)
	UpdateCommentCount(ctx context.Context, commentid, change int64) error
	QueryCommentByParentId(ctx context.Context, req *model.InteractReq) ([]*model.Comment, error)
}
type InteractCache interface {
	NewVideoLike(ctx context.Context, videoid, userid int64) error
	NewCommentLike(ctx context.Context, commentid, userid int64) error
	IsVideoLikeExist(ctx context.Context, videoid, userid int64) (bool, error)
	IsCommentLikeExist(ctx context.Context, commentid, userid int64) (bool, error)
	QueryVideoLikeData(ctx context.Context) ([]*model.VideoLikeCountKey, error)
	QueryCommentLikeData(ctx context.Context) ([]*model.CommentLikeCountKey, error)
	QueryAllUserLike(ctx context.Context) ([]*model.LikeKey, error)
	QueryUserLikeById(ctx context.Context, userid int64) ([]*model.LikeKey, error)
}
type RpcPort interface {
	IsVideoExist(ctx context.Context, videoID int64) (bool, error)
	IsUserExist(ctx context.Context, userId int64) (bool, error)
	UpdateVideoCommentCount(ctx context.Context, videoID, count int64) error
	UpdateVideoLikeCount(ctx context.Context, videoID, count int64) error
	AddCount(ctx context.Context, videoID, t int64) error
	QueryVideoList(ctx context.Context, videoID []int64) ([]*model.Video, error)
}
