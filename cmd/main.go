package main

import (
	"flag"
	"fmt"
	"github.com/sevlyar/go-daemon"
	"github.com/sirupsen/logrus"
	"jct/common/config"
	"jct/common/cron"
	"jct/internal/service"
	"log"
	"os"
	"syscall"
	"time"
)

var (
	signal = flag.String("s", "", `Send signal to the daemon:
  quit — graceful shutdown
  stop — fast shutdown`)
)

var janction service.JanctionService

func main() {
	ParseCommandArgs()
	err := janction.InitLogin()
	if err != nil {
		fmt.Println("[Init Error]", err)
	}
	// 注册节点
	err = janction.InitController()
	if err != nil {
		fmt.Println("[Join Controller Error]", err)
	}

	err = cron.Run()
	if err != nil {
		logrus.Fatal(err)
	}

	flag.Parse()

	//daemon.AddCommand(daemon.StringFlag(signal, "reload"), syscall.SIGHUP, reloadHandler)

	cntxt := &daemon.Context{
		PidFileName: "jct_node.pid",
		PidFilePerm: 0644,
		LogFileName: "jct_node.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[jct_node]"},
	}

	if len(daemon.ActiveFlags()) > 0 {
		d, err := cntxt.Search()
		if err != nil {
			logrus.Fatalf("Unable send signal to the daemon: %s", err.Error())
		}
		daemon.SendCommands(d)
		return
	}

	d, err := cntxt.Reborn()
	if err != nil {
		logrus.Fatalln(err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	logrus.Println("Janction Node Started")

	go worker() //启动主循环函数

	err = daemon.ServeSignals() //信号处理函数
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}

	logrus.Println("Jacntion Daemon terminated")
}

var (
	stop = make(chan struct{})
	done = make(chan struct{})
)

func init() {
	janction = service.JanctionService{}
}

func worker() {
LOOP:
	for {
		// TODO ping
		logrus.Infof("waiting")
		time.Sleep(5 * time.Second) // this is work to be done by worker.
		select {
		case <-stop:
			break LOOP
		default:
		}
	}
	done <- struct{}{}
}

func termHandler(sig os.Signal) error {
	log.Println("Janction Node Terminating...")
	stop <- struct{}{}
	if sig == syscall.SIGQUIT {
		<-done
	}
	return daemon.ErrStop
}

//func reloadHandler(sig os.Signal) error {
//	log.Println("configuration reloaded")
//	return nil
//}

func ParseCommandArgs() *config.JanctionConf {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.json", "config path")
	flag.Parse()

	daemon.AddCommand(daemon.StringFlag(signal, "quit"), syscall.SIGQUIT, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "stop"), syscall.SIGTERM, termHandler)

	conf, err := config.Read(configPath)
	if err != nil {
		panic(err)
	}
	config.InitConfig(conf)
	return conf
}

func ping() {
	//var header = map[string]string{}
	//params := url.Values{
	//	"version": []string{"0.0.1"},
	//}
	//_, err := utils.GetWithTimeout("http://127.0.0.1:8767/ping", header, params, 30*time.Second)
	//if err != nil {
	//	return
	//}
}
