package mysql

import (
	"TikTok-rpc/app/user/domain/model"
	"TikTok-rpc/app/user/domain/repository"
	"TikTok-rpc/pkg/constants"
	"context"
	"gorm.io/gorm"
)

type userDB struct {
	client *gorm.DB
}

func NewUserDB(client *gorm.DB) repository.UserDB {
	return &userDB{client: client}
}

func (db *userDB) IsUserExist(ctx context.Context, user *model.User) (bool, error) {
	var count int64
	var err error
	if user.UserName != "" {
		err = db.client.WithContext(ctx).
			Table(constants.TableUser).
			Where("BINARY username = ?", user.UserName).
			Count(&count).
			Error
	} else {
		err = db.client.WithContext(ctx).
			Table(constants.TableUser).
			Where("BINARY id = ?", user.Uid).
			Count(&count).
			Error
	}
	//找不到会有错误
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

func (db *userDB) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	var userResp *User
	userResp = &User{
		Username: user.UserName,
		Password: user.Password,
	}
	err := db.client.WithContext(ctx).
		Table(constants.TableUser).
		Create(userResp).
		Error
	if err != nil {
		return 0, err
	}
	return userResp.Id, nil
}
func (db *userDB) GetUserInfo(ctx context.Context, user *model.User) (*model.User, error) {
	var userResp *User
	var err error
	if user.UserName != "" {
		err = db.client.WithContext(ctx).
			Table(constants.TableUser).
			Where("username = ?", user.UserName).
			Find(&userResp).
			Error
	} else {
		err = db.client.WithContext(ctx).
			Table(constants.TableUser).
			Where("id = ?", user.Uid).
			Find(&userResp).
			Error
	}
	if err != nil {
		return nil, err
	}

	return &model.User{
		Uid:       userResp.Id,
		UserName:  userResp.Username,
		Password:  userResp.Password,
		AvatarUrl: userResp.AvatarUrl,
		UpdateAT:  userResp.UpdatedAt.Unix(),
		CreateAT:  userResp.CreatedAt.Unix(),
		DeleteAT:  0,
	}, err
}

func (db *userDB) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	var userResp *User
	var err error
	err = db.client.WithContext(ctx).
		Table(constants.TableUser).
		Where("BINARY id = ?", user.Uid).
		Find(&userResp).
		Error
	if err != nil {
		return nil, err
	}

	if user.UserName != "" {
		userResp.Username = user.UserName
	}
	if user.Password != "" {
		userResp.Password = user.Password
	}
	if user.AvatarUrl != "" {
		userResp.AvatarUrl = user.AvatarUrl
	}

	err = db.client.WithContext(ctx).
		Table(constants.TableUser).
		Where("BINARY id = ?", user.Uid).
		Updates(&userResp).
		Error
	if err != nil {
		return nil, err
	}
	return &model.User{
		Uid:       userResp.Id,
		UserName:  userResp.Username,
		AvatarUrl: userResp.AvatarUrl,
		UpdateAT:  userResp.UpdatedAt.Unix(),
		CreateAT:  userResp.CreatedAt.Unix(),
		DeleteAT:  0,
	}, nil
}

func (db *userDB) UpdateMFA(ctx context.Context, user *model.User, mfa *model.MFAMessage) error {
	err := db.client.WithContext(ctx).
		Table(constants.TableUser).
		Where("BINARY id = ?", user.Uid).
		Update("opt_secret", mfa.Secret).
		Update("mfa_status", mfa.Status).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (db *userDB) CheckMFA(ctx context.Context, user *model.User) (*model.MFAMessage, error) {
	var userResp *User
	err := db.client.WithContext(ctx).
		Table(constants.TableUser).
		Where("BINARY id = ?", user.Uid).
		Find(&userResp).
		Error
	if err != nil {
		return nil, err
	}
	return &model.MFAMessage{
		Secret: userResp.OptSecret,
		Status: userResp.MfaStatus,
	}, nil

}
