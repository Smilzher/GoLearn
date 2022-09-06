package main

import "fmt"

//要求统计1-8000的数字中，哪些是素数？

//向intChan放入1-8000个数
func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	//关闭intChan
	close(intChan)
}

//从intChan取出数据，并判断是否为素数，如果是，则放入primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	//使用for循环取出
	var flag bool
	for {
		num, ok := <-intChan
		if !ok { //取不到
			break
		}
		flag = true
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //不是素数
				flag = false
				break
			}
		}
		if flag {
			//将这个数放入primeChan
			primeChan <- num
		}
	}
	fmt.Println("有一个primeNum协程协程因为取不到数据，退出")
	//这里不能关闭primeChan, 因为不知道其他协程会不会还在写入
	//向exitChan写入true
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) //放入结果
	//标识退出的管道
	exitChan := make(chan bool, 4) //4个管道处理，

	//开启一个协程，向intChan 放入1-8000个数据
	go putNum(intChan)

	//开启4个协程，4核cpu并行处理， 从intChan取出数据，并判断是否素数，如果是，则放入到primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	//主线程处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		//当从exitChan中取出4个结果，则说明处理完了数据，可以关闭primeChan
		close(primeChan)
	}()

	//遍历primeChan, 把结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		//将结果输出
		fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("main线程退出")
}
