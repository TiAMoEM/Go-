package main

import "fmt"
import "sort"

func main() {
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 10
	map1[4] = 40
	map1[8] = 80
	fmt.Println(map1)

	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	fmt.Println(keys)

	for _, v := range keys {
		fmt.Printf("map1[%v] = %v\n", v, map1[v])
	}
}
