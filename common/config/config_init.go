package config

import (
	"jct/utils/cache"
	"jct/utils/cache/memcache"
	"time"
)

func InitConfig(conf *JanctionConf) {
	initENV(conf)
	InitMemCache()
}

var (
	OsType       string
	TestnetUrl   string
	Architecture string
	UseGPU       int
	UseCPU       int
	PrivateKey   string
	Path         string
)

var MemCache cache.Cache

func initENV(conf *JanctionConf) {
	OsType = conf.GetString("os_type", "")
	TestnetUrl = conf.GetString("testnet_url", "")
	Architecture = conf.GetString("architecture", "")
	UseCPU = conf.GetInt("useCPU", 1)
	UseGPU = conf.GetInt("UseGPU", 0)
	PrivateKey = conf.GetString("private_key", "")
	Path = conf.GetString("path", "./")
}

func InitMemCache() {
	MemCache = memcache.New(5*time.Minute, 30*time.Minute)
}
