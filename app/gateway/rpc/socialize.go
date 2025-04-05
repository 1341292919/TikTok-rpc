package rpc

import (
	"TikTok-rpc/pkg/base/client"
	"github.com/bytedance/gopkg/util/logger"
)

func InitSocializeRPC() {
	c, err := client.InitSocializeRPC()
	if err != nil {
		logger.Fatalf("api.rpc.socialize InitSocializeRPC failed, err is %v", err)
	}
	socializeClient = *c
}
