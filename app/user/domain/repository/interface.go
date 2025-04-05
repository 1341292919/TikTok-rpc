package repository

import (
	"TikTok-rpc/app/user/domain/model"
	"context"
)

type UserDB interface {
	IsUserExist(ctx context.Context, user *model.User) (bool, error)
	CreateUser(ctx context.Context, user *model.User) (int64, error)
	GetUserInfo(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateMFA(ctx context.Context, user *model.User, mfa *model.MFAMessage) error
}
