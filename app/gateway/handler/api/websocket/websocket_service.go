package websocket

import (
	"TikTok-rpc/app/gateway/pack"
	"TikTok-rpc/app/gateway/service"
	"TikTok-rpc/pkg/constants"
	"TikTok-rpc/pkg/errno"
	"context"
	"encoding/json"
	"strconv"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
)

var upgrade = websocket.HertzUpgrader{}

// Chat .
// @router / [GET]
func Chat(ctx context.Context, c *app.RequestContext) {
	var err error
	err = upgrade.Upgrade(c, func(conn *websocket.Conn) {
		uid := strconv.FormatInt(service.GetUserIDFromContext(c), 10)
		s := service.NewWebSocketService(ctx, c, conn)
		//登录
		if err := s.Login(); err != nil {
			conn.WriteJSON(pack.BuildFailResponse(errno.NewErrNo(errno.InternalWebSocketError, "Login Failed"+err.Error())))
			logger.Infof("Chat user Login failed:" + err.Error())
			return
		}
		conn.WriteJSON(pack.BuildResponse(errno.Success, "Welcome! "+uid))
		defer s.Logout()
		//读取离线信息
		if err := s.ReadOfflineMessage(); err != nil {
			conn.WriteJSON(pack.BuildFailResponse(errno.NewErrNo(errno.InternalWebSocketError, "ReadOfflineMessage"+err.Error())))
			logger.Infof("Chat ReadOfflineMessage failed:" + err.Error())
			return
		}
		//接收在线信息 以及监听请求
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				logger.Infof("Chat ReadMessage failed:" + err.Error())
				conn.WriteJSON(pack.BuildFailResponse(errno.NewErrNo(errno.InternalWebSocketError, "ReadMessage failed"+err.Error())))
				return
			}

			var msg pack.Request
			if err := json.Unmarshal(message, &msg); err != nil {
				logger.Infof("Chat ReadMessage failed:" + err.Error())
				conn.WriteJSON(pack.BuildFailResponse(errno.NewErrNo(errno.InternalWebSocketError, "Wrong JSON:"+err.Error())))
				continue
			}
			switch msg.Type {
			case constants.PrivateChat:
				err = s.SendPrivateMessage(msg.Data)
				if err != nil {
					logger.Infof("SendPrivateMessage failed:" + err.Error())
					conn.WriteJSON(pack.BuildFailResponse(errno.NewErrNo(errno.InternalWebSocketError, "SendPrivateMessage:"+err.Error())))
					continue
				}
			case constants.GroupChat:
				err = s.SendGroupMessage(msg.Data)
				if err != nil {
					logger.Infof("SendGroupMessage failed:" + err.Error())
					conn.WriteJSON(pack.BuildFailResponse(errno.NewErrNo(errno.InternalWebSocketError, "SendGroupMessage:"+err.Error())))
					continue
				}
			case constants.PrivateMessage:
				err = s.ReadPrivateMessage(msg)
				if err != nil {
					logger.Infof("ReadPrivateMessage failed:" + err.Error())
					conn.WriteJSON(pack.BuildFailResponse(errno.NewErrNo(errno.InternalWebSocketError, "ReadPrivateMessage:"+err.Error())))
					continue
				}
			case constants.GroupMessage:
				err = s.ReadGroupMessage(msg)
				if err != nil {
					logger.Infof("ReadGroupMessage failed:" + err.Error())
					conn.WriteJSON(pack.BuildFailResponse(errno.NewErrNo(errno.InternalWebSocketError, "ReadGroupMessage:"+err.Error())))
					continue
				}
			default:
				conn.WriteJSON(pack.BuildFailResponse(errno.NewErrNo(errno.InternalWebSocketError, "UnKnow Type:"+err.Error())))
				continue
			}
		}
	})
	if err != nil {
		pack.SendFailResponse(c, errno.WebSocketError)
		return
	}
}
