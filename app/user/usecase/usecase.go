package usecase

import (
	"TikTok-rpc/app/user/domain/model"
	"TikTok-rpc/app/user/domain/repository"
	"TikTok-rpc/app/user/domain/service"
	"context"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, user *model.User) (uid int64, err error)
	Login(ctx context.Context, user *model.User) (*model.User, error)
	GetUserInfo(ctx context.Context, user *model.User) (*model.User, error)
	UploadAvatar(ctx context.Context, user *model.User) (*model.User, error)
	GetMFACode(ctx context.Context, user *model.User) (*model.MFA, error)
	MFABind(ctx context.Context, user *model.User, code, secret string) error
	QueryUserIdByUsername(ctx context.Context, user *model.User) (int64, error)
}

// svc下辖db层
type useCase struct {
	db  repository.UserDB
	svc *service.UserService //想通过svc构建service层 作为db的上一层，所有的参数检验应该在这一层实现
}

func NewUserCase(db repository.UserDB, svc *service.UserService) *useCase {
	return &useCase{
		db:  db,
		svc: svc,
	}
}
