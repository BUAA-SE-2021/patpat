package main

import (
	"flag"
	"fmt"
	"patpat/config"
	"patpat/initialize"
	"strings"
)

func main() {
	judgePtr := flag.String("judge", "1-12345-hanhan", "待测文件夹")
	flag.Parse()
	folderName := *judgePtr
	tmpParamList := strings.Split(folderName, "-")
	num := tmpParamList[0]
	sid := tmpParamList[1]
	name := tmpParamList[2]
	fmt.Println("Lab:", num, "学号:", sid, "姓名:", name)
	test := config.FetchConfig(num, sid, name)
	fmt.Println("测试文件:", test)

	initialize.Compile(folderName)
	initialize.Execute(folderName)
}
