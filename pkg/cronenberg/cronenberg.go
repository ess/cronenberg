// Package cronenberg provides interfaces and types for running external
// processes repeatedly at specific times.
package cronenberg

// Job is a data value that models a process that we want to run repeatedly.
type Job struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description,omitempty"`
	Command     string            `yaml:"command"`
	When        string            `yaml:"when"`
	Lock        bool              `yaml:"lock,omitempty"`
	Env         map[string]string `yaml:"env,omitempty"`
}

// JobService is an interface that describes an API for loading Jobs.
type JobService interface {
	All() []*Job
}

// Runner is an interface that wraps a Job and allows it to be run.
type Runner interface {
	Run()
	Spec() string
	Name() string
	Env() map[string]string
}

// Cron is an interface that manages a collection of Runners, running them
// according to their Jobs' schedules.
type Cron interface {
	Start()
	Stop()
	Manage(Runner) error
	Count() int
}

// Logger is an interface that describes a simple logging API.
type Logger interface {
	Info(string, string)
	Error(string, string)
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
