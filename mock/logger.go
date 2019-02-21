package mock

import (
	"fmt"
	"time"
)

type Logger struct {
	Lines []string
}

func NewLogger() *Logger {
	return &Logger{Lines: make([]string, 0)}
}

func (l *Logger) Info(context string, message string) {
	l.Lines = append(l.Lines, fmt.Sprintf("INFO %s %s", context, message))
}

func (l *Logger) Error(context string, message string) {
	l.Lines = append(l.Lines, fmt.Sprintf("ERR %s %s", context, message))
}

func (l *Logger) Wait() {
	time.Sleep(100 * time.Millisecond)
}
