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

func unmarshalStruct() {
	str := "{\"Name\": \"IronMan\", \"Age\": 30, \"Birthday\": \"1988-11-11\", \"Sal\": 80000, \"Skill\": \"Money\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster = %v, monster.Name = %v\n", monster, monster.Name)
}

func main() {
	unmarshalStruct()
}
