package usecase

import (
	"TikTok-rpc/app/interact/domain/model"
	"TikTok-rpc/app/interact/domain/repository"
	"TikTok-rpc/app/interact/domain/service"
	"context"
)

type InteractUseCase interface {
	Like(ctx context.Context, req *model.InteractReq) error
	QueryLikeList(ctx context.Context, req *model.InteractReq) ([]*model.Video, int64, error)
	Comment(ctx context.Context, req *model.InteractReq) error
	DeleteComment(ctx context.Context, req *model.InteractReq) error
	QueryCommentList(ctx context.Context, req *model.InteractReq) ([]*model.Comment, int64, error)
}
type useCase struct {
	db    repository.InteractDB
	cache repository.InteractCache
	svc   *service.InteractService
	Rpc   repository.RpcPort
	MQ    repository.MqPort
}

func NewInteractUseCase(db repository.InteractDB, svc *service.InteractService, cache repository.InteractCache, rpc repository.RpcPort, mq repository.MqPort) *useCase {
	return &useCase{
		db:    db,
		cache: cache,
		svc:   svc,
		Rpc:   rpc,
		MQ:    mq,
	}
}
