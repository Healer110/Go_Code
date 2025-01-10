package main

import (
	"fmt"
	"net"
	"server/model"
	"time"
)

// 处理和客户端的通讯
func controller(conn net.Conn) {
	defer conn.Close()

	// 调用总控
	processor := &Processor{
		Conn: conn,
	}

	// 启动控制器
	err := processor.ProcessController()
	if err != nil {
		return
	}

}

// 编写一个函数，完成对UserDao的初始化任务
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	// 服务区启动时，就初始化redis的连接池
	initPool("0.0.0.0:6379", 16, 0, time.Second*300)
	// 初始化
	initUserDao()

	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("服务端监听失败...", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("服务器建立连接异常...", err)
			continue
		}

		// 连接成功，启动协程,跟客户端保持通讯
		go controller(conn)

	}

}
