package config

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
cmd: ["echo a", "echo b"]
test: ["echo 1","echo 2"]
`

type Config struct {
	Cmd  []string `yaml:"cmd"`
	Test []string `yaml:"test"`
}

func FetchConfig() {
	t := Config{}
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
	fmt.Printf("------------------------------\n")
	fmt.Println(len(t.Cmd))
	fmt.Println(t.Cmd[0])
}
