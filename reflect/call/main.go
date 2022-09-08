package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (cal Cal) GetSub(name string) {
	res := cal.Num1 - cal.Num2
	fmt.Printf("%s完成了减法运算, %d - %d = %d", name, cal.Num1, cal.Num2, res)
}

func TestStruct(a interface{}) {
	//获取reflect.Type类型
	//typ := reflect.TypeOf(a)
	//获取reflect.Value类型
	val := reflect.ValueOf(a)
	//获取a对应的类别
	kd := val.Kind()
	//如果传入的不是struct 就退出
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取结构体字段
	num := val.NumField()
	for i := 0; i < num; i++ {
		fmt.Printf("field %d: 值为=%v\n", i, val.Field(i))
	}
	//获取结构体方法数
	methodNum := val.NumMethod()
	fmt.Printf("struct has %d methods\n", methodNum)
	//调用方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf("tom"))
	val.Method(0).Call(params)
}

//编写一个Cal结构体，有两个字段Num1和Num2
//方法GetSub(name string)
//使用反射遍历Cal结构体所有的字段信息
//使用反射机制完成对GetSub的调用，输出形式
//"tom完成了减法运算,8 - 3 = 5"
func main() {
	cal := Cal{
		Num1: 8,
		Num2: 3,
	}
	TestStruct(cal)
}
