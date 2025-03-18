// Code generated by hertz generator.

package interact

import (
	"context"

	interact "TikTok-rpc/biz/model/interact"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Like .
// @router /like/action [POST]
func Like(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.LikeRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interact.LikeResponse)

	c.JSON(consts.StatusOK, resp)
}

// QueryLikeList .
// @router /like/list [GET]
func QueryLikeList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.QueryLikeListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interact.QueryLikeListResponse)

	c.JSON(consts.StatusOK, resp)
}

// CommentVideo .
// @router /comment/publish [POST]
func CommentVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.CommentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interact.CommentResponse)

	c.JSON(consts.StatusOK, resp)
}

// QueryCommentList .
// @router /comment/list [GET]
func QueryCommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.QueryCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interact.QueryCommentListResponse)

	c.JSON(consts.StatusOK, resp)
}

// DeleteComment .
// @router /comment/delete [DELETE]
func DeleteComment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.DeleteCommentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interact.DeleteCommentResponse)

	c.JSON(consts.StatusOK, resp)
}
