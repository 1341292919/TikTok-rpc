package service

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"
)

func (svc *InteractService) UpdateDB(ctx context.Context) error {
	vcount, err := svc.cache.QueryVideoLikeData(ctx)
	if err != nil {
		return err
	}
	Ccount, err := svc.cache.QueryCommentLikeData(ctx)
	if err != nil {
		return err
	}
	userlikes, err := svc.cache.QueryAllUserLike(ctx)
	for _, v := range vcount {
		err = svc.Rpc.UpdateVideoLikeCount(ctx, v.Id, v.Count)
		if err != nil {
			return err
		}
	}
	for _, c := range Ccount {
		err = svc.db.UpdateCommentLikeCount(ctx, c.Id, c.Count)
		if err != nil {
			return err
		}
	}
	for _, u := range userlikes {
		if u.Status == 0 {
			err = svc.db.CreateNewUserLike(ctx, u.VideoId, u.Uid, u.Type)
			if err != nil {
				return err
			}
		} else if u.Status == 1 {
			err = svc.db.DeleteUserLike(ctx, u.Uid, u.CommentId, u.Type)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func SyncDB() {
	for {
		time.Sleep(30 * time.Second)
		err := svc.UpdateDB(context.Background())
		if err != nil {
			logger.Infof("failed to update db: %v", err)
			continue
		}
		hlog.Info("UpdateDB Success!")
	}
}
