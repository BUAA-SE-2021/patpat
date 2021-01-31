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
		if strings.Compare(strings.TrimRight(testOutputLines[i], "\r\n"), strings.TrimRight(actualOutputLines[i], "\r\n")) != 0 {
			// fmt.Println(strings.Compare(strings.TrimRight(testOutputLines[i], "\r\n"), strings.TrimRight(actualOutputLines[i], "\r\n")))
			compareResult = mapTable[i] // 输出第几个输入时，出现错误
			// fmt.Println("Error\n", len(strings.TrimRight(testOutputLines[i], "\r\n")), "\n", len(strings.TrimRight(actualOutputLines[i], "\r\n")))
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
