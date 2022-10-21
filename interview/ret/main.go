package main

import "fmt"

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i //或者直接写成return
}

func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return &i
}

func main() {
	fmt.Println("return b():", b())
	fmt.Println("return c():", *(c()))

	str := `test \n ttt %d`
	fmt.Printf(str, 100)
}
