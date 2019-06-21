package main

import "fmt"

/*
输出如下图形
0 0 0 0 0 0
0 0 1 0 0 0
0 2 0 3 0 0
0 0 0 0 0 0
*/

func main() {
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Println(arr2)
	fmt.Printf("arr2[0]的地址是%p\n", &arr2[0])
	fmt.Printf("arr2[1]的地址是%p\n", &arr2[1])
	fmt.Printf("arr2[0][0]的地址是%p\n", &arr2[0][0])
	fmt.Printf("arr2[1][0]的地址是%p\n", &arr2[1][0])

}
