package judge

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Compare(testOutputLines []string, actualOutput string, mapTable []int) (result bool, wrongPos int) {
	actualOutputLines := strings.Split(strings.TrimRight(actualOutput, "\r\n"), "\n")

	actual := ""
	for _, v := range actualOutputLines {
		actual = actual + v + "\n"
	}
	content := []byte(actual)
	if err := ioutil.WriteFile("actualOutput2.txt", content, 0644); err != nil {
		panic(err)
	}

	lenTest := len(testOutputLines)
	lenActual := len(actualOutputLines)
	fmt.Println(lenTest, lenActual)
	var minLen int
	if lenTest < lenActual {
		minLen = lenTest
	} else {
		minLen = lenActual
	}
	result = true
	i := 0
	for i < minLen {
		if strings.Compare(strings.TrimRight(testOutputLines[i], "\r\n"), strings.TrimRight(actualOutputLines[i], "\r\n")) != 0 {
			fmt.Println(strings.Compare(strings.TrimRight(testOutputLines[i], "\r\n"), strings.TrimRight(actualOutputLines[i], "\r\n")))
			result = false
			wrongPos = mapTable[i]
			fmt.Println("Error\n", len(strings.TrimRight(testOutputLines[i], "\r\n")), "\n", len(strings.TrimRight(actualOutputLines[i], "\r\n")))
			break
		}
		i++
	}
	if lenActual < lenTest {
		result = false
		wrongPos = -1
	} else if lenActual > lenTest {
		result = false
		wrongPos = -2
	}
	return result, wrongPos
}
