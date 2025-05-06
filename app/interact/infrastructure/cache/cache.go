package cache

import (
	"TikTok-rpc/app/interact/domain/model"
	"TikTok-rpc/app/interact/domain/repository"
	"TikTok-rpc/pkg/constants"
	"TikTok-rpc/pkg/errno"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type interactCache struct {
	UserLike  *redis.Client
	LikeCount *redis.Client
}

func NewInteractCache(userlike, count *redis.Client) repository.InteractCache {
	return &interactCache{
		UserLike:  userlike,
		LikeCount: count,
	}
}

// 函数返回时调用update函数失去了原子性 仍需优化
func (cache *interactCache) NewCommentLike(ctx context.Context, commentid, userid int64) error {
	userKey := fmt.Sprintf("uvlk:%d", userid)
	field := fmt.Sprintf("comment:%d", commentid)
	value := fmt.Sprintf("%d|%d", time.Now().Unix(), 1)
	_, err := cache.UserLike.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, userKey, field, value)
		pipe.Expire(ctx, userKey, 8*time.Hour)
		return nil
	})
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "NewCommentLike:"+err.Error())
	}
	return cache.UpdateLikeCount(ctx, commentid, 1, 1)
}
func (cache *interactCache) UnlikeComment(ctx context.Context, commentid, userid int64) error {
	userKey := fmt.Sprintf("uvlk:%d", userid)
	field := fmt.Sprintf("comment:%d", commentid)
	value := fmt.Sprintf("%d|%d", time.Now().Unix(), 0)
	_, err := cache.UserLike.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, userKey, field, value)
		pipe.Expire(ctx, userKey, 8*time.Hour)
		return nil
	})
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "NewCommentLike:"+err.Error())
	}
	return cache.UpdateLikeCount(ctx, commentid, -1, 1)
}
func (cache *interactCache) NewVideoLike(ctx context.Context, videoid, userid int64) error {
	userKey := fmt.Sprintf("uvlk:%d", userid)
	field := fmt.Sprintf("video:%d", videoid)
	value := fmt.Sprintf("%d|%d", time.Now().Unix(), 1)
	_, err := cache.UserLike.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, userKey, field, value)
		pipe.Expire(ctx, userKey, 8*time.Hour)
		return nil
	})
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "NewVideoLike:"+err.Error())
	}
	return cache.UpdateLikeCount(ctx, videoid, 1, 0)
}
func (cache *interactCache) UnlikeVideo(ctx context.Context, videoid, userid int64) error {
	userKey := fmt.Sprintf("uvlk:%d", userid)
	field := fmt.Sprintf("video:%d", videoid)
	value := fmt.Sprintf("%d|%d", time.Now().Unix(), 0)
	_, err := cache.UserLike.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, userKey, field, value)
		pipe.Expire(ctx, userKey, 8*time.Hour)
		return nil
	})
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "NewVideoLike:"+err.Error())
	}
	return cache.UpdateLikeCount(ctx, videoid, -1, 0)
}
func (cache *interactCache) UpdateLikeCount(ctx context.Context, id, value, t int64) error {
	score, err := cache.LikeCount.ZScore(ctx, constants.VideoKey, strconv.FormatInt(id, 10)).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "UpdateLikeCountk :get score failed"+err.Error())
	}
	if errors.Is(err, redis.Nil) {
		score = 0
	}
	v := redis.Z{
		Score:  score + float64(value),
		Member: id,
	}
	if t == 0 {
		_, err = cache.LikeCount.ZAdd(ctx, constants.VideoLikeKey, v).Result()
	} else if t == 1 {
		_, err = cache.LikeCount.ZAdd(ctx, constants.CommentLikeKey, v).Result()
	}
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "UpdateLikeCount : add failed"+err.Error())
	}
	return nil
}
func (cache *interactCache) IsVideoLikeExist(ctx context.Context, videoid, userid int64) (bool, error) {
	userKey := fmt.Sprintf("uvlk:%d", userid)
	field := fmt.Sprintf("video:%d", videoid)
	value, err := cache.UserLike.HGet(ctx, userKey, field).Result()
	if err != nil {
		// 键或字段不存在表示未点赞
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		return false, errno.NewErrNo(errno.InternalRedisErrorCode, "IsLikeExist: "+err.Error())
	}
	parts := strings.Split(value, "|")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid data format in redis")
	}
	// 检查状态值
	status, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, fmt.Errorf("parse status failed: %v", err)
	}
	return status == 1, nil
}
func (cache *interactCache) IsCommentLikeExist(ctx context.Context, commentid, userid int64) (bool, error) {
	userKey := fmt.Sprintf("uvlk:%d", userid)
	field := fmt.Sprintf("comment:%d", commentid)
	value, err := cache.UserLike.HGet(ctx, userKey, field).Result()
	if err != nil {
		// 键或字段不存在表示未点赞
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		return false, errno.NewErrNo(errno.InternalRedisErrorCode, "IsLikeExist: "+err.Error())
	}
	parts := strings.Split(value, "|")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid data format in redis")
	}
	// 检查状态值
	status, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, fmt.Errorf("parse status failed: %v", err)
	}
	return status == 1, nil
}

func (cache *interactCache) QueryUserLikeByUid(ctx context.Context, userid int64) ([]*model.UserLike, error) {
	userKey := fmt.Sprintf("uvlk:%d", userid)

	fields, err := cache.UserLike.HGetAll(ctx, userKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, fmt.Sprintf("QueryUserLikeByUid: failed to get user likes: %v", err))
	}
	results := make([]*model.UserLike, 0)
	for field, value := range fields {
		if !strings.HasPrefix(field, "video:") {
			continue
		}
		videoIDStr := strings.TrimPrefix(field, "video:")
		videoID, err := strconv.ParseInt(videoIDStr, 10, 64)
		if err != nil {
			continue
		}

		parts := strings.Split(value, "|")
		if len(parts) != 2 {
			continue
		}

		timestamp, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			continue
		}

		status, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}

		// 只返回点赞状态为1的记录
		if status == 1 {
			results = append(results, &model.UserLike{
				Uid:     userid,
				VideoId: videoID,
				Status:  int64(status),
				Time:    time.Unix(timestamp, 0).Unix(),
				Type:    0,
			})
		}
	}
	return results, nil
}

// 该接口函数用于查询redis内的所有内容 用于同步mysql
func (cache *interactCache) GetUserLikeMessage(ctx context.Context) ([]*model.UserLike, []*model.LikeCount, []*model.LikeCount, error) {
	userLikes := make([]*model.UserLike, 0)
	userIter := cache.UserLike.Scan(ctx, 0, "uvlk:*", 0).Iterator()
	for userIter.Next(ctx) {
		userKey := userIter.Val()
		userID, err := strconv.ParseInt(strings.TrimPrefix(userKey, "uvlk:"), 10, 64)
		if err != nil {
			continue
		}
		fields, err := cache.UserLike.HGetAll(ctx, userKey).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return nil, nil, nil, fmt.Errorf("get user likes failed: %v", err)
		}
		for field, value := range fields {
			if strings.HasPrefix(field, "video:") {
				videoID, err := strconv.ParseInt(strings.TrimPrefix(field, "video:"), 10, 64)
				if err != nil {
					continue
				}
				parts := strings.Split(value, "|")
				if len(parts) != 2 {
					continue
				}
				status, _ := strconv.Atoi(parts[1])
				userLike := &model.UserLike{
					Uid:     userID,
					VideoId: videoID,
					Status:  int64(status),
					Type:    0, //代表是对视频的点赞
				}
				userLikes = append(userLikes, userLike)
			} else if strings.HasPrefix(field, "comment:") {
				commentID, err := strconv.ParseInt(strings.TrimPrefix(field, "comment:"), 10, 64)
				if err != nil {
					continue
				}
				parts := strings.Split(value, "|")
				if len(parts) != 2 {
					continue
				}
				status, _ := strconv.Atoi(parts[1])
				userLike := &model.UserLike{
					Uid:       userID,
					CommentId: commentID,
					Status:    int64(status),
					Type:      1, //代表是对评论的点赞
				}
				userLikes = append(userLikes, userLike)
			}
		}
	}
	//接下来对LikeCount的信息进行统计
	commentLikeCounts := make([]*model.LikeCount, 0)
	videoLikeCounts := make([]*model.LikeCount, 0)
	Vrank, err := cache.LikeCount.ZRevRangeWithScores(ctx, constants.VideoLikeKey, 0, 100).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, nil, nil, errno.NewErrNo(errno.InternalRedisErrorCode, "GetVideoIdByRank :"+err.Error())
	}
	for _, item := range Vrank {
		id, ok := item.Member.(string)
		if !ok {
			continue
		}
		videoID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			continue
		}
		videoLikeCounts = append(videoLikeCounts, &model.LikeCount{
			Id:    videoID,
			Count: int64(item.Score),
			Type:  0,
		})
	}
	Crank, err := cache.LikeCount.ZRevRangeWithScores(ctx, constants.CommentLikeKey, 0, 100).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, nil, nil, errno.NewErrNo(errno.InternalRedisErrorCode, "GetCommentLikeCount :"+err.Error())
	}
	for _, item := range Crank {
		id, ok := item.Member.(string)
		if !ok {
			continue
		}
		commentID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			continue
		}
		commentLikeCounts = append(commentLikeCounts, &model.LikeCount{
			Id:    commentID,
			Count: int64(item.Score),
			Type:  1,
		})
	}
	return userLikes, videoLikeCounts, commentLikeCounts, nil
}

// 由于redis内的内容会丢失，以下接口函数用于转载mysql内的内容
func (cache *interactCache) UploadUserLike(ctx context.Context, data []*model.UserLike) error {
	pipe := cache.UserLike.TxPipeline()
	for _, item := range data {
		userKey := fmt.Sprintf("uvlk:%d", item.Uid)
		field := fmt.Sprintf("video:%d", item.VideoId)
		value := fmt.Sprintf("%d|%d", item.Time, item.Status)
		pipe.HSet(ctx, userKey, field, value)
		pipe.Expire(ctx, userKey, 8*time.Hour)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, fmt.Sprintf("UploadUserLike failed: %v", err))
	}
	return nil
}
func (cache *interactCache) UploadLikeCount(ctx context.Context, data []*model.LikeCount) error {
	pipe := cache.LikeCount.TxPipeline()
	for _, item := range data {
		v := redis.Z{
			Score:  float64(item.Count),
			Member: item.Id,
		}
		if item.Type == 0 {
			pipe.ZAdd(ctx, constants.VideoLikeKey, v)
		} else if item.Type == 1 {
			pipe.ZAdd(ctx, constants.CommentLikeKey, v)
		}
	}
	pipe.Expire(ctx, constants.VideoLikeKey, 8*time.Hour)
	pipe.Expire(ctx, constants.CommentLikeKey, 8*time.Hour)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, fmt.Sprintf("UploadUserLike failed: %v", err))
	}
	return nil
}
