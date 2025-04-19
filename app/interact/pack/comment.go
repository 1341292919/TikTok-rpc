package pack

import (
	rpc "TikTok-rpc/app/interact/domain/model"
	"TikTok-rpc/kitex_gen/model"
	"strconv"
)

func BuildComment(data *rpc.Comment) *model.Comment {
	return &model.Comment{
		UserId:     strconv.FormatInt(data.Uid, 10),
		Id:         strconv.FormatInt(data.Id, 10),
		ParentId:   strconv.FormatInt(data.ParentId, 10),
		Content:    data.Content,
		CreatedAt:  strconv.FormatInt(data.CreateAT, 10),
		UpdatedAt:  strconv.FormatInt(data.UpdateAT, 10),
		LikeCount:  data.LikeCount,
		ChildCount: data.ChildCount,
	}
}
func BuildCommentList(data []*rpc.Comment, count int64) *model.CommentList {
	commentList := make([]*model.Comment, 0)
	for _, v := range data {
		commentList = append(commentList, BuildComment(v))
	}
	return &model.CommentList{
		Items: commentList,
		Total: count,
	}
}
