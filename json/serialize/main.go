package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"monster_name"` //tag标签，用于输出json指定的字段
	Age      int     `json:"monster_age"`
	Birthday string  `json:"monster_birthday"`
	Sal      float64 `json:"monster_sal"`
	Skill    string  `json:"monster_skill"`
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
