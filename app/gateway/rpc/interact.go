package rpc

import (
	"TikTok-rpc/pkg/base/client"
	"github.com/bytedance/gopkg/util/logger"
)

func InitInteractRPC() {
	c, err := client.InitInteractRPC()
	if err != nil {
		logger.Fatalf("api.rpc.interact InitInteractRPC failed, err is %v", err)
	}
	interactClient = *c
}
