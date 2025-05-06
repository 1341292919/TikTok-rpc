package service

import (
	"TikTok-rpc/app/websocket/domain/model"
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func (svc *WebSocketService) NewMessage(ctx context.Context, req *model.Message) error {
	req.CreatedAT = time.Now().Unix()
	req.Id = 1
	err := svc.cache.NewMessage(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (svc *WebSocketService) QueryOfflineMessage(ctx context.Context, tagetid int64) ([]*model.Message, error) {
	data, err := svc.db.QueryTargetMessage(ctx, tagetid)
	if err != nil {
		return nil, err
	}
	targetMessage := make([]*model.Message, 0)
	for _, v := range data {
		hlog.Info(v)
		if v.Status == 0 {
			targetMessage = append(targetMessage, v)
		}
	}
	return targetMessage, nil
}

func (svc *WebSocketService) QueryPrivateMessage(ctx context.Context, req *model.ChatReq) ([]*model.Message, error) {
	data, err := svc.db.QueryPrivateMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (svc *WebSocketService) QueryGroupMessage(ctx context.Context, req *model.ChatReq) ([]*model.Message, error) {
	data, err := svc.db.QueryGroupMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	return data, nil
}
