package main

import (
	"errors"
	"fmt"
	_ "time"
)

//使用defer + recover 来捕获和处理异常
func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err =", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res =", res)
}

//自定义错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}

func test2() {
	err := readConf("config2.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test2()继续执行")
}

func main() {
	// test()
	// for {
	// 	fmt.Println("main()函数下的输出")
	// 	time.Sleep(time.Second)
	// }
	test2()
	fmt.Println("main()的执行")
}
