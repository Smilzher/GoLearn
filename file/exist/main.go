package main

import (
	"fmt"
	"os"
)

//如果返回的错误为nil,说明文件和文件夹存在
//如果返回的错误使用os.IsNotExist()判断为true,说明文件和文件夹不存在
//如果返回的错误为其他类型，则不确定是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	file := "/data/yyy.txt"
	res, err := pathExists(file)
	fmt.Println(res, err)
}
