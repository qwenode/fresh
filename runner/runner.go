package runner

import (
	"github.com/qwenode/gogo/ee"
	"io"
	"os/exec"
	"strings"
)

func run() bool {
	runnerLog("Running...卧槽")
	p := buildPath()
	s := ee.GetString("run_args")
	split := strings.Split(s, " ")
	cmd := exec.Command(p, split...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		fatal(err)
	}

	go io.Copy(appLogWriter{}, stderr)
	go io.Copy(appLogWriter{}, stdout)

	go func() {
		<-stopChannel
		pid := cmd.Process.Pid
		runnerLog("Killing PID %d", pid)
		cmd.Process.Kill()
	}()

	return true
}
