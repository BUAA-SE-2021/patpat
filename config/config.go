package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

// var data = `
// num: 1 # 第几次作业
// sid: 18373722 # 学号
// name: "朱英豪" # 姓名
// test: ["testfile1.yaml","testfile2.yaml"] # 可有更多，这是个列表
// `

type Config struct {
	Num  int      `yaml:"num"`
	Sid  int      `yaml:"sid"`
	Name string   `yaml:"name"`
	Test []string `yaml:"test"`
}

func FetchConfig(num string, sid string, name string) (test []string) {
	t := Config{}

	judgeFile := num + "-" + sid + "-" + name
	fin, err := os.Open(judgeFile + "/" + "judge.yaml")
	if err != nil {
		fmt.Println(err)
	}
	cin, _ := ioutil.ReadAll(fin)
	err = yaml.Unmarshal([]byte(cin), &t)
	if err != nil {
		fmt.Println(err)
	}

	if num != strconv.Itoa(t.Num) || sid != strconv.Itoa(t.Sid) || name != t.Name {
		fmt.Println("与judge.yaml中的参数不匹配，请检查配置！")
	}
	if len(t.Test) == 0 {
		fmt.Println("没有测试文件的配置！")
	}

	test = t.Test
	return test
}
