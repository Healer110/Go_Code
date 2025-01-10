package message

import "server/model"

// 定义消息类型
const (
	LoginMsgType            = "LoginMsg"
	LoginResMsgType         = "LoginResMsg"
	RegisterMsgType         = "RegisterMsg"
	RegisterResMsgType      = "RegisterResMsg"
	NotifyUserStatusMsgType = "NotifyUserStatusMsg"
	SmsMesType              = "SmsMes"
)

// 这里定义几个常亮表示用户在线状态
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

// 要发送包的封装结构体
type Message struct {
	Type string `json:"type"` // 消息的类型
	Data string `json:"data"` // 消息的内容
}

// 定义两个消息 登录消息，登录成功与否的回复消息
type LoginMsg struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"usePwd"`
	UserName string `json:"username"`
}

type LonginResMsg struct {
	Code    int    `json:"code"`  // 返回状态码，500表示用户未注册，200表示登录成功
	UsersId []int  `json:usersId` // 增加字段，保存用户ID的切片
	Error   string `json:"error"` // 返回错误信息
}

type RegisterMsg struct {
	User model.User `json:"user"`
}

type RegisterResMsg struct {
	Code  int    `json:"code"`  // 返回状态码，400表示用户已占用，200表示注册成功
	Error string `json:"error"` // 返回错误信息
}

// 为了配合服务器端推送通知用户上线离线，定义一个新的消息类型
type NotifyUserStatusMsg struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

// 增加一个SmsMes
type SmsMes struct {
	Content    string `json:"content"`
	model.User        // 匿名结构体
}
