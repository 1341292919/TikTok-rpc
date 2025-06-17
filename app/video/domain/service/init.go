package service

import "TikTok-rpc/app/video/domain/repository"

type VideoService struct {
	db    repository.VideoDB
	cache repository.VideoCache
	Rpc   repository.VideoRpc
}

func NewVideoService(db repository.VideoDB, cache repository.VideoCache, Rpc repository.VideoRpc) *VideoService {
	if db == nil {
		panic("userService`s db should not be nil")
	}
	if cache == nil {
		panic("userService`s cache should not be nil")
	}
	svc := &VideoService{
		db:    db,
		cache: cache,
		Rpc:   Rpc,
	}
	return svc
}
