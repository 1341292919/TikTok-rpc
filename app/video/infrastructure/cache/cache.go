package cache

import (
	"TikTok-rpc/app/video/domain/model"
	"TikTok-rpc/app/video/domain/repository"
	"TikTok-rpc/pkg/constants"
	"TikTok-rpc/pkg/errno"
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type videoCache struct {
	Video   *redis.Client
	VideoId *redis.Client
}

func NewVideoCache(v *redis.Client, vid *redis.Client) repository.VideoCache {
	return &videoCache{
		Video:   v,
		VideoId: vid,
	}
}

func (ca *videoCache) NewIdToRank(ctx context.Context, vid int64) error {
	v := redis.Z{
		Score:  0, //初始分数设为0
		Member: vid,
	}
	_, err := ca.VideoId.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.ZAdd(ctx, constants.VideoIdKey, v)
		pipe.Expire(ctx, constants.VideoIdKey, 5*time.Minute)
		return nil
	})
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "AddIdToRank :"+err.Error())
	}
	return nil
}

func (ca *videoCache) RemoveIdFromRank(ctx context.Context, vid int64) error {
	_, err := ca.VideoId.ZRem(ctx, constants.VideoIdKey, vid).Result()
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "RemoveIdFromRank :"+err.Error())
	}
	return nil
}

func (ca *videoCache) UpdateIdRank(ctx context.Context, vid int64) error {
	score, err := ca.VideoId.ZScore(ctx, constants.VideoIdKey, strconv.FormatInt(vid, 10)).Result()
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "UpdateIdRank :get score failed"+err.Error())
	}
	v := redis.Z{
		Score:  score + 1,
		Member: vid,
	}
	_, err = ca.VideoId.ZAdd(ctx, constants.VideoIdKey, v).Result()
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "UpdateIdRank : add failed"+err.Error())
	}
	return nil
}
func (ca *videoCache) GetVideoIdByRank(ctx context.Context, count int64) ([]string, error) {
	rank, err := ca.VideoId.ZRevRange(ctx, constants.VideoIdKey, 0, count).Result()
	//这里可以将没有视频不定义为错误
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, "GetVideoIdByRank :"+err.Error())
	}
	return rank, nil
}

func (ca *videoCache) AddVideoToRank(ctx context.Context, video []*model.Video) error {
	pipe := ca.Video.TxPipeline()
	for _, v := range video {
		//可以考虑将这一步骤放到svc层
		videoJSON, err := json.Marshal(v)
		if err != nil {
			return err
		}
		pipe.ZAdd(ctx, constants.VideoKey, redis.Z{
			Score:  float64(v.VisitCount),
			Member: videoJSON,
		})
	}
	pipe.Expire(ctx, constants.VideoKey, 10*time.Minute)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "Execute failed :"+err.Error())
	}
	return nil
}

// 可以考虑json翻译svc层来做
func (ca *videoCache) GetVideoByRank(ctx context.Context, count int64) ([]*model.Video, error) {
	videoJSON, err := ca.Video.ZRevRange(ctx, constants.VideoKey, 0, count).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, "GetVideoByRank :"+err.Error())
	}
	videos := make([]*model.Video, len(videoJSON))
	for i, v := range videoJSON {
		var video *model.Video
		err = json.Unmarshal([]byte(v), &video)
		if err != nil {
			return nil, err
		}
		videos[i] = video
	}
	return videos, nil
}
func (ca *videoCache) DeleteVideoRank(ctx context.Context) error {
	// 删除视频排行Key
	_, err := ca.Video.Del(ctx, constants.VideoKey).Result()
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "DeleteVideoRank failed: "+err.Error())
	}
	return nil
}

func (ca *videoCache) DeleteVideoIdRank(ctx context.Context) error {
	// 删除视频ID排行Key
	_, err := ca.VideoId.Del(ctx, constants.VideoIdKey).Result()
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "DeleteVideoIdRank failed: "+err.Error())
	}
	return nil
}
