package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//通过反射获取传入变量的type, kind值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)
	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)
	//下面将rVal转成interface{}
	iV := rVal.Interface()
	//将interface{}通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

//对结构体的反射
func reflectTest02(b interface{}) {
	//通过反射获取的传入的变量的type, kind,值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	//将rVal转换成interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T\n", iV, iV)
	//将interface{}通过断言转成需要的类型
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	//演示对（基本数据类型、interface{}、reflect.Value）进行反射的基本操作
	//var num int = 100
	//reflectTest01(num)

	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
}
