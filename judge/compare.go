package judge

import (
	"strings"
)

func Compare(testOutputLines []string, actualOutputLines []string, mapTable []int) (compareResult int, smallerLen int, wrongOutputPos int) {
	compareResult = -3 // pass
	lenTest := len(testOutputLines)
	lenActual := len(actualOutputLines)
	// fmt.Println(lenTest, lenActual)
	if lenTest < lenActual {
		smallerLen = lenTest
	} else {
		smallerLen = lenActual
	}

	i := 0
	for i < smallerLen {
		if strings.Compare(strings.Trim(testOutputLines[i], "\r\n\t "), strings.Trim(actualOutputLines[i], "\r\n\t ")) != 0 {

			compareResult = mapTable[i] // 输出第几个输入时，出现错误

			return compareResult, smallerLen, i
		}
		i++
	}
	if lenActual < lenTest {
		compareResult = -1
	} else if lenActual > lenTest {
		compareResult = -2
	}
	return compareResult, smallerLen, -1
}
