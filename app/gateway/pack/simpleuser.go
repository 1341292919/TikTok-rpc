package pack

import (
	"TikTok-rpc/app/gateway/model/model"
	rpc "TikTok-rpc/kitex_gen/model"
)

func SimpleUser(data *rpc.SimpleUser) *model.SimpleUser {
	return &model.SimpleUser{
		ID:        data.Id,
		Username:  data.Username,
		AvatarURL: data.AvatarUrl,
	}
}
func SimpleUserList(data *rpc.SimpleUserList) *model.SimpleUserList {
	uList := make([]*model.SimpleUser, 0)
	for _, v := range data.Items {
		uList = append(uList, SimpleUser(v))
	}
	return &model.SimpleUserList{
		Total: data.Total,
		Items: uList,
	}
}
