package pack

import (
	"TikTok-rpc/kitex_gen/model"
	"TikTok-rpc/pkg/errno"
)

type Base struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type Response struct {
	Base `json:"base"`
}

func BuildBaseResp(err errno.ErrNo) *model.BaseResp {
	return &model.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
