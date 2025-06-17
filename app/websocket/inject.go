package websocket

import (
	"TikTok-rpc/app/websocket/controllers/rpc"
	"TikTok-rpc/app/websocket/domain/service"
	"TikTok-rpc/app/websocket/infrastructure/cache"
	"TikTok-rpc/app/websocket/infrastructure/mysql"
	"TikTok-rpc/app/websocket/usecase"
	"TikTok-rpc/kitex_gen/websocket"
	"TikTok-rpc/pkg/base/client"
	"TikTok-rpc/pkg/constants"

	"github.com/bytedance/gopkg/util/logger"
)

func InjectWebsocketHandler() websocket.WebsocketService {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}
	m, err := client.NewRedisClient(constants.RedisDBWebsocket)
	if err != nil {
		panic(err)
	}
	db := mysql.NewWebsocketDB(gormDB)
	ca := cache.NewWebsocketCache(m)
	if err != nil {
		logger.Errorf("Failed to init user rpc client: %v", err)
	}
	svc := service.NewWebSocketService(db, ca)
	uc := usecase.NewWebsocketUseCase(db, ca, svc)
	return rpc.NewWebsocketServiceImpl(uc)
}
