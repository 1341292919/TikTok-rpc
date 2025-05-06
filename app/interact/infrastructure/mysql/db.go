package mysql

import (
	"TikTok-rpc/app/interact/domain/model"
	"TikTok-rpc/app/interact/domain/repository"
	"TikTok-rpc/pkg/constants"
	"TikTok-rpc/pkg/errno"
	"context"
	"errors"

	"gorm.io/gorm"
)

type interactDB struct {
	client *gorm.DB
}

func NewInteractDB(client *gorm.DB) repository.InteractDB {
	return &interactDB{client: client}
}

func (db *interactDB) IsCommentExist(ctx context.Context, id int64) (bool, error) {
	var c *Comment
	err := db.client.WithContext(ctx).
		Table(constants.TableComment).
		Where("id = ?", id).
		First(&c).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errno.NewErrNo(errno.InternalDatabaseErrorCode, err.Error())
	}
	return true, nil
}

func (db *interactDB) IsVideoLikeExist(ctx context.Context, id, uid int64) (bool, error) {
	return false, nil
}

func (db *interactDB) IsCommentLikeExist(ctx context.Context, id, uid int64) (bool, error) {
	return false, nil
}

func (db *interactDB) UpdateCommentLikeCount(ctx context.Context, cid, newcount int64) error {
	err := db.client.WithContext(ctx).
		Table(constants.TableComment).
		Where("id = ?", cid).
		Update("like_count", newcount).
		Error
	return err
}

func (db *interactDB) CreateNewUserLike(ctx context.Context, targetid, uid, t int64) error {
	var u *UserLike
	err := db.client.WithContext(ctx).
		Table(constants.TableUserLike).
		Where("user_id = ? and target_id = ? and type = ?", uid, targetid, t).
		First(&u).
		Error
	//如果找到了 应该返回空类型让更新继续下去 如果没找到要继续执行函数 如果其他错误则应该返回
	//考虑将这些检验放在外层
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if u.UserId != 0 { //真找到了
			return nil
		}
		if err != nil {
			return err
		}
	}
	u = &UserLike{
		UserId:   uid,
		Type:     t,
		TargetId: targetid,
	}
	err = db.client.WithContext(ctx).
		Table(constants.TableUserLike).
		Create(u).
		Error
	if err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "CreateNewUserLike:"+err.Error())
	}
	return nil
}
func (db *interactDB) DeleteUserLike(ctx context.Context, targetid, uid, t int64) error {
	var u *UserLike
	err := db.client.WithContext(ctx).
		Table(constants.TableUserLike).
		Where("user_id = ? and target_id = ? and type = ?", uid, targetid, t).
		Find(&u).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "DeleteUserLike:"+err.Error())
	}
	u = &UserLike{
		UserId:   uid,
		Type:     t,
		TargetId: targetid,
	}
	err = db.client.WithContext(ctx).
		Table(constants.TableUserLike).
		Where("user_id = ? and target_id = ? and type = ?", uid, targetid, t).
		Delete(&u).
		Error
	if err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "DeleteUserLike:"+err.Error())
	}
	return nil
}

func (db *interactDB) QueryAllUserLike(ctx context.Context) ([]*model.UserLike, error) {
	var userLikes []*UserLike
	err := db.client.WithContext(ctx).
		Table(constants.TableUserLike).
		Find(&userLikes).
		Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "QueryAllUserLike:"+err.Error())
	}
	return buildUserLikeList(userLikes), nil
}
func (db *interactDB) QueryCommentLikeCount(ctx context.Context) ([]*model.LikeCount, error) {
	var commentData []*Comment
	err := db.client.WithContext(ctx).
		Table(constants.TableComment).
		Find(&commentData).
		Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "QueryAllUserLike:"+err.Error())
	}
	return buildLikeCountList(commentData), nil
}

func (db *interactDB) QueryUserLikeByUid(ctx context.Context, uid int64) ([]*model.UserLike, error) {
	var userLikes []*UserLike
	err := db.client.WithContext(ctx).
		Table(constants.TableUserLike).
		Where("user_id = ?", uid).
		Find(&userLikes).
		Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "QueryUserLikeByUid:"+err.Error())
	}
	return buildUserLikeList(userLikes), nil
}

func (db *interactDB) CreateNewComment(ctx context.Context, req *model.InteractReq) (int64, error) {
	var commentData *Comment
	if req.Type == 0 {
		commentData = &Comment{
			UserId:   req.Uid,
			ParentId: req.VideoId,
			Content:  req.Content,
			Type:     req.Type,
		}
	} else if req.Type == 1 {
		commentData = &Comment{
			UserId:   req.Uid,
			ParentId: req.CommentId,
			Content:  req.Content,
			Type:     req.Type,
		}
	}
	err := db.client.WithContext(ctx).
		Table(constants.TableComment).
		Create(&commentData).
		Error
	if err != nil {
		return -1, errno.NewErrNo(errno.InternalDatabaseErrorCode, "CreateNewComment:"+err.Error())
	}
	return commentData.Id, nil
}

func (db *interactDB) DeleteComment(ctx context.Context, req *model.InteractReq) (*model.Comment, error) {
	var commentData *Comment

	commentData = &Comment{
		ParentId: req.CommentId,
		Type:     req.Type,
	}

	err := db.client.WithContext(ctx).
		Table(constants.TableComment).
		Where("id = ?", req.CommentId).
		First(&commentData).
		Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "DeleteComment:"+err.Error())
	}
	if commentData.UserId != req.Uid {
		return nil, errno.NewErrNo(errno.ServiceNoAuthority, "Can not delete other user`s comment")
	}
	err = db.client.WithContext(ctx).
		Table(constants.TableComment).
		Where("id = ?", req.CommentId).
		Delete(&commentData).
		Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "DeleteComment:"+err.Error())
	}
	return &model.Comment{
		Id:       commentData.Id,
		Uid:      commentData.UserId,
		ParentId: commentData.ParentId,
		Type:     commentData.Type,
	}, nil
}

func (db *interactDB) UpdateCommentCount(ctx context.Context, commentid, change int64) error {
	err := db.client.WithContext(ctx).
		Table(constants.TableComment).
		Where("id = ?", commentid).
		Update("child_count", gorm.Expr("child_count + ?", change)).
		Error
	if err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "UpdateCommentCount:"+err.Error())
	}
	return nil
}

func (db *interactDB) QueryCommentByParentId(ctx context.Context, req *model.InteractReq) ([]*model.Comment, error) {
	var err error
	var commentData []*Comment
	var count int64
	var id int64
	//这个感觉要放在外面svc
	if req.Type == 0 {
		id = req.VideoId
	} else {
		id = req.CommentId
	}
	err = db.client.WithContext(ctx).
		Table(constants.TableComment).
		Where("parent_id = ? and type = ?", id, req.Type).
		Limit(int(req.PageSize)).
		Offset(int((req.PageNum - 1) * req.PageSize)).
		Count(&count).
		Find(&commentData).
		Error
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, "QueryCommentByParentId:"+err.Error())
	}
	return buildCommentList(commentData), nil
}

func buildUserLike(data *UserLike) *model.UserLike {
	return &model.UserLike{
		Uid:       data.UserId,
		VideoId:   data.TargetId,
		CommentId: data.TargetId,
		Type:      data.Type,
	}
}
func buildUserLikeList(data []*UserLike) []*model.UserLike {
	var list []*model.UserLike
	for _, item := range data {
		list = append(list, buildUserLike(item))
	}
	return list
}

func buildComment(data *Comment) *model.Comment {
	return &model.Comment{
		Id:         data.Id,
		ParentId:   data.ParentId,
		Type:       data.Type,
		Content:    data.Content,
		ChildCount: data.ChildCount,
		CreateAT:   data.CreatedAt.Unix(),
		UpdateAT:   data.UpdatedAt.Unix(),
		LikeCount:  data.LikeCount,
		Uid:        data.UserId,
	}
}

func buildCommentList(data []*Comment) []*model.Comment {
	commenlist := make([]*model.Comment, 0)
	for _, item := range data {
		commenlist = append(commenlist, buildComment(item))
	}
	return commenlist
}

func buildLikeCount(data *Comment) *model.LikeCount {
	return &model.LikeCount{
		Count: data.LikeCount,
		Id:    data.Id,
		Type:  1,
	}
}
func buildLikeCountList(data []*Comment) []*model.LikeCount {
	list := make([]*model.LikeCount, 0)
	for _, item := range data {
		list = append(list, buildLikeCount(item))
	}
	return list
}
