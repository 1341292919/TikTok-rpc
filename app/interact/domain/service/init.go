package service

import (
	"TikTok-rpc/app/interact/domain/repository"
	"context"
)

type InteractService struct {
	db    repository.InteractDB
	cache repository.InteractCache
	Rpc   repository.RpcPort
	Mq    repository.MqPort
}

func NewInteractService(db repository.InteractDB, cache repository.InteractCache, rpc repository.RpcPort, mq repository.MqPort) *InteractService {
	if db == nil {
		panic("interactService`s db should not be nil")
	}
	if cache == nil {
		panic("interactService`s cache should not be nil")
	}
	if rpc == nil {
		panic("interactService`s rpc should not be nil")
	}
	if mq == nil {
		panic("interactService`s mq should not be nil")
	}
	svc := &InteractService{
		db:    db,
		cache: cache,
		Rpc:   rpc,
		Mq:    mq,
	}
	svc.init()
	return svc
}

func (svc *InteractService) init() {
	svc.initLikeConsumer()
	svc.initCommentConsumer()
	svc.initSync()
}
func (svc *InteractService) initSync() {
	go svc.SyncDB()
}
func (svc *InteractService) initLikeConsumer() {
	go svc.ConsumeLikes(context.Background())
}
func (svc *InteractService) initCommentConsumer() {
	go svc.ConsumeComments(context.Background())
}
