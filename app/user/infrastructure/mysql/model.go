package mysql

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	AvatarUrl string
	OptSecret string
	MfaStatus int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
