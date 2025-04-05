// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	api_interact "TikTok-rpc/app/gateway/router/api/interact"
	api_socialize "TikTok-rpc/app/gateway/router/api/socialize"
	api_user "TikTok-rpc/app/gateway/router/api/user"
	api_video "TikTok-rpc/app/gateway/router/api/video"
	model "TikTok-rpc/app/gateway/router/model"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	model.Register(r)

	api_video.Register(r)

	api_user.Register(r)

	api_socialize.Register(r)

	api_interact.Register(r)
}
