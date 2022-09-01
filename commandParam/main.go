package main

import (
	"flag"
	"fmt"
	"os"
)

func receiveArgs() {
	fmt.Println("命令行的参数有", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v", i, v)
	}
}

//flag包用来解析命令行参数
func receiveByFlag() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "", "localhost")
	flag.IntVar(&port, "port", 3306, "端口")
	//此操作非常重要，必须调用该方法
	flag.Parse()
	//输出结果
	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}

func main() {
	//receiveArgs()
	receiveByFlag()
}
