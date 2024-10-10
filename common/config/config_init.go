package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"jct/utils/cache"
	"jct/utils/cache/memcache"
	"jct/utils/snowflake"
	"os"
	"strings"
	"time"
)

func InitConfig(conf *JanctionConf) {
	initENV(conf)
	initMemCache()
	initUtils()
	initNodeId()

}

var (
	OsType        string
	TestnetUrl    string
	ControllerUrl string
	Architecture  string
	JCT_CPU       string
	JCT_GPU       string
	JCT_GPU_ID    string
	UseGPU        int
	UseCPU        int
	Path          string
	Task          string
	TaskType      string
	NodeID        string
)

var MemCache cache.Cache

func initENV(conf *JanctionConf) {
	OsType = conf.GetString("os_type", "")
	TestnetUrl = conf.GetString("testnet_url", "")
	ControllerUrl = conf.GetString("controller_url", "")
	Architecture = conf.GetString("architecture", "")
	UseCPU = 1
	UseGPU = 0
	useDevice := os.Getenv("JCT_USE_DEVICE")
	if useDevice == "gpu" {
		UseCPU = 0
		UseGPU = 1
	}
	JCT_CPU = os.Getenv("JCT_CPU")
	Task = os.Getenv("JCT_TASK")
	TaskType = os.Getenv("JCT_TASK_TYPE")

	Path = conf.GetString("path", "./")
	JCT_GPU = os.Getenv("JCT_GPU")
	JCT_GPU_ID = os.Getenv("JCT_GPU_ID")
}

func initUtils() {
	if err := snowflake.Init("2024-07-01", 1); err != nil {
		fmt.Println("Init snowflake failed, ", err)
	}
}

func initMemCache() {
	MemCache = memcache.New(5*time.Minute, 30*time.Minute)
}

func initNodeId() {
	_, err := os.Stat(".id") //os.Stat获取文件信息
	if os.IsNotExist(err) {
		f, err := os.Create(".id")
		defer f.Close()
		if err != nil {
			logrus.Error(err)
		} else {
			_, err = f.Write((snowflake.GenID()))
			if err != nil {
				logrus.Error(err)
			}
		}
	}

	uuid, err := os.ReadFile(".id")
	if err != nil {
		logrus.Error(err)
	}
	nodeId := strings.Replace(string(uuid), "\n", "", -1)
	NodeID = nodeId
}
