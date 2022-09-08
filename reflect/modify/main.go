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

func testFloat(b interface{}) {
	rTyp := reflect.TypeOf(b)
	val := reflect.ValueOf(b)
	fmt.Println("rTyp=", rTyp)
	fmt.Println("kind=", val.Kind())
	iVal := val.Interface()
	floatVal := iVal.(float64)
	fmt.Printf("floatVal=%v, Type=%T\n", iVal, floatVal)
}

func main() {
	var num int = 20
	testInt(&num)
	fmt.Println("num=", num)
	fmt.Println()
	var num2 float64 = 1.2
	testFloat(num2)
}
