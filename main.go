package main

import (
	"flag"
	"fmt"
	"patpat/config"
	"patpat/initialize"
	"patpat/util"
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

	testName, testData := util.FetchTestCase("test/" + tests[0])
	fmt.Println(testName)
	// fmt.Println(testData)
	testIn := ""
	for _, v := range testData {
		testIn = testIn + v[0] + "\n"
	}
	fmt.Println(testIn)

	initialize.CompileJava("javac", folderName+"/src/*.java")
	actualOutput := initialize.RunJava(2, testIn, "java", "-classpath", folderName+"/src", "Test")
	fmt.Print(actualOutput)
}
