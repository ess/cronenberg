package cron

import (
	"github.com/ess/cronenberg/pkg/cronenberg"
)

func NewRunner(job *cronenberg.Job, logger cronenberg.Logger) cronenberg.Runner {
	return &Runner{job: job, logger: logger, running: false}
}

type Runner struct {
	job     *cronenberg.Job
	logger  cronenberg.Logger
	running bool
}

func (runner *Runner) Name() string {
	return runner.job.Name
}

func (runner *Runner) Spec() string {
	return "0 " + runner.job.When
}

func (runner *Runner) Env() map[string]string {
	return runner.job.Env
}

func (runner *Runner) Run() {
	// skip the run if the job is specified as locking and already running
	if runner.locked() {
		runner.logger.Error(runner.job.Name, "== LOCKED ==")
	} else {
		runner.logger.Info(runner.job.Name, "Running '"+runner.job.Command+"'")

		executor := cronenberg.NewLoggedRunner(runner.job.Name, runner.logger)

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
