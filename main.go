package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"patpat/config"
	"patpat/initialize"
	"patpat/judge"
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
	for _, t := range tests {
		testName, testData := util.FetchTestCase("test/" + t)
		fmt.Println(testName)

		testInputList, testInput, testOutputLines, mapTable := util.ParseData(testData)

		initialize.CompileJava("javac", folderName+"/src/*.java")
		runStatus, actualOutput, actualOutputLines := initialize.RunJava(2, testInput, "java", "-classpath", folderName+"/src", "Test")

		compareResult, smallerLen, wrongOutputPos := judge.Compare(testOutputLines, actualOutputLines, mapTable)

		content := []byte(actualOutput)
		if err := ioutil.WriteFile("actualOutput.txt", content, 0644); err != nil {
			panic(err)
		}

		judge.ReportGen(t[0:len(tests[0])-5]+"_result", runStatus, compareResult, smallerLen, wrongOutputPos, testInputList, testOutputLines, actualOutputLines)
	}

}
