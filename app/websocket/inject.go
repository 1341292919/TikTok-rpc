package websocket

import (
	"TikTok-rpc/app/websocket/controllers/rpc"
	"TikTok-rpc/app/websocket/domain/service"
	"TikTok-rpc/app/websocket/infrastructure/cache"
	"TikTok-rpc/app/websocket/infrastructure/mysql"
	"TikTok-rpc/app/websocket/usecase"
	"TikTok-rpc/kitex_gen/websocket"
	"github.com/bytedance/gopkg/util/logger"
)

func InjectWebsocketHandler() websocket.WebsocketService {
	gormDB, err := mysql.InitMySQL()
	if err != nil {
		panic(err)
	}
	m, err := cache.Init()
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
