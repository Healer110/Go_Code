package main

import "github.com/garyburd/redigo/redis"

// 定义一个全局的pool
var pool *redis.Pool

// 程序启动时，初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0, // 0表示没有限制
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	// 从连接池获取一个连接
	conn := pool.Get()
	defer conn.Close()

	// 连接池关闭后，就无法取出连接了
}
