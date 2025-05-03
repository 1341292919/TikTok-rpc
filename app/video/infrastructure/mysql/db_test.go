package mysql

import (
	"TikTok-rpc/app/video/domain/repository"
	"TikTok-rpc/config"
	"TikTok-rpc/pkg/base/client"
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

var _db repository.VideoDB

func initDB() {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}
	_db = NewVideoDB(gormDB)
}

func initConfig() {
	config.Init("video-test")
	initDB()
}

func TestVideoDB_CreateVideo(t *testing.T) {
	initConfig()
	ctx := context.Background()
	video := buildTestModelVideo(t)
	Convey("TestVideoDB_CreateVideo", t, func() {
		Convey("TestVideoDB_CreateVideo_normal", func() {
			id, err := _db.CreateVideo(ctx, video)
			So(err, ShouldBeNil)
			getVideo, err := _db.QueryVideoById(ctx, strconv.FormatInt(id, 10))
			So(err, ShouldBeNil)
			So(getVideo.Id, ShouldEqual, getVideo.Id)
			So(getVideo.Uid, ShouldEqual, getVideo.Uid)
			So(getVideo.Title, ShouldEqual, getVideo.Title)
			So(getVideo.CoverUrl, ShouldEqual, getVideo.CoverUrl)
			So(getVideo.VideoUrl, ShouldEqual, getVideo.VideoUrl)
			So(getVideo.Description, ShouldEqual, getVideo.Description)
		})
	})
}

func TestVideoDB_QueryVideoByKeyWord(t *testing.T) {
	initConfig()
	ctx := context.Background()
	videoReq := buildTestModelVideoReq(t)
	Convey("TestVideoDB_QueryVideoByKeyWord", t, func() {
		Convey("TestVideoDB_QueryVideoByKeyWord_normal", func() {
			_, _, err := _db.QueryVideoByKeyWord(ctx, videoReq)
			So(err, ShouldBeNil)
		})
	})
}

func TestVideoDB_UpdateLikeCount(t *testing.T) {
	initConfig()
	ctx := context.Background()
	video := buildTestModelVideo(t)
	Convey("TestVideoDB_UpdateLikeCount", t, func() {
		Convey("TestVideoDB_UpdateLikeCount_normal", func() {
			id, err := _db.CreateVideo(ctx, video)
			So(err, ShouldBeNil)
			getVideo, err := _db.QueryVideoById(ctx, strconv.FormatInt(id, 10))
			So(err, ShouldBeNil)
			So(getVideo.Uid, ShouldEqual, getVideo.Uid)
			err = _db.UpdateLikeCount(ctx, getVideo.Uid, 1)
			So(err, ShouldBeNil)
			getVideo2, err := _db.QueryVideoById(ctx, strconv.FormatInt(id, 10))
			So(err, ShouldBeNil)
			So(getVideo.LikeCount, ShouldEqual, getVideo2.LikeCount)
		})
	})
}
func TestVideoDB_UpdateCommentCount(t *testing.T) {
	initConfig()
	ctx := context.Background()
	video := buildTestModelVideo(t)
	Convey("TestVideoDB_UpdateCommentCount", t, func() {
		Convey("TestVideoDB_UpdateCommentCount_normal", func() {
			id, err := _db.CreateVideo(ctx, video)
			So(err, ShouldBeNil)
			getVideo, err := _db.QueryVideoById(ctx, strconv.FormatInt(id, 10))
			So(err, ShouldBeNil)
			So(getVideo.Uid, ShouldEqual, getVideo.Uid)
			err = _db.UpdateCommentCount(ctx, getVideo.Uid, 1)
			So(err, ShouldBeNil)
			getVideo2, err := _db.QueryVideoById(ctx, strconv.FormatInt(id, 10))
			So(err, ShouldBeNil)
			So(getVideo.CommentCount, ShouldEqual, getVideo2.CommentCount)
		})
	})
}
