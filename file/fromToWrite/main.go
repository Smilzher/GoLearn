package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filePath1 := "/data/abc.txt"
	filePath2 := "/data/ddd.txt"
	data, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Printf("read file err=%v\n", err)
		return
	}
	err = ioutil.WriteFile(filePath2, data, 0777)
	if err != nil {
		fmt.Printf("write file error=%v\n", err)
	}
}
