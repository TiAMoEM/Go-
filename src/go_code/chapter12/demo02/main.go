package main

import (
	"fmt"
)

func main() {
	cities := make(map[string]string)
	cities["no1"] = "BeiJing"
	cities["no2"] = "ShangHai"
	cities["no3"] = "HangZhou"
	fmt.Println(cities)

	cities["no3"] = "TianJin"
	fmt.Println(cities)
	delete(cities, "no4")
	delete(cities, "no3")
	fmt.Println(cities)
}
