package main

import (
	"fmt"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) area2() float64 {
	(*c).radius = 10.0
	return 3.14 * (*c).radius * (*c).radius
}

func main() {
	var c Circle
	c.radius = 4.0
	res := c.area()
	fmt.Println(res)
	fmt.Println(c.area2())
}
