package main

import "fmt"

func main() {
	var age int
	fmt.Println("please input your age")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("111111")
	} else {
		fmt.Println("222222")
	}

	//判断两个数之和既能被3整除也能被5整除

	var num1 int32 = 10
	var num2 int32 = 5
	if (num1+num2)%3 == 0 && (num1+num2)%5 == 0 {
		fmt.Println("既能被3整除也能被5整除")
	}

	//判断年份是否为闰年

	var year int = 2019
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println(year, "是闰年")
	} else {
		fmt.Println(year, "不是闰年")
	}

}
