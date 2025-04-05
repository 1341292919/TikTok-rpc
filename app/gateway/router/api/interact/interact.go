// Code generated by hertz generator. DO NOT EDIT.

package interact

import (
	interact "TikTok-rpc/app/gateway/handler/api/interact"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_comment := root.Group("/comment", _commentMw()...)
		_comment.DELETE("/delete", append(_deletecommentMw(), interact.DeleteComment)...)
		_comment.GET("/list", append(_querycommentlistMw(), interact.QueryCommentList)...)
		_comment.POST("/publish", append(_commentvideoMw(), interact.CommentVideo)...)
	}
	{
		_like := root.Group("/like", _likeMw()...)
		_like.POST("/action", append(_like0Mw(), interact.Like)...)
		_like.GET("/list", append(_querylikelistMw(), interact.QueryLikeList)...)
	}
}
