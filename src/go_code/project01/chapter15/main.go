package main

import (
	"fmt"
)

func main() {
	key := ""
	loop := true
	for {
		fmt.Println("-----------家庭收支记账软件-----------")
		fmt.Println("           1 收支明细")
		fmt.Println("           2 登记收入")
		fmt.Println("           3 登记支出")
		fmt.Println("           4 退出")
		fmt.Println("Please select (1~4):")
		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("-----------当前收支明细记录-----------")
		case "2":
			fmt.Println("登记收入")
		case "3":
			fmt.Println("登记支出")
		case "4":
			loop = false
		default:
			fmt.Println("input error")
		}
		if !loop {
			break
		}
	}
	fmt.Println("已退出家庭记账软件的使用...")
}
