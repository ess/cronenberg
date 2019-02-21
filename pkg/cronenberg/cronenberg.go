package cronenberg

type Job struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description,omitempty"`
	Command     string `yaml:"command"`
	When        string `yaml:"when"`
	Lock        bool   `yaml:"lock,omitempty"`
}

type JobService interface {
	All() []*Job
}

type Runner interface {
	Run()
	Spec() string
	Name() string
}

type Cron interface {
	Start()
	Stop()
	Manage(Runner) error
	Count() int
}

type Logger interface {
	Info(string, string)
	Error(string, string)
}
