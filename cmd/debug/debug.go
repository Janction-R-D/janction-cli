package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"jct/common/config"
	"jct/common/cron"
	"jct/internal/service"
	"os"
)

var janction service.JanctionService

func main() {
	fmt.Println("[Janction Node]\t", "version v0.0.2")
	ParseCommandArgs()
	fmt.Println("[Os Type]\t", config.OsType)
	fmt.Println("[TestnetUrl]\t", config.TestnetUrl)
	fmt.Println("[Architecture]\t", config.Architecture)
	fmt.Println("[Use CPU]\t", config.UseGPU)
	fmt.Println("[Use GPU]\t", config.UseCPU)
	fmt.Println("[Private Key]\t", os.Getenv("PRIVATE_KEY"))
	err := janction.InitLogin()
	if err != nil {
		fmt.Println("[Init Error]", err)
	}
	// 注册节点
	//err = janction.InitController()
	//if err != nil {
	//	fmt.Println("[Join Controller Error]", err)
	//}
	err = cron.Run()
	if err != nil {
		logrus.Fatal(err)
	}
	for {
	}
}

func init() {
	janction = service.JanctionService{}

}

func ParseCommandArgs() *config.JanctionConf {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.json", "config path")
	flag.Parse()
	conf, err := config.Read(configPath)
	if err != nil {
		panic(err)
	}
	config.InitConfig(conf)
	return conf
}
