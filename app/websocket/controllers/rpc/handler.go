package rpc

import (
	"TikTok-rpc/app/websocket/domain/model"
	"TikTok-rpc/app/websocket/pack"
	"TikTok-rpc/app/websocket/usecase"
	"TikTok-rpc/kitex_gen/websocket"
	"TikTok-rpc/pkg/errno"
	"context"
)

type WebsocketServiceImpl struct {
	useCase usecase.WebSocketUseCase
}

func NewWebsocketServiceImpl(useCase usecase.WebSocketUseCase) *WebsocketServiceImpl {
	return &WebsocketServiceImpl{useCase: useCase}
}

func (s *WebsocketServiceImpl) AddMessage(ctx context.Context, req *websocket.AddMessageRequest) (resp *websocket.AddMessageResponse, err error) {
	resp = new(websocket.AddMessageResponse)
	chat := &model.Message{
		UserId:   req.Id,
		Content:  req.Content,
		TargetId: req.TargetId,
		Status:   req.Status,
		Type:     req.Status,
	}
	e := s.useCase.NewMessage(ctx, chat)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

func (s *WebsocketServiceImpl) QueryOfflineMessage(ctx context.Context, req *websocket.QueryOfflineMessageRequest) (resp *websocket.QueryOfflineMessageResponse, err error) {
	resp = new(websocket.QueryOfflineMessageResponse)
	data, e := s.useCase.QueryOffLineMessage(ctx, req.Id)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Data = pack.BuildMessageList(data)
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}
func (s *WebsocketServiceImpl) QueryPrivateHistoryMessage(ctx context.Context, req *websocket.QueryPrivateHistoryMessageRequest) (resp *websocket.QueryPrivateHistoryMessageResponse, err error) {
	resp = new(websocket.QueryPrivateHistoryMessageResponse)
	data, e := s.useCase.QueryPrivateMessage(ctx, &model.ChatReq{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		UserId:   req.UserId,
		TargetId: req.TargetId,
	})
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Data = pack.BuildMessageList(data)
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}
func (s *WebsocketServiceImpl) QueryGroupHistoryMessage(ctx context.Context, req *websocket.QueryGroupHistoryMessageRequest) (resp *websocket.QueryGroupHistoryMessageResponse, err error) {
	resp = new(websocket.QueryGroupHistoryMessageResponse)
	data, e := s.useCase.QueryGroupMessage(ctx, &model.ChatReq{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		UserId:   req.UserId,
		TargetId: req.TargetId,
	})
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Data = pack.BuildMessageList(data)
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}
