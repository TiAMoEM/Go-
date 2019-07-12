package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "E:\\test\\abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err = %v", err)
		return
	}
	defer file.Close()
	str := "Hello World\n\r"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
