package logger

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/ess/cronenberg/pkg/cronenberg"
)

type Logger struct {
	log *logrus.Logger
}

func New() cronenberg.Logger {
	log := logrus.New()
	log.Formatter = &Formatter{}

	log.Out = os.Stdout

	return &Logger{log: log}
}

func (logger *Logger) Info(context string, message string) {
	logger.log.WithField("context", context).Info(message)
}

func (logger *Logger) Error(context string, message string) {
	logger.log.WithField("context", context).Error(message)
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
