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

	testName, testData := util.FetchTestCase("test/" + tests[0])
	fmt.Println(testName)

	testInputList, testInput, testOutputLines, mapTable := util.ParseData(testData)

	initialize.CompileJava("javac", folderName+"/src/*.java")
	actualOutput := initialize.RunJava(2, testInput, "java", "-classpath", folderName+"/src", "Test")

	result, wrongPos := judge.Compare(testOutputLines, actualOutput, mapTable)
	fmt.Println(result, wrongPos)
	fmt.Println(len(testInputList), len(mapTable), len(testOutputLines), len(actualOutput))

	content := []byte(actualOutput)
	if err := ioutil.WriteFile("actualOutput.txt", content, 0644); err != nil {
		panic(err)
	}

}
