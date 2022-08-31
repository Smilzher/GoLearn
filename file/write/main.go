package main

import (
	"bufio"
	"fmt"
	"os"
)

func write1() {
	filePath := "/data/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	str := "hello world\r\n"
	//写入时。使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush() //将缓冲的数据真正写入到文件中
}

func main() {
	write1()
}
