package initialize

import (
	"io"
	"os"
	"patpat/global"
	"strings"

	"gopkg.in/yaml.v3"
)

type JudgeConfig struct {
	Num   int      `yaml:"num"`
	Sid   int      `yaml:"sid"`
	Name  string   `yaml:"name"`
	Tests []string `yaml:"tests"`
}

type TestCase struct {
	Name string     `yaml:"name"`
	Data [][]string `yaml:"data"`
}

type FormalTestCase struct {
	Num      int    `gorm:"num"`
	FileName string `gorm:"file_name"`
	Yaml     string `gorm:"yaml"`
}

type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func FetchJudgeConfig(configAddr string) (tests []string) {
	t := JudgeConfig{}
	fin, err := os.Open(configAddr)
	if err != nil {
		panic(err)
	}
	cin, _ := io.ReadAll(fin)
	err = yaml.Unmarshal([]byte(cin), &t)
	if err != nil {
		panic(err)
	}

	if len(t.Tests) == 0 {
		panic("No test cases!")
	}

	tests = t.Tests
	return tests
}

func FetchFormalTestCase(num int) (formalTestCases []FormalTestCase) {
	global.DB.Where("num = ?", num).Find(&formalTestCases)
	return formalTestCases
}

func ParseFormalTestCase(formalTestCase FormalTestCase) (num int, fileName string, name string, testData *[][]string) {
	t := TestCase{}
	err := yaml.Unmarshal([]byte(formalTestCase.Yaml), &t)
	if err != nil {
		panic(err)
	}
	num = formalTestCase.Num
	fileName = formalTestCase.FileName
	name = t.Name
	testData = &t.Data
	return num, fileName, name, testData
}

func FetchTestCase(addr string) (name string, testData *[][]string) {
	t := TestCase{}

	fin, err := os.Open(addr)
	if err != nil {
		panic(err)
	}
	cin, _ := io.ReadAll(fin)
	err = yaml.Unmarshal([]byte(cin), &t)
	if err != nil {
		panic(err)
	}

	name = t.Name
	testData = &t.Data
	return name, testData
}

func ParseTestData(testData *[][]string) (testInputList []string, testInput string, testOutputLines []string, testOutput string, mapTable []int) {
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

func FetchMySQLConfig() (host string, port string, username string, password string, db string) {
	t := MySQLConfig{}
	fin, err := os.Open("mysql.yaml")
	if err != nil {
		panic(err)
	}
	cin, _ := io.ReadAll(fin)
	err = yaml.Unmarshal([]byte(cin), &t)
	if err != nil {
		panic(err)
	}
	return t.Host, t.Port, t.Username, t.Password, t.Database
}
