package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//全局pool
var pool *redis.Pool

//当启动程序时，初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			option := redis.DialPassword("123456")
			return redis.Dial("tcp", "127.0.0.1:6379", option)
		},
	}
}

func main() {
	//redis连接池
	//事先初始化一定数量的链接，放入到连接池
	//当Go需要操作redis时，直接从redis连接池中取出连接即可
	//这样节省获取redis链接的时间，从而提高效率

	//先从pool取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "汤姆猫~~")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	//取出
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println("r=", r)
	//pool.Close()
	//如果我们要从pool取出连接，一定保证连接池是没有关闭的
	//...
}
