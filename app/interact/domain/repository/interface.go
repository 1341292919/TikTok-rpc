package repository

import (
	"TikTok-rpc/app/interact/domain/model"
	"TikTok-rpc/pkg/kafka"
	"context"
)

type InteractDB interface {
	IsCommentExist(ctx context.Context, id int64) (bool, error)
	IsVideoLikeExist(ctx context.Context, id, uid int64) (bool, error)
	IsCommentLikeExist(ctx context.Context, id, uid int64) (bool, error)
	UpdateCommentLikeCount(ctx context.Context, cid, newcount int64) error
	CreateNewUserLike(ctx context.Context, targetid, uid, t int64) error
	DeleteUserLike(ctx context.Context, targetid, uid, t int64) error
	QueryAllUserLike(ctx context.Context) ([]*model.UserLike, error)
	QueryUserLikeByUid(ctx context.Context, uid int64) ([]*model.UserLike, error)
	CreateNewComment(ctx context.Context, req *model.CommentMessage) (int64, error)
	DeleteComment(ctx context.Context, req *model.CommentMessage) (*model.Comment, error)
	UpdateCommentCount(ctx context.Context, commentid, change int64) error
	QueryCommentByParentId(ctx context.Context, req *model.InteractReq) ([]*model.Comment, error)
	QueryCommentLikeCount(ctx context.Context) ([]*model.LikeCount, error)
}
type InteractCache interface {
	NewCommentLike(ctx context.Context, commentid, userid int64) error
	UnlikeComment(ctx context.Context, commentid, userid int64) error
	NewVideoLike(ctx context.Context, videoid, userid int64) error
	UnlikeVideo(ctx context.Context, videoid, userid int64) error
	UpdateLikeCount(ctx context.Context, id, value, t int64) error
	IsVideoLikeExist(ctx context.Context, videoid, userid int64) (bool, error)
	IsCommentLikeExist(ctx context.Context, commentid, userid int64) (bool, error)
	GetUserLikeMessage(ctx context.Context) ([]*model.UserLike, []*model.LikeCount, []*model.LikeCount, error)
	UploadUserLike(ctx context.Context, data []*model.UserLike) error
	UploadLikeCount(ctx context.Context, data []*model.LikeCount) error
	QueryUserLikeByUid(ctx context.Context, userid int64) ([]*model.UserLike, error)
}
type RpcPort interface {
	IsVideoExist(ctx context.Context, videoID int64) (bool, error)
	IsUserExist(ctx context.Context, userId int64) (bool, error)
	UpdateVideoCommentCount(ctx context.Context, videoID, count int64) error
	UpdateVideoLikeCount(ctx context.Context, videoID, count int64) error
	AddCount(ctx context.Context, videoID, t int64) error
	QueryVideoList(ctx context.Context, videoID []int64) ([]*model.Video, error)
	QueryVideoLikeCount(ctx context.Context) ([]*model.LikeCount, error)
}
type MqPort interface {
	SendLikeMessage(ctx context.Context, like *model.UserLike) error
	SendCommentMessage(ctx context.Context, comment *model.CommentMessage) error
	ConsumeLikeMessage(ctx context.Context) <-chan *kafka.Message
	ConsumeCommentMessage(ctx context.Context) <-chan *kafka.Message
}
