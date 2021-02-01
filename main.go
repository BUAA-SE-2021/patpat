package main

import (
	"flag"
	"fmt"
	"patpat/config"
	"patpat/initialize"
	"patpat/judge"
	"patpat/util"
	"strconv"
	"strings"
)

func main() {
	initialize.InitMySQL()
	judgePtr := flag.String("judge", "0-12345-hanhan", "Please specify the name of the folder containing Java files to judge")
	flag.Parse()
	folderName := *judgePtr
	tmpParamList := strings.Split(folderName, "-")
	num, err := strconv.Atoi(tmpParamList[0])
	if err != nil {
		panic("Cannot parse num!")
	}
	sid, err := strconv.Atoi(tmpParamList[1])
	if err != nil {
		panic("Cannot parse sid!")
	}
	name := tmpParamList[2]
	fmt.Println("Lab:", num, "SID:", sid, "Name:", name)
	tests := config.FetchJudgeConfig("test/judge.yaml")
	fmt.Println("Test cases:", tests)
	for _, t := range tests {
		testName, testData := util.FetchTestCase("test/" + t)
		fmt.Println(testName)
		testInputList, testInput, testOutputLines, testOutput, mapTable := util.ParseData(testData)
		initialize.CompileJava("javac", folderName+"/src/*.java")
		runStatus, actualOutput, actualOutputLines := initialize.RunJava(2, testInput, "java", "-classpath", folderName+"/src", "Test")
		compareResult, smallerLen, wrongOutputPos := judge.Compare(testOutputLines, actualOutputLines, mapTable)
		judge.ReportGen(t[0:len(tests[0])-5], runStatus, compareResult, smallerLen, wrongOutputPos, testInputList, testOutputLines, actualOutputLines, testOutput, actualOutput)
		judge.GradeUpload(num, sid, name, t, judge.CalcGrade(runStatus, compareResult))
	}
}
