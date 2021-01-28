package initialize

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Execute(folderName string) {
	log, _ := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
	// os.Stdout = log
	// os.Stderr = log
	subProcess := exec.Command("java", "-classpath", folderName+"/src", "Test")
	stdin, err := subProcess.StdinPipe()
	if err != nil {
		fmt.Println(err) //replace with logger, or anything you want
	}
	defer stdin.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line
	subProcess.Stdout = log
	subProcess.Stderr = log
	fmt.Println("START RUN JAVA")             //for debug
	if err = subProcess.Start(); err != nil { //Use start, not run
		fmt.Println("An error occured: ", err) //replace with logger, or anything you want
	}
	io.WriteString(stdin, "QUIT\n")
	subProcess.Wait()
	fmt.Println("END RUN JAVA") //for debug
}
