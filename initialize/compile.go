package initialize

import (
	"fmt"
	"os"
	"os/exec"
)

func Compile(folderName string) {
	subProcess := exec.Command("javac", folderName+"/src/*.java")
	stdin, err := subProcess.StdinPipe()
	if err != nil {
		panic(err) //replace with logger, or anything you want
	}
	defer stdin.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line
	subProcess.Stdout = os.Stdout
	subProcess.Stderr = os.Stderr
	fmt.Println("START COMPILE")              //for debug
	if err = subProcess.Start(); err != nil { //Use start, not run
		panic(err) //replace with logger, or anything you want
	}
	subProcess.Wait()
	fmt.Println("END COMPILE") //for debug
}
