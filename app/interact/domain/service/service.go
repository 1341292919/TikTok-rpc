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
	data, _, _, err := svc.cache.GetUserLikeMessage(ctx)
	//如果redis刚刚启动 向redis引入数据
	if err != nil {
		return err
	}
	if len(data) == 0 {
		data, err := svc.db.QueryAllUserLike(ctx)
		if err != nil {
			return err
		}
		err = svc.cache.UploadUserLike(ctx, data)
		cLikeCount, err := svc.db.QueryCommentLikeCount(ctx)
		if err != nil {
			return err
		}
		err = svc.cache.UploadLikeCount(ctx, cLikeCount)
		if err != nil {
			return err
		}
		vLikeCount, err := svc.Rpc.QueryVideoLikeCount(ctx)
		if err != nil {
			return err
		}
		err = svc.cache.UploadLikeCount(ctx, vLikeCount)
		if err != nil {
			return err
		}
	}
	//有video_id
	if req.VideoId != 0 {
		return svc.LikeVideo(ctx, req)
	} else {
		return svc.LikeComment(ctx, req)
	}
}

func (svc *InteractService) LikeVideo(ctx context.Context, req *model.InteractReq) error {
	exist, err := svc.cache.IsVideoLikeExist(ctx, req.VideoId, req.Uid)
	if err != nil {
		return err
	}
	if req.ActionType == 1 { //点赞操作 阻挡已经点赞
		if exist {
			return errno.NewErrNo(errno.ServiceRepeatOperation, "like exist")
		}
		err = svc.cache.NewVideoLike(ctx, req.VideoId, req.Uid)
		if err != nil {
			return err
		}
	} else if req.ActionType == 0 { //取消点赞操作，阻挡未曾点赞
		if !exist {
			return errno.NewErrNo(errno.ServiceRepeatOperation, "like not exist")
		}
		err = svc.cache.UnlikeVideo(ctx, req.VideoId, req.Uid)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *InteractService) LikeComment(ctx context.Context, req *model.InteractReq) error {
	exist, err := svc.cache.IsCommentLikeExist(ctx, req.CommentId, req.Uid)
	if err != nil {
		return err
	}
	if req.ActionType == 1 { //点赞操作 阻挡已经点赞
		if exist {
			return errno.NewErrNo(errno.ServiceRepeatOperation, "like exist")
		}
		err = svc.cache.NewCommentLike(ctx, req.CommentId, req.Uid)
		if err != nil {
			return err
		}
	} else if req.ActionType == 0 { //取消点赞操作，阻挡未曾点赞
		if !exist {
			return errno.NewErrNo(errno.ServiceRepeatOperation, "like not exist")
		}
		err = svc.cache.UnlikeComment(ctx, req.CommentId, req.Uid)
		if err != nil {
			return err
		}
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

func (svc *InteractService) Comment(ctx context.Context, req *model.InteractReq) (int64, error) {
	//有video_id
	if req.VideoId != 0 {
		return svc.CommentToVideo(ctx, req)
	} else {
		return svc.CommentToComment(ctx, req)
	}
}

func (svc *InteractService) CommentToVideo(ctx context.Context, req *model.InteractReq) (int64, error) {
	req.Type = 0
	id, err := svc.db.CreateNewComment(ctx, req)
	if err != nil {
		return 0, err
	}
	//调用rpc的服务更新视频的评论数
	err = svc.Rpc.UpdateVideoCommentCount(ctx, req.VideoId, 1)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (svc *InteractService) CommentToComment(ctx context.Context, req *model.InteractReq) (int64, error) {
	req.Type = 1
	id, err := svc.db.CreateNewComment(ctx, req)
	if err != nil {
		return 0, err
	}
	//更新评论的评论数
	err = svc.db.UpdateCommentCount(ctx, req.CommentId, 1)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (svc *InteractService) DeleteComment(ctx context.Context, req *model.InteractReq) error {
	c, err := svc.db.DeleteComment(ctx, req)
	if err != nil {
		return err
	}
	if c.Type == 0 {
		err = svc.Rpc.UpdateVideoCommentCount(ctx, c.ParentId, -1)
		if err != nil {
			return err
		}
	} else if c.Type == 1 {
		err = svc.db.UpdateCommentCount(ctx, c.ParentId, -1)
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
