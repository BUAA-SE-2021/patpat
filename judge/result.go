package judge

import (
	"os"
	"patpat/global"
	"patpat/model"
)

func ReportGen(reportName string, runStatus int, compareResult int, smallerLen int, wrongOutputPos int, testInputList []string, testOutputLines []string, actualOutputLines []string, testOutput string, actualOutput string) {
	content := "# " + reportName + " 评测情况\n\n" + "## 通过情况\n\n"
	if runStatus == 0 && compareResult == -3 {
		content += "Congratulations, AC!\n"
	} else if runStatus == 1 && compareResult == -3 {
		content += "TLE，输出结果正确\n"
	} else if runStatus == 2 && compareResult == -3 {
		content += "RE，输出结果正确\n"
	} else {
		if runStatus == 0 {
			content += "WA，输出结果错误\n\n"
		} else if runStatus == 1 {
			content += "TLE，输出结果错误\n\n"
		} else if runStatus == 2 {
			content += "RE，输出结果错误\n\n"
		}
		content += "## 输出比较\n\n"
		if compareResult == -1 {
			content += "实际输出行数 < 期望输出行数。多余输出如下：\n\n```java\n"
			lenTest := len(testOutputLines)
			i := smallerLen
			for i < lenTest {
				content += testOutputLines[i] + "\n"
				i++
			}
			content += "```\n\n"
		} else if compareResult == -2 {
			content += "实际输出行数 > 期望输出行数。多余输出如下：\n\n```java\n"
			lenActual := len(actualOutputLines)
			i := smallerLen
			for i < lenActual {
				content += actualOutputLines[i] + "\n"
				i++
			}
			content += "```\n\n"
		} else {
			content += "### 期望输出行\n\n```java\n" + testOutputLines[wrongOutputPos] + "\n```\n\n"
			content += "### 实际输出行\n\n```java\n" + actualOutputLines[wrongOutputPos] + "\n```\n\n"
		}

		content += "## 请复制以下行辅助调试\n\n注：由于可能存在某些命令无输出，所以定位不一定完全准确\n\n```java\n"
		if compareResult == -1 || compareResult == -2 {
			for _, v := range testInputList {
				content += v + "\n"
			}
			content += "```\n"
		} else {
			i := 0
			for i <= compareResult {
				content += testInputList[i] + "\n"
				i++
			}
			content += "```\n"
		}
	}
	content += "\n## 更多信息\n\n### 完整期望输出\n\n```java\n" + testOutput + "```\n\n"
	content += "### 完整实际输出\n\n```java\n" + actualOutput + "```\n"
	if err := os.WriteFile(reportName+"_result"+".md", []byte(content), 0644); err != nil {
		panic(err)
	}
}

func TaJudgeReportGen(reportName string, runStatus int, compareResult int, smallerLen int, wrongOutputPos int, testInputList []string, testOutputLines []string, actualOutputLines []string, testOutput string, actualOutput string) {
	content := "# " + reportName + " 评测情况\n\n" + "## 通过情况\n\n"
	if runStatus == 0 && compareResult == -3 {
		content += "Congratulations, AC!\n"
	} else if runStatus == 1 && compareResult == -3 {
		content += "TLE，输出结果正确\n"
	} else if runStatus == 2 && compareResult == -3 {
		content += "RE，输出结果正确\n"
	} else {
		if runStatus == 0 {
			content += "WA，输出结果错误\n\n"
		} else if runStatus == 1 {
			content += "TLE，输出结果错误\n\n"
		} else if runStatus == 2 {
			content += "RE，输出结果错误\n\n"
		}
		content += "## 输出比较\n\n"
		if compareResult == -1 {
			content += "实际输出行数 < 期望输出行数。\n\n"
		} else if compareResult == -2 {
			content += "实际输出行数 > 期望输出行数。\n\n"
		} else {
			content += "### 期望输出行\n\n```java\n" + testOutputLines[wrongOutputPos] + "\n```\n\n"
			content += "### 实际输出行\n\n```java\n" + actualOutputLines[wrongOutputPos] + "\n```\n"
		}
	}
	if err := os.WriteFile(reportName+"_result"+".md", []byte(content), 0644); err != nil {
		panic(err)
	}
}

func CalcGrade(runStatus int, compareResult int) (result int) {
	// AC 完全正确 1
	// TLE 超时 -1
	// WA 输出错误 -2
	// CE 编译错误 -3
	// RE 运行时错误 -4
	if runStatus == 0 && compareResult == -3 {
		result = 1
	} else if runStatus == 1 {
		result = -1
	} else if runStatus == 2 {
		result = -4
	} else {
		result = -2
	}
	return result
}

func GradeUpload(num int, sid int, name string, test string, result int) {
	judgeResult := model.JudgeResultUsual{Num: num, Sid: sid, Name: name, Test: test, Result: result}
	// fmt.Println(judgeResult)
	global.DB.Create(&judgeResult)
}

func GradeUploadFormal(num int, sid int, name string, test string, result int, tag string) {
	judgeResult := model.JudgeResultFormal{Num: num, Sid: sid, Name: name, Test: test, Result: result, Tag: tag}
	global.DB.Create(&judgeResult)
}
