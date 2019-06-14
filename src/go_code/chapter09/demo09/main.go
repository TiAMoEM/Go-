package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "hello杯"
	fmt.Println("str len=", len(str))

	//字符串遍历
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("字符=%c, \n", str2[i])
	}

	//字符串转整数
	n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("error ", err)
	} else {
		fmt.Println("result ", n)
	}

	//整数转字符串
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v, str=%T", str, str)

	//字符串转[]byte
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	//[]byte转字符串
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

}
