package service

import (
	"TikTok-rpc/app/gateway/model/model"
	"TikTok-rpc/app/gateway/pack"
	"TikTok-rpc/app/gateway/rpc"
	"TikTok-rpc/kitex_gen/user"
	web "TikTok-rpc/kitex_gen/websocket"
	"TikTok-rpc/pkg/errno"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"strconv"
	"sync"
	"time"
)

type WebSocketService struct {
	ctx  context.Context
	c    *app.RequestContext
	conn *websocket.Conn
}
type _user struct {
	username string
	conn     *websocket.Conn
}

var (
	userMapMutex sync.RWMutex //信号量限制userMap的访问
	userMap      = make(map[string]*_user)
)

func NewWebSocketService(ctx context.Context, c *app.RequestContext, conn *websocket.Conn) *WebSocketService {
	return &WebSocketService{
		ctx:  ctx,
		c:    c,
		conn: conn,
	}
}

func (s *WebSocketService) Login() error {
	uid := strconv.FormatInt(GetUserIDFromContext(s.c), 10)
	data, err := s.QueryUser(GetUserIDFromContext(s.c))
	if err != nil {
		return err
	}
	userMapMutex.Lock()
	userMap[uid] = &_user{conn: s.conn, username: data.Username}
	userMapMutex.Unlock()
	return nil
}

func (s *WebSocketService) Logout() {
	uid := strconv.FormatInt(GetUserIDFromContext(s.c), 10)
	userMapMutex.Lock()
	delete(userMap, uid)
	userMapMutex.Unlock()
}

func (s *WebSocketService) SendPrivateMessage(req pack.MessageReq) error {
	id := GetUserIDFromContext(s.c)
	//确认用户存在
	data, err := s.QueryUser(req.TargetId)
	if err != nil {
		return err
	}
	if data == nil {
		return errno.NewErrNo(errno.ServiceUserNotExistCode, "SendMessage: user not found")
	}
	to := strconv.FormatInt(req.TargetId, 10)

	userMapMutex.RLock()
	toConn := userMap[to]
	userMapMutex.RUnlock()

	switch toConn {
	case nil: //离线
		{
			err = s.SaveMessage(id, req.TargetId, req.Content, true)
			if err != nil {
				return errno.NewErrNo(errno.ServiceUserNotExistCode, "Offline:SendMessage:"+err.Error())
			}
		}
	default: // 在线
		{
			err = toConn.conn.WriteJSON(pack.BuildResponse(errno.Success, &model.Message{
				Content:  req.Content,
				UserID:   strconv.FormatInt(id, 10),
				TargetID: strconv.FormatInt(req.TargetId, 10),

				CreatedAt: pack.ChangeFormat(strconv.FormatInt(time.Now().Unix(), 10)),
				Type:      0,
			}))
			if err != nil {
				return errno.NewErrNo(errno.ServiceUserNotExistCode, "SendMessage:"+err.Error())
			}
			err = s.SaveMessage(id, req.TargetId, req.Content, false)
			if err != nil {
				return errno.NewErrNo(errno.ServiceUserNotExistCode, "Offline:SendMessage:"+err.Error())
			}
		}
	}
	return nil
}

func (s *WebSocketService) ReadOfflineMessage() error {
	id := GetUserIDFromContext(s.c)
	req := &web.QueryOfflineMessageRequest{
		Id: id,
	}
	messageResp, err := rpc.ReadOfflineMessageRPC(s.ctx, req)
	if err != nil {
		return err
	}
	s.conn.WriteJSON(pack.BuildResponse(errno.Success, pack.BuildMessageList(messageResp)))
	return nil
}

func (s *WebSocketService) SendGroupMessage(req pack.MessageReq) error {
	id := GetUserIDFromContext(s.c)
	err := rpc.AddMessageRPC(s.ctx, &web.AddMessageRequest{
		Id:       id,
		TargetId: req.TargetId,
		Content:  req.Content,
		Status:   1,
		Type:     1,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *WebSocketService) ReadPrivateMessage(req pack.Request) error {
	id := GetUserIDFromContext(s.c)
	data, err := s.QueryUser(req.Data.TargetId)
	if err != nil {
		return err
	}
	if data == nil {
		return errno.NewErrNo(errno.ServiceUserNotExistCode, "SendMessage: user not found")
	}
	messageResp, err := rpc.ReadPrivateHistoryMessageRPC(s.ctx, &web.QueryPrivateHistoryMessageRequest{
		UserId:   id,
		TargetId: req.Data.TargetId,
		PageSize: req.Param.PageSize,
		PageNum:  req.Param.PageNum,
	})
	if err != nil {
		return err
	}
	s.conn.WriteJSON(pack.BuildResponse(errno.Success, pack.BuildMessageList(messageResp)))
	return nil
}

func (s *WebSocketService) ReadGroupMessage(req pack.Request) error {
	id := GetUserIDFromContext(s.c)
	messageResp, err := rpc.ReadGroupHistoryMessageRPC(s.ctx, &web.QueryGroupHistoryMessageRequest{
		UserId:   id,
		TargetId: req.Data.TargetId,
		PageSize: req.Param.PageSize,
		PageNum:  req.Param.PageNum,
	})
	if err != nil {
		return err
	}
	s.conn.WriteJSON(pack.BuildResponse(errno.Success, pack.BuildMessageList(messageResp)))
	return nil
}

func (s *WebSocketService) QueryUser(id int64) (*model.User, error) {
	req := &user.GetUserInformationRequest{
		UserId: id,
	}
	resp, err := rpc.GetUserMessagesRPC(s.ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (s *WebSocketService) SaveMessage(id, targetid int64, content string, off bool) error {
	var err error
	if off {
		err = rpc.AddMessageRPC(s.ctx, &web.AddMessageRequest{
			Id:       id,
			TargetId: targetid,
			Content:  content,
			Status:   0,
		})
		if err != nil {
			return err
		}
	} else {
		err = rpc.AddMessageRPC(s.ctx, &web.AddMessageRequest{
			Id:       id,
			TargetId: targetid,
			Content:  content,
			Status:   1,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
