package util

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type TestCase struct {
	Name string     `yaml:"name"`
	Data [][]string `yaml:"data"`
}

func FetchTestCase(addr string) (name string, data [][]string) {
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
	data = t.Data
	return name, data
}
