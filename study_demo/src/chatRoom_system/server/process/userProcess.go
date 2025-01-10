package process

import (
	"encoding/json"
	"fmt"
	"net"
	"server/common/message"
	"server/model"
	"server/utils"
)

// 处理跟用户相关的操作
type UserProcess struct {
	Conn net.Conn
	// 增加一个字段，表示该connect是哪个用户的
	UserId int
}

// 编写通知所有用户在线的方法
// 每个用户上线后，推送一次，这里需要传入上线用户的Id
func (up *UserProcess) NotifyOthersOnlineUser(userId int) {
	// 遍历所有已经上线的用户,然后一个一个的发送NotifyUserStatusMsg消息
	for id, otherUp := range userMgr.onlineUsers {
		// 过滤掉自己
		if id == userId {
			continue
		}
		otherUp.NotifyMeOnline(userId)
	}
}

// 发送通知的逻辑
func (up *UserProcess) NotifyMeOnline(userId int) {
	var msg message.Message
	msg.Type = message.NotifyUserStatusMsgType
	var notifyStatusMsg message.NotifyUserStatusMsg = message.NotifyUserStatusMsg{
		UserId: userId,
		Status: message.UserOnline,
	}

	data, err := json.Marshal(notifyStatusMsg)
	if err != nil {
		fmt.Println("[server/process/userProcess.go] NotifyMeOnline() err: ", err)
		return
	}
	msg.Data = string(data)

	// 将总的message序列化，然后发送
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("[server/process/userProcess.go] NotifyMeOnline() err: ", err)
		return
	}

	// 构造Transfer发送
	// 发送数据，封装到一个函数中
	tf := &utils.Transfer{
		Conn: up.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("[server/process/userProcess.go] NotifyMeOnline()-WritePkg err:", err)
		return
	}

}

// 编写一个方法，处理注册
func (up *UserProcess) ServerProcessRegister(msg *message.Message) (err error) {
	// 从传入的消息中，取出消息的具体内容data,并反序列化成RegisterMsg
	var registerMsg message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data), &registerMsg)
	if err != nil {
		fmt.Println("register 消息反序列化error：", err)
		return
	}

	var resMsg message.Message
	resMsg.Type = message.RegisterResMsgType
	var registerResMsg message.RegisterResMsg

	// 去数据库完成注册
	err = model.MyUserDao.Register(&registerMsg.User)
	if err != nil {
		registerResMsg.Code = 400
		registerResMsg.Error = fmt.Sprint(err)
	} else {
		registerResMsg.Code = 200
	}

	// 序列化消息内容
	data, err := json.Marshal(registerResMsg)
	if err != nil {
		fmt.Println("返回消息序列化失败")
		return
	}

	resMsg.Data = string(data)
	// 序列化消息体
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("返回消息序列化失败")
		return
	}

	// 发送数据，封装到一个函数中
	tf := &utils.Transfer{
		Conn: up.Conn,
	}

	tf.WritePkg(data)
	return
}

// 编写一个函数ServerProcessLogin函数，专门处理登录请求
func (up *UserProcess) ServerProcessLogin(msg *message.Message) (err error) {
	// 从传入的消息中，取出消息的具体内容data,并反序列化成LoginMsg
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("login 消息反序列化error：", err)
		return
	}

	// fmt.Println(msg.Data)

	// 声明一个返回的消息类型
	var resMsg message.Message
	resMsg.Type = message.LoginResMsgType

	var loginResMsg message.LonginResMsg

	// 如果用户的id=100, 密码=123456 ，认为用户是合法的，否则不合法
	// if loginMsg.UserId == 100 && loginMsg.UserPwd == "123456" {
	// 	loginResMsg.Code = 200
	// } else {
	// 	loginResMsg.Code = 500 // 500 表示用户不存在
	// 	loginResMsg.Error = "用户不存在，请注册..."
	// }

	user, err := model.MyUserDao.LoginCheck(loginMsg.UserId, loginMsg.UserPwd)
	fmt.Println("登录用户信息：", "[server/process/userProcess.go]", *user)
	if err != nil {
		loginResMsg.Code = 500
		loginResMsg.Error = fmt.Sprint(err)
	} else {
		loginResMsg.Code = 200
		// 登录成功，将登录成功的用户放入到userMgr中,并将登录成功的userId赋值给up结构体
		up.UserId = loginMsg.UserId
		userMgr.AddOnlineUser(up)
		// 上线后，立即通知其他用户
		up.NotifyOthersOnlineUser(loginMsg.UserId)
		// 将当前在线用户的Id放入到LonginResMsg.UsersId中
		for idx, _ := range userMgr.onlineUsers {
			loginResMsg.UsersId = append(loginResMsg.UsersId, idx)
		}

	}

	// 序列化消息内容
	data, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("返回消息序列化失败")
		return
	}

	resMsg.Data = string(data)

	// 序列化消息体
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("返回消息序列化失败")
		return
	}

	// 发送数据，封装到一个函数中
	tf := &utils.Transfer{
		Conn: up.Conn,
	}

	tf.WritePkg(data)
	return
}
