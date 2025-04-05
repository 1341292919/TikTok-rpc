package rpc

import (
	"TikTok-rpc/pkg/base/client"
	"github.com/bytedance/gopkg/util/logger"
)

func InitVideoRPC() {
	c, err := client.InitVideoRPC()
	if err != nil {
		logger.Fatalf("api.rpc.video InitVideoRPC failed, err is %v", err)
	}
	videoClient = *c
}
