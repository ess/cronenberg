package logger

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/ess/cronenberg"
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
