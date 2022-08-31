package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/data/ab.txt")
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	fmt.Printf("file=%v", file)
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}
