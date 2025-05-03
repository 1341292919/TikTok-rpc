package websock

import (
	"TikTok-rpc/app/gateway/handler/api/websocket"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func register(h *server.Hertz) {
	h.GET(`/ws`, append(_homeMW(), websocket.Chat)...)
}
