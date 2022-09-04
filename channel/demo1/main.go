package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	intChan <- 10
	intChan <- 20
	intChan <- 30
	//intChan的容量为3，再放会报告deadlock
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	fmt.Printf("num1=%v num2=%v num3=%v", num1, num2, num3)
}
