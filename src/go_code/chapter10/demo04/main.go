package main

import (
	"fmt"
)

func main() {
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	slice := intArr[1:3]
	fmt.Println("intArr =", intArr)
	fmt.Println("slice =", slice)
	fmt.Println("slice 的元素个数 =", len(slice))
	fmt.Println("slice 的容量", cap(slice))

	//切片的使用方式，方式一如上
	//方式二 var name []type = make([]type, len, cap)
	var slice1 []float64 = make([]float64, 5, 10)
	slice1[1] = 10
	slice1[3] = 20
	fmt.Println(slice1)
	fmt.Println("slice1 size =", len(slice1))
	fmt.Println("slice1 cap =", cap(slice1))

	//方式三：定义一个切片，直接就置顶具体数组
	var strslice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strslice =", strslice)
	fmt.Println("strslice size =", len(strslice))
	fmt.Println("strslice cap =", cap(strslice))

}
