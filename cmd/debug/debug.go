package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"jct/common/config"
	"jct/common/cron"
	"jct/internal/service"
	"os"
	"time"
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
	fmt.Println("[Task]\t", os.Getenv("JCT_TASK"))
	fmt.Println("[GPU]\t", os.Getenv("JCT_GPU"))
	fmt.Println("[GPU ID]\t", os.Getenv("JCT_GPU_ID"))

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
		time.Sleep(time.Second * 3)
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
