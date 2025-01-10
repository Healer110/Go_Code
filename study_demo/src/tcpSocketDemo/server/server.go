package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 循环接收客户端发送的数据
	defer conn.Close()
	fmt.Println("服务器等待客户端的输入............", conn.RemoteAddr().String())

	for {
		// 创建新的切片
		buf := make([]byte, 1024)

		n, err := conn.Read(buf) // 读取的时候会阻塞
		if err != nil {
			fmt.Println("服务端read err:", err)
			return
		}
		// 显示客户端发送的内容到终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("server listening...")
	ln, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen error, ", err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept encounter error...", err)
			continue
		}
		fmt.Println("client ip address =", conn.RemoteAddr().String())
		// 启动一个协程为该客户端服务
		go process(conn)
	}
}
