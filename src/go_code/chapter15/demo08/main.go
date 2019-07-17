package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

//struct序列化
func testStruct() {
	monster := Monster{
		Name:     "SpiderMan",
		Age:      16,
		Birthday: "2000-11-11",
		Sal:      8000.0,
		Skill:    "toast",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误, err = %v\n", err)
	}
	fmt.Printf("monster序列化后=%v\n", string(data))
}

//map序列化
func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "IronMan"
	a["age"] = 30
	a["address"] = "America"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误, err = %v\n", err)
	}
	fmt.Printf("a map 序列化后=%v\n", string(data))
}

//切片的序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "Bill"
	m1["age"] = 25
	m1["address"] = "BeiJing"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "Rose"
	m2["age"] = 20
	m2["address"] = [2]string{"America", "TianJin"}
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误, err = %v\n", err)
	}
	fmt.Printf("slice 序列化后=%v\n", string(data))
}

func main() {
	testStruct()
	testMap()
	testSlice()
}
