package websock

import (
	"TikTok-rpc/app/gateway/router/api/auth"

	"github.com/cloudwego/hertz/pkg/app"
)

func _wsAuth() []app.HandlerFunc {
	return append(make([]app.HandlerFunc, 0),
		tokenAuthFunc(),
	)
}

func tokenAuthFunc() app.HandlerFunc {
	return auth.DoubleTokenAuthFunc()
}
