package mysql

import (
	"TikTok-rpc/app/interact/domain/model"
	"TikTok-rpc/app/interact/domain/repository"
	"TikTok-rpc/config"
	"TikTok-rpc/pkg/base/client"
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var _db repository.InteractDB

func initDB() {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}
	_db = NewInteractDB(gormDB)
}

func initConfig() {
	config.Init("video-test")
	initDB()
}

func TestInteractDB_CreateNewComment(t *testing.T) {
	initConfig()
	ctx := context.Background()
	req := buildTestModelInteractReq(t)
	Convey("TestInteractDB_CreateNewComment", t, func() {
		Convey("TestInteractDB_CreateNewComment_normal", func() {
			id, err := _db.CreateNewComment(ctx, req)
			So(err, ShouldBeNil)
			getcomment, err := _db.QueryCommentByParentId(ctx, req)
			So(err, ShouldBeNil)
			So(getcomment, ShouldNotBeEmpty)
			found := false
			for _, comment := range getcomment {
				if comment.Id == id {
					found = true
					break
				}
			}
			So(found, ShouldBeTrue)
		})
	})
}

func TestInteractDB_DeleteComment(t *testing.T) {
	initConfig()
	ctx := context.Background()
	req := buildTestModelInteractReq(t)
	Convey("TestInteractDB_DeleteComment", t, func() {
		Convey("TestInteractDB_DeleteComment_normal", func() {
			id, err := _db.CreateNewComment(ctx, req)
			So(err, ShouldBeNil)
			getcomment, err := _db.QueryCommentByParentId(ctx, req)
			So(err, ShouldBeNil)
			So(getcomment, ShouldNotBeEmpty)
			found := false
			for _, comment := range getcomment {
				if comment.Id == id {
					found = true
					break
				}
			}
			So(found, ShouldBeTrue)
			comment, err := _db.DeleteComment(ctx, &model.InteractReq{CommentId: id, Type: req.Type, Uid: req.Uid})
			So(err, ShouldBeNil)
			getcomment, err = _db.QueryCommentByParentId(ctx, req)
			So(err, ShouldBeNil)
			So(getcomment, ShouldNotBeEmpty)
			found = false
			for _, c := range getcomment {
				if c.Id == comment.Id {
					found = true
					break
				}
			}
			So(found, ShouldBeFalse)
		})
	})
}

func TestInteractDB_CreateNewUserLike(t *testing.T) {
	initConfig()
	ctx := context.Background()
	like := buildTestModelUserLike(t)
	var err error
	Convey("TestInteractDB_CreateNewUserLike", t, func() {
		Convey("TestInteractDB_CreateNewUserLike_normal", func() {
			if like.Type == 0 {
				err = _db.CreateNewUserLike(ctx, like.VideoId, like.Uid, like.Type)
			} else if like.Type == 1 {
				err = _db.CreateNewUserLike(ctx, like.CommentId, like.Uid, like.Type)
			}
			So(err, ShouldBeNil)
			likeData, err := _db.QueryUserLikeByUid(ctx, like.Uid)
			So(err, ShouldBeNil)
			So(likeData, ShouldNotBeEmpty)
			found := false
			for _, l := range likeData {
				if l.Uid == like.Uid && l.Type == like.Type && (l.CommentId == like.CommentId || l.VideoId == like.VideoId) {
					found = true
					break
				}
			}
			So(found, ShouldBeTrue)
		})
		Convey("TestInteractDB_DeleteUserLike_normal", func() {
			if like.Type == 0 {
				err = _db.DeleteUserLike(ctx, like.VideoId, like.Uid, like.Type)
			} else if like.Type == 1 {
				err = _db.DeleteUserLike(ctx, like.CommentId, like.Uid, like.Type)
			}
			So(err, ShouldBeNil)
			likeData, err := _db.QueryUserLikeByUid(ctx, like.Uid)
			So(err, ShouldBeNil)
			found := false
			for _, l := range likeData {
				if l.Uid == like.Uid && l.Type == like.Type && (l.CommentId == like.CommentId || l.VideoId == like.VideoId) {
					found = true
					break
				}
			}
			So(found, ShouldBeFalse)
		})
	})
}
