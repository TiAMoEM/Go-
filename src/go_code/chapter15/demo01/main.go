package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("D:\\Chrome下载\\resume.html")
	if err != nil {
		fmt.Println("open file err =", err)
	}
	fmt.Printf("file = %v", file)
	err = file.Close()
	if err != nil {
		fmt.Println("close file err =", err)
	}
}
