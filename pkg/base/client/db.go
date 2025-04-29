package client

import (
	"TikTok-rpc/pkg/errno"
	"TikTok-rpc/pkg/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitMySQL() (db *gorm.DB, err error) {
	dsn, err := utils.GetMysqlDSN()
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, fmt.Sprintf("dal.InitMySQL get mysql DSN error: %v", err))
	}
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, fmt.Sprintf("dal.InitMySQL mysql connect error: %v", err))
	}
	sqlDB, err := DB.DB() // 尝试获取 DB 实例对象
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, fmt.Sprintf("get generic database object error: %v", err))
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(20 * time.Second)
	// 进行连通性测试
	if err = sqlDB.Ping(); err != nil {
		return nil, errno.NewErrNo(errno.InternalDatabaseErrorCode, fmt.Sprintf("ping database error: %v", err))
	}
	return DB, nil
}
