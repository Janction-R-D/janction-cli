package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"jct/common/config"
	"jct/common/cron"
	"jct/internal/service"
	"time"
)

var janction service.JanctionService

func main() {
	fmt.Println("[Janction Node]\t", "version v0.0.3")
	ParseCommandArgs()
	fmt.Println("[Os Type]\t", config.OsType)
	fmt.Println("[TestnetUrl]\t", config.TestnetUrl)
	fmt.Println("[Architecture]\t", config.Architecture)
	fmt.Println("[Use CPU]\t", config.UseGPU)
	fmt.Println("[Use GPU]\t", config.UseCPU)
	fmt.Println("[Task]\t\t", config.Task)
	fmt.Println("[CPU]\t\t", config.JCT_CPU)
	fmt.Println("[GPU]\t\t", config.JCT_GPU)
	fmt.Println("[GPU ID]\t", config.JCT_GPU_ID)

	err := janction.InitLogin()
	if err != nil {
		fmt.Println("[Init Error]", err)
	}
	err = cron.Run()
	if err != nil {
		logrus.Fatal(err)
	}
	for {
		janction.ExecTask()
		time.Sleep(time.Second * 5)
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
