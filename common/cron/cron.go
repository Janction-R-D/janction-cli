package cron

import (
	"github.com/pkg/errors"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	syncOnce = &sync.Once{}
	c        *crontab
)

func init() {
	c = &crontab{
		janctionTask: JanctionCron(),
	}
}

func Run() error {
	syncOnce.Do(func() {
		c.start()
	})
	return c.err
}

func Stop() {
	c.stop()
}

type crontab struct {
	janctionTask *janctionCron
	err          error
}

func (c *crontab) start() {
	err := c.janctionTask.start()
	if err != nil {
		c.err = errors.Wrap(err, "Failed start")
	}
	return
}

func (c *crontab) stop() {
	c.janctionTask.stop()
	logrus.Info("Stopped all cron jobs")
}
