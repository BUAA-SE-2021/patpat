package run

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

func RunJava(timeout int, testInput string, command string, args ...string) (int, string, []string) {
	// instantiate new command
	runStatus := 0
	cmd := exec.Command(command, args...)
	// get pipe to standard output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic("cmd.StdoutPipe() error: " + err.Error())
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic("cmd.StdinPipe() error: " + err.Error())
	}

	// start process via command
	fmt.Println("START RUN JAVA")
	if err := cmd.Start(); err != nil {
		panic("cmd.Start() error: " + err.Error())
	}

	io.WriteString(stdin, testInput)

	// setup a buffer to capture standard output
	var buf bytes.Buffer

	// create a channel to capture any errors from wait
	done := make(chan error)
	go func() {
		if _, err := buf.ReadFrom(stdout); err != nil {
			panic("buf.Read(stdout) error: " + err.Error())
		}
		done <- cmd.Wait()
	}()

	// block on select, and switch based on actions received
	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		if err := cmd.Process.Kill(); err != nil {
			panic("failed to kill: " + err.Error())
		}
		fmt.Println("timeout reached, process killed")
		runStatus = 1 // TLE
	case err := <-done:
		if err != nil {
			close(done)
			// panic("process done, with error: " + err.Error())
			fmt.Println("process done, with error: " + err.Error())
			runStatus = 2 // error when running
		}
		fmt.Println("END RUN JAVA")
	}
	actualOutput := buf.String()
	actualOutputLines := strings.Split(strings.TrimRight(actualOutput, "\r\n"), "\n")
	return runStatus, actualOutput, actualOutputLines
}
