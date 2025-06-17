package repository

import (
	"TikTok-rpc/app/websocket/domain/model"
	"context"
)

type WebsocketDB interface {
	UpdateMessageList(ctx context.Context, message []*model.Message) error
	QueryTargetMessage(ctx context.Context, targetid int64) ([]*model.Message, error)
	QueryPrivateMessage(ctx context.Context, req *model.ChatReq) ([]*model.Message, error)
	QueryGroupMessage(ctx context.Context, req *model.ChatReq) ([]*model.Message, error)
}
type WebsocketCache interface {
	NewMessage(ctx context.Context, message *model.Message) error
	GetMessage(ctx context.Context, count int64) ([]*model.Message, error)
	NewMessageList(ctx context.Context, m []*model.Message) error
}
