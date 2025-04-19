package cache

import (
	"TikTok-rpc/app/interact/domain/model"
	"TikTok-rpc/app/interact/domain/repository"
	"TikTok-rpc/pkg/errno"
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 修改后的Lua脚本 - 支持点赞/取消点赞操作
const (
	videoLikeScript = `
	local userKey = KEYS[1]
	local countKey = KEYS[2]
	local opLogKey = KEYS[3]
	local itemID = ARGV[1]
	local opType = ARGV[2]  -- 0:点赞 1:取消
	local timestamp = ARGV[3]
	
	-- 记录操作流水
	redis.call("XADD", opLogKey, "*", "uid", ARGV[4], "vid", itemID, "op", opType, "ts", timestamp)
	
	-- 更新当前状态
	if opType == "0" then
		redis.call("HSET", userKey, "video:"..itemID, timestamp)
		return redis.call("INCR", countKey)
	else
		redis.call("HDEL", userKey, "video:"..itemID)
		return redis.call("DECR", countKey)
	end
	`

	commentLikeScript = `
	local userKey = KEYS[1]
	local countKey = KEYS[2]
	local opLogKey = KEYS[3]
	local itemID = ARGV[1]
	local opType = ARGV[2]
	local timestamp = ARGV[3]
	
	-- 记录操作流水
	redis.call("XADD", opLogKey, "*", "uid", ARGV[4], "cid", itemID, "op", opType, "ts", timestamp)
	
	-- 更新当前状态
	if opType == "0" then
		redis.call("HSET", userKey, "comment:"..itemID, timestamp)
		return redis.call("INCR", countKey)
	else
		redis.call("HDEL", userKey, "comment:"..itemID)
		return redis.call("DECR", countKey)
	end
	`
)

type interactCache struct {
	UserLike  *redis.Client
	Likecount *redis.Client
}

func NewInteractCache(userlike, count *redis.Client) repository.InteractCache {
	return &interactCache{
		UserLike:  userlike,
		Likecount: count,
	}
}

func (cache *interactCache) IsVideoLikeExist(ctx context.Context, videoid, userid int64) (bool, error) {
	userKey := fmt.Sprintf("uvlk:%d", userid)
	videoField := fmt.Sprintf("video:%d", videoid)
	exists, err := cache.UserLike.HExists(ctx, userKey, videoField).Result()
	if err != nil {
		return false, errno.NewErrNo(errno.InternalRedisErrorCode, "Check Like exist failed:"+err.Error())
	}
	return exists, nil
}

func (cache *interactCache) IsCommentLikeExist(ctx context.Context, commentid, userid int64) (bool, error) {
	userKey := fmt.Sprintf("uvlk:%d", userid)
	videoField := fmt.Sprintf("comment:%d", commentid)
	exists, err := cache.UserLike.HExists(ctx, userKey, videoField).Result()
	if err != nil {
		return false, errno.NewErrNo(errno.InternalRedisErrorCode, "Check Like exist failed:"+err.Error())
	}
	return exists, nil
}

func (cache *interactCache) NewVideoLike(ctx context.Context, videoid, userid int64) error {
	hlog.Info(videoid, userid)
	_, err := cache.UserLike.Eval(ctx, videoLikeScript, []string{
		fmt.Sprintf("uvlk:%d", userid),
		fmt.Sprintf("video:likes:%d", videoid),
		fmt.Sprintf("video:ops:%d", videoid), // 操作日志流
	},
		videoid,
		"0", // 点赞操作
		time.Now().Unix(),
		userid,
	).Result()
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "NewVideoLike:"+err.Error())
	}
	return nil
}

func (cache *interactCache) UnlikeVideoLike(ctx context.Context, videoid, userid int64) error {
	_, err := cache.UserLike.Eval(ctx, videoLikeScript, []string{
		fmt.Sprintf("uvlk:%d", userid),
		fmt.Sprintf("video:likes:%d", videoid),
		fmt.Sprintf("video:ops:%d", videoid),
	},
		videoid,
		"1", // 取消点赞
		time.Now().Unix(),
		userid,
	).Result()
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "UnlikeVideoLike:"+err.Error())
	}
	return nil
}

func (cache *interactCache) NewCommentLike(ctx context.Context, commentid, userid int64) error {
	_, err := cache.UserLike.Eval(ctx, commentLikeScript, []string{
		fmt.Sprintf("uvlk:%d", userid),
		fmt.Sprintf("comment:likes:%d", commentid),
		fmt.Sprintf("comment:ops:%d", commentid),
	},
		commentid,
		"0",
		time.Now().Unix(),
		userid,
	).Result()
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "NewCommentLike:"+err.Error())
	}
	return nil
}

func (cache *interactCache) UnlikeCommentLike(ctx context.Context, commentid, userid int64) error {
	_, err := cache.UserLike.Eval(ctx, commentLikeScript, []string{
		fmt.Sprintf("uvlk:%d", userid),
		fmt.Sprintf("comment:likes:%d", commentid),
		fmt.Sprintf("comment:ops:%d", commentid),
	},
		commentid,
		"1",
		time.Now().Unix(),
		userid,
	).Result()
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, "UnlikeCommentLike:"+err.Error())
	}
	return nil
}

func (cache *interactCache) QueryVideoLikeData(ctx context.Context) ([]*model.VideoLikeCountKey, error) {
	counts := make([]*model.VideoLikeCountKey, 0)

	// 使用SCAN迭代所有视频点赞计数键，避免阻塞Redis
	iter := cache.UserLike.Scan(ctx, 0, "video:likes:*", 100).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		// 提取视频ID
		vidStr := strings.TrimPrefix(key, "video:likes:")

		videoId, err := strconv.ParseInt(vidStr, 10, 64)

		if err != nil {
			continue
		}

		// 获取点赞数
		count, err := cache.UserLike.Get(ctx, key).Int64()
		if err != nil && !errors.Is(err, redis.Nil) {
			return nil, errno.NewErrNo(errno.InternalRedisErrorCode, "获取视频点赞数失败:"+err.Error())
		}

		counts = append(counts, &model.VideoLikeCountKey{
			Id:    videoId,
			Count: count,
		})
	}

	if err := iter.Err(); err != nil {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, "扫描视频点赞键失败:"+err.Error())
	}

	return counts, nil
}

func (cache *interactCache) QueryCommentLikeData(ctx context.Context) ([]*model.CommentLikeCountKey, error) {
	counts := make([]*model.CommentLikeCountKey, 0)

	// 使用SCAN迭代所有评论点赞计数键
	iter := cache.UserLike.Scan(ctx, 0, "comment:likes:*", 100).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		// 提取评论ID
		cidStr := strings.TrimPrefix(key, "comment:likes:")
		commentId, err := strconv.ParseInt(cidStr, 10, 64)
		if err != nil {
			continue
		}

		// 获取点赞数
		count, err := cache.UserLike.Get(ctx, key).Int64()
		if err != nil && !errors.Is(err, redis.Nil) {
			return nil, errno.NewErrNo(errno.InternalRedisErrorCode, "获取评论点赞数失败:"+err.Error())
		}

		counts = append(counts, &model.CommentLikeCountKey{
			Id:    commentId,
			Count: count,
		})
	}

	if err := iter.Err(); err != nil {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, "扫描评论点赞键失败:"+err.Error())
	}

	return counts, nil
}

func (cache *interactCache) QueryAllUserLike(ctx context.Context) ([]*model.LikeKey, error) {
	var likeKeys []*model.LikeKey

	// 1. 获取所有用户点赞键（使用SCAN避免阻塞）
	iter := cache.UserLike.Scan(ctx, 0, "uvlk:*", 100).Iterator()
	for iter.Next(ctx) {
		userKey := iter.Val()
		// 2. 提取用户ID
		uid, err := strconv.ParseInt(strings.TrimPrefix(userKey, "uvlk:"), 10, 64)
		if err != nil {
			continue
		}

		// 3. 获取该用户的所有点赞记录
		likeItems, err := cache.UserLike.HGetAll(ctx, userKey).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return nil, errno.NewErrNo(errno.InternalRedisErrorCode, "获取用户点赞记录失败:"+err.Error())
		}

		// 4. 构建LikeKey结构
		for field, timestampStr := range likeItems {
			timestamp, _ := strconv.ParseInt(timestampStr, 10, 64)
			likeKey := &model.LikeKey{
				Uid:  uid,
				Time: timestamp,
				Type: 0, // 点赞状态
			}

			// 区分视频和评论点赞
			if strings.HasPrefix(field, "video:") {
				vid, _ := strconv.ParseInt(strings.TrimPrefix(field, "video:"), 10, 64)
				likeKey.VideoId = vid
				likeKey.Status = 0 // 视频类型
			} else if strings.HasPrefix(field, "comment:") {
				cid, _ := strconv.ParseInt(strings.TrimPrefix(field, "comment:"), 10, 64)
				likeKey.CommentId = cid
				likeKey.Status = 1 // 评论类型
			}

			likeKeys = append(likeKeys, likeKey)
		}
	}

	if err := iter.Err(); err != nil {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, "扫描用户点赞键失败:"+err.Error())
	}

	// 5. 补充取消点赞的记录（从操作日志中获取）
	// 获取所有操作日志流键
	opsIter := cache.UserLike.Scan(ctx, 0, "*:ops:*", 100).Iterator()
	for opsIter.Next(ctx) {
		opKey := opsIter.Val()
		// 解析流类型（video或comment）
		var itemType int32
		var itemId int64
		if strings.Contains(opKey, "video:ops:") {
			vidStr := strings.TrimPrefix(opKey, "video:ops:")
			vid, err := strconv.ParseInt(vidStr, 10, 64)
			if err != nil {
				continue
			}
			itemType = 0 // 视频类型
			itemId = vid
		} else if strings.Contains(opKey, "comment:ops:") {
			cidStr := strings.TrimPrefix(opKey, "comment:ops:")
			cid, err := strconv.ParseInt(cidStr, 10, 64)
			if err != nil {
				continue
			}
			itemType = 1 // 评论类型
			itemId = cid
		} else {
			continue
		}

		// 读取整个操作流
		ops, err := cache.UserLike.XRange(ctx, opKey, "-", "+").Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			continue // 跳过读取失败的操作流
		}

		// 处理每条操作记录
		for _, op := range ops {
			uidStr, ok := op.Values["uid"].(string)
			if !ok {
				continue
			}
			uid, err := strconv.ParseInt(uidStr, 10, 64)
			if err != nil {
				continue
			}

			opTypeStr, ok := op.Values["op"].(string)
			if !ok {
				continue
			}
			opType, err := strconv.ParseInt(opTypeStr, 10, 32)
			if err != nil {
				continue
			}

			timestampStr, ok := op.Values["ts"].(string)
			if !ok {
				continue
			}
			timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
			if err != nil {
				continue
			}

			// 只处理取消点赞操作（opType=1）
			if opType == 1 {
				likeKey := &model.LikeKey{
					Uid:  uid,
					Time: timestamp,
					Type: 1, // 取消点赞状态
				}

				if itemType == 0 {
					likeKey.VideoId = itemId
					likeKey.Status = 0 // 视频类型
				} else {
					likeKey.CommentId = itemId
					likeKey.Status = 1 // 评论类型
				}

				likeKeys = append(likeKeys, likeKey)
			}
		}
	}

	if err := opsIter.Err(); err != nil {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, "扫描操作日志键失败:"+err.Error())
	}

	return likeKeys, nil
}

func (cache *interactCache) QueryUserLikeById(ctx context.Context, userid int64) ([]*model.LikeKey, error) {
	var likeKeys []*model.LikeKey
	userKey := fmt.Sprintf("uvlk:%d", userid)

	// 1. 获取该用户的所有点赞记录（当前状态）
	likeItems, err := cache.UserLike.HGetAll(ctx, userKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, "获取用户点赞记录失败:"+err.Error())
	}

	// 2. 构建当前点赞状态
	for field, timestampStr := range likeItems {
		timestamp, _ := strconv.ParseInt(timestampStr, 10, 64)
		likeKey := &model.LikeKey{
			Uid:  userid,
			Time: timestamp,
			Type: 0, // 点赞状态
		}

		// 区分视频和评论点赞
		if strings.HasPrefix(field, "video:") {
			vid, _ := strconv.ParseInt(strings.TrimPrefix(field, "video:"), 10, 64)
			likeKey.VideoId = vid
			likeKey.Status = 0 // 视频类型
		} else if strings.HasPrefix(field, "comment:") {
			cid, _ := strconv.ParseInt(strings.TrimPrefix(field, "comment:"), 10, 64)
			likeKey.CommentId = cid
			likeKey.Status = 1 // 评论类型
		}

		likeKeys = append(likeKeys, likeKey)
	}

	// 3. 获取该用户的所有取消点赞记录（从操作日志中）
	// 查找用户参与过的所有操作流
	opsKeys, err := cache.UserLike.Keys(ctx, "*:ops:*").Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, "获取操作日志键失败:"+err.Error())
	}

	for _, opKey := range opsKeys {
		// 解析流类型（video或comment）
		var itemType int64
		var itemId int64

		if strings.Contains(opKey, "video:ops:") {
			vidStr := strings.TrimPrefix(opKey, "video:ops:")
			vid, err := strconv.ParseInt(vidStr, 10, 64)
			if err != nil {
				continue
			}
			itemType = 0 // 视频类型
			itemId = vid
		} else if strings.Contains(opKey, "comment:ops:") {
			cidStr := strings.TrimPrefix(opKey, "comment:ops:")
			cid, err := strconv.ParseInt(cidStr, 10, 64)
			if err != nil {
				continue
			}
			itemType = 1 // 评论类型
			itemId = cid
		} else {
			continue
		}

		// 读取操作流中该用户的操作记录
		ops, err := cache.Likecount.XRange(ctx, opKey, "-", "+").Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			continue
		}

		for _, op := range ops {
			// 只处理当前用户的操作
			if uidStr, ok := op.Values["uid"].(string); ok {
				if opUid, _ := strconv.ParseInt(uidStr, 10, 64); opUid == userid {
					opTypeStr, _ := op.Values["op"].(string)
					opType, _ := strconv.ParseInt(opTypeStr, 10, 64)
					timestampStr, _ := op.Values["ts"].(string)
					timestamp, _ := strconv.ParseInt(timestampStr, 10, 64)

					// 只收集取消点赞操作
					if opType == 1 {
						likeKey := &model.LikeKey{
							Uid:  userid,
							Time: timestamp,
							Type: 1, // 取消点赞状态
						}

						if itemType == 0 {
							likeKey.VideoId = itemId
							likeKey.Status = 0 // 视频类型
						} else {
							likeKey.CommentId = itemId
							likeKey.Status = 1 // 评论类型
						}

						likeKeys = append(likeKeys, likeKey)
					}
				}
			}
		}
	}

	// 4. 按时间排序所有记录
	sort.Slice(likeKeys, func(i, j int) bool {
		return likeKeys[i].Time < likeKeys[j].Time
	})

	return likeKeys, nil
}
