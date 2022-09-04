package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	//多种数据类型放入管道
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)

	cat1 := Cat{Name: "tome", Age: 10}
	cat2 := Cat{Name: "tom~", Age: 18}
	allChan <- cat1
	allChan <- cat2
	allChan <- 10
	allChan <- "jack"
	//取出
	newCat := <-allChan
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v", a.Name)
}
