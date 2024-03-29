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
	// stuPkgPtr := stuCmd.Bool("pkg", false, "Use package or not.")

	taCmd := flag.NewFlagSet("ta", flag.ExitOnError)
	taJudgePtr := taCmd.String("judge", "0-12345-hanhan", "Please specify the name of the folder containing Java files to judge.")
	taPwdPtr := taCmd.String("pwd", "888888", "Please enter your password.")
	tagPtr := taCmd.String("tag", "test", "Tag for this judge.")
	// taPkgPtr := taCmd.Bool("pkg", false, "Use package or not.")

	regCmd := flag.NewFlagSet("reg", flag.ExitOnError)
	regSidPtr := regCmd.Int("sid", 123456, "Please specify your SID.")
	regPwdPtr := regCmd.String("pwd", "888888", "Please enter your password.")

	queCmd := flag.NewFlagSet("query", flag.ExitOnError)
	queSidPtr := queCmd.Int("sid", 123456, "Please specify your SID.")
	quePwdPtr := queCmd.String("pwd", "888888", "Please enter your password.")

	switch os.Args[1] {
	case "stu":
		stuCmd.Parse(os.Args[2:])
		if *onlineModePtr {
			initialize.InitMySQL()
		}
		folderName := *judgePtr
		paramList := strings.Split(folderName, "-")
		num, err := strconv.Atoi(paramList[0])
		if err != nil {
			fmt.Println("请确保文件夹命名为 'num-sid-name' 的格式!")
			return
			// panic("Cannot parse num!")
		}
		sid, err := strconv.Atoi(paramList[1])
		if err != nil {
			fmt.Println("请确保文件夹命名为 'num-sid-name' 的格式!")
			return
			// panic("Cannot parse sid!")
		}
		name := paramList[2]
		fmt.Println("Lab: " + strconv.Itoa(num) + " | SID: " + strconv.Itoa(sid) + " | Name: " + name)
		tests := initialize.FetchJudgeConfig("test/judge.yaml")
		fmt.Println("Test cases:")
		for _, t := range tests {
			fmt.Println("\t" + t)
		}

		var exitCode int
		switch goos {
		case "windows":
			// if *stuPkgPtr {
			// 	exitCode = run.CompileJava("javac", "-encoding", "UTF-8", "-cp", "./"+folderName+"/src", "-d", "./"+folderName+"/out", "./"+folderName+"/src"+"/Test.java")
			// } else {
			// 	exitCode = run.CompileJava("javac", "-encoding", "UTF-8", folderName+"/src/*.java")
			// }
			exitCode = run.CompileJava("javac", "-encoding", "UTF-8", "-cp", "./"+folderName+"/src", "-d", "./"+folderName+"/out", "./"+folderName+"/src"+"/*.java")
		case "darwin", "linux":
			// if *stuPkgPtr {
			// 	exitCode = run.CompileJava("javac", "-encoding", "UTF-8", "-cp", "./"+folderName+"/src", "-d", "./"+folderName+"/out", "./"+folderName+"/src"+"/Test.java")
			// } else {
			// 	exitCode = run.CompileJava("/bin/sh", "-c", "javac -encoding UTF-8 "+folderName+"/src/*.java")
			// }
			exitCode = run.CompileJava("/bin/sh", "-c", "javac -encoding UTF-8 -cp ./"+folderName+"/src -d ./"+folderName+"/out ./"+folderName+"/src/*.java")
		}
		if exitCode != 0 {
			fmt.Println("Compile Error!")
			if *onlineModePtr {
				judge.GradeUpload(num, sid, name, "testcase", -3)
			}
		} else {
			fmt.Println(strings.Repeat("_", 80))
			fmt.Println()
			fmt.Println("您本次自测的评测情况如下: ")
			for _, t := range tests {
				fmt.Println()

				testName, testData, err := initialize.FetchTestCase("test/" + t)
				if err != nil {
					fmt.Println("Cannot fetch test case: " + t)
					fmt.Println("Caused by:")
					fmt.Println(err)
					continue
				}

				PrintBanner(testName)

				testInputList, testInput, testOutputLines, testOutput, mapTable := initialize.ParseTestData(testData)
				// var runStatus int
				// var actualOutput string
				// var actualOutputLines []string
				// if *stuPkgPtr {
				// 	runStatus, actualOutput, actualOutputLines = run.RunJava(2, testInput, "java", "-classpath", "./"+folderName+"/out", "Test")
				// } else {
				// 	runStatus, actualOutput, actualOutputLines = run.RunJava(2, testInput, "java", "-classpath", folderName+"/src", "Test")
				// }
				runStatus, actualOutput, actualOutputLines := run.RunJava(2, testInput, "java", "-classpath", "./"+folderName+"/out", "Test")
				compareResult, smallerLen, wrongOutputPos := judge.Compare(testOutputLines, actualOutputLines, mapTable)
				resultMessage := "Num = " + strconv.Itoa(num) + ", 评测点 = " + t[0:len(t)-5] + ", Grade = " + GradeToString(judge.CalcGrade(runStatus, compareResult))
				fmt.Println(resultMessage)
				judge.ReportGen(t[0:len(t)-5], runStatus, compareResult, smallerLen, wrongOutputPos, testInputList, testOutputLines, actualOutputLines, testOutput, actualOutput)
				if *onlineModePtr {
					judge.GradeUpload(num, sid, name, t[0:len(t)-5], judge.CalcGrade(runStatus, compareResult))
				}
			}
		}
	case "ta":
		initialize.InitMySQL()
		taCmd.Parse(os.Args[2:])
		folderName := *taJudgePtr
		pwd := *taPwdPtr
		paramList := strings.Split(folderName, "-")
		num, err := strconv.Atoi(paramList[0])
		if err != nil {
			fmt.Println("请确保文件夹命名为 'num-sid-name' 的格式!")
			return
			// panic("Cannot parse num!")
		}
		sid, err := strconv.Atoi(paramList[1])
		if err != nil {
			fmt.Println("请确保文件夹命名为 'num-sid-name' 的格式!")
			return
			// panic("Cannot parse sid!")
		}
		name := paramList[2]
		fmt.Println("Lab: " + strconv.Itoa(num) + " | SID: " + strconv.Itoa(sid) + " | Name: " + name)
		if v1.Login(sid, pwd) {
			formalTestCases := initialize.FetchFormalTestCase(num)
			fmt.Println("Test cases:")
			for _, formalTestCase := range formalTestCases {
				fmt.Println("\t" + formalTestCase.FileName)
			}

			fmt.Println()
			var exitCode int
			switch goos {
			case "windows":
				// if *taPkgPtr {
				// 	exitCode = run.CompileJava("javac", "-encoding", "UTF-8", "-cp", "./"+folderName+"/src", "-d", "./"+folderName+"/out", "./"+folderName+"/src"+"/Test.java")
				// } else {
				// 	exitCode = run.CompileJava("javac", "-encoding", "UTF-8", folderName+"/src/*.java")
				// }
				exitCode = run.CompileJava("javac", "-encoding", "UTF-8", "-cp", "./"+folderName+"/src", "-d", "./"+folderName+"/out", "./"+folderName+"/src"+"/*.java")
			case "darwin", "linux":
				// if *taPkgPtr {
				// 	exitCode = run.CompileJava("javac", "-encoding", "UTF-8", "-cp", "./"+folderName+"/src", "-d", "./"+folderName+"/out", "./"+folderName+"/src"+"/Test.java")
				// } else {
				// 	exitCode = run.CompileJava("/bin/sh", "-c", "javac -encoding UTF-8 "+folderName+"/src/*.java")
				// }
				exitCode = run.CompileJava("/bin/sh", "-c", "javac -encoding UTF-8 -cp ./"+folderName+"/src -d ./"+folderName+"/out ./"+folderName+"/src/*.java")
			}
			// exitCode := run.CompileJava("javac", "-encoding", "UTF-8", "-cp", "./"+folderName+"/src", "-d", "./"+folderName+"/out", "./"+folderName+"/src"+"/Test.java")
			if exitCode != 0 {
				fmt.Println("Compile Error!")
				judge.GradeUploadFormal(num, sid, name, "testcase", -3, *tagPtr)
			} else {
				fmt.Println(strings.Repeat("_", 80))
				fmt.Println()
				fmt.Println("您本次黑盒测试的评测情况如下：")
				for _, t := range formalTestCases {
					_, _, testName, testData := initialize.ParseFormalTestCase(t)
					// testName, testData := initialize.FetchTestCase("test/" + t)

					fmt.Println()
					PrintBanner(testName)

					testInputList, testInput, testOutputLines, testOutput, mapTable := initialize.ParseTestData(testData)
					// var runStatus int
					// var actualOutput string
					// var actualOutputLines []string
					// if *taPkgPtr {
					// 	runStatus, actualOutput, actualOutputLines = run.RunJava(2, testInput, "java", "-classpath", "./"+folderName+"/out", "Test")
					// } else {
					// 	runStatus, actualOutput, actualOutputLines = run.RunJava(2, testInput, "java", "-classpath", folderName+"/src", "Test")
					// }
					runStatus, actualOutput, actualOutputLines := run.RunJava(2, testInput, "java", "-classpath", "./"+folderName+"/out", "Test")
					compareResult, smallerLen, wrongOutputPos := judge.Compare(testOutputLines, actualOutputLines, mapTable)
					resultMessage := "Num = " + strconv.Itoa(num) + ", 评测点 = " + t.FileName + ", Grade = " + GradeToString(judge.CalcGrade(runStatus, compareResult))
					fmt.Println(resultMessage)
					judge.TaJudgeReportGen(t.FileName, runStatus, compareResult, smallerLen, wrongOutputPos, testInputList, testOutputLines, actualOutputLines, testOutput, actualOutput)
					judge.GradeUploadFormal(num, sid, name, t.FileName, judge.CalcGrade(runStatus, compareResult), *tagPtr)
				}
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
			if err := os.WriteFile("ta_judge_result.txt", []byte(result), 0644); err != nil {
				panic(err)
			}
		}
	default:
		fmt.Println("Expected 'stu' or 'ta' or 'reg' or 'query' subcommands!")
	}

	fmt.Println()
}

func PrintBanner(title string) {
	extra := int((80 - 2 - len(title)) / 2)
	if extra < 0 {
		extra = 0
	}

	fmt.Print(strings.Repeat("=", extra))
	fmt.Print(" \033[33m" + title + "\033[0m ")
	fmt.Println(strings.Repeat("=", 80-extra-2-len(title)))
}

// (AC: 1, TLE: -1, WA: -2, CE: -3, RE: -4)
func GradeToString(grade int) string {
	switch grade {
	case 1:
		// Green
		return "\033[92mAC\033[0m"
	case -1:
		// Blue
		return "\033[36mTLE\033[0m"
	case -2:
		// Red
		return "\033[91mWA\033[0m"
	case -3:
		// Yellow
		return "\033[33mCE\033[0m"
	case -4:
		// Magenta
		return "\033[95mRE\033[0m"
	default:
		// Skyblue
		return "\033[96mN/A\033[0m"
	}
}
