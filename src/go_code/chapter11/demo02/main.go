package main

import "fmt"

func main() {
	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var hereName = ""
	fmt.Println("请输入要查找的人名...")
	fmt.Scanln(&hereName)

	//方法一
	for i := 0; i < len(names); i++ {
		if hereName == names[i] {
			fmt.Printf("找到%v， 下标%v \n", hereName, i)
			break
		} else if i == (len(names) - 1) {
			fmt.Printf("没有找到%v \n", hereName)
		}
	}

	//方法二
	index := -1
	for i := 0; i < len(names); i++ {
		if hereName == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v， 下标%v \n", hereName, index)
	} else {
		fmt.Printf("没有找到%v \n", hereName)
	}
}
