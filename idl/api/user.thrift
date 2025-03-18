namespace go user

include "model.thrift"
//注册
struct RegisterRequest{
    1: required string username (api.form="username"),
    2: required string password (api.form="password"),
}

struct RegisterResponse{
    1: model.BaseResp base,
}
//登录
struct LoginRequest{
    1: required string username(api.form="username"),
    2: required string password(api.form="password"),
    3:optional string code (api.form="code"),
}

struct LoginResponse{
    1: model.BaseResp base,
    2: model.User data,
}

//以图搜图
struct SearchImagesRequest{
    1: binary data (api.form="data"),
}

struct SearchImagesResponse{
    1:model.BaseResp base,
    2:string data //返回图片存储的url
}

//上传头像
struct UploadAvatarRequest{
     1: binary data (api.form="data"),
}
struct UploadAvatarResponse{
    1: model.BaseResp base,
    2: model.User data,
}
//获取用户信息
struct GetUserInformationRequest{
    1: required i64  user_id (api.form="user_id"),
}
struct GetUserInformationResponse{
    1: model.BaseResp base,
    2:model.User data,
}
//获取 MFA qrcode
struct GetMFARequest{

}
struct GetMFAResponse{
    1:model.BaseResp base,
    2:model.MFAMessage data,
}
//绑定多因素身份认证(MFA)
struct MFABindRequest{
    1:required string code (api.form="code"),
    2:required string secret (api.form="secret"),
}

struct MFABindResponse{
    1:model.BaseResp base,
}
service UserService {
    RegisterResponse Register (1: RegisterRequest req) (api.post="/user/register"),
    LoginResponse Login(1: LoginRequest req) (api.post="/user/login"),
    UploadAvatarResponse UploadAvatar(1:UploadAvatarRequest req)(api.put="/user/avatar/upload"),
    GetUserInformationResponse GetInformation(1:GetUserInformationRequest req)(api.get="/user/info")
    SearchImagesResponse SearchImage(1:SearchImagesRequest req)(api.post="/user/image/search")//怎么不是get
    GetMFAResponse GetMFA(1:GetMFARequest req)(api.get="/user/image/search")
    MFABindResponse MindBind(1:MFABindRequest req)(api.post="/user/image/search")
}
