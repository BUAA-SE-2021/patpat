package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Num   int      `yaml:"num"`
	Sid   int      `yaml:"sid"`
	Name  string   `yaml:"name"`
	Tests []string `yaml:"tests"`
}

func FetchConfig(num int, sid int, name string, folderName string) (test []string) {
	t := Config{}
	fin, err := os.Open(folderName + "/" + "judge.yaml")
	if err != nil {
		panic(err)
	}
	cin, _ := ioutil.ReadAll(fin)
	err = yaml.Unmarshal([]byte(cin), &t)
	if err != nil {
		panic(err)
	}

	if num != t.Num || sid != t.Sid || name != t.Name {
		panic("Inconsistent params! Please check your judge.yaml")
	}
	if len(t.Tests) == 0 {
		panic("No test cases!")
	}

	test = t.Tests
	return test
}
