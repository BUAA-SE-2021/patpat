package main

import (
	"patpat/config"
)

func main() {
	// cmdPtr := flag.String("cmd", "pwd", "command to run Java file")
	// testPtr := flag.String("test", "zoo", "specify test file")
	// flag.Parse()
	// fmt.Println("cmd:", *cmdPtr)
	// fmt.Println("test:", *testPtr)
	// fin, _ := os.Open("in.txt")
	// fans, _ := os.Open("ans.txt")
	// cin, _ := ioutil.ReadAll(fin)
	// cans, _ := ioutil.ReadAll(fans)
	// fmt.Println(string(cin), string(cans))
	config.FetchConfig()
}
