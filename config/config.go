package config

import (
	"errors"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Etcd         *etcd
	Mysql        *mySQL
	Redis        *redis
	Service      *service
	Oss          *oss
	runtimeViper = viper.New()
)

const (
	File     = "./config/config.yaml"
	FileTest = "/media/yang/OS/code/go/tiktok-rpc/config/config.yaml"
	FileType = "yaml"
)

func Init(service string) {
	runtimeViper.SetConfigFile(FileTest)
	runtimeViper.SetConfigType(FileType)
	if err := runtimeViper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			logger.Fatal("config.Init: could not find config files")
		}
		logger.Fatalf("config.Init: read config error: %v", err)
	}
	configMapping(service)
	// 设置持续监听
	runtimeViper.OnConfigChange(func(e fsnotify.Event) {
		logger.Infof("config: notice config changed: %v\n", e.String())
		configMapping(service)
	})
	runtimeViper.WatchConfig()
}

func configMapping(srv string) {
	c := new(config)
	if err := runtimeViper.Unmarshal(&c); err != nil {
		// 由于这个函数会在配置重载时被再次触发，所以需要判断日志记录方式
		logger.Fatalf("config.configMapping: config: unmarshal error: %v", err)
	}
	Mysql = &c.MySQL
	Redis = &c.Redis
	Oss = &c.OSS
	Etcd = &c.Etcd
	Service = getService(srv)
}
func getService(name string) *service {
	addrList := runtimeViper.GetStringSlice("services." + name + ".addr")

	return &service{
		Name:     runtimeViper.GetString("services." + name + ".name"),
		AddrList: addrList,
		LB:       runtimeViper.GetBool("services." + name + ".load-balance"),
	}
}
