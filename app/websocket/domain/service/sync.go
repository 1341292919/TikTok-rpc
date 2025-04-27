package service

import (
	"TikTok-rpc/app/websocket/domain/model"
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"
)

func (svc *WebSocketService) UpdateDB(ctx context.Context) error {
	data, err := svc.cache.GetMessage(ctx, 1000)
	if err != nil {
		return err
	}
	valid := make([]*model.Message, 0)
	for _, v := range data {
		if v.Id == 1 {
			valid = append(valid, v)
		}
	}
	err = svc.db.UpdateMessageList(ctx, valid)
	if err != nil {
		return err
	}
	return nil
}

func SyncDB() {

	for {
		time.Sleep(15 * time.Second)
		err := svc.UpdateDB(context.Background())
		if err != nil {
			logger.Infof("failed to update db: %v", err)
			continue
		}
		hlog.Info("UpdateDB Success!")
	}

}
