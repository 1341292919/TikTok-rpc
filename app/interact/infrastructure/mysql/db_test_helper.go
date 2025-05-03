package mysql

import (
	"TikTok-rpc/app/interact/domain/model"
	"math/rand"
	"testing"
)

func buildTestModelComment(t *testing.T) *model.Comment {
	t.Helper()
	return &model.Comment{
		Uid:      int64(1),
		ParentId: int64(2),
		Content:  "test comment",
	}
}
func buildTestModelInteractReq(t *testing.T) *model.InteractReq {
	t.Helper()
	return &model.InteractReq{
		Uid:       int64(1),
		VideoId:   int64(2),
		CommentId: int64(3),
		Content:   "test interact request",
		Type:      int64(rand.Intn(100) % 2),
		PageNum:   1,
		PageSize:  10,
	}
}

func buildTestModelUserLike(t *testing.T) *model.UserLike {
	t.Helper()
	return &model.UserLike{
		Uid:       int64(10000),
		Status:    1,
		Type:      int64(rand.Intn(100) % 2),
		CommentId: int64(3),
		VideoId:   int64(2),
	}
}
