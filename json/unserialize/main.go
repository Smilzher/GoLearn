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

func unmarshalStruct() {
	str := "{\"monster_name\":\"牛魔王\",\"monster_age\":300,\"monster_birthday\":\"2022-08-08\", \"monster_sal\":8000,\"monster_skill\":\"牛魔拳\"}"
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
		return
	}
	fmt.Printf("反序列化后 monster=%v\n", monster)
}

func unmarshalMap() {
	str := "{\"address\":\"北京\", \"name\":\"test\"}"

	var data map[string]interface{}
	//反序列化map,不需要make,因为make操作被封装到unmarshal函数中
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
		return
	}
	fmt.Printf("反序列化后 data=%v\n", data)
}

//切片的代码类似

func main() {
	unmarshalStruct()
	unmarshalMap()
}
