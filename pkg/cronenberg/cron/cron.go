// Package cron provides mechanisms for using the upstream cron package.
package cron

import (
	cron "gopkg.in/robfig/cron.v2"

	"github.com/ess/cronenberg/pkg/cronenberg"
)

// New creates a new instance of Cron.
func New() cronenberg.Cron {
	return &Cron{cron.New()}
}

// Cron is an object that uses the upstream cron package to manage a collection
// of Runners.
type Cron struct {
	*cron.Cron
}

// Manage takes a Runner and schedules it as a Cron entry. If there are issues
// adding the Runner to the schedule, an error is returned. Otherwise, nil is
// returned.
func (c *Cron) Manage(runner cronenberg.Runner) error {
	_, err := c.AddJob(runner.Spec(), runner)

	return err
}

// Count is the number of Runners scheduled in the Cron at any given time.
func (c *Cron) Count() int {
	return len(c.Entries())
}
