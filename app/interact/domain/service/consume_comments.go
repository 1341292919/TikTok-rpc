package service

import (
	"TikTok-rpc/app/interact/domain/model"
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/bytedance/sonic"
)

func (svc *InteractService) ConsumeComments(ctx context.Context) {
	msgCh := svc.Mq.ConsumeCommentMessage(ctx)
	go func() {
		for msg := range msgCh {
			req := new(model.CommentMessage)
			err := sonic.Unmarshal(msg.V, req)
			if err != nil {
				logger.Errorf("InteractService: Consume Unmarshal msg err: %v", err)
			}
			if req.Delete == 1 {
				err = svc.UpdateDeleteComment(ctx, req)
				if err != nil {
					logger.Errorf("InteractService: Consume NewDeleteComment err: %v", err)
				}
			}
			_, err = svc.UpdateComment(ctx, req)
			if err != nil {
				logger.Errorf("InteractService: Consume NewComment err: %v", err)
			}
		}
	}()
}

func (svc *InteractService) UpdateComment(ctx context.Context, req *model.CommentMessage) (int64, error) {
	if req.Type == 0 {
		return svc.CommentToVideo(ctx, req)
	} else {
		return svc.CommentToComment(ctx, req)
	}
}

func (svc *InteractService) CommentToVideo(ctx context.Context, req *model.CommentMessage) (int64, error) {
	id, err := svc.db.CreateNewComment(ctx, req)
	if err != nil {
		return 0, err
	}
	//调用rpc的服务更新视频的评论数
	err = svc.Rpc.UpdateVideoCommentCount(ctx, req.TargetId, 1)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (svc *InteractService) CommentToComment(ctx context.Context, req *model.CommentMessage) (int64, error) {
	id, err := svc.db.CreateNewComment(ctx, req)
	if err != nil {
		return 0, err
	}
	//更新评论的评论数
	err = svc.db.UpdateCommentCount(ctx, req.TargetId, 1)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (svc *InteractService) UpdateDeleteComment(ctx context.Context, req *model.CommentMessage) error {
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
