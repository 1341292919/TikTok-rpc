package mysql

import (
	"TikTok-rpc/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitMySQL() (db *gorm.DB, err error) {
	DB, err := gorm.Open(mysql.Open(constants.MySQLDSN), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(20 * time.Second)
	return DB, nil
}
