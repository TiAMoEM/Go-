package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	// Scores [5]float64
	// ptr    *int
	// slice  []int
	// map1   map[string]string
}

func main() {
	//方式一
	var p1 Person
	fmt.Println(p1)
	//方式二
	var p2 Person = Person{"Bill", 18}
	fmt.Println(p2)
	//方式三
	var p3 *Person = new(Person)
	// (*p3).Name = "john"
	p3.Name = "Rose"
	p3.Age = 18
	// (*p3).Age = 20
	fmt.Println(*p3)

	//方式四
	var p4 *Person = &Person{"John", 19}
	fmt.Println(*p4)
}
