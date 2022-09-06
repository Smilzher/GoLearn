package main

import "fmt"

func main() {
	//默认管道是双向的，可读可写
	//var chan1 chan int

	//声明只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 10
	//num := <-chan2 //err
	fmt.Println("chan2=", chan2)

	//声明只读
	var chan3 <-chan int
	num2 := <-chan3 //会报错，因为管道中没有内容
	fmt.Println("num2", num2)
}
