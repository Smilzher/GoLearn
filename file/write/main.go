package main

import (
	"bufio"
	"fmt"
	"io"
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

func readWrite() {
	filePath := "/data/abc.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	//先读取原来文件的内容，并显示在终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //如果读取到文件末尾
			break
		}
		fmt.Print(str)
	}

	str := "hello,北京\r\n" //\r\n表示换行
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

func main() {
	//write1()
	readWrite()
}
