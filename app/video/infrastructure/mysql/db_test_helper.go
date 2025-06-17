package mysql

import (
	"TikTok-rpc/app/video/domain/model"
	"testing"
	"time"
)

func buildTestModelVideo(t *testing.T) *model.Video {
	t.Helper()
	return &model.Video{
		Uid:         1,
		VideoUrl:    "test",
		CoverUrl:    "test",
		Title:       "test",
		Description: "test",
	}
}

func buildTestModelVideoReq(t *testing.T) *model.VideoReq {
	t.Helper()
	return &model.VideoReq{
		Keyword:  "test",
		PageNum:  1,
		PageSize: 10,
		FromDate: 0,
		ToDate:   time.Now().Unix(),
		Username: "test",
	}
}
