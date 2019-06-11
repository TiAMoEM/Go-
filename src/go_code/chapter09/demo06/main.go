package main

import (
	"fmt"
	"strings"
)

/*
1) AddUpper()是一个函数，返回的数据类型是 func (int) int
2) 返回的是一个匿名函数，但是这个匿名函数会引用到函数外的n，这个匿名函数和n共同形成一个整体，称为闭包

*/

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f2("winter"))
	fmt.Println("文件名处理后=", f2("bird.jpg"))

}

//例子

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}

}
