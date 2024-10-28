package cron

import (
	"jct/internal/service"
	"log"
)

var janction service.JanctionService

func init() {
	janction = service.JanctionService{}
}

func Heartbeat() {
	err := janction.Run()
	log.Printf("Ping Janction Service Success...")
	if err != nil {
		log.Println(err)
		return
	}
}
