// Code generated by hertz generator.

package main

import (
	"TikTok-rpc/app/gateway/middleware/jwt"
	"TikTok-rpc/app/gateway/rpc"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init() {
	rpc.Init()
	jwt.Init()
}
func main() {
	Init()
	h := server.Default()
	register(h)

	h.Spin()
}
