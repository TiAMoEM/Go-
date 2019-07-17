package main

import (
	"fmt"
	"os"
)

// 命令行输入 go run main.go xxx xxx xxx

func main() {
	fmt.Println("命令行的参数有", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%v] = %v\n", i, v)
	}
}
