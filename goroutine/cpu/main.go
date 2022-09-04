package main

import (
	"fmt"
	"runtime"
)

func main() {
	//获取当前系统cpu的数量
	num := runtime.NumCPU()
	//设置num-1的cpu运行go程序
	runtime.GOMAXPROCS(num)
	fmt.Println("num=", num)
}
