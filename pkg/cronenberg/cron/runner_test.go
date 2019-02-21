package cron

import (
	"fmt"
	"testing"
	"time"

	"github.com/ess/cronenberg/mock"
	"github.com/ess/cronenberg/pkg/cronenberg"
)

func TestCron_Run(t *testing.T) {
	name := "test-job"
	when := "* * * * *"
	command := "echo blah"
	job := &cronenberg.Job{
		When:    when,
		Command: command,
		Name:    name,
	}

	t.Run("it runs the command", func(t *testing.T) {
		log := mock.NewLogger()
		runner := NewRunner(job, log)

		runner.Run()

		time.Sleep(10 * time.Millisecond) // to give the logger time to catch up

		found := false
		expected := fmt.Sprintf("INFO %s %s", name, "blah")
		for _, line := range log.Lines {
			if line == expected {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expected blah to be logged")
		}
	})
}
