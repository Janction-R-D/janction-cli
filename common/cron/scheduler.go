package cron

import (
	"github.com/go-co-op/gocron"
	"github.com/pkg/errors"
	"time"
)

type janctionCron struct {
	scheduler *gocron.Scheduler
}

const (
	TagHeartbeat = "heartbeat"
)

func JanctionCron() *janctionCron {
	return &janctionCron{scheduler: gocron.NewScheduler(time.Local)}
}

func (c *janctionCron) start() error {
	var err error
	_, err = c.scheduler.Tag(TagHeartbeat).Every(55).Seconds().Do(Heartbeat)
	if err != nil {
		return errors.Wrapf(err, "Scheduler init err, tag: %s", TagHeartbeat)
	}
	c.scheduler.StartAsync()
	return nil
}

func (c *janctionCron) stop() {
	c.scheduler.Stop()
}
