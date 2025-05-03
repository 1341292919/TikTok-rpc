package interact

import (
	"TikTok-rpc/app/interact/controllers/rpc"
	"TikTok-rpc/app/interact/domain/service"
	"TikTok-rpc/app/interact/infrastructure/cache"
	"TikTok-rpc/app/interact/infrastructure/mysql"
	rpccli "TikTok-rpc/app/interact/infrastructure/rpc"
	"TikTok-rpc/app/interact/usecase"
	"TikTok-rpc/kitex_gen/interact"
	"TikTok-rpc/pkg/base/client"
	"TikTok-rpc/pkg/constants"
	"github.com/bytedance/gopkg/util/logger"
)

func InjectInteractHandler() interact.InteractService {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}
	Ulike, err := client.NewRedisClient(constants.RedisDBInteract)
	if err != nil {
		panic(err)
	}
	Lcount, err := client.NewRedisClient(constants.RedisDBInteract)
	db := mysql.NewInteractDB(gormDB)
	ca := cache.NewInteractCache(Ulike, Lcount)
	vClient, err := client.InitVideoRPC()
	if err != nil {
		logger.Errorf("Failed to init video rpc client: %v", err)
	}
	uClient, err := client.InitUserRPC()
	if err != nil {
		logger.Errorf("Failed to init user rpc client: %v", err)
	}
	rpcImpl := rpccli.NewInteractRpcImpl(*vClient, *uClient)
	svc := service.NewInteractService(db, ca, rpcImpl)
	serviceAdapter := usecase.NewInteractUseCase(db, svc, ca, rpcImpl)
	return rpc.NewInteractServiceImpl(serviceAdapter)
}
