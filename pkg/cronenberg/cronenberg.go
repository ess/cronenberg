package cronenberg

type Job struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description,omitempty"`
	Command     string            `yaml:"command"`
	When        string            `yaml:"when"`
	Lock        bool              `yaml:"lock,omitempty"`
	Env         map[string]string `yaml:"env,omitempty"`
}

type JobService interface {
	All() []*Job
}

type Runner interface {
	Run()
	Spec() string
	Name() string
	Env() map[string]string
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
