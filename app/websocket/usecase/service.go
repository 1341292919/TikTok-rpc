package usecase

import (
	"TikTok-rpc/app/websocket/domain/model"
	"context"
)

func (uc *useCase) NewMessage(ctx context.Context, message *model.Message) error {
	return uc.svc.NewMessage(ctx, message)
}
func (uc *useCase) QueryOffLineMessage(ctx context.Context, targetid int64) ([]*model.Message, error) {
	data, err := uc.svc.QueryOfflineMessage(ctx, targetid)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (uc *useCase) QueryGroupMessage(ctx context.Context, req *model.ChatReq) ([]*model.Message, error) {
	return uc.svc.QueryGroupMessage(ctx, req)
}
func (uc *useCase) QueryPrivateMessage(ctx context.Context, req *model.ChatReq) ([]*model.Message, error) {
	return uc.svc.QueryPrivateMessage(ctx, req)
}
