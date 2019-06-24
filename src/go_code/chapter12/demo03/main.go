package main

import (
	"fmt"
)

func main() {
	cities := make(map[string]string)
	cities["no1"] = "BeiJing"
	cities["no2"] = "ShangHai"
	cities["no3"] = "HangZhou"

	for key, value := range cities {
		fmt.Printf("key = %v, value = %v\n", key, value)
	}

	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "Tom"
	studentMap["stu01"]["sex"] = "male"
	studentMap["stu01"]["age"] = "18"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "Mary"
	studentMap["stu02"]["sex"] = "female"
	studentMap["stu02"]["age"] = "16"

	for k1, v1 := range studentMap {
		fmt.Println("k1 =", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2 = %v, v2 = %v\n", k2, v2)
		}
		fmt.Println()
	}

}
