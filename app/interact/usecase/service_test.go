package usecase

import (
	"TikTok-rpc/app/interact/domain/model"
	"TikTok-rpc/app/interact/domain/service"
	"TikTok-rpc/app/interact/infrastructure/cache"
	"TikTok-rpc/app/interact/usecase/mocks"
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUseCase_CommentVideo(t *testing.T) {
	mockDB := new(mocks.InteractDB)
	mockCache := new(redis.Client)
	mockC := new(redis.Client)
	c := cache.NewInteractCache(mockCache, mockC)
	mockRpc := new(mocks.RpcPort)
	mockService := service.NewInteractService(mockDB, c, mockRpc)
	uc := NewInteractUseCase(mockDB, mockService, c, mockRpc)

	interactReq := &model.InteractReq{
		Uid:     int64(1),
		Content: "test",
		VideoId: int64(1),
		Type:    0,
	}
	mockRpc.On("IsVideoExist", mock.Anything, interactReq.VideoId).Return(true, nil)
	mockRpc.On("UpdateVideoCommentCount", mock.Anything, interactReq.VideoId, int64(1)).Return(nil)
	mockDB.On("CreateNewComment", mock.Anything, interactReq).Return(int64(1), nil)
	cid, err := uc.Comment(context.Background(), interactReq)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), cid)
	mockDB.AssertExpectations(t)
}
func TestUseCase_CommentComment(t *testing.T) {
	mockDB := new(mocks.InteractDB)
	mockCache := new(redis.Client)
	mockC := new(redis.Client)
	c := cache.NewInteractCache(mockCache, mockC)
	mockRpc := new(mocks.RpcPort)
	mockService := service.NewInteractService(mockDB, c, mockRpc)
	uc := NewInteractUseCase(mockDB, mockService, c, mockRpc)

	interactReq := &model.InteractReq{
		Uid:       int64(1),
		Content:   "test",
		CommentId: int64(1),
		Type:      1,
	}
	mockDB.On("CreateNewComment", mock.Anything, interactReq).Return(int64(1), nil)
	mockDB.On("IsCommentExist", mock.Anything, interactReq.CommentId).Return(true, nil)
	mockDB.On("UpdateCommentCount", mock.Anything, interactReq.CommentId, int64(1)).Return(nil)
	cid, err := uc.Comment(context.Background(), interactReq)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), cid)
	mockDB.AssertExpectations(t)
}
func TestUseCase_DeleteVideoComment(t *testing.T) {
	mockDB := new(mocks.InteractDB)
	mockCache := new(redis.Client)
	mockC := new(redis.Client)
	c := cache.NewInteractCache(mockCache, mockC)
	mockRpc := new(mocks.RpcPort)
	mockService := service.NewInteractService(mockDB, c, mockRpc)
	uc := NewInteractUseCase(mockDB, mockService, c, mockRpc)
	interactReq := &model.InteractReq{
		CommentId: int64(1),
	}
	comment := &model.Comment{
		ParentId: int64(1),
		Type:     int64(0),
	}
	mockDB.On("IsCommentExist", mock.Anything, interactReq.CommentId).Return(true, nil)
	mockDB.On("DeleteComment", mock.Anything, interactReq).Return(comment, nil)
	mockRpc.On("UpdateVideoCommentCount", mock.Anything, comment.ParentId, int64(-1)).Return(nil)
	err := uc.DeleteComment(context.Background(), interactReq)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}
func TestUseCase_DeleteComment(t *testing.T) {
	mockDB := new(mocks.InteractDB)
	mockCache := new(redis.Client)
	mockC := new(redis.Client)
	c := cache.NewInteractCache(mockCache, mockC)
	mockRpc := new(mocks.RpcPort)
	mockService := service.NewInteractService(mockDB, c, mockRpc)
	uc := NewInteractUseCase(mockDB, mockService, c, mockRpc)
	interactReq := &model.InteractReq{
		CommentId: int64(1),
	}
	comment := &model.Comment{
		ParentId: int64(2),
		Type:     int64(1),
	}
	mockDB.On("IsCommentExist", mock.Anything, interactReq.CommentId).Return(true, nil)
	mockDB.On("DeleteComment", mock.Anything, interactReq).Return(comment, nil)
	mockDB.On("UpdateCommentCount", mock.Anything, comment.ParentId, int64(-1)).Return(nil)
	err := uc.DeleteComment(context.Background(), interactReq)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}
