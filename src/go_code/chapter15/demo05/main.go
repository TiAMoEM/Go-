package main

import (
	"fmt"
	"os"
)

func main() {
	path := "d:\\test.txt"
	_, err := os.Stat(path)
	if err == nil {
		fmt.Println("is exist")
	}
	if os.IsNotExist(err) {
		fmt.Println("not exist")
	}
}
