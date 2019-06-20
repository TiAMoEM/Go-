package main

import "fmt"

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("Could not find!!!!")
		return
	}
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("find it!, index = %v \n", middle)
	}
}

//自己写的非递归算法
func BinaryFind1(arr *[6]int, data int) (index int) {
	left := 0
	right := len(*arr)
	for {
		if left > right {
			return -1
		}
		middle := (left + right) / 2
		if (*arr)[middle] > data {
			right = middle - 1
		} else if (*arr)[middle] < data {
			left = middle + 1
		} else {
			return middle
		}
	}
}

func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	BinaryFind(&arr, 0, len(arr)-1, 1234)
	fmt.Println(BinaryFind1(&arr, 1234))

}
