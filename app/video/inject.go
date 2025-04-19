package video

import (
	"TikTok-rpc/app/video/controllers/rpc"
	"TikTok-rpc/app/video/domain/service"
	"TikTok-rpc/app/video/infrastructure/cache"
	"TikTok-rpc/app/video/infrastructure/mysql"
	rpccli "TikTok-rpc/app/video/infrastructure/rpc"
	"TikTok-rpc/app/video/usecase"
	"TikTok-rpc/kitex_gen/video"
	"TikTok-rpc/pkg/base/client"
	"github.com/bytedance/gopkg/util/logger"
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
	uClient, err := client.InitUserRPC()
	if err != nil {
		logger.Errorf("Failed to init user rpc client: %v", err)
	}
	rpcImpl := rpccli.NewVideoRpcImpl(*uClient)
	svc := service.NewVideoService(db, ca, rpcImpl)
	uc := usecase.NewVideoUseCase(db, svc, ca, rpcImpl)
	return rpc.NewVideoServiceImpl(uc)
}
