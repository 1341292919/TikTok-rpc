package video

import (
	"TikTok-rpc/app/video/controllers/rpc"
	"TikTok-rpc/app/video/domain/service"
	"TikTok-rpc/app/video/infrastructure/cache"
	"TikTok-rpc/app/video/infrastructure/mysql"
	"TikTok-rpc/app/video/usecase"
	"TikTok-rpc/kitex_gen/video"
)

func InjectVideoHandler() video.VideoService {
	gormDB, err := mysql.InitMySQL()
	if err != nil {
		panic(err)
	}
	videoRd, videoIdRd, err := cache.Init()
	if err != nil {
		panic(err)
	}
	db := mysql.NewVideoDB(gormDB)
	ca := cache.NewVideoCache(videoRd, videoIdRd)
	svc := service.NewVideoService(db, ca)
	uc := usecase.NewVideoUseCase(db, svc, ca)
	return rpc.NewVideoServiceImpl(uc)
}
