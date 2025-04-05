package utils

import (
	"TikTok-rpc/kitex_gen/model"
	"TikTok-rpc/pkg/errno"
)

func IsRPCSuccess(resp *model.BaseResp) bool {
	if resp.Code == errno.SuccessCode && resp.Msg == errno.SuccessMsg {
		return true
	}
	return false
}
