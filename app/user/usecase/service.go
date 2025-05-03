package usecase

import (
	"TikTok-rpc/app/user/domain/model"
	"TikTok-rpc/pkg/errno"
	"context"
	"fmt"
)

//所有service层出现的错误都应该正确由errno封装
//由于并没有对用户名和密码作要求所以省略了部分的参数检验

func (uc *useCase) RegisterUser(ctx context.Context, u *model.User) (uid int64, err error) {
	//这边应该完成用户注册的几个步骤 1.参数检验、2.用户存在检验、3.密码哈希、4.db create new user
	exist, err := uc.db.IsUserExist(ctx, u)
	if err != nil {
		return 0, fmt.Errorf("check user exist failed: %w", err)
	}
	if exist {
		return 0, errno.NewErrNo(errno.ServiceUserExistCode, "user already exist")
	}
	u.Password, err = uc.svc.PasswordHash(u.Password)
	if err != nil {
		return 0, fmt.Errorf("hash password failed: %w", err)
	}
	uid, err = uc.svc.CreateUser(ctx, u)
	if err != nil {
		return 0, fmt.Errorf("create user failed: %w", err)
	}
	return uid, nil
}

// 登录应该完成的几个步骤 1.参数检验、2.用户存在检验、3密码检验 4.MFA
func (uc *useCase) Login(ctx context.Context, user *model.User) (*model.User, error) {
	//用户存在检验
	exist, err := uc.db.IsUserExist(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("check user exist failed: %w", err)
	}
	if !exist {
		return nil, errno.NewErrNo(errno.ServiceUserNotExistCode, "user not exist")
	}
	var u *model.User
	u, err = uc.db.GetUserInfo(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("get user Info failed: %w", err)
	}
	//密码检验
	if !uc.svc.PasswordVerify(user.Password, u.Password) {
		return nil, errno.Errorf(errno.ServiceUserPasswordError, "password not match")
	}
	//以下三次调用应该合并吗？
	check, err := uc.svc.IsRequiredMFA(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("check MFA required failed: %w", err)
	}
	if check {
		//对code的存在进行检验
		if err = uc.svc.Verify(uc.svc.VerifyMFACode(user.Code)); err != nil {
			return nil, err
		}
		//code匹配与否
		flag, err := uc.svc.MFACheck(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("check MFA required failed: %w", err)
		}
		if !flag {
			return nil, errno.NewErrNo(errno.ParamLogicalErrorCode, "invaild MFA code")
		}
	}
	return u, nil
}

func (uc *useCase) UploadAvatar(ctx context.Context, user *model.User) (*model.User, error) {
	u, err := uc.svc.UploadAvatar(ctx, user)
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
		return nil, errno.NewErrNo(errno.ServiceUserNotExistCode, "user not exist")
	}
	var u *model.User
	u, err = uc.svc.GetUserInfoById(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("get user Info failed: %w", err)
	}
	return u, nil
}

func (uc *useCase) GetMFACode(ctx context.Context, user *model.User) (*model.MFA, error) {
	userData, err := uc.svc.GetUserInfoById(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("check user meassage failed: %w", err)
	}
	MFA, err := uc.svc.GetMFAQCode(ctx, userData)
	if err != nil {
		return nil, fmt.Errorf("generate mfa meassage failed: %w", err)
	}
	return MFA, nil
}

func (uc *useCase) MFABind(ctx context.Context, user *model.User, code, secret string) error {
	//我们是不是应该检验传入的code，secret与用户id是否匹配？
	flag := uc.svc.TotpValidate(code, secret)
	if !flag {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "Invalid code and secret")
	}
	MFA := &model.MFAMessage{
		Secret: secret,
		Status: 1,
	}
	err := uc.svc.UpdateMFA(ctx, user, MFA)
	if err != nil {
		return err
	}
	return nil
}

func (uc *useCase) QueryUserIdByUsername(ctx context.Context, user *model.User) (int64, error) {
	return uc.svc.QueryUserIdByUsername(ctx, user)
}
