package cache

import (
	"TikTok-rpc/app/video/domain/repository"
	"TikTok-rpc/config"
	"TikTok-rpc/pkg/base/client"
	"TikTok-rpc/pkg/constants"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	. "github.com/smartystreets/goconvey/convey"
	"sort"
	"strconv"
	"testing"
)

func initTest(t *testing.T) repository.VideoCache {
	t.Helper()
	config.Init("video-test")
	v, err := client.NewRedisClient(constants.RedisDBVideo)
	if err != nil {
		panic(err)
	}
	vid, err := client.NewRedisClient(constants.RedisDBVideo)
	if err != nil {
		panic(err)
	}
	return NewVideoCache(v, vid)
}
func TestVideoCache_NewIdToRank(t *testing.T) {
	ca := initTest(t)
	ctx := context.Background()
	vid := buildVideoId(t)
	Convey("TestVideoCache_NewIdToRank", t, func() {
		Convey("TestVideoCache_NewIdToRank_normal", func() {
			err := ca.NewIdToRank(ctx, vid)
			So(err, ShouldBeNil)
			idList, err := ca.GetVideoIdByRank(ctx, 100)
			So(err, ShouldBeNil)
			So(contains(idList, vid), ShouldBeTrue)
		})
		Convey("TestVideoCache_NewIdToRank_delete", func() {
			err := ca.DeleteVideoIdRank(ctx)
			So(err, ShouldBeNil)
			idList, err := ca.GetVideoIdByRank(ctx, 100)
			So(err, ShouldBeNil)
			So(contains(idList, vid), ShouldBeFalse)
		})
	})
}
func TestVideoCache_AddVideoToRank(t *testing.T) {
	ca := initTest(t)
	ctx := context.Background()
	vList := buildTestModelVideoList(t)
	Convey("TestVideoCache_AddVideoToRank", t, func() {
		Convey("TestVideoCache_AddVideoToRank_normal", func() {
			err := ca.DeleteVideoRank(ctx)
			So(err, ShouldBeNil)
			err = ca.AddVideoToRank(ctx, vList)
			So(err, ShouldBeNil)
			sort.Slice(vList, func(i, j int) bool {
				return vList[i].VisitCount > vList[j].VisitCount
			})
			getList, err := ca.GetVideoByRank(ctx, 100)
			So(err, ShouldBeNil)
			for i, v := range vList {
				So(getList[i].VisitCount, ShouldEqual, v.VisitCount)
				So(getList[i].Id, ShouldEqual, v.Id)
			}
		})
		Convey("TestVideoCache_AddVideoToRank_delete", func() {
			err := ca.DeleteVideoRank(ctx)
			So(err, ShouldBeNil)
			getList, err := ca.GetVideoByRank(ctx, 100)
			So(err, ShouldBeNil)
			So(len(getList), ShouldEqual, 0)
		})
	})
}
func contains(slice []string, target int64) bool {
	hlog.Info(len(slice), target)
	targetid := strconv.FormatInt(target, 10)
	for _, item := range slice {
		if item == targetid {
			return true
		}
	}
	return false
}
