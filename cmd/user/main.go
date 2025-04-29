package main

import (
	"TikTok-rpc/app/user"
	"TikTok-rpc/config"
	"TikTok-rpc/kitex_gen/user/userservice"
	"TikTok-rpc/pkg/constants"
	"TikTok-rpc/pkg/utils"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

var serviceName = constants.UserServiceName

func init() {
	config.Init(serviceName)
}

func main() {
	// 应该把etcd 以及可用的端口号 在代码中也设置管理和调配
	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		logger.Fatalf("User: new etcd registry failed, err: %v", err)
	}

	listenAddr, err := utils.GetAvailablePort()
	if err != nil {
		logger.Fatalf("User: get available port failed, err: %v", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", listenAddr) // 服务监听端口
	if err != nil {
		logger.Fatalf("User: resolve tcp addr failed, err: %v", err)
	}
	svr := userservice.NewServer(
		//只能注入一个handler
		user.InjectUserHandler(),
		server.WithServiceAddr(addr),
		server.WithRegistry(r), // 关键：注册到 ETCD
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: serviceName, // 关键：设置服务名称
		}),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
