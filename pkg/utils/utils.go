package utils

import (
	"TikTok-rpc/config"
	"errors"
	"net"
	"strings"

	"github.com/bytedance/gopkg/util/logger"
)

// GetMysqlDSN 会拼接 mysql 的 DSN
func GetMysqlDSN() (string, error) {
	if config.Mysql == nil {
		return "", errors.New("config not found")
	}

	dsn := strings.Join([]string{
		config.Mysql.Username, ":", config.Mysql.Password,
		"@tcp(", config.Mysql.Addr, ")/",
		config.Mysql.Database, "?charset=" + config.Mysql.Charset + "&parseTime=true",
	}, "")

	return dsn, nil
}

// AddrCheck 会检查当前的监听地址是否已被占用
func AddrCheck(addr string) bool {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return false
	}
	defer func() {
		if err := l.Close(); err != nil {
			logger.Errorf("utils.AddrCheck: failed to close listener: %v", err.Error())
		}
	}()
	return true
}

// GetAvailablePort 会尝试获取可用的监听地址
func GetAvailablePort() (string, error) {
	if config.Service.AddrList == nil {
		return "", errors.New("utils.GetAvailablePort: config.Service.AddrList is nil")
	}
	for _, addr := range config.Service.AddrList {
		if ok := AddrCheck(addr); ok {
			return addr, nil
		}
	}
	return "", errors.New("utils.GetAvailablePort: not available port from config")
}
