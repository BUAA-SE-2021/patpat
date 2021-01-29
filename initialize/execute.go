package initialize

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func Execute(folderName string, testInput string) {
	log, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
	if err != nil {
		panic(err) //replace with logger, or anything you want
	}
	subProcess := exec.Command("java", "-classpath", folderName+"/src", "Test")
	stdin, err := subProcess.StdinPipe()
	if err != nil {
		panic(err) //replace with logger, or anything you want
	}
	defer stdin.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line
	subProcess.Stdout = log
	subProcess.Stderr = log

	fmt.Println("START RUN JAVA")             //for debug
	if err = subProcess.Start(); err != nil { //Use start, not run
		panic(err) //replace with logger, or anything you want
	}

	io.WriteString(stdin, testInput)
	subProcess.Wait()
	fmt.Println("END RUN JAVA") //for debug

}

func Run(timeout int, testInput string, command string, args ...string) {

	// instantiate new command
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
	case err := <-done:
		if err != nil {
			close(done)
			panic("process done, with error: " + err.Error())
		}
		fmt.Println("END RUN JAVA")
	}
	fmt.Println(buf.String())
}
