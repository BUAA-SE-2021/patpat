package config

import (
	"io/ioutil"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Num   int      `yaml:"num"`
	Sid   int      `yaml:"sid"`
	Name  string   `yaml:"name"`
	Tests []string `yaml:"tests"`
}

func FetchConfig(num string, sid string, name string) (test []string) {
	t := Config{}

	judgeFile := num + "-" + sid + "-" + name
	fin, err := os.Open(judgeFile + "/" + "judge.yaml")
	if err != nil {
		panic(err)
	}
	cin, _ := ioutil.ReadAll(fin)
	err = yaml.Unmarshal([]byte(cin), &t)
	if err != nil {
		panic(err)
	}

	if num != strconv.Itoa(t.Num) || sid != strconv.Itoa(t.Sid) || name != t.Name {
		panic("Inconsistent params! Please check your judge.yaml")
	}
	if len(t.Tests) == 0 {
		panic("No test cases!")
	}

	test = t.Tests
	return test
}
