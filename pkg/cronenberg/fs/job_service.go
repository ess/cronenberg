package fs

import (
	yaml "gopkg.in/yaml.v2"

	"github.com/ess/cronenberg/pkg/cronenberg"
)

type JobService struct {
	path   string
	logger cronenberg.Logger
}

func NewJobService(path string, logger cronenberg.Logger) cronenberg.JobService {
	return &JobService{path: path, logger: logger}
}

func (service *JobService) All() []*cronenberg.Job {
	data, err := ReadFile(service.path)
	if err != nil {
		service.logger.Error(
			"cronenberg",
			"Could not load jobs file "+service.path+": "+err.Error(),
		)
	}

	jobs := make([]*cronenberg.Job, 0)

	err = yaml.Unmarshal(data, &jobs)
	if err != nil {
		service.logger.Error(
			"cronenberg",
			"Could not parse jobs file "+service.path+": "+err.Error(),
		)
	}

	return jobs
}
