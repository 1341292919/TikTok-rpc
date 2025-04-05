namespace go user

include "model.thrift"
//注册
struct RegisterRequest{
    1: required string username ,
    2: required string password ,
}

struct RegisterResponse{
    1: model.BaseResp base,
    2: required i64 user_id,
}
//登录
struct LoginRequest{
    1: required string username
    2: required string password,
    3: optional string code ,
}

struct LoginResponse{
    1: model.BaseResp base,
    2: model.User data,
}

//以图搜图
struct SearchImagesRequest{
    1: binary data ,
}

struct SearchImagesResponse{
    1:model.BaseResp base,
    2:string data //返回图片存储的url
}

//上传头像
struct UploadAvatarRequest{
      1:required string avatar_url,
      2:required i64 user_id,
}
struct UploadAvatarResponse{
    1: model.BaseResp base,
    2: model.User data,
}
//获取用户信息
struct GetUserInformationRequest{
    1: required i64  user_id ,
}
struct GetUserInformationResponse{
    1: model.BaseResp base,
    2:model.User data,
}
//获取 MFA qrcode
struct GetMFARequest{
    1:required i64 user_id,
}
struct GetMFAResponse{
    1:model.BaseResp base,
    2:model.MFAMessage data,
}
//绑定多因素身份认证(MFA)
struct MFABindRequest{
    1:required string code ,
    2:required string secret ,
    3:required i64 user_id
}

struct MFABindResponse{
    1:model.BaseResp base,
}
service UserService {
    RegisterResponse Register (1: RegisterRequest req),
    LoginResponse Login(1: LoginRequest req),
    UploadAvatarResponse UploadAvatar(1:UploadAvatarRequest req),
    GetUserInformationResponse GetInformation(1:GetUserInformationRequest req),
    SearchImagesResponse SearchImage(1:SearchImagesRequest req),
    GetMFAResponse GetMFA(1:GetMFARequest req),
    MFABindResponse MindBind(1:MFABindRequest req)
}
