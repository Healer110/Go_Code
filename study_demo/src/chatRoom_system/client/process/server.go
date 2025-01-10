package process

import (
	"bufio"
	"chatRoom_system/client/utils"
	"chatRoom_system/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

// 显示登录成功后的提示信息
func ShowMenu() {
	fmt.Println("---------- status:登录成功 ----------")
	fmt.Println("--------1. 显示在线用户列表--------")
	fmt.Println("--------2. 发送信息        --------")
	fmt.Println("--------3. 信息列表        --------")
	fmt.Println("--------4. 退出系统        --------")
	fmt.Println("请输入(1-4)：")
	var key int
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanln(&key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("请输入要发送的消息内容：")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		content = scanner.Text()
		// fmt.Scanln(&content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("显示信息列表")
	case 4:
		fmt.Println("退出系统...")
		os.Exit(0)
	default:
		fmt.Println("输入错误，请重新输入...")
	}
}

// 和服务器端保持通信
func ServerProcessMsg(conn net.Conn) {
	// 创建Transfer实例，不停的读取服务端的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("[clent/server.go] 客户端读取服务端消息异常...", err)
			return
		}

		// 接收到了消息
		// fmt.Println("msg =", msg)
		switch msg.Type {
		case message.NotifyUserStatusMsgType: // 有人上线了
			// 取出NotifyUserStatusMsg
			var notifyUserStatusMsg message.NotifyUserStatusMsg
			err := json.Unmarshal([]byte(msg.Data), &notifyUserStatusMsg)
			if err != err {
				fmt.Println("[client/process/server.go] error: ", err)
			}
			// 将用户的状态，保存到客户端的map中
			updateUserStatus(&notifyUserStatusMsg)
		case message.SmsMesType:
			// 有人群发消息了
			outputGroupMsg(&msg)

		default:
			fmt.Println("服务器端返回了未知类型的消息...")

		}
	}

}
