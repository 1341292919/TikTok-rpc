package cache

import (
	"TikTok-rpc/app/video/domain/model"
	"math/rand"
	"testing"
)

func buildVideoId(t *testing.T) int64 {
	t.Helper()
	return int64(rand.Uint32())
}
func buildTestModelVideoList(t *testing.T) []*model.Video {
	t.Helper()
	videolist := make([]*model.Video, 10)
	for i := 0; i < 10; i++ {
		videolist[i] = &model.Video{
			Uid:         int64(rand.Uint32()),
			VideoUrl:    "test",
			CoverUrl:    "test",
			Title:       "test",
			Description: "test",
			VisitCount:  rand.Int63n(1000),
		}
	}
	return videolist
}
