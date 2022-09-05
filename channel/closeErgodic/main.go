package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	//管道关闭，不再能写入数据；可继续读数据
	//intChan <- 300
	fmt.Println("ok~")
	n1 := <-intChan
	fmt.Println("n1=", n1)

	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 //放入100个数据到管道
	}
	//遍历管道不能使用普通的for循环
	//在遍历时，如果channel没有关闭，则会出现deadlock的错误
	//在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完成后，正常退出遍历
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
