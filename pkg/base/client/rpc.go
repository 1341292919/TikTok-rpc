package client

import (
	"TikTok-rpc/kitex_gen/interact/interactservice"
	"TikTok-rpc/kitex_gen/socialize/socializeservice"
	"TikTok-rpc/kitex_gen/user/userservice"
	"TikTok-rpc/kitex_gen/video/videoservice"
	"TikTok-rpc/kitex_gen/websocket/websocketservice"
	"TikTok-rpc/pkg/constants"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var serviceToETCD = map[string][]string{
	"user":      []string{constants.UserETCD},
	"video":     []string{constants.VideoETCD},
	"interact":  []string{constants.InteractETCD},
	"socialize": []string{constants.SocializeETCD},
	"websocket": []string{constants.WebsocketETCD},
}

func initRpcClient[T any](serviceName string, newClientFunc func(string, ...client.Option) (T, error),
) (*T, error) {
	etcdAddrs, ok := serviceToETCD[serviceName]
	if !ok {
		return nil, fmt.Errorf("no ETCD address configured for service %s", serviceName)
	}
	r, err := etcd.NewEtcdResolver(etcdAddrs)
	if err != nil {
		return nil, fmt.Errorf("init ETCD resolver failed: %w", err)
	}
	client, err := newClientFunc(serviceName, client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: fmt.Sprintf("%s-client", serviceName)}),
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
