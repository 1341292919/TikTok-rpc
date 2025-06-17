package rpc

import (
	"TikTok-rpc/app/user/domain/model"
	"TikTok-rpc/app/user/pack"
	"TikTok-rpc/app/user/usecase"
	"TikTok-rpc/kitex_gen/user"
	"TikTok-rpc/pkg/errno"
	"context"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	useCase usecase.UserUseCase
}

func NewUserServiceImpl(useCase usecase.UserUseCase) *UserServiceImpl {
	return &UserServiceImpl{useCase: useCase}
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = new(user.RegisterResponse)
	u := &model.User{
		UserName: req.Username,
		Password: req.Password,
	}
	uid, e := s.useCase.RegisterUser(ctx, u)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.UserId = &uid
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)
	u := new(model.User)
	//实际上这边code应该放到service检验 但是空指针报错？
	if req.Code == nil {
		u = &model.User{
			UserName: req.Username,
			Password: req.Password,
		}
	} else {
		u = &model.User{
			UserName: req.Username,
			Password: req.Password,
			Code:     *req.Code,
		}
	}
	userData, e := s.useCase.Login(ctx, u)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Data = pack.BuildUser(userData)
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

// UploadAvatar implements the UserServiceImpl interface.
func (s *UserServiceImpl) UploadAvatar(ctx context.Context, req *user.UploadAvatarRequest) (resp *user.UploadAvatarResponse, err error) {
	resp = new(user.UploadAvatarResponse)
	u := &model.User{
		AvatarUrl: req.AvatarUrl,
		Uid:       req.UserId,
	}
	userData, e := s.useCase.UploadAvatar(ctx, u)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Data = pack.BuildUser(userData)
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

// GetInformation implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetInformation(ctx context.Context, req *user.GetUserInformationRequest) (resp *user.GetUserInformationResponse, err error) {
	resp = new(user.GetUserInformationResponse)
	u := &model.User{
		Uid: req.UserId,
	}
	userData, e := s.useCase.GetUserInfo(ctx, u)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Data = pack.BuildUser(userData)
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

// SearchImage implements the UserServiceImpl interface.
func (s *UserServiceImpl) SearchImage(ctx context.Context, req *user.SearchImagesRequest) (resp *user.SearchImagesResponse, err error) {
	// TODO: Your code here...
	return
}

// GetMFA implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetMFA(ctx context.Context, req *user.GetMFARequest) (resp *user.GetMFAResponse, err error) {
	resp = new(user.GetMFAResponse)
	u := &model.User{
		Uid: req.UserId,
	}
	userData, e := s.useCase.GetMFACode(ctx, u)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Data = pack.BuildMFA(userData)
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

// MindBind implements the UserServiceImpl interface.
func (s *UserServiceImpl) MindBind(ctx context.Context, req *user.MFABindRequest) (resp *user.MFABindResponse, err error) {
	resp = new(user.MFABindResponse)
	u := &model.User{
		Uid: req.UserId,
	}
	e := s.useCase.MFABind(ctx, u, req.Code, req.Secret)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}
func (s *UserServiceImpl) QueryUserIdByUsername(ctx context.Context, req *user.QueryUserIdByUsernameRequest) (resp *user.QueryUserIdByUsernameResponse, err error) {
	resp = new(user.QueryUserIdByUsernameResponse)
	u := &model.User{
		UserName: req.Username,
	}
	id, e := s.useCase.QueryUserIdByUsername(ctx, u)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.Id = &id
	return
}
