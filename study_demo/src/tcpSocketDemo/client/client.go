package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 目标服务器的地址
	serverAddress := "localhost:8888"

	// 连接到服务器
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer conn.Close()
	// 功能一，客户端可以发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin)

	// 从终端读取用户的输入发送给服务器
	fmt.Println("请输入要发送的内容：")
	for {
		content, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("read string err =", err)
		}

		if strings.Trim(content, " \r\n") == "exit" {
			fmt.Println("client exit...")
			break
		}

		// 将读取到的内容发送给服务器
		n, err := conn.Write([]byte(content))
		if err != nil {
			fmt.Println("conn write err =", err)
		}
		fmt.Println("客户端发送的字节数：", n)
	}

}
