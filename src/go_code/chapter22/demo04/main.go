package main

import "fmt"

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

func Map(f func(int) int, l []int) []int {
	j := make([]int, len(l))
	for k, v := range l {
		j[k] = f(v)
	}
	return j
}

func main() {
	var s stack
	s.push(25)
	s.push(14)
	fmt.Printf("stack %v\n", s)

	m := []int{1, 3, 4}
	f := func(i int) int {
		return i * i
	}
	fmt.Printf("%v", (Map(f, m)))

}
