// Code generated by hertz generator.

package socialize

import (
	"context"

	socialize "TikTok-rpc/app/gateway/model/api/socialize"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Follow .
// @router /relation/action [POST]
func Follow(ctx context.Context, c *app.RequestContext) {
	var err error
	var req socialize.FollowRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(socialize.FollowResponse)

	c.JSON(consts.StatusOK, resp)
}

// QueryFollowList .
// @router /following/list [GET]
func QueryFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req socialize.QueryFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(socialize.QueryFollowListResponse)

	c.JSON(consts.StatusOK, resp)
}

// QueryFollowerList .
// @router /follower/list [GET]
func QueryFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req socialize.QueryFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(socialize.QueryFollowerListResponse)

	c.JSON(consts.StatusOK, resp)
}

// QueryFriendList .
// @router /friends/list [GET]
func QueryFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req socialize.QueryFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(socialize.QueryFriendListResponse)

	c.JSON(consts.StatusOK, resp)
}
