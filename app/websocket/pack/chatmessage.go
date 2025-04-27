package pack

import (
	rpc "TikTok-rpc/app/websocket/domain/model"
	"TikTok-rpc/kitex_gen/model"
	"strconv"
)

func BuildMessage(data *rpc.Message) *model.ChatMessage {
	return &model.ChatMessage{
		UserId:    strconv.FormatInt(data.UserId, 10),
		TargetId:  strconv.FormatInt(data.TargetId, 10),
		Content:   data.Content,
		CreatedAt: strconv.FormatInt(data.CreatedAT, 10),
		Type:      data.Type,
	}
}

func BuildMessageList(data []*rpc.Message) *model.ChatMessageList {
	list := make([]*model.ChatMessage, 0)
	for _, v := range data {
		list = append(list, BuildMessage(v))
	}
	return &model.ChatMessageList{
		Items: list,
		Total: int64(len(data)),
	}
}
