package main

import (
	interact "TikTok-rpc/rpc/interact/kitex_gen/interact"
	"context"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// Like implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) Like(ctx context.Context, req *interact.LikeRequest) (resp *interact.LikeResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryLikeList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) QueryLikeList(ctx context.Context, req *interact.QueryLikeListRequest) (resp *interact.QueryLikeListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentVideo implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentVideo(ctx context.Context, req *interact.CommentRequest) (resp *interact.CommentResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryCommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) QueryCommentList(ctx context.Context, req *interact.QueryCommentListRequest) (resp *interact.QueryCommentListResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteComment implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) DeleteComment(ctx context.Context, req *interact.DeleteCommentRequest) (resp *interact.DeleteCommentResponse, err error) {
	// TODO: Your code here...
	return
}
