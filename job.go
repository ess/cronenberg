package cronenberg

import (
	//"os/exec"

	"github.com/ess/cronenberg/fs"
	yaml "gopkg.in/yaml.v2"
)

type Job struct {
	Name        string  `yaml:"name"`
	Description string  `yaml:"description,omitempty"`
	Command     string  `yaml:"command"`
	When        string  `yaml:"when"`
	Lock        bool    `yaml:"lock,omitempty"`
	running     bool    `yaml:"-"`
	logger      *Logger `yaml:"-"`
}

func (j *Job) Run() {
	// skip the run if the job is specified as locking and already running
	if j.locked() {
		j.logger.Error(j.Name, "== LOCKED ==")
	} else {
		j.logger.Info(j.Name, "Running '"+j.Command+"'")

		runner := NewLoggedRunner(j.Name, j.logger)
		j.lock()
		j.running = true
		runner.Execute(j.Command)
		j.unlock()
		j.running = false

		j.logger.Info(j.Name, "Finished")
	}
}

func (j *Job) locked() bool {
	if j.Lock {
		return j.running
	}

	return false
}

func (j *Job) lock() {
	if j.Lock {
		j.running = true
	}
}

func (j *Job) unlock() {
	if j.Lock {
		j.running = false
	}
}

func LoadJobs(jobFile string, logger *Logger) []*Job {
	data, err := fs.ReadFile(jobFile)
	if err != nil {
		logger.Error("cronenberg", "Could not load jobs file "+jobFile+": "+err.Error())
	}

	jobs := make([]*Job, 0)

	err = yaml.Unmarshal(data, &jobs)
	if err != nil {
		logger.Error("cronenberg", "Could not parse jobs file "+jobFile+": "+err.Error())
	}

	for _, j := range jobs {
		j.logger = logger
	}

	return jobs
}
