package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() //关闭conn
	//循环接收客户端发送的数据
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//等待客户端通过conn发送信息
		//如果客户端没有write[发送]，那么协程就会阻塞在这里
		fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn读取
		if err != nil {
			fmt.Printf("客户端退出 err=%v", err)
			return //!!!
		}
		//显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close()

	//循环等待客户端来链接
	for {
		fmt.Println("等待客户端来链接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//协程，为客户端服务
		go process(conn)
	}
}
