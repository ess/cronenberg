package cron

import (
	"fmt"
	"os"
	"testing"

	"github.com/ess/cronenberg/mock"
	"github.com/ess/cronenberg/pkg/cronenberg"
)

var (
	name    string = "test-job"
	when    string = "* * * * *"
	command string = "echo blah"
)

func setFlibberty(flibberty string) {
	os.Setenv("flibberty", flibberty)
}

func unsetFlibberty() {
	os.Unsetenv("flibberty")
}

func TestRunner_Name(t *testing.T) {
	job := &cronenberg.Job{
		When:    when,
		Name:    name,
		Command: command,
	}

	log := mock.NewLogger()
	runner := NewRunner(job, log)

	actual := runner.Name()

	if actual != name {
		t.Errorf("Expected '%s', got '%s'", name, actual)
	}
}

func TestRunner_Spec(t *testing.T) {
	job := &cronenberg.Job{
		When:    when,
		Name:    name,
		Command: command,
	}

	log := mock.NewLogger()
	runner := NewRunner(job, log)

	actual := runner.Spec()

	if actual != "0 "+when {
		t.Errorf("Expected '0 %s', got '%s'", when, actual)
	}
}

func TestRunner_Run(t *testing.T) {
	t.Run("for locking jobs", func(t *testing.T) {
		job := &cronenberg.Job{
			When:    when,
			Command: command,
			Name:    name,
			Lock:    true,
		}

		t.Run("when the job is already running", func(t *testing.T) {
			log := mock.NewLogger()
			runner := &Runner{job: job, logger: log, running: true}

			t.Run("it reports that the job is locked", func(t *testing.T) {
				runner.Run()

				log.Wait()

				found := false
				expected := fmt.Sprintf("ERR %s %s", name, "== LOCKED ==")
				for _, line := range log.Lines {
					if line == expected {
						found = true
						break
					}
				}

				if !found {
					t.Errorf("Expected the job to be reported as locked")
				}

			})
		})

		t.Run("when the job is not running", func(t *testing.T) {
			log := mock.NewLogger()
			runner := &Runner{job: job, logger: log, running: false}

			t.Run("it runs the command", func(t *testing.T) {
				runner.Run()

				log.Wait()

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
		})
	})

	t.Run("for standard jobs", func(t *testing.T) {
		job := &cronenberg.Job{
			When:    when,
			Command: command,
			Name:    name,
		}

		t.Run("it runs the command", func(t *testing.T) {
			log := mock.NewLogger()
			runner := NewRunner(job, log)

			runner.Run()

			log.Wait()

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
	})

	t.Run("with incoming env vars", func(t *testing.T) {
		flibberty := "gibbets"
		setFlibberty(flibberty)
		defer unsetFlibberty()

		command := "echo $flibberty"
		job := &cronenberg.Job{
			When:    when,
			Command: command,
			Name:    name,
		}

		t.Run("it passes the incoming env vars to the command's shell", func(t *testing.T) {
			log := mock.NewLogger()
			runner := NewRunner(job, log)

			runner.Run()

			log.Wait()

			found := false
			expected := fmt.Sprintf("INFO %s %s", name, flibberty)
			for _, line := range log.Lines {
				if line == expected {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("Expected flibberty to be logged")
			}
		})
	})

	t.Run("with job-scoped env vars", func(t *testing.T) {
		flibberty := "upstream"
		env := make(map[string]string, 0)
		env["flibberty"] = "gibbets"

		command := "echo $flibberty"
		job := &cronenberg.Job{
			When:    when,
			Command: command,
			Name:    name,
			Env:     env,
		}

		t.Run("when the env var is also provided by the calling shell", func(t *testing.T) {
			setFlibberty(flibberty)
			defer unsetFlibberty()

			t.Run("the job-scoped value is used", func(t *testing.T) {
				log := mock.NewLogger()
				runner := NewRunner(job, log)

				runner.Run()

				log.Wait()

				found := false
				expected := fmt.Sprintf("INFO %s %s", name, env["flibberty"])
				for _, line := range log.Lines {
					if line == expected {
						found = true
						break
					}
				}

				if !found {
					t.Errorf("Expected the job's flibberty to be logged")
				}
			})
		})

		t.Run("when the env var is not provided by the calling shell", func(t *testing.T) {
			unsetFlibberty()

			t.Run("the job-scoped value is used", func(t *testing.T) {
				log := mock.NewLogger()
				runner := NewRunner(job, log)

				runner.Run()

				log.Wait()

				found := false
				expected := fmt.Sprintf("INFO %s %s", name, env["flibberty"])
				for _, line := range log.Lines {
					if line == expected {
						found = true
						break
					}
				}

				if !found {
					t.Errorf("Expected the job's flibberty to be logged")
				}

			})
		})
	})
}
