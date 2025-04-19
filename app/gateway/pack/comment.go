package pack

import (
	"TikTok-rpc/app/gateway/model/model"
	rpc "TikTok-rpc/kitex_gen/model"
)

func Comment(data *rpc.Comment) *model.Comment {
	return &model.Comment{
		ID:         data.Id,
		UserID:     data.UserId,
		VideoID:    data.VideoId,
		ParentID:   data.ParentId,
		Content:    data.Content,
		CreatedAt:  ChangeFormat(data.CreatedAt),
		UpdatedAt:  ChangeFormat(data.UpdatedAt),
		ChildCount: data.ChildCount,
		LikeCount:  data.LikeCount,
	}
}

func CommentList(data *rpc.CommentList) *model.CommentList {
	commentlist := make([]*model.Comment, 0)
	for _, v := range data.Items {
		commentlist = append(commentlist, Comment(v))
	}
	return &model.CommentList{
		Items: commentlist,
		Total: data.Total,
	}
}
