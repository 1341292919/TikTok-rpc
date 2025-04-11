package service

import (
	"TikTok-rpc/app/video/domain/model"
	"context"
	"time"
)

func (svc *VideoService) CreateVideo(ctx context.Context, video *model.Video) (int64, error) {
	return svc.db.CreateVideo(ctx, video)
}

func (svc *VideoService) QueryPublishList(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error) {
	return svc.db.QueryVideoByUid(ctx, req)
}
func (svc *VideoService) SearchVideoByKeyWord(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error) {
	//检查参数 如果没有传入两个日期就把两个日期设置为 0~现在的日期
	//这里又要检验一次参数？有没有好的解决办法
	if req.ToDate == 0 && req.FromDate == 0 {
		req.ToDate = time.Now().Unix()
		req.FromDate = 0
	}
	return svc.db.QueryVideoByKeyWord(ctx, req)
}

func (svc *VideoService) QueryPopularVideoList(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error) {
	//cache层返回的错误是errno包装的，但实际上最好是让基建层不会出现业务错误，还需要改进！
	var count int64
	data, err := svc.cache.GetVideoByRank(ctx, 100)
	if err != nil {
		return nil, -1, err
	}
	if len(data) == 0 {
		//从缓存中获取VID-从MySQL中获取视频信息-构建redis视频列表
		vId, err := svc.cache.GetVideoIdByRank(ctx, 100)
		//当vId也没有信息时
		if len(vId) == 0 {
			data, count, err = svc.db.QueryPopularVideo(ctx, req)
			if err != nil {
				return nil, -1, err
			}
			for _, v := range data {
				err = svc.cache.NewIdToRank(ctx, v.Id)
				if err != nil {
					return nil, -1, err
				}
			}
		} else {
			if err != nil {
				return nil, -1, err
			}
			data, err = svc.db.QueryVideoListById(ctx, vId)
			if err != nil {
				return nil, -1, err
			}
			err = svc.cache.AddVideoToRank(ctx, data)
			if err != nil {
				return nil, -1, err
			}
		}
	}
	//按页分好
	startIndex := (req.PageNum - 1) * req.PageSize
	endIndex := startIndex + req.PageSize

	count = int64(len(data))
	if startIndex >= count {
		return nil, 0, nil
	}

	if endIndex > count {
		endIndex = count
	}
	return data[startIndex:endIndex], count, nil

}

func (svc *VideoService) VideoStream(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error) {
	//暂时没有什么好的想法-要处理last-time时间逻辑
	return svc.db.QueryVideoByKeyWord(ctx, req)
}
