package main

import "fmt"

func SelectSort(arr *[6]int) {
	for j := 0; j < len(arr)-1; j++ {
		min := arr[j]
		minIndex := j
		for i := j + 1; i < len(arr); i++ {
			if min > arr[i] {
				min = arr[i]
				minIndex = i
			}
		}
		if minIndex != j {
			arr[j], arr[minIndex] = arr[minIndex], arr[j]
		}
		fmt.Printf("第%d次 %v \n", j+1, *arr)
	}
}

func main() {
	arr := [...]int{8, 3, 7, 1, 6, 9}
	SelectSort(&arr)
}
