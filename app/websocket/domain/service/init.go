package service

import "TikTok-rpc/app/websocket/domain/repository"

type WebSocketService struct {
	cache repository.WebsocketCache
	db    repository.WebsocketDB
}

var svc *WebSocketService

func NewWebSocketService(db repository.WebsocketDB, cache repository.WebsocketCache) *WebSocketService {
	if db == nil {
		panic("WebsocketService`s db should not be nil")
	}
	if cache == nil {
		panic("WebsocketService`s cache should not be nil")
	}
	svc = &WebSocketService{
		db:    db,
		cache: cache,
	}
	return svc
}
