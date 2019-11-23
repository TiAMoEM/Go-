package main

import "fmt"

func SetWay(myMap *[8][7]int, i int, j int) bool {
	if myMap[6][5] == 2 {
		return true
	} else {
		if myMap[i][j] == 0 {
			myMap[i][j] = 2
			if SetWay(myMap, i+1, j) {
				return true
			} else if SetWay(myMap, i, j+1) {
				return true
			} else if SetWay(myMap, i-1, j) {
				return true
			} else {
				myMap[i][j] = 3
				return false
			}
		} else {
			return false
		}
	}
}

func main() {
	var myMap [8][7]int
	//将地图的四周用墙围起来
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}

	//设置障碍
	myMap[3][1] = 1
	myMap[3][2] = 1
	// myMap[1][2] = 1
	// myMap[2][2] = 1

	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

	SetWay(&myMap, 1, 1)
	fmt.Println("探测完毕...")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}
