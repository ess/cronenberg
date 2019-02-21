package cron

import (
	cron "gopkg.in/robfig/cron.v2"

	"github.com/ess/cronenberg/pkg/cronenberg"
)

func New() cronenberg.Cron {
	return &Cron{cron.New()}
}

type Cron struct {
	*cron.Cron
}

func (c *Cron) Manage(runner cronenberg.Runner) error {
	_, err := c.AddJob(runner.Spec(), runner)

	return err
}

func (c *Cron) Count() int {
	return len(c.Entries())
}
