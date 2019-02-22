package cron

import (
	"github.com/ess/cronenberg/pkg/cronenberg"
	"github.com/ess/cronenberg/pkg/cronenberg/os"
)

// NewRunner takes a Job and a Logger, using them to configure the returned
// Runner.
func NewRunner(job *cronenberg.Job, logger cronenberg.Logger) cronenberg.Runner {
	return &Runner{job: job, logger: logger, running: false}
}

// Runner is an object that knows how to run a Job.
type Runner struct {
	job     *cronenberg.Job
	logger  cronenberg.Logger
	running bool
}

// Name returns the name of the Job that the Runner wraps.
func (runner *Runner) Name() string {
	return runner.job.Name
}

// Spec returns the cron schedule spec for the Job that the Runner wraps.
func (runner *Runner) Spec() string {
	return "0 " + runner.job.When
}

// Env returns the environment map for the Job that the Runner wraps.
func (runner *Runner) Env() map[string]string {
	return runner.job.Env
}

// Run runs the wrapped Job's command. If the job is configured as a locking
// process, only one instance of that process will ever be run by a given
// Runner.
func (runner *Runner) Run() {
	// skip the run if the job is specified as locking and already running
	if runner.locked() {
		runner.logger.Error(runner.job.Name, "== LOCKED ==")
	} else {
		runner.logger.Info(runner.job.Name, "Running '"+runner.job.Command+"'")

		executor := os.NewLoggedRunner(runner.job.Name, runner.logger)

		runner.lock()
		executor.Execute(runner.job.Command, runner.job.Env)
		runner.unlock()

		runner.logger.Info(runner.job.Name, "Finished")
	}
}

func (runner *Runner) locked() bool {
	if runner.lockingJob() {
		return runner.running
	}

	return false
}

func (runner *Runner) lock() {
	if runner.lockingJob() {
		runner.running = true
	}
}

func (runner *Runner) unlock() {
	if runner.lockingJob() {
		runner.running = false
	}
}

func (runner *Runner) lockingJob() bool {
	return runner.job.Lock
}
