package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}

//Less方法是年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
	//可以用以下方法替换交换操作,类似python的写法
	// hs[i], hs[j] = hs[j], hs[i]
}

type Student struct {
	Name  string
	Age   int
	Score float64
}

func main() {
	var intSlice = []int{0, -1, 10, 7, 90}
	fmt.Println(intSlice)
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄|%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	for _, v := range heroes {
		fmt.Println(v)
	}
	sort.Sort(heroes)
	fmt.Println("--------------------")
	for _, v := range heroes {
		fmt.Println(v)
	}

}
