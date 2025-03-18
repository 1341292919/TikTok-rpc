package main

import (
	user "TikTok-rpc/rpc/user/kitex_gen/user"
	"context"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UploadAvatar implements the UserServiceImpl interface.
func (s *UserServiceImpl) UploadAvatar(ctx context.Context, req *user.UploadAvatarRequest) (resp *user.UploadAvatarResponse, err error) {
	// TODO: Your code here...
	return
}

// GetInformation implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetInformation(ctx context.Context, req *user.GetUserInformationRequest) (resp *user.GetUserInformationResponse, err error) {
	// TODO: Your code here...
	return
}

// SearchImage implements the UserServiceImpl interface.
func (s *UserServiceImpl) SearchImage(ctx context.Context, req *user.SearchImagesRequest) (resp *user.SearchImagesResponse, err error) {
	// TODO: Your code here...
	return
}

// GetMFA implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetMFA(ctx context.Context, req *user.GetMFARequest) (resp *user.GetMFAResponse, err error) {
	// TODO: Your code here...
	return
}

// MindBind implements the UserServiceImpl interface.
func (s *UserServiceImpl) MindBind(ctx context.Context, req *user.MFABindRequest) (resp *user.MFABindResponse, err error) {
	// TODO: Your code here...
	return
}
