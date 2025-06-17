package rpc

import (
	"TikTok-rpc/app/interact/domain/model"
	"TikTok-rpc/app/interact/pack"
	"TikTok-rpc/app/interact/usecase"
	interact "TikTok-rpc/kitex_gen/interact"
	"TikTok-rpc/pkg/errno"
	"context"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct {
	useCase usecase.InteractUseCase
}

func NewInteractServiceImpl(useCase usecase.InteractUseCase) *InteractServiceImpl {
	return &InteractServiceImpl{
		useCase: useCase,
	}
}

// Like implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) Like(ctx context.Context, req *interact.LikeRequest) (resp *interact.LikeResponse, err error) {
	resp = new(interact.LikeResponse)
	interactReq := new(model.InteractReq)
	if req.CommentId == nil {
		interactReq = &model.InteractReq{
			VideoId:    *req.VideoId,
			ActionType: req.ActionType,
			Uid:        req.UserId,
		}
	} else if req.VideoId == nil {
		interactReq = &model.InteractReq{
			CommentId:  *req.CommentId,
			ActionType: req.ActionType,
			Uid:        req.UserId,
		}
	} else {
		interactReq = &model.InteractReq{
			VideoId:    *req.VideoId,
			CommentId:  *req.CommentId,
			ActionType: req.ActionType,
			Uid:        req.UserId,
		}
	}
	e := s.useCase.Like(ctx, interactReq)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

// QueryLikeList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) QueryLikeList(ctx context.Context, req *interact.QueryLikeListRequest) (resp *interact.QueryLikeListResponse, err error) {
	resp = new(interact.QueryLikeListResponse)
	interactReq := &model.InteractReq{
		Uid:      req.UserId,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
	}
	data, count, e := s.useCase.QueryLikeList(ctx, interactReq)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Data = pack.BuildVideoList(data, count)
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

// CommentVideo implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) Comment(ctx context.Context, req *interact.CommentRequest) (resp *interact.CommentResponse, err error) {
	resp = new(interact.CommentResponse)
	interactReq := new(model.InteractReq)
	if req.CommentId == nil {
		interactReq = &model.InteractReq{
			VideoId: *req.VideoId,
			Content: req.Content,
			Uid:     req.UserId,
		}
	} else if req.VideoId == nil {
		interactReq = &model.InteractReq{
			CommentId: *req.CommentId,
			Content:   req.Content,
			Uid:       req.UserId,
		}
	}
	e := s.useCase.Comment(ctx, interactReq)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

// QueryCommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) QueryCommentList(ctx context.Context, req *interact.QueryCommentListRequest) (resp *interact.QueryCommentListResponse, err error) {
	resp = new(interact.QueryCommentListResponse)
	interactReq := new(model.InteractReq)
	if req.CommentId == nil {
		interactReq = &model.InteractReq{
			VideoId:  *req.VideoId,
			PageSize: req.PageSize,
			PageNum:  req.PageNum,
		}
	} else if req.VideoId == nil {
		interactReq = &model.InteractReq{
			CommentId: *req.CommentId,
			PageNum:   req.PageNum,
			PageSize:  req.PageSize,
		}
	}
	data, count, e := s.useCase.QueryCommentList(ctx, interactReq)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Data = pack.BuildCommentList(data, count)
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}

// DeleteComment implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) DeleteComment(ctx context.Context, req *interact.DeleteCommentRequest) (resp *interact.DeleteCommentResponse, err error) {
	resp = new(interact.DeleteCommentResponse)
	interactReq := &model.InteractReq{
		CommentId: *req.CommentId,
		Uid:       req.UserId,
		VideoId:   *req.VideoId,
	}
	e := s.useCase.DeleteComment(ctx, interactReq)
	if e != nil {
		resp.Base = pack.BuildBaseResp(errno.ConvertErr(e))
		return
	}
	resp.Base = pack.BuildBaseResp(errno.Success)
	return
}
