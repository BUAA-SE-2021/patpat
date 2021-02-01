package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type JudgeConfig struct {
	Num   int      `yaml:"num"`
	Sid   int      `yaml:"sid"`
	Name  string   `yaml:"name"`
	Tests []string `yaml:"tests"`
}

func FetchJudgeConfig(configAddr string) (tests []string) {
	t := JudgeConfig{}
	fin, err := os.Open(configAddr)
	if err != nil {
		panic(err)
	}
	cin, _ := ioutil.ReadAll(fin)
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
