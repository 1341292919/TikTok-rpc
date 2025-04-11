package main

import (
	"TikTok-rpc/app/video"
	"TikTok-rpc/kitex_gen/video/videoservice"
	"TikTok-rpc/pkg/constants"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

var serviceName = constants.VideoServiceName

func main() {

	r, err := etcd.NewEtcdRegistry([]string{constants.VideoETCD})
	if err != nil {
		logger.Fatalf("Video: new etcd registry failed, err: %v", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:9998") // 服务监听端口
	if err != nil {
		logger.Fatalf("User: resolve tcp addr failed, err: %v", err)
	}
	svr := videoservice.NewServer(
		//只能注入一个handler
		video.InjectVideoHandler(),
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
