package main

import (
	"fmt"
)

func BubbleSort(arr *[5]int) {
	temp := 0
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j+1]
				(*arr)[j+1] = (*arr)[j]
				(*arr)[j] = temp
			}
		}
	}
}

func main() {
	arr := [...]int{5, 7, 9, 2, 3}
	fmt.Println(arr)
	BubbleSort(&arr)
	fmt.Println(arr)
}
