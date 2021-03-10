package main

import (
	"flag"
	"fmt"
	"os"
	v1 "patpat/api/v1"
	"patpat/initialize"
	"patpat/judge"
	"patpat/run"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	goos := runtime.GOOS

	stuCmd := flag.NewFlagSet("stu", flag.ExitOnError)
	judgePtr := stuCmd.String("judge", "0-12345-hanhan", "Please specify the name of the folder containing Java files to judge.")
	onlineModePtr := stuCmd.Bool("online", true, "Online or offline mode")

	taCmd := flag.NewFlagSet("ta", flag.ExitOnError)
	taJudgePtr := taCmd.String("judge", "0-12345-hanhan", "Please specify the name of the folder containing Java files to judge.")
	tagPtr := taCmd.String("tag", "test", "Tag for this judge.")

	regCmd := flag.NewFlagSet("reg", flag.ExitOnError)
	regSidPtr := regCmd.Int("sid", 123456, "Please specify your SID.")
	regPwdPtr := regCmd.String("pwd", "888888", "Please enter your password.")

	queCmd := flag.NewFlagSet("query", flag.ExitOnError)
	queSidPtr := queCmd.Int("sid", 123456, "Please specify your SID.")
	quePwdPtr := queCmd.String("pwd", "888888", "Please enter your password.")

	switch os.Args[1] {
	case "stu":
		stuCmd.Parse(os.Args[2:])
		if *onlineModePtr == true {
			initialize.InitMySQL()
		}
		folderName := *judgePtr
		paramList := strings.Split(folderName, "-")
		num, err := strconv.Atoi(paramList[0])
		if err != nil {
			panic("Cannot parse num!")
		}
		sid, err := strconv.Atoi(paramList[1])
		if err != nil {
			panic("Cannot parse sid!")
		}
		name := paramList[2]
		fmt.Println("Lab:", num, "SID:", sid, "Name:", name)
		tests := initialize.FetchJudgeConfig("test/judge.yaml")
		fmt.Println("Test cases:", tests)
		var exitCode int
		switch goos {
		case "windows":
			exitCode = run.CompileJava("javac", "-encoding", "UTF-8", folderName+"/src/*.java")
		case "darwin", "linux":
			exitCode = run.CompileJava("/bin/sh", "-c", "javac -encoding UTF-8 "+folderName+"/src/*.java")
		}
		if exitCode != 0 {
			fmt.Println("Compile Error!")
			if *onlineModePtr == true {
				judge.GradeUpload(num, sid, name, "testcase", -3)
			}
		} else {
			fmt.Println("您本次自测的评测情况如下:")
			for _, t := range tests {
				testName, testData := initialize.FetchTestCase("test/" + t)
				fmt.Println(testName)
				testInputList, testInput, testOutputLines, testOutput, mapTable := initialize.ParseTestData(testData)
				runStatus, actualOutput, actualOutputLines := run.RunJava(2, testInput, "java", "-classpath", folderName+"/src", "Test")
				compareResult, smallerLen, wrongOutputPos := judge.Compare(testOutputLines, actualOutputLines, mapTable)
				resultMessage := "Num = " + strconv.Itoa(num) + ", 评测点 = " + t[0:len(t)-5] + ", Grade = " + strconv.Itoa(judge.CalcGrade(runStatus, compareResult))
				fmt.Println(resultMessage)
				judge.ReportGen(t[0:len(t)-5], runStatus, compareResult, smallerLen, wrongOutputPos, testInputList, testOutputLines, actualOutputLines, testOutput, actualOutput)
				if *onlineModePtr == true {
					judge.GradeUpload(num, sid, name, t[0:len(t)-5], judge.CalcGrade(runStatus, compareResult))
				}
			}
		}
	case "ta":
		initialize.InitMySQL()
		taCmd.Parse(os.Args[2:])
		folderName := *taJudgePtr
		paramList := strings.Split(folderName, "-")
		num, err := strconv.Atoi(paramList[0])
		if err != nil {
			panic("Cannot parse num!")
		}
		sid, err := strconv.Atoi(paramList[1])
		if err != nil {
			panic("Cannot parse sid!")
		}
		name := paramList[2]
		fmt.Println("Lab:", num, "SID:", sid, "Name:", name)
		tests := initialize.FetchJudgeConfig("test/judge.yaml")
		fmt.Println("Test cases:", tests)
		var exitCode int
		switch goos {
		case "windows":
			exitCode = run.CompileJava("javac", "-encoding", "UTF-8", strconv.Itoa(num)+folderName+"/src/*.java")
		case "darwin", "linux":
			exitCode = run.CompileJava("/bin/sh", "-c", "javac -encoding UTF-8 "+strconv.Itoa(num)+folderName+"/src/*.java")
		}
		if exitCode != 0 {
			fmt.Println("Compile Error!")
			judge.GradeUploadFormal(num, sid, name, "testcase", -3, *tagPtr)
		} else {
			fmt.Println("您本次自测的评测情况如下:")
			for _, t := range tests {
				testName, testData := initialize.FetchTestCase("test/" + t)
				fmt.Println(testName)
				_, testInput, testOutputLines, _, mapTable := initialize.ParseTestData(testData)
				runStatus, _, actualOutputLines := run.RunJava(2, testInput, "java", "-classpath", strconv.Itoa(num)+"/"+folderName+"/src", "Test")
				compareResult, _, _ := judge.Compare(testOutputLines, actualOutputLines, mapTable)
				resultMessage := "Num = " + strconv.Itoa(num) + ", 评测点 = " + t[0:len(t)-5] + ", Grade = " + strconv.Itoa(judge.CalcGrade(runStatus, compareResult))
				fmt.Println(resultMessage)
				judge.GradeUploadFormal(num, sid, name, t[0:len(t)-5], judge.CalcGrade(runStatus, compareResult), *tagPtr)
			}
		}
	case "reg":
		initialize.InitMySQL()
		regCmd.Parse(os.Args[2:])
		sid := *regSidPtr
		pwd := *regPwdPtr
		result := v1.Register(sid, pwd)
		fmt.Println(result)
	case "query":
		initialize.InitMySQL()
		queCmd.Parse(os.Args[2:])
		sid := *queSidPtr
		pwd := *quePwdPtr
		if v1.Login(sid, pwd) {
			result := v1.QueryResult(sid)
			fmt.Print(result)
		}
	default:
		fmt.Println("Expected 'stu' or 'reg' or 'query' subcommands!")
	}
}
