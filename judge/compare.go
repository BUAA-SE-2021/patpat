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

	// 2023/10/17 TS:
	// If the actual output only different from the standard output in line number,
	// then this -1 (wrongOutputPos) will cause runtime error (index out of range [-1])
	// when we try to print the wrong output line.
	// The solution is to check this -1 in report generation. :)

	return compareResult, smallerLen, -1
}
