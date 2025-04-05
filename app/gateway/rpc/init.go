package rpc

import (
	"TikTok-rpc/kitex_gen/interact/interactservice"
	"TikTok-rpc/kitex_gen/socialize/socializeservice"
	"TikTok-rpc/kitex_gen/user/userservice"
	"TikTok-rpc/kitex_gen/video/videoservice"
)

// 全局变量
var (
	userClient      userservice.Client
	videoClient     videoservice.Client
	interactClient  interactservice.Client
	socializeClient socializeservice.Client
)

func Init() {
	InitUserRPC()
	InitVideoRPC()
	InitSocializeRPC()
	InitInteractRPC()
}
