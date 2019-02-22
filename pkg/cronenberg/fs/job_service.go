package fs

import (
	yaml "gopkg.in/yaml.v2"

	"github.com/ess/cronenberg/pkg/cronenberg"
)

// JobService is an object that knows how to load job definitions from a yaml
// file.
type JobService struct {
	path   string
	logger cronenberg.Logger
}

// NewJobService takes a file path and a logger, using them to configure the
// returned JobService.
func NewJobService(path string, logger cronenberg.Logger) cronenberg.JobService {
	return &JobService{path: path, logger: logger}
}

// All reads the jobs from the service's file path and returns them as an array
// of Job objects. If there are issues along the way, the result is an empty
// array.
func (service *JobService) All() []*cronenberg.Job {
	jobs := make([]*cronenberg.Job, 0)

	data, err := ReadFile(service.path)
	if err != nil {
		service.logger.Error(
			"cronenberg",
			"Could not load jobs file "+service.path+": "+err.Error(),
		)

		return jobs
	}

	err = yaml.Unmarshal(data, &jobs)
	if err != nil {
		service.logger.Error(
			"cronenberg",
			"Could not parse jobs file "+service.path+": "+err.Error(),
		)
	}

	return jobs
}
