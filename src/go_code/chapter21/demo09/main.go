package main

import "fmt"

func InsertSort(arr *[7]int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1

		for insertIndex >= 0 && arr[insertIndex] > insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d次插入后 %v \n", i, *arr)
	}
}

func main() {
	arr := [7]int{6, 3, 5, 2, 4, 1, 9}
	InsertSort(&arr)
}
