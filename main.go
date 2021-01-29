package main

import (
	"flag"
	"fmt"
	"patpat/config"
	"patpat/initialize"
	"strings"
)

func main() {
	judgePtr := flag.String("judge", "0-12345-hanhan", "please specify the name of the folder containing Java files to judge")
	flag.Parse()
	folderName := *judgePtr
	tmpParamList := strings.Split(folderName, "-")
	num := tmpParamList[0]
	sid := tmpParamList[1]
	name := tmpParamList[2]
	fmt.Println("Lab:", num, "SID:", sid, "Name:", name)
	tests := config.FetchConfig(num, sid, name)
	fmt.Println("Test cases:", tests)

	initialize.Compile(folderName)
	// initialize.Execute(folderName, "QUIT\n")
	initialize.Run(2, "SUDO", "java", "-classpath", folderName+"/src", "Test")
}
