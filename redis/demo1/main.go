package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	pwd := redis.DialPassword("123456")
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", pwd)
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("Set", "name", "tome")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	//通过go向redis读取数据string
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}
	fmt.Printf("r = %v", r)

	//r, err := conn.Do("Get", "name")
	//if err != nil {
	//	fmt.Println("get err=", err)
	//	return
	//}
	//res := string(r.([]uint8))
	//fmt.Println("r=", res)
}
