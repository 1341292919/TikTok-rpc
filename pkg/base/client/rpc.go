package client

import (
	"TikTok-rpc/config"
	"TikTok-rpc/kitex_gen/interact/interactservice"
	"TikTok-rpc/kitex_gen/socialize/socializeservice"
	"TikTok-rpc/kitex_gen/user/userservice"
	"TikTok-rpc/kitex_gen/video/videoservice"
	"TikTok-rpc/kitex_gen/websocket/websocketservice"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"

	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func initRpcClient[T any](serviceName string, newClientFunc func(string, ...client.Option) (T, error),
) (*T, error) {
	if config.Etcd == nil || config.Etcd.Addr == "" {
		return nil, errors.New("config.Etcd.Addr is nil")
	}
	// 初始化Etcd Resolver
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})
	if err != nil {
		return nil, fmt.Errorf("initRPCClient etcd.NewEtcdResolver failed: %w", err)
	}
	client, err := newClientFunc(serviceName, client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: fmt.Sprintf("%s-client", serviceName)}),
		client.WithSuite(tracing.NewClientSuite()),
	)
	if err != nil {
		return nil, fmt.Errorf("init RPC client failed: %w", err)
	}
	return &client, nil
}

func InitUserRPC() (*userservice.Client, error) { return initRpcClient("user", userservice.NewClient) }
func InitVideoRPC() (*videoservice.Client, error) {
	return initRpcClient("video", videoservice.NewClient)
}
func InitInteractRPC() (*interactservice.Client, error) {
	return initRpcClient("interact", interactservice.NewClient)
}
func InitSocializeRPC() (*socializeservice.Client, error) {
	return initRpcClient("socialize", socializeservice.NewClient)
}
func InitWebsocketRPC() (*websocketservice.Client, error) {
	return initRpcClient("websocket", websocketservice.NewClient)
}
