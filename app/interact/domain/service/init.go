package service

import "TikTok-rpc/app/interact/domain/repository"

type InteractService struct {
	db    repository.InteractDB
	cache repository.InteractCache
	Rpc   repository.RpcPort
}

var svc *InteractService

func NewInteractService(db repository.InteractDB, cache repository.InteractCache, rpc repository.RpcPort) *InteractService {
	if db == nil {
		panic("interactService`s db should not be nil")
	}
	if cache == nil {
		panic("interactService`s cache should not be nil")
	}
	if rpc == nil {
		panic("interactService`s rpc should not be nil")
	}
	svc = &InteractService{
		db:    db,
		cache: cache,
		Rpc:   rpc,
	}
	return svc
}
