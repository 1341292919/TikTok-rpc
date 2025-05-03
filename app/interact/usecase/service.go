package usecase

import (
	"TikTok-rpc/app/interact/domain/model"
	"TikTok-rpc/pkg/errno"
	"context"
	"fmt"
)

// video_id、commenrt_id 必须存在其一（其中一个可以） 检验其是否存在的逻辑较复杂 交给svc
// 1.检验参数-点赞操作，2.video_id、comment_id存在与否2.创建喜欢关系3.表进行更新
func (u *useCase) Like(ctx context.Context, req *model.InteractReq) error {
	err := u.svc.Verify(u.svc.VerifyActionType(req.ActionType))
	if err != nil {
		return err
	}
	//id存在检验
	err = u.svc.IsIdOk(ctx, req.VideoId, req.CommentId)
	if err != nil {
		return err
	}
	//创建喜欢关系-逻辑有点复杂
	err = u.svc.NewLike(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (u *useCase) QueryLikeList(ctx context.Context, req *model.InteractReq) ([]*model.Video, int64, error) {
	err := u.svc.Verify(u.svc.VerifyPageParam(req.PageNum, req.PageSize))
	if err != nil {
		return nil, -1, err
	}
	exist, err := u.Rpc.IsUserExist(ctx, req.Uid)
	if err != nil {
		return nil, -1, err
	}
	if !exist {
		return nil, -1, errno.NewErrNo(errno.ServiceUserNotExistCode, "uid not exist")
	}
	data, err := u.svc.QueryLikeList(ctx, req)
	count := int64(len(data))
	return data, count, nil
}

func (u *useCase) Comment(ctx context.Context, req *model.InteractReq) (int64, error) {
	err := u.svc.IsIdOk(ctx, req.VideoId, req.CommentId)
	if err != nil {
		return -1, err
	}
	id, err := u.svc.Comment(ctx, req)
	return id, err
}

func (u *useCase) DeleteComment(ctx context.Context, req *model.InteractReq) error {
	exist, err := u.db.IsCommentExist(ctx, req.CommentId)
	if err != nil {
		return fmt.Errorf("check comment exist failed: %w", err)
	}
	if !exist {
		return errno.NewErrNo(errno.ServiceCommentNotExist, "comment not exist")
	}
	err = u.svc.DeleteComment(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (u *useCase) QueryCommentList(ctx context.Context, req *model.InteractReq) ([]*model.Comment, int64, error) {
	err := u.svc.Verify(u.svc.VerifyPageParam(req.PageNum, req.PageSize))
	if err != nil {
		return nil, -1, err
	}
	err = u.svc.IsIdOk(ctx, req.VideoId, req.CommentId)
	if err != nil {
		return nil, -1, err
	}
	data, count, err := u.svc.QueryCommentList(ctx, req)
	if err != nil {
		return nil, -1, err
	}
	return data, count, nil
}
