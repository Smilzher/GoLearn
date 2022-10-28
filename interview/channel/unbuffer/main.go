package main

import (
	"fmt"
	"time"
)

func loop(ch chan int) {
	for {
		select {
		case i := <-ch:
			fmt.Println("this  value of unbuffer channel", i)
		}
	}
}

func main() {
	//1. 会产生阻塞，从而报错
	//ch := make(chan int)
	//ch <- 1
	//go loop(ch)
	//time.Sleep(1 * time.Millisecond)

	//2. 不会产生阻塞
	//ch := make(chan int)
	//go loop(ch)
	//ch <- 1
	//time.Sleep(1 * time.Millisecond)

	// 3. 缓冲channel, 会产生阻塞
	//解决：
	//(1)把 channel 长度调大一点
	//(2)把 channel 的信息发送者 ch <- 1 这些代码移动到 go loop(ch) 下面 ，让 channel 实时消费就不会导致阻塞了
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	go loop(ch)
	time.Sleep(1 * time.Millisecond)
}
