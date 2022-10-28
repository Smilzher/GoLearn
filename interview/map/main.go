package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//test1()
	//unsafe()
	//safeLock()
	syncMap()
}

// 2个协程同时读和写，以下程序会出现致命错误：fatal error: concurrent map writes
func test1() {
	var m1 map[string]string
	m1["1"] = "1" //不能对未初始化的map进行赋值，这样将会抛出一个异常
	fmt.Println(m1)
}

func unsafe() {
	s := make(map[int]int)
	for i := 0; i < 100; i++ {
		go func(i int) {
			s[i] = i
		}(i)
	}
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Printf("map第%d个元素值是%d\n", i, s[i])
		}(i)
	}
	time.Sleep(1 * time.Second)
}

//实现map线程安全，有两种方式
//1. 使用读写锁 map + sync.RWMutex
func safeLock() {
	var lock sync.RWMutex
	s := make(map[int]int)
	for i := 0; i < 100; i++ {
		go func(i int) {
			lock.Lock()
			s[i] = i
			lock.Unlock()
		}(i)
	}
	for i := 0; i < 100; i++ {
		go func(i int) {
			lock.RLock()
			fmt.Printf("map第%d个元素值是%d\n", i, s[i])
			lock.RUnlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
}

//使用Go提供的 sync.Map
func syncMap() {
	var m sync.Map
	for i := 0; i < 100; i++ {
		go func(i int) {
			m.Store(i, i)
		}(i)
	}
	for i := 0; i < 100; i++ {
		go func(i int) {
			v, ok := m.Load(i)
			fmt.Printf("Load: %v, %v\n", v, ok)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
