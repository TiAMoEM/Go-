package main

import "fmt"

func QuickSort(left int, right int, arr *[9]int) {
	l := left
	r := right
	middle := arr[(left+right)/2]
	temp := 0

	for l < r {
		for arr[l] < middle {
			l++
		}

		for arr[r] > middle {
			r--
		}

		if l > r {
			break
		}

		temp = arr[l]
		arr[l] = arr[r]
		arr[r] = temp

		// if arr[l] == middle {
		// 	r--
		// }
		// if arr[r] == middle {
		// 	l++
		// }
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort(left, r, arr)
	}
	if right > l {
		QuickSort(l, right, arr)
	}
}

func main() {
	arr := [9]int{-9, 78, 0, 23, -567, 70, 123, 90, -23}
	fmt.Println("初始数组", arr)
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println("main.......")
	fmt.Println(arr)
}
