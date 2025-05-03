package mysql

import (
	"TikTok-rpc/app/video/domain/model"
	"TikTok-rpc/app/video/domain/repository"
	"TikTok-rpc/pkg/constants"
	"TikTok-rpc/pkg/errno"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type videoDB struct {
	client *gorm.DB
}

func NewVideoDB(client *gorm.DB) repository.VideoDB {
	return &videoDB{client: client}
}

func (db *videoDB) CreateVideo(ctx context.Context, video *model.Video) (int64, error) {
	videoResp := &Video{
		UserId:      video.Uid,
		Title:       video.Title,
		Description: video.Description,
		VideoUrl:    video.VideoUrl,
		CoverUrl:    video.CoverUrl,
	}
	err := db.client.
		WithContext(ctx).
		Table(constants.TableVideo).
		Create(&videoResp).
		Error
	if err != nil {
		return 0, err
	}
	return videoResp.Id, nil
}
func (db *videoDB) IsVideoExist(ctx context.Context, id int64) (bool, error) {
	var video *Video
	err := db.client.
		WithContext(ctx).
		Table(constants.TableVideo).
		Where("id = ?", id).
		First(&video).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (db *videoDB) QueryVideoByUid(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error) {
	var videoResp []*Video
	var count int64
	var err error
	err = db.client.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 查询视频列表
		err = tx.
			Table(constants.TableVideo).
			Where("user_id = ?", req.Uid).
			Limit(int(req.PageSize)).
			Offset(int((req.PageNum - 1) * req.PageNum)).
			Count(&count).
			Find(&videoResp).
			Error
		if err != nil {
			return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get videoInfo: %v", err)
		}
		// 遍历查询到的视频，将每个视频的 visit_count 加一
		for _, video := range videoResp {
			err = updateVisitCount(tx, video.Id, 1)
			if err != nil {
				log.Printf("Failed to update visit_count for video %d: %v", video.Id, err)
				return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to update visit_count: %v", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, -1, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql:Transaction falied: %v", err)
	}
	return buildVideoList(videoResp), count, nil
}
func (db *videoDB) QueryVideoById(ctx context.Context, id string) (*model.Video, error) {
	//让每次调用该函数时都先检查Video存在与否，这样能保证这一层函数里不会有业务错误，但是会频繁访问数据库吧？
	//目前还是需要返回业务错误，如id不存在
	var videoResp *Video
	var err error
	err = db.client.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err = tx.Table(constants.TableVideo).
			Where("id = ?", id).
			First(&videoResp).
			Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) { //业务错误
				return errno.Errorf(errno.InternalServiceErrorCode, "video id not exist")
			}
			return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get videoInfo: %v", err)
		}
		err = updateVisitCount(tx, videoResp.Id, 1)
		if err != nil {
			log.Printf("Failed to update visit_count for video %d: %v", videoResp.Id, err)
			return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to update visit_count: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql:Transaction falied: %v", err)
	}
	return buildVideo(videoResp), nil
}
func (db *videoDB) QueryVideoListById(ctx context.Context, id []string) ([]*model.Video, error) {
	var videoResp []*model.Video
	for _, i := range id {
		v, err := db.QueryVideoById(ctx, i)
		if err != nil {
			return nil, err
		}
		videoResp = append(videoResp, v)
	}
	return videoResp, nil
}

func (db *videoDB) QueryVideoByKeyWord(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error) {
	keyword := "%" + req.Keyword + "%"
	var videoResp []*Video
	var count int64
	from_date := time.Unix(req.FromDate, 0)
	to_date := time.Unix(req.ToDate, 0)
	err := db.client.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var err error
		if req.Uid == -1 {
			err = tx.
				Table(constants.TableVideo).
				Where("created_at >= ? AND created_at <= ? ", from_date, to_date).
				Where("title LIKE ? OR description LIKE ?", keyword, keyword).
				Limit(int(req.PageSize)).
				Offset(int((req.PageNum - 1) * req.PageSize)).
				Count(&count).
				Find(&videoResp).
				Error
		} else {
			err = tx.
				Table(constants.TableVideo).
				Where("created_at >= ? AND created_at <= ? ", from_date, to_date).
				Where("title LIKE ? OR description LIKE ?", keyword, keyword).
				Where("id = ?", req.Uid).
				Limit(int(req.PageSize)).
				Offset(int((req.PageNum - 1) * req.PageSize)).
				Count(&count).
				Find(&videoResp).
				Error
		}
		if err != nil {
			return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get videoInfo: %v", err)
		}
		for _, video := range videoResp {
			err = updateVisitCount(tx, video.Id, 1)
			if err != nil {
				log.Printf("Failed to update visit_count for video %d: %v", video.Id, err)
				return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to update visit_count: %v", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, -1, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql:Transaction falied: %v", err)
	}
	return buildVideoList(videoResp), count, nil
}
func (db *videoDB) QueryPopularVideo(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error) {
	var videoResp []*Video
	var count int64
	var err error
	err = db.client.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err = tx.
			Table(constants.TableVideo).
			Order("visit_count DESC").
			Limit(int(req.PageSize)).
			Offset(int((req.PageNum - 1) * req.PageSize)).
			Count(&count).
			Find(&videoResp).
			Error
		if err != nil {
			return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get videoInfo: %v", err)
		}
		for _, video := range videoResp {
			err = updateVisitCount(tx, video.Id, 1)
			if err != nil {
				log.Printf("Failed to update visit_count for video %d: %v", video.Id, err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, -1, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql:Transaction falied: %v", err)
	}
	return buildVideoList(videoResp), count, nil
}
func (db *videoDB) UpdateCommentCount(ctx context.Context, videoid, changecount int64) error {
	err := db.client.WithContext(ctx).
		Table(constants.TableVideo).
		Where("id = ?", videoid).
		Update("comment_count", gorm.Expr("comment_count + ?", changecount)).
		Error
	if err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to update comment count: %v", err)
	}
	return nil
}
func (db *videoDB) UpdateLikeCount(ctx context.Context, videoid, likecount int64) error {
	err := db.client.WithContext(ctx).
		Table(constants.TableVideo).
		Where("id = ?", videoid).
		Update("like_count", likecount).
		Error
	if err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to update like count: %v", err)
	}
	return nil
}

func (db *videoDB) QueryVideoDuringTime(ctx context.Context, req *model.VideoReq) ([]*model.Video, int64, error) {
	var videoResp []*Video
	var count int64
	from_date := time.Unix(req.FromDate, 0)
	to_date := time.Unix(req.ToDate, 0)
	err := db.client.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.
			Table(constants.TableVideo).
			Where("created_at >= ? AND created_at <= ? ", from_date, to_date).
			Limit(int(req.PageSize)).
			Offset(int((req.PageNum - 1) * req.PageSize)).
			Count(&count).
			Find(&videoResp).
			Error
		if err != nil {
			return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get videoInfo: %v", err)
		}
		for _, video := range videoResp {
			err = updateVisitCount(tx, video.Id, 1)
			if err != nil {
				log.Printf("Failed to update visit_count for video %d: %v", video.Id, err)
				return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to update visit_count: %v", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, -1, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql:Transaction falied: %v", err)
	}
	return buildVideoList(videoResp), count, nil
}
func (db *videoDB) QueryLikeCount(ctx context.Context) ([]*model.LikeCount, error) {
	var videoResp []*Video
	err := db.client.WithContext(ctx).
		Table(constants.TableVideo).
		Find(&videoResp).Error
	if err != nil {
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get videoInfo: %v", err)
	}
	return buildLikeCountList(videoResp), nil
}

func updateVisitCount(tx *gorm.DB, videoid int64, delta int) error {
	return tx.Table(constants.TableVideo).
		Where("id = ?", videoid).
		Update("visit_count", gorm.Expr("visit_count + ?", delta)).
		Error
}

func buildVideo(data *Video) *model.Video {
	return &model.Video{
		Id:           data.Id,
		Title:        data.Title,
		Description:  data.Description,
		VideoUrl:     data.VideoUrl,
		CoverUrl:     data.CoverUrl,
		Uid:          data.UserId,
		CreateAT:     data.CreatedAt.Unix(),
		UpdateAT:     data.UpdatedAt.Unix(),
		CommentCount: data.CommentCount,
		VisitCount:   data.VisitCount,
		LikeCount:    data.LikeCount,
	}
}
func buildVideoList(data []*Video) []*model.Video {
	videoList := make([]*model.Video, 0)
	for _, video := range data {
		videoList = append(videoList, buildVideo(video))
	}
	return videoList
}
func buildLikeCount(data *Video) *model.LikeCount {
	return &model.LikeCount{
		Count: data.LikeCount,
		Id:    data.Id,
	}
}
func buildLikeCountList(data []*Video) []*model.LikeCount {
	list := make([]*model.LikeCount, 0)
	for _, item := range data {
		list = append(list, buildLikeCount(item))
	}
	return list
}
