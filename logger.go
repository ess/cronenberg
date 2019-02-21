package cronenberg

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Logger
}

func NewLogger() *Logger {
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
