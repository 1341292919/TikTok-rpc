package service

import (
	"TikTok-rpc/app/user/domain/repository"
)

type UserService struct {
	db repository.UserDB
}

func NewUserService(db repository.UserDB) *UserService {
	if db == nil {
		panic("userService`s db should not be nil")
	}

	svc := &UserService{
		db: db,
	}
	return svc
}
