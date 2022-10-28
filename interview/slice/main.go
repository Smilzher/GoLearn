package main

import "fmt"

//go tool compile -S test.go | grep CALL 得到汇编代码
func main() {
	//test1()

	//deepCopy()
	//shallowCopy()

	expend1()
}

func test1() {
	slice := make([]int, 0)
	slice = append(slice, 1)
	fmt.Println(slice, len(slice), cap(slice))
}

//深拷贝和浅拷贝 案例
//1.copy
func deepCopy() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3, 5)
	fmt.Printf("slice1 %v, %p\n", slice1, slice1)
	copy(slice2, slice1)
	fmt.Printf("slice2 %v, %p\n", slice2, slice2)
	slice3 := make([]int, 0, 5)
	fmt.Printf("slice3 before:%v, %p\n", slice3, slice3)
	for _, v := range slice1 {
		slice3 = append(slice3, v)
	}
	fmt.Printf("slice3 after:%v, %p\n", slice3, slice3)
}

//引用类型的变量，默认赋值操作就是浅拷贝
func shallowCopy() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("shallowCopy slice1:%v, %p\n", slice1, slice1)
	slice2 := slice1
	fmt.Printf("shallowCopy slice2:%v, %p\n", slice2, slice2)
}

//slice扩容机制
//1.如果新申请容量比两倍原有容量大，那么扩容后容量为新申请容量
//2.如果原有slice长度小于1024,那么每次扩容为原来的2倍
//3.如果原slice长度大于等于1024，那么每次扩容就扩为原来的1.25倍

func expend1() {
	//2（老容量）+ 3（新添加的元素）= 5，超出 4 （老容量的两倍），即预估容量为 5
	//如果原有slice长度小于1024,那么每次扩容为原来的2倍
	//故而才可能容量为6
	s := []int64{1, 2}
	s = append(s, 3, 4, 5)
	fmt.Printf("expend1 len=%d, cap=%d", len(s), cap(s))
}
