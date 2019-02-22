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

/*
Copyright 2019 Dennis Walters

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
