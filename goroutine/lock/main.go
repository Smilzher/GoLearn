package main

import (
	"fmt"
	"sync"
	"time"
)

//需求：计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中
//最后显示出来，要求使用goroutine完成

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock 是一个全局的互斥锁
	//sync是包：synchornized 同步
	//Mutex: 是互斥结构体
	lock sync.Mutex
)

//test函数计算n!, 放入myMap中
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//加锁
	lock.Lock()
	myMap[n] = res //map writes
	//解锁
	lock.Unlock()
}

func main() {
	for i := 1; i <= 10; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 10)

	//此处也要加锁，因为10s后，有的协程还没执行完(在写操作)，此时若读，会出现冲突
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
