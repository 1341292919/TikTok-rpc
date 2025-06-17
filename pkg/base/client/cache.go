package client

import (
	"TikTok-rpc/config"
	"TikTok-rpc/pkg/errno"
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(db int) (*redis.Client, error) {
	ctx := context.Background()
	if config.Redis == nil {
		return nil, errors.New("redis config is nil")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Username: config.Redis.Username,
		Password: config.Redis.Password,
		DB:       db,
	})
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, "Connect falied"+err.Error())
	}
	return client, nil
}
