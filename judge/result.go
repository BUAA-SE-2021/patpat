package judge

import (
	"io/ioutil"
	"patpat/global"
	"patpat/model"
)

func ReportGen(reportName string, runStatus int, compareResult int, smallerLen int, wrongOutputPos int, testInputList []string, testOutputLines []string, actualOutputLines []string) {
	content := "# " + reportName + " 评测情况\n\n" + "## 通过情况\n\n"
	if runStatus == 0 && compareResult == -3 {
		content += "Congratulations, AC!\n"
	} else if runStatus == 1 && compareResult == -3 {
		content += "TLE，输出结果正确\n"
	} else {
		if runStatus == 0 {
			content += "WA，输出结果错误\n\n"
		} else if runStatus == 1 {
			content += "TLE，输出结果错误\n\n"
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
			content += "### 期望输出\n\n```java\n" + testOutputLines[wrongOutputPos] + "\n```\n\n"
			content += "### 实际输出\n\n```java\n" + actualOutputLines[wrongOutputPos] + "\n```\n\n"
		}

		content += "## 请复制以下行辅助调试\n\n```java\n"
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
	if err := ioutil.WriteFile(reportName+"_result"+".md", []byte(content), 0644); err != nil {
		panic(err)
	}
}

func CalcGrade(runStatus int, compareResult int) (result int) {
	// AC 0
	// TLE 输出正确 1
	// WA(TLE) 输出错误 2
	if runStatus == 0 && compareResult == -3 {
		result = 0
	} else if runStatus == 1 && compareResult == -3 {
		result = 1
	} else {
		result = 2
	}
	return result
}

func GradeUpload(num int, sid int, name string, test string, result int) {
	judgeResult := model.JudgeResultUsual{Num: num, Sid: sid, Name: name, Test: test, Result: result}
	// fmt.Println(judgeResult)
	global.DB.Create(&judgeResult)
}
