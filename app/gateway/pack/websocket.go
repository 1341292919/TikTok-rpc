package pack

import (
	"TikTok-rpc/app/gateway/model/model"
	rpc "TikTok-rpc/kitex_gen/model"
	"TikTok-rpc/pkg/errno"
)

type MessageReq struct {
	Content  string `json:"content"`
	TargetId int64  `json:"target_id"`
}
type Request struct {
	Type  string     `json:"type"`
	Data  MessageReq `json:"data"`
	Param Param      `json:"param"`
}
type Param struct {
	PageSize int64 `json:"page_size"`
	PageNum  int64 `json:"page_num"`
}
type WebsocketResponse struct {
	Base `json:"base"`
	Data interface{} `json:"data"`
}

func BuildResponse(err errno.ErrNo, data interface{}) WebsocketResponse {
	return WebsocketResponse{
		Base: Base{
			Code: err.ErrorCode,
			Msg:  err.ErrorMsg,
		},
		Data: data,
	}
}
func BuildFailResponse(err errno.ErrNo) Response {
	return Response{
		Base: Base{
			Code: err.ErrorCode,
			Msg:  err.ErrorMsg,
		},
	}
}

func BuildMessage(data *rpc.ChatMessage) *model.Message {
	return &model.Message{
		UserID:    data.UserId,
		Type:      data.Type,
		Content:   data.Content,
		TargetID:  data.TargetId,
		CreatedAt: ChangeFormat(data.CreatedAt),
	}
}
func BuildMessageList(data *rpc.ChatMessageList) *model.MessageList {
	list := make([]*model.Message, 0)
	for _, v := range data.Items {
		list = append(list, BuildMessage(v))
	}
	return &model.MessageList{
		Items: list,
		Total: data.Total,
	}
}
