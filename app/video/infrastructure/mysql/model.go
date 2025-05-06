package mysql

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	Id           int64
	UserId       int64
	VideoUrl     string
	CoverUrl     string
	Title        string
	Description  string
	VisitCount   int64
	LikeCount    int64
	CommentCount int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
