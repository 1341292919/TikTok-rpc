package pack

import (
	user "TikTok-rpc/app/user/domain/model"
	"TikTok-rpc/kitex_gen/model"
	"strconv"
)

// 将 框架中运行的user类型转换为idl中的
func BuildUser(data *user.User) *model.User {
	return &model.User{
		Id:        strconv.FormatInt(data.Uid, 10),
		Username:  data.UserName,
		AvatarUrl: data.AvatarUrl,
		CreatedAt: strconv.FormatInt(data.CreateAT, 10),
		UpdatedAt: strconv.FormatInt(data.UpdateAT, 10),
		DeletedAt: strconv.FormatInt(data.DeleteAT, 10),
	}
}
func BuildMFA(data *user.MFA) *model.MFAMessage {
	return &model.MFAMessage{
		Secret: data.Secret,
		Qrcode: data.Qrcode,
	}
}
