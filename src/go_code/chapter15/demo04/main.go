package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//将一个文件的内容，写入到另一个文件
	filePath1 := "E:\\test\\test.py"
	filePath2 := "E:\\test\\abc.py"
	data, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Printf("read file error = %v", err)
		return
	}
	err = ioutil.WriteFile(filePath2, data, 0666)
	if err != nil {
		fmt.Printf("write file error = %v", err)
		return
	}
}
