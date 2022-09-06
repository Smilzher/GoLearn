package main

import "fmt"

//使用select解决从管道取数据的阻塞问题

func main() {
	// 定义一个管道10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	//管道5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统方法遍历管道时，如果不关闭会阻塞导致deadlock
	//实际开发中，可能我们不好确定什么时候关闭该管道
	//可以使用select方式解决
	//label:
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取数据%s\n", v)
		default:
			fmt.Printf("都取不到了，不玩了。。。\n")
			return
			//break label
		}
	}
}
