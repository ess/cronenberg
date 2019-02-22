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
