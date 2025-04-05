package usecase

import (
	"TikTok-rpc/app/user/domain/model"
	"TikTok-rpc/app/user/domain/service"
	"TikTok-rpc/pkg/crypt"
	"TikTok-rpc/pkg/errno"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func (uc *useCase) RegisterUser(ctx context.Context, u *model.User) (uid int64, err error) {
	//这边应该完成用户注册的几个步骤 1.参数检验、2.用户存在检验、3.密码哈希、4.db create new user
	exist, err := uc.db.IsUserExist(ctx, u)

	if err != nil {
		return 0, fmt.Errorf("check user exist failed: %w", err)
	}
	if exist {
		return 0, errno.NewErrNo(errno.ServiceUserExist, "user already exist")
	}
	u.Password, err = crypt.PasswordHash(u.Password)
	if err != nil {
		return 0, fmt.Errorf("hash password failed: %w", err)
	}
	uid, err = uc.db.CreateUser(ctx, u)
	if err != nil {
		return 0, fmt.Errorf("create user failed: %w", err)
	}
	return uid, nil
}

func (uc *useCase) Login(ctx context.Context, user *model.User) (*model.User, error) {
	//登录应该完成的几个步骤 1.参数检验、2.用户存在检验、3密码检验 4.MFA
	exist, err := uc.db.IsUserExist(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("check user exist failed: %w", err)
	}
	if !exist {
		return nil, errno.NewErrNo(errno.ServiceUserExist, "user not exist")
	}
	var u *model.User
	u, err = uc.db.GetUserInfo(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("get user Info failed: %w", err)
	}
	//密码检验
	if !crypt.VerifyPassword(user.Password, u.Password) {
		return nil, errno.Errorf(errno.ServiceUserPasswordError, "password not match")
	}
	//事实上这边错误应该分为两种-1.不合规的code 2.code不符
	if u.OptSecret != "" {
		flag := service.TotpValidate(user.Code, u.OptSecret)
		if !flag {
			return nil, errno.NewErrNo(errno.ServiceUserExist, "invaild MFA code")
		}
	}
	return u, nil
}
func (uc *useCase) UploadAvatar(ctx context.Context, user *model.User) (*model.User, error) {
	//1.参数检验、
	u, err := uc.db.UpdateUser(ctx, user)
	hlog.Info(user.AvatarUrl)
	if err != nil {
		return nil, fmt.Errorf("update user failed: %w", err)
	}
	return u, nil
}
func (uc *useCase) GetUserInfo(ctx context.Context, user *model.User) (*model.User, error) {
	exist, err := uc.db.IsUserExist(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("check user exist failed: %w", err)
	}
	if !exist {
		return nil, errno.NewErrNo(errno.ServiceUserExist, "user not exist")
	}
	var u *model.User
	u, err = uc.db.GetUserInfo(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("get user Info failed: %w", err)
	}
	return u, nil
}

func (uc *useCase) GetMFACode(ctx context.Context, user *model.User) (*model.MFA, error) {
	//还应该有一层检验是否已经绑定了？
	userData, err := uc.GetUserInfo(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("check user meassage failed: %w", err)
	}
	MFA, err := service.OptSecret(userData)
	if err != nil {
		return nil, fmt.Errorf("generate mfa meassage failed: %w", err)
	}
	return MFA, nil
}

func (uc *useCase) MFABind(ctx context.Context, user *model.User, code, secret string) error {
	//检验code与secret
	//我们是不是应该检验传入的code，secret与用户id是否匹配？
	flag := service.TotpValidate(code, secret)
	if !flag {
		return errno.NewErrNo(errno.ServiceUserExist, "Invalid code and secret")
	}
	MFA := &model.MFAMessage{
		Secret: secret,
		Status: 1,
	}
	err := uc.db.UpdateMFA(ctx, user, MFA)
	if err != nil {
		return err
	}
	return nil
}
