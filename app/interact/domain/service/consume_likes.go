package service

import (
	"TikTok-rpc/app/interact/domain/model"
	"TikTok-rpc/pkg/errno"
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/bytedance/sonic"
)

func (svc *InteractService) ConsumeLikes(ctx context.Context) {
	msgCh := svc.Mq.ConsumeLikeMessage(ctx)
	go func() {
		for msg := range msgCh {
			req := new(model.UserLike)
			err := sonic.Unmarshal(msg.V, req)
			if err != nil {
				logger.Errorf("InteractService: Consume Unmarshal msg err: %v", err)
			}
			err = svc.UpdateLike(ctx, req)
			if err != nil {
				logger.Errorf("InteractService: Consume NewLike err: %v", err)
			}
		}
	}()
}

func (svc *InteractService) UpdateLike(ctx context.Context, req *model.UserLike) error {
	data, _, _, err := svc.cache.GetUserLikeMessage(ctx)
	//如果redis刚刚启动 向redis引入数据 -这里本应该进行对数据的定向转载
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
	if req.Type == 0 {
		return svc.LikeVideo(ctx, req)
	} else {
		return svc.LikeComment(ctx, req)
	}
}

func (svc *InteractService) LikeVideo(ctx context.Context, req *model.UserLike) error {
	exist, err := svc.cache.IsVideoLikeExist(ctx, req.VideoId, req.Uid)
	if err != nil {
		return err
	}
	if req.Status == 1 { //点赞操作 阻挡已经点赞
		if exist {
			return errno.NewErrNo(errno.ServiceRepeatOperation, "like exist")
		}
		err = svc.cache.NewVideoLike(ctx, req.VideoId, req.Uid)

		if err != nil {
			return err
		}
	} else if req.Status == 0 { //取消点赞操作，阻挡未曾点赞
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

func (svc *InteractService) LikeComment(ctx context.Context, req *model.UserLike) error {
	exist, err := svc.cache.IsCommentLikeExist(ctx, req.CommentId, req.Uid)
	if err != nil {
		return err
	}
	if req.Status == 1 { //点赞操作 阻挡已经点赞
		if exist {
			return errno.NewErrNo(errno.ServiceRepeatOperation, "like exist")
		}
		err = svc.cache.NewCommentLike(ctx, req.CommentId, req.Uid)
		if err != nil {
			return err
		}
	} else if req.Status == 0 { //取消点赞操作，阻挡未曾点赞
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
