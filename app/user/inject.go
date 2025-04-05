package user

import (
	"TikTok-rpc/app/user/controllers/rpc"
	"TikTok-rpc/app/user/infrastructure/mysql"
	"TikTok-rpc/app/user/usecase"
	"TikTok-rpc/kitex_gen/user"
)

func InjectUserHandler() user.UserService {
	gormDB, err := mysql.InitMySQL()
	if err != nil {
		panic(err)
	}
	db := mysql.NewUserDB(gormDB)
	uc := usecase.NewUserCase(db)
	return rpc.NewUserServiceImpl(uc)
}
