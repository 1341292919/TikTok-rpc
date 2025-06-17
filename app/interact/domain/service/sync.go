package service

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"
)

func (svc *InteractService) UpdateDB(ctx context.Context) error {
	userLikes, vCount, cCount, err := svc.cache.GetUserLikeMessage(ctx)
	if err != nil {
		return err
	}
	for _, v := range vCount {
		hlog.Info(v.Count)
		err = svc.Rpc.UpdateVideoLikeCount(ctx, v.Id, v.Count)
		if err != nil {
			logger.Infof(err.Error())
			continue
		}
	}
	for _, c := range cCount {
		err = svc.db.UpdateCommentLikeCount(ctx, c.Id, c.Count)
		if err != nil {
			logger.Infof(err.Error())
			continue
		}
	}
	for _, u := range userLikes {
		if u.Status == 1 {
			if u.Type == 0 {
				err = svc.db.CreateNewUserLike(ctx, u.VideoId, u.Uid, u.Type)
			} else if u.Type == 1 {
				err = svc.db.CreateNewUserLike(ctx, u.CommentId, u.Uid, u.Type)
			}
			if err != nil {
				logger.Infof(err.Error())
				continue
			}
		} else if u.Status == 0 {
			if u.Type == 0 {
				err = svc.db.DeleteUserLike(ctx, u.VideoId, u.Uid, u.Type)
			} else if u.Type == 1 {
				err = svc.db.DeleteUserLike(ctx, u.CommentId, u.Uid, u.Type)
			}
			if err != nil {
				logger.Infof(err.Error())
				continue
			}
		}
	}
	return nil
}
func (svc *InteractService) SyncDB() {
	defer svc.UpdateDB(context.Background())
	for {
		err := svc.UpdateDB(context.Background())
		if err != nil {
			logger.Infof("failed to update db: %v", err)
			continue
		}
		time.Sleep(30 * time.Minute)
	}
}
