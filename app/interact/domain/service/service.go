package service

import (
	"TikTok-rpc/app/interact/domain/model"
	"TikTok-rpc/pkg/errno"
	"context"
	"fmt"
)

func (svc *InteractService) IsIdOk(ctx context.Context, videoId, commentId int64) error {
	if videoId == 0 {
		//videoID不存在 检验comment_id
		exist, err := svc.db.IsCommentExist(ctx, commentId)
		if err != nil {
			return fmt.Errorf("check comment exist failed: %w", err)
		}
		if !exist {
			return errno.NewErrNo(errno.ServiceUserExistCode, "comment not exist")
		}
	} else {
		exist, err := svc.Rpc.IsVideoExist(ctx, videoId)

		if err != nil {
			return fmt.Errorf("check video exist failed: %w", err)
		}
		if !exist {
			return errno.NewErrNo(errno.ServiceUserExistCode, "video not exist")
		}
	}
	return nil
}
func (svc *InteractService) NewLike(ctx context.Context, req *model.InteractReq) error {
	var userLike *model.UserLike
	userLike = &model.UserLike{
		Uid:       req.Uid,
		Status:    req.ActionType,
		Type:      req.Type,
		VideoId:   req.VideoId,
		CommentId: req.CommentId,
	}
	err := svc.Mq.SendLikeMessage(ctx, userLike)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "SendNewLike :"+err.Error())
	}
	return nil
}
func (svc *InteractService) QueryLikeList(ctx context.Context, req *model.InteractReq) ([]*model.Video, error) {
	var vids []int64
	//点赞信息先从cache内访问
	data, err := svc.cache.QueryUserLikeByUid(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		likes, err := svc.db.QueryUserLikeByUid(ctx, req.Uid)
		if err != nil {
			return nil, err
		}
		for _, v := range likes {
			if v.Type == 1 {
				continue
			} //评论类型
			vids = append(vids, v.VideoId)
		}
	} else {
		for _, v := range data {
			vids = append(vids, v.VideoId)
		}
	}
	videoData, err := svc.Rpc.QueryVideoList(ctx, vids)
	if err != nil {
		return nil, err
	}
	count := int64(len(videoData))
	//按页分好
	startIndex := (req.PageNum - 1) * req.PageSize
	endIndex := startIndex + req.PageSize

	count = int64(len(data))
	if startIndex > count {
		return nil, nil
	}

	if endIndex > count {
		endIndex = count
	}
	return videoData[startIndex:endIndex], nil
}
func (svc *InteractService) Comment(ctx context.Context, req *model.InteractReq) error {
	//有video_id
	var c *model.CommentMessage
	if req.VideoId != 0 {
		req.Type = 0
		c = &model.CommentMessage{
			UId:      req.Uid,
			Content:  req.Content,
			Type:     req.Type,
			TargetId: req.VideoId,
			Delete:   0,
		}
	} else {
		req.Type = 1
		c = &model.CommentMessage{
			UId:      req.Uid,
			Content:  req.Content,
			Type:     req.Type,
			TargetId: req.CommentId,
			Delete:   0,
		}
	}
	err := svc.Mq.SendCommentMessage(ctx, c)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "SendCommentMessage :"+err.Error())
	}
	return nil
}
func (svc *InteractService) DeleteComment(ctx context.Context, req *model.InteractReq) error {
	//有video_id
	var c *model.CommentMessage
	if req.VideoId != 0 {
		req.Type = 0
		c = &model.CommentMessage{
			UId:      req.Uid,
			Content:  req.Content,
			Type:     req.Type,
			TargetId: req.VideoId,
			Delete:   1,
		}
	} else {
		req.Type = 1
		c = &model.CommentMessage{
			UId:      req.Uid,
			Content:  req.Content,
			Type:     req.Type,
			TargetId: req.CommentId,
			Delete:   1,
		}
	}
	err := svc.Mq.SendCommentMessage(ctx, c)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "SendDeleteCommentMessage :"+err.Error())
	}
	return nil
}
func (svc *InteractService) QueryCommentList(ctx context.Context, req *model.InteractReq) ([]*model.Comment, int64, error) {
	if req.VideoId != 0 {
		req.Type = 0
	} else {
		req.Type = 1
	}
	data, err := svc.db.QueryCommentByParentId(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	count := int64(len(data))
	return data, count, nil
}
