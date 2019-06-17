package main

import (
	"fmt"
)

func main() {
	var intArr [3]int
	fmt.Println(intArr)
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)
	fmt.Printf("intArr的地址=%p\n, intArr[0]的地址是%p\n, intArr[1]的地址是%p\n, intArr[2]的地址是%p\n", &intArr, &intArr[0], &intArr[1], &intArr[2])

	//四种初始化数组的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01 =", numArr01)

	var numArr02 = [3]int{5, 6, 7}
	fmt.Println("numArr02 =", numArr02)

	var numArr03 = [...]int{8, 9, 10}
	fmt.Println("numArr03 =", numArr03)

	var numArr04 = [...]int{1: 800, 0: 900, 2: 999}
	fmt.Println("numArr04 =", numArr04)

	strArr05 := [...]string{1: "tom", 0: "jack", 2: "mary"}
	fmt.Println("strArr05 =", strArr05)
}
