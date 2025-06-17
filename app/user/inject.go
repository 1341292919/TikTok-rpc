package user

import (
	"TikTok-rpc/app/user/controllers/rpc"
	"TikTok-rpc/app/user/domain/service"
	"TikTok-rpc/app/user/infrastructure/mysql"
	"TikTok-rpc/app/user/usecase"
	"TikTok-rpc/kitex_gen/user"
	"TikTok-rpc/pkg/base/client"
)

func InjectUserHandler() user.UserService {
	gormDB, err := client.InitMySQL()
	if err != nil {
		panic(err)
	}
	db := mysql.NewUserDB(gormDB)
	svc := service.NewUserService(db)
	uc := usecase.NewUserUseCase(db, svc)
	return rpc.NewUserServiceImpl(uc)
}
