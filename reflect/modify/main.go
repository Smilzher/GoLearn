package main

import (
	"fmt"
	"reflect"
)

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	fmt.Printf("val type=%T\n", val)
	val.Elem().SetInt(120)
	fmt.Printf("val=%v\n", val)
}

func main() {
	var num int = 20
	testInt(&num)
	fmt.Println("num=", num)
}
