package usecase_test

import (
	"TikTok-rpc/app/user/domain/model"
	"TikTok-rpc/app/user/domain/service"
	"TikTok-rpc/app/user/usecase"
	"TikTok-rpc/app/user/usecase/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUseCase_RegisterUser(t *testing.T) {
	mockDB := new(mocks.UserDB)
	mockService := service.NewUserService(mockDB)
	uc := usecase.NewUserUseCase(mockDB, mockService)
	user := &model.User{
		UserName: "testuser",
		Password: "password",
	}
	mockDB.On("IsUserExist", mock.Anything, user).Return(false, nil)
	mockDB.On("CreateUser", mock.Anything, user).Return(int64(1), nil)

	uid, err := uc.RegisterUser(context.Background(), user)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), uid)

	mockDB.AssertExpectations(t)
}

func TestUseCase_GetUserMessage(t *testing.T) {
	mockDB := new(mocks.UserDB)
	mockService := service.NewUserService(mockDB)
	uc := usecase.NewUserUseCase(mockDB, mockService)
	user := &model.User{
		Uid:       1,
		UserName:  "testuser",
		AvatarUrl: "test",
	}
	mockDB.On("IsUserExist", mock.Anything, user).Return(true, nil)
	mockDB.On("GetUserInfo", mock.Anything, user).Return(user, nil)

	userResp, err := uc.GetUserInfo(context.Background(), user)
	assert.Nil(t, err)
	assert.NotNil(t, userResp)
	mockDB.AssertExpectations(t)
}
