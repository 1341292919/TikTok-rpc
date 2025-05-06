package pack

import (
	"TikTok-rpc/app/gateway/model/model"
	"TikTok-rpc/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Base struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type Response struct {
	Base `json:"base"`
}

// 数据类型多样-用interface
func SendResponse(c *app.RequestContext, data interface{}) {
	c.JSON(consts.StatusOK, data)
}

func BuildBaseResp(err errno.ErrNo) *model.BaseResp {
	return &model.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}

func SendFailResponse(c *app.RequestContext, err errno.ErrNo) {
	baseResp := Base{
		Code: err.ErrorCode,
		Msg:  err.Error(),
	}
	response := Response{
		Base: baseResp,
	}
	SendResponse(c, response)
}
