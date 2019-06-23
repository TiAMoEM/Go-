package main

import (
	"fmt"
	"math/rand"
)

//第一题
func exercise1() {
	var arr [10]int
	sum := 0
	temp := 0
	max := 0
	maxIndex := 0
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
		sum += arr[i]
	}
	fmt.Println("原序列~~~", arr)
	for j := 0; j < len(arr)/2; j++ {
		temp = arr[len(arr)-1-j]
		arr[len(arr)-1-j] = arr[j]
		arr[j] = temp
	}
	fmt.Println("反序输出~~~", arr)
	fmt.Println("平均值", float64(sum)/float64(len(arr)))
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
	}
	fmt.Println("最大的值是~~~", max)
	fmt.Println("最大值位置~~~", maxIndex+1)
}

func exercise2() {
	arr := [9]int{1, 2, 3, 4, 5, 6, 8, 9, 10}
	var num int
	fmt.Println("please input number~~~~~~~~~~~")
	fmt.Scanln(&num)
	slice := arr[:]
	slice1 := append(slice, num)
	fmt.Println(slice1)
	for i := 0; i < len(slice1); i++ {
		if slice1[i] > num {
			for j := len(slice1) - 1; j >= i; j-- {
				slice1[j] = slice1[j-1]
			}
			slice1[i] = num
			break
		}
	}
	fmt.Println(slice1)
}

func main() {
	exercise1()
	exercise2()
}
