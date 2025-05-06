package main

import (
	"TikTok-rpc/app/interact"
	"TikTok-rpc/app/interact/domain/service"
	"TikTok-rpc/config"
	interactservice "TikTok-rpc/kitex_gen/interact/interactservice"
	"TikTok-rpc/pkg/constants"
	"TikTok-rpc/pkg/utils"
	"log"
	"net"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var serviceName = constants.InteractServiceName

func init() {
	config.Init(serviceName)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		logger.Fatalf("Interact: new etcd registry failed, err: %v", err)
	}
	listenAddr, err := utils.GetAvailablePort()
	if err != nil {
		logger.Fatalf("Interact: get available port failed, err: %v", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", listenAddr) // 服务监听端口
	if err != nil {
		logger.Fatalf("Interact: resolve tcp addr failed, err: %v", err)
	}
	svr := interactservice.NewServer(
		//只能注入一个handler
		interact.InjectInteractHandler(),
		server.WithServiceAddr(addr),
		server.WithRegistry(r), // 关键：注册到 ETCD
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: serviceName, // 关键：设置服务名称
		}),
	)
	go service.SyncDB()
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
