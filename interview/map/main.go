package main

import "fmt"

func main() {
	var m1 map[string]string
	m1["1"] = "1" //不能对未初始化的map进行赋值，这样将会抛出一个异常
	fmt.Println(m1)
}
