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
