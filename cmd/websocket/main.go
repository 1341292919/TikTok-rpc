package main

import (
	"TikTok-rpc/app/websocket"
	"TikTok-rpc/config"
	"TikTok-rpc/kitex_gen/websocket/websocketservice"
	"TikTok-rpc/pkg/base"
	"TikTok-rpc/pkg/constants"
	"TikTok-rpc/pkg/pprof"
	"TikTok-rpc/pkg/utils"
	"context"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"log"
	"net"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var serviceName = constants.WebsocketServiceName

func init() {
	config.Init(serviceName)
}
func main() {
	pprof.Load(serviceName)
	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		logger.Fatalf("Websocket: new etcd registry failed, err: %v", err)
	}
	listenAddr, err := utils.GetAvailablePort()
	if err != nil {
		logger.Fatalf("Websocket: get available port failed, err: %v", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", listenAddr) // 服务监听端口
	if err != nil {
		logger.Fatalf("Websocketr: resolve tcp addr failed, err: %v", err)
	}

	p := base.TelemetryProvider(serviceName, config.Otel.CollectorAddr)
	defer func() {
		err := p.Shutdown(context.Background())
		logger.Fatalf("shutdown telemetry provider failed, err: %v", err)
	}()

	svr := websocketservice.NewServer(
		//只能注入一个handler
		websocket.InjectWebsocketHandler(),
		server.WithSuite(tracing.NewServerSuite()),
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
