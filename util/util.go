package util

import (
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type TestCase struct {
	Name string     `yaml:"name"`
	Data [][]string `yaml:"data"`
}

func FetchTestCase(addr string) (name string, testData *[][]string) {
	t := TestCase{}

	fin, err := os.Open(addr)
	if err != nil {
		panic(err)
	}
	cin, _ := ioutil.ReadAll(fin)
	err = yaml.Unmarshal([]byte(cin), &t)
	if err != nil {
		panic(err)
	}

	name = t.Name
	testData = &t.Data
	return name, testData
}

func ParseData(testData *[][]string) (testInputList []string, testInput string, testOutputLines []string, testOutput string, mapTable []int) {
	linesInOutput := 0
	for i, v := range *testData {
		lenSinglePoint := len(v)
		if lenSinglePoint == 0 || lenSinglePoint > 2 {
			panic("Wrong Test Case Format!")
		}
		testInputList = append(testInputList, strings.TrimRight(v[0], "\r\n"))
		if lenSinglePoint == 2 {
			curOutput := strings.TrimRight(v[1], "\r\n")
			curOutputLines := strings.Split(curOutput, "\n")
			for _, s := range curOutputLines {
				testOutputLines = append(testOutputLines, s)
				mapTable = append(mapTable, i)
				linesInOutput++
			}
		}
	}
	testInput = ""
	for _, v := range testInputList {
		testInput = testInput + v + "\n"
	}
	testOutput = ""
	for _, v := range testOutputLines {
		testOutput = testOutput + v + "\n"
	}
	return testInputList, testInput, testOutputLines, testOutput, mapTable
}
