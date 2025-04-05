package pack

import (
	"TikTok-rpc/app/gateway/model/model"
	rpc "TikTok-rpc/kitex_gen/model"
)

func User(data *rpc.User) *model.User {
	return &model.User{
		ID:        data.Id,
		Username:  data.Username,
		AvatarURL: data.AvatarUrl,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
