package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile1() {
	name := "/data/a.txt"
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	defer file.Close()
	//创建一个*Reader, 是带缓冲的,
	reader := bufio.NewReader(file)
	//循环读取文件的内容
	for {
		str, err := reader.ReadString('\n') //读取到一个换行就结束
		if err == io.EOF {                  //io.EOF表示文件末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}

func readFile2() {
	file := "/data/a.txt"
	//一次性将文件读取到位
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	fmt.Printf("%v", string(content))
}

func main() {
	//readFile1()
	readFile2()
}
