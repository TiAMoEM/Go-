package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//方式一
	file, err := os.Open("D:\\Chrome下载\\resume.html")
	if err != nil {
		fmt.Println("open file err =", err)
	}
	defer file.Close()
	const (
		defaultBufSize = 4096
	)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")

	//方式二，该方法适用于文件不大的情况
	//使用ioutil一次将整个文件读入到内存中
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err = %v", err)
	}
	fmt.Printf("%v", string(content))
}
