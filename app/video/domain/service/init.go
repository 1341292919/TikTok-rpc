package service

import "TikTok-rpc/app/video/domain/repository"

type VideoService struct {
	db    repository.VideoDB
	cache repository.VideoCache
}

func NewVideoService(db repository.VideoDB, cache repository.VideoCache) *VideoService {
	if db == nil {
		panic("userService`s db should not be nil")
	}
	if cache == nil {
		panic("userService`s cache should not be nil")
	}
	svc := &VideoService{
		db:    db,
		cache: cache,
	}
	return svc
}
