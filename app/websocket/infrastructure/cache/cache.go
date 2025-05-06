package cache

import (
	"TikTok-rpc/app/websocket/domain/model"
	"TikTok-rpc/app/websocket/domain/repository"
	"TikTok-rpc/pkg/constants"
	"TikTok-rpc/pkg/errno"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type websocketCache struct {
	Message *redis.Client
}

func NewWebsocketCache(m *redis.Client) repository.WebsocketCache {
	return &websocketCache{
		Message: m,
	}
}

func (ca *websocketCache) NewMessage(ctx context.Context, message *model.Message) error {
	pipe := ca.Message.TxPipeline()
	messageJson, err := json.Marshal(message)
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "NewMessage:"+err.Error())
	}
	pipe.ZAdd(ctx, constants.MessageKey, redis.Z{
		Score:  float64(message.CreatedAT),
		Member: messageJson,
	})
	pipe.Expire(ctx, constants.MessageKey, 30*time.Minute)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "Execute failed :"+err.Error())
	}
	return nil
}

func (ca *websocketCache) GetMessage(ctx context.Context, count int64) ([]*model.Message, error) {
	messageJSON, err := ca.Message.ZRevRange(ctx, constants.MessageKey, 0, count).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, "GetVideoByRank :"+err.Error())
	}
	messages := make([]*model.Message, len(messageJSON))
	for i, v := range messageJSON {
		var message *model.Message
		err = json.Unmarshal([]byte(v), &message)
		if err != nil {
			return nil, err
		}
		messages[i] = message
	}
	return messages, nil
}

func (ca *websocketCache) NewMessageList(ctx context.Context, m []*model.Message) error {
	pipe := ca.Message.TxPipeline()
	for _, message := range m {
		messageJson, err := json.Marshal(message)
		if err != nil {
			return errno.NewErrNo(errno.InternalRedisErrorCode, "NewMessageList:"+err.Error())
		}
		pipe.ZAdd(ctx, constants.MessageKey, redis.Z{
			Score:  float64(message.CreatedAT),
			Member: messageJson,
		})
	}
	pipe.Expire(ctx, constants.MessageKey, 30*time.Minute)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "NewMessageList:"+err.Error())
	}
	return nil
}
