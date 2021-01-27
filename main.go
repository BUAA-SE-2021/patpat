package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	// cmdPtr := flag.String("cmd", "pwd", "command to run Java file")
	// testPtr := flag.String("test", "zoo", "specify test file")
	// flag.Parse()
	// fmt.Println("cmd:", *cmdPtr)
	// fmt.Println("test:", *testPtr)
	// fin, _ := os.Open("in.txt")
	// fans, _ := os.Open("ans.txt")
	// cin, _ := ioutil.ReadAll(fin)
	// cans, _ := ioutil.ReadAll(fans)
	// fmt.Println(string(cin), string(cans))
	// config.FetchConfig()

	f, _ := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
	os.Stdout = f
	os.Stderr = f
	subProcess := exec.Command("java", "-classpath", "src", "Test") //Just for testing, replace with your subProcess

	stdin, err := subProcess.StdinPipe()
	if err != nil {
		fmt.Println(err) //replace with logger, or anything you want
	}
	defer stdin.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line

	// subProcess.Stdout = os.Stdout
	// subProcess.Stderr = os.Stderr
	subProcess.Stdout = f
	subProcess.Stderr = f

	fmt.Println("START COMPILE")              //for debug
	if err = subProcess.Start(); err != nil { //Use start, not run
		fmt.Println("An error occured: ", err) //replace with logger, or anything you want
	}

	io.WriteString(stdin, "QUIT\n")
	subProcess.Wait()
	fmt.Println("END COMPILE") //for debug
}
