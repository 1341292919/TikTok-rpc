package usecase

import (
	"TikTok-rpc/app/user/domain/model"
	"TikTok-rpc/app/user/domain/repository"
	"context"
)

// UserUseCase 接口应该不应该定义在 domain 中，这属于 use case 层
type UserUseCase interface {
	RegisterUser(ctx context.Context, user *model.User) (uid int64, err error)
	Login(ctx context.Context, user *model.User) (*model.User, error)
	GetUserInfo(ctx context.Context, user *model.User) (*model.User, error)
	UploadAvatar(ctx context.Context, user *model.User) (*model.User, error)
	GetMFACode(ctx context.Context, user *model.User) (*model.MFA, error)
	MFABind(ctx context.Context, user *model.User, code, secret string) error
}

// useCase 实现了 domain.UserUseCase
// 只会以接口的形式被调用, 所以首字母小写改为私有类型
type useCase struct {
	db repository.UserDB
	//svc *service.UserService
	//cache repository.UserCache
}

func NewUserCase(db repository.UserDB /*svc *service.UserService, re repository.UserCache*/) *useCase {
	return &useCase{
		db: db,
		//svc:   svc,
		//cache: re,
	}
}
