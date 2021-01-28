package initialize

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Compile(folderName string) {
	subProcess := exec.Command("javac", folderName+"/src/*.java")
	stdin, err := subProcess.StdinPipe()
	if err != nil {
		fmt.Println(err) //replace with logger, or anything you want
	}
	defer stdin.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line
	subProcess.Stdout = os.Stdout
	subProcess.Stderr = os.Stderr
	fmt.Println("START COMPILE")              //for debug
	if err = subProcess.Start(); err != nil { //Use start, not run
		fmt.Println("An error occured: ", err) //replace with logger, or anything you want
	}
	io.WriteString(stdin, "QUIT\n")
	subProcess.Wait()
	fmt.Println("END COMPILE") //for debug
}
