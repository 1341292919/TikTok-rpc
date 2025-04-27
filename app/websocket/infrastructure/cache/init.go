package cache

import (
	"TikTok-rpc/pkg/constants"
	"TikTok-rpc/pkg/errno"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
)

func Init() (*redis.Client, error) {
	ctx := context.Background()

	redisD := redis.NewClient(&redis.Options{
		Addr:     constants.RedisHost + ":" + constants.RedisPort,
		Username: constants.RedisUserName,
		Password: constants.RedisPassWord,
		DB:       1,
	})

	if _, err := redisD.Ping(ctx).Result(); err != nil {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, "videoId Init falied"+err.Error())
	}

	hlog.Info("Redis连接成功")

	return redisD, nil
}
