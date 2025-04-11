package pack

import (
	"TikTok-rpc/app/gateway/model/model"
	rpc "TikTok-rpc/kitex_gen/model"
	"strconv"
	"time"
)

func Video(data *rpc.Video) *model.Video {
	return &model.Video{
		ID:           data.Id,
		UserID:       data.UserId,
		VideoURL:     data.VideoUrl,
		CoverURL:     data.CoverUrl,
		Title:        data.Title,
		VisitCount:   data.VisitCount,
		LikeCount:    data.LikeCount,
		CommentCount: data.CommentCount,
		CreatedAt:    ChangeFormat(data.CreatedAt),
		UpdatedAt:    ChangeFormat(data.UpdatedAt),
	}
}
func ChangeFormat(timeStr string) string {
	timestamp, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return ""
	}
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
func VideoList(data *rpc.VideoList) *model.VideoList {
	videoList := make([]*model.Video, 0)
	for _, v := range data.Items {
		videoList = append(videoList, Video(v)) // 正确使用 append
	}
	return &model.VideoList{
		Items: videoList,
		Total: data.Total,
	}
}
