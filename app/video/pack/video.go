package pack

import (
	rpc "TikTok-rpc/app/video/domain/model"
	"TikTok-rpc/kitex_gen/model"
	"strconv"
)

func BuildVideo(data *rpc.Video) *model.Video {
	return &model.Video{
		UserId:      strconv.FormatInt(data.Uid, 10),
		Id:          strconv.FormatInt(data.Id, 10),
		VideoUrl:    data.VideoUrl,
		CoverUrl:    data.CoverUrl,
		CreatedAt:   strconv.FormatInt(data.CreateAT, 10),
		UpdatedAt:   strconv.FormatInt(data.UpdateAT, 10),
		Title:       data.Title,
		Description: data.Description,
		VisitCount:  data.VisitCount,
	}
}

func BuildVideoList(data []*rpc.Video, count int64) *model.VideoList {
	videoList := make([]*model.Video, 0)
	for _, v := range data {
		videoList = append(videoList, BuildVideo(v))
	}
	return &model.VideoList{
		Items: videoList,
		Total: count,
	}
}
func BuildLikeCount(data *rpc.LikeCount) *model.LikeCount {
	return &model.LikeCount{
		Count:   data.Count,
		VideoId: data.Id,
	}
}
func BuildLikeCountList(data []*rpc.LikeCount) *model.LikeCountList {
	likeCountList := make([]*model.LikeCount, 0)
	for _, v := range data {
		likeCountList = append(likeCountList, BuildLikeCount(v))
	}
	return &model.LikeCountList{
		Items: likeCountList,
		Total: int64(len(data)),
	}
}
