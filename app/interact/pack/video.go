package pack

import (
	rpc "TikTok-rpc/app/interact/domain/model"
	"TikTok-rpc/kitex_gen/model"
)

func BuildVideo(data *rpc.Video) *model.Video {
	return &model.Video{
		UserId:      data.Uid,
		Id:          data.Id,
		VideoUrl:    data.VideoUrl,
		CoverUrl:    data.CoverUrl,
		CreatedAt:   data.CreateAT,
		UpdatedAt:   data.UpdateAT,
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
