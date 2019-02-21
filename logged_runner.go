package cronenberg

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
)

type LoggedRunner struct {
	context string
	logger  *Logger
}

var NewLoggedRunner = func(context string, logger *Logger) *LoggedRunner {
	return &LoggedRunner{context: context, logger: logger}
}

func (runner *LoggedRunner) Execute(command string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", command)

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
	log     *Logger
	context string
	level   string
	output  *bytes.Buffer
}

func newPassThrough(log *Logger, context string, level string, output *bytes.Buffer) *passThrough {
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
