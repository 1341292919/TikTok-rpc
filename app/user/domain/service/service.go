package service

import (
	"TikTok-rpc/app/user/domain/model"
	"context"
)

func (svc *UserService) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	return svc.db.CreateUser(ctx, user)
}

func (svc *UserService) IsRequiredMFA(ctx context.Context, user *model.User) (bool, error) {
	MFAMessage, err := svc.db.CheckMFA(ctx, user)
	if err != nil {
		return false, err
	}
	if MFAMessage.Secret == "" && MFAMessage.Status == 0 {
		return false, nil
	}
	return true, nil
}
func (svc *UserService) GetMFAQCode(ctx context.Context, user *model.User) (*model.MFA, error) {
	MFAMessage, err := svc.db.CheckMFA(ctx, user)
	if err != nil {
		return nil, err
	}
	MFA, err := svc.OptSecret(user.UserName, MFAMessage)
	if err != nil {
		return nil, err
	}
	return MFA, nil
}
func (svc *UserService) MFACheck(ctx context.Context, user *model.User) (bool, error) {
	MFAMessage, err := svc.db.CheckMFA(ctx, user)
	if err != nil {
		return false, err
	}
	return svc.TotpValidate(user.Code, MFAMessage.Secret), nil

}
func (svc *UserService) UploadAvatar(ctx context.Context, user *model.User) (*model.User, error) {
	return svc.db.UpdateUser(ctx, user)
}
func (svc *UserService) GetUserInfoById(ctx context.Context, user *model.User) (*model.User, error) {
	return svc.db.GetUserInfo(ctx, user)
}

func (svc *UserService) UpdateMFA(ctx context.Context, user *model.User, mfa *model.MFAMessage) error {
	return svc.db.UpdateMFA(ctx, user, mfa)
}
func (svc *UserService) QueryUserIdByUsername(ctx context.Context, user *model.User) (int64, error) {
	data, err := svc.db.QueryUserIdByUsername(ctx, user)
	if err != nil {
		return -1, err
	}
	return data.Uid, nil
}
