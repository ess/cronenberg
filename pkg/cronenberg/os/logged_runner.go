package os

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/ess/cronenberg/pkg/cronenberg"
)

// LoggedRunner is an object that knows how to run an external process and
// stream both its standard output and standard error to a Logger.
type LoggedRunner struct {
	context string
	logger  cronenberg.Logger
}

// NewLoggedRunner takes a context and a Logger, using them to configure the
// returned LoggedRunner.
var NewLoggedRunner = func(context string, logger cronenberg.Logger) *LoggedRunner {
	return &LoggedRunner{context: context, logger: logger}
}

func (runner *LoggedRunner) incomingVars() map[string]string {
	env := make(map[string]string, 0)

	for _, line := range os.Environ() {
		parts := strings.SplitN(line, "=", 2)
		env[parts[0]] = parts[1]
	}

	return env
}

func (runner *LoggedRunner) resolveVars(vars map[string]string) []string {
	env := runner.incomingVars()

	for key, value := range vars {
		env[key] = value
	}

	resolved := make([]string, 0)

	for key, value := range env {
		resolved = append(resolved, fmt.Sprintf("%s=%s", key, value))
	}

	return resolved
}

// Execute takes a command string and an environment variable map, executes the
// command with the env vars applied, and returns a byte array of output as well
// as an error. If the command returns cleanly, the error is nil. Otherwise, the
// error is non-nil.
func (runner *LoggedRunner) Execute(command string, vars map[string]string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", command)

	cmd.Env = runner.resolveVars(vars)

	output := make([]byte, 0)
	buf := bytes.NewBuffer(output)

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	stdout := newPassThrough(runner.logger, runner.context, "info", buf)
	stderr := newPassThrough(runner.logger, runner.context, "error", buf)

	cmd.Start()

	go func() {
		io.Copy(stdout, stdoutIn)
	}()

	go func() {
		io.Copy(stderr, stderrIn)
	}()

	err := cmd.Wait()

	//b := buf.Bytes()

	return buf.Bytes(), err
}

type passThrough struct {
	log     cronenberg.Logger
	context string
	level   string
	output  *bytes.Buffer
}

func newPassThrough(log cronenberg.Logger, context string, level string, output *bytes.Buffer) *passThrough {
	return &passThrough{
		log:     log,
		context: context,
		level:   level,
		output:  output,
	}
}

func (p *passThrough) Write(d []byte) (int, error) {
	p.output.Write(d)

	//line := strings.TrimSpace(string(d))
	lines := strings.Split(string(d), "\n")

	for _, line := range lines {
		if len(line) > 0 {
			switch p.level {
			case "error":
				p.log.Error(p.context, line)
			default:
				p.log.Info(p.context, line)
			}
		}
	}

	return len(d), nil
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
