package config

import (
	"fmt"
	"jct/utils/cache"
	"jct/utils/cache/memcache"
	"jct/utils/snowflake"
	"time"
)

func InitConfig(conf *JanctionConf) {
	initENV(conf)
	initMemCache()
	initUtils()

}

var (
	OsType        string
	TestnetUrl    string
	ControllerUrl string
	Architecture  string
	UseGPU        int
	UseCPU        int
	//PrivateKey   string
	Path string
)

var MemCache cache.Cache

func initENV(conf *JanctionConf) {
	OsType = conf.GetString("os_type", "")
	TestnetUrl = conf.GetString("testnet_url", "")
	ControllerUrl = conf.GetString("controller_url", "")
	Architecture = conf.GetString("architecture", "")
	UseCPU = conf.GetInt("useCPU", 1)
	UseGPU = conf.GetInt("UseGPU", 0)
	//PrivateKey = conf.GetString("private_key", "")
	Path = conf.GetString("path", "./")
}

func initUtils() {
	if err := snowflake.Init("2024-07-01", 1); err != nil {
		fmt.Println("Init snowflake failed, ", err)
	}
}

func initMemCache() {
	MemCache = memcache.New(5*time.Minute, 30*time.Minute)
}
