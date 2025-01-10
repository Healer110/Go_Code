package process

import (
	"chatRoom_system/client/utils"
	"chatRoom_system/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type UserProcess struct {
	// 字段...
}

func (up *UserProcess) Login(userId int, userPwd string) (err error) {
	// fmt.Printf("userId = %d, userPwd = %s \n", userId, userPwd)
	// return nil

	// 连接到服务端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("Login 连接客户端失败：", err)
		return
	}

	defer conn.Close()

	// 发送消息给服务器
	var msg message.Message
	// 定义消息类型
	msg.Type = message.LoginMsgType
	// 创建data
	var data message.LoginMsg
	data.UserId = userId
	data.UserPwd = userPwd
	// 对消息数据进行序列化，序列化后的字符串赋值给msg--> data
	res, err := json.Marshal(data)
	if err != nil {
		fmt.Println("序列化消息数据异常，", err)
		return
	}
	msg.Data = string(res)

	// 对发送的消息体再次进行序列化，序列化后由客户端转化为[]byte进行传送
	res, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("序列化消息体异常，", err)
		return
	}

	// 发送数据到服务端
	// 先发送一个数据长度，再发送数据，这是自定义发送的消息格式
	var pkgLen uint32 = uint32(len(res))
	var bytes [4]byte
	// 将一个数字转换为字节切片
	binary.BigEndian.PutUint32(bytes[:], pkgLen)
	n, err := conn.Write(bytes[:])
	if n != 4 || err != nil {
		fmt.Println("conn.Write 失败：", err)
		return
	}

	_, err = conn.Write(res)
	if err != nil {
		fmt.Println("conn.Write 发送data失败：", err)
		return
	}

	// 接收服务端的数据, 构造一个接收包的结构体，并调用其方法
	tf := &utils.Transfer{
		Conn: conn,
	}
	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err: ", err)
		return
	}
	// 反序列化data
	var responseLoginMsg message.LonginResMsg
	err = json.Unmarshal([]byte(msg.Data), &responseLoginMsg)
	if err != nil {
		fmt.Println("登录消息反序列化失败....")
		return
	}

	if responseLoginMsg.Code == 200 {
		// fmt.Println("登录成功...")

		// 初始化CurUser
		CurUser.Conn = conn
		CurUser.User.UserId = userId
		CurUser.User.UserStatus = message.UserOnline

		// 登录成功后，显示一下在线用户列表
		fmt.Println("当前在线用户列表：")
		for _, v := range responseLoginMsg.UsersId {
			// 过滤掉自己的ID
			if v == userId {
				continue
			}
			fmt.Printf("用户Id: %d\n", v)

			// 完成客户端onlineUsers初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println()

		// 这里启动一个隐藏的协程，
		// 该协程保持跟服务端的通信，及时接收服务器发送过来的消息
		go ServerProcessMsg(conn)

		// 显示登录成功后的消息，并调用登录成功后的菜单
		for {
			ShowMenu()
		}
	} else {
		// fmt.Printf("登录失败：%s \n", responseLoginMsg.Error)
		err = errors.New(responseLoginMsg.Error)

	}
	return
}

// 完成注册
func (up *UserProcess) Register(userId int, userPwd string, username string) (err error) {
	// 连接到服务端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("Login 连接客户端失败：", err)
		return
	}

	defer conn.Close()

	// 发送消息给服务器
	var msg message.Message
	// 定义消息类型
	msg.Type = message.RegisterMsgType
	// 创建data
	var data message.RegisterMsg
	data.User.UserId = userId
	data.User.UserPwd = userPwd
	data.User.UserName = username
	// 对消息数据进行序列化，序列化后的字符串赋值给msg--> data
	res, err := json.Marshal(data)
	if err != nil {
		fmt.Println("序列化消息数据异常，", err)
		return
	}
	msg.Data = string(res)

	// 对发送的消息体再次进行序列化，序列化后由客户端转化为[]byte进行传送
	res, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("序列化消息体异常，", err)
		return
	}

	// 接收服务端的数据, 构造一个接收包的结构体，并调用其方法
	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(res)
	if err != nil {
		fmt.Println("[client/process/userProcess.go] conn.Write 发送data失败：", err)
		return
	}

	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err: ", err)
		return
	}
	// 反序列化data
	var responseRegisterMsg message.RegisterResMsg
	err = json.Unmarshal([]byte(msg.Data), &responseRegisterMsg)
	if err != nil {
		fmt.Println("登录消息反序列化失败....")
		return
	}

	if responseRegisterMsg.Code == 200 {
		fmt.Println("注册成功...")
	} else {
		// fmt.Printf("登录失败：%s \n", responseLoginMsg.Error)
		err = errors.New(responseRegisterMsg.Error)

	}
	return
}
