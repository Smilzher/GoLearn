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

func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      300,
		Birthday: "2022-08-08",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}
	//将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("monster 序列化后=%v\n", string(data))
	//fmt.Println(monster)
}

func main() {
	testStruct()
}
