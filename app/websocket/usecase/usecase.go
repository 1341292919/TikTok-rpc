package usecase

import (
	"TikTok-rpc/app/websocket/domain/model"
	"TikTok-rpc/app/websocket/domain/repository"
	"TikTok-rpc/app/websocket/domain/service"
	"context"
)

type WebSocketUseCase interface {
	NewMessage(ctx context.Context, message *model.Message) error
	QueryOffLineMessage(ctx context.Context, targetid int64) ([]*model.Message, error)
	QueryGroupMessage(ctx context.Context, req *model.ChatReq) ([]*model.Message, error)
	QueryPrivateMessage(ctx context.Context, req *model.ChatReq) ([]*model.Message, error)
}
type useCase struct {
	cache repository.WebsocketCache
	db    repository.WebsocketDB
	svc   *service.WebSocketService
}

func NewWebsocketUseCase(db repository.WebsocketDB, cache repository.WebsocketCache, svc *service.WebSocketService) *useCase {
	return &useCase{
		db:    db,
		svc:   svc,
		cache: cache,
	}
}
