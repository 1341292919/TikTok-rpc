package cache

import (
	"TikTok-rpc/pkg/constants"
	"TikTok-rpc/pkg/errno"
	"context"
	"github.com/redis/go-redis/v9"
)

func InitCache() (*redis.Client, *redis.Client, error) {
	ctx := context.Background()

	redisCommentLike := redis.NewClient(&redis.Options{
		Addr:     constants.RedisHost + ":" + constants.RedisPort,
		Username: constants.RedisUserName,
		Password: constants.RedisPassWord,
		DB:       0,
	})

	redisVideoLike := redis.NewClient(&redis.Options{
		Addr:     constants.RedisHost + ":" + constants.RedisPort,
		Username: constants.RedisUserName,
		Password: constants.RedisPassWord,
		DB:       1,
	})
	if _, err := redisVideoLike.Ping(ctx).Result(); err != nil {
		return nil, nil, errno.NewErrNo(errno.InternalRedisErrorCode, "video like Init falied"+err.Error())
	}
	if _, err := redisCommentLike.Ping(ctx).Result(); err != nil {
		return nil, nil, errno.NewErrNo(errno.InternalRedisErrorCode, "comment like Init falied"+err.Error())
	}

	return redisCommentLike, redisVideoLike, nil
}
