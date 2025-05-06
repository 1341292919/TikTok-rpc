package cache

import (
	"TikTok-rpc/app/interact/domain/repository"
	"TikTok-rpc/config"
	"TikTok-rpc/pkg/base/client"
	"TikTok-rpc/pkg/constants"
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func initTest(t *testing.T) repository.InteractCache {
	t.Helper()
	config.Init("test")
	c1, err := client.NewRedisClient(constants.RedisDBInteract)
	if err != nil {
		panic(err)
	}
	c2, err := client.NewRedisClient(constants.RedisDBInteract)
	if err != nil {
		panic(err)
	}
	return NewInteractCache(c1, c2)
}
func TestInteractCache_NewVideoLike(t *testing.T) {
	ca := initTest(t)
	ctx := context.Background()
	vid := buildTestModelVideoId(t)
	uid := buildTestModelUserId(t)
	Convey("TestInteractCache_NewVideoLike", t, func() {
		Convey("TestInteractCache_NewVideoLike_normal", func() {
			err := ca.NewVideoLike(ctx, vid, uid)
			So(err, ShouldBeNil)
			userlikes, err := ca.QueryUserLikeByUid(ctx, uid)
			So(err, ShouldBeNil)
			found := false
			for _, l := range userlikes {
				if l.Uid == uid && l.VideoId == vid {
					found = true
					break
				}
			}
			So(found, ShouldBeTrue)
		})
		Convey("TestInteractCache_UnlikeVideo", func() {
			err := ca.UnlikeVideo(ctx, vid, uid)
			So(err, ShouldBeNil)
			userlikes, err := ca.QueryUserLikeByUid(ctx, uid)
			So(err, ShouldBeNil)
			found := false
			for _, l := range userlikes {
				if l.Uid == uid && l.VideoId == vid {
					found = true
					break
				}
			}
			So(found, ShouldBeFalse)
		})
	})
}
func TestInteractCache_NewCommentLike(t *testing.T) {
	ca := initTest(t)
	ctx := context.Background()
	cid := buildTestModelCommentId(t)
	uid := buildTestModelUserId(t)
	Convey("TestInteractCache_NewCommentLike", t, func() {
		Convey("TestInteractCache_NewCommentLike_normal", func() {
			err := ca.NewCommentLike(ctx, cid, uid)
			So(err, ShouldBeNil)
			userlikes, _, _, err := ca.GetUserLikeMessage(ctx)
			So(err, ShouldBeNil)
			found := false
			for _, l := range userlikes {
				if l.Uid == uid && l.CommentId == cid && l.Type == 1 {
					found = true
					break
				}
			}
			So(found, ShouldBeTrue)
		})
		Convey("TestInteractCache_UnlikeComment", func() {
			err := ca.UnlikeComment(ctx, cid, uid)
			So(err, ShouldBeNil)
			userlikes, _, _, err := ca.GetUserLikeMessage(ctx)
			So(err, ShouldBeNil)
			found := false
			for _, l := range userlikes {
				if l.Uid == uid && l.CommentId == cid && l.Type == 1 && l.Status == 1 {
					found = true
					break
				}
			}
			So(found, ShouldBeFalse)
		})
	})
}
