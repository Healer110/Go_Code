package process

import (
	"chatRoom_system/client/model"
	"chatRoom_system/common/message"
	"fmt"
)

// 客户端维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser // 在用户登录成功后，完成对该变量的初始化工作

// 在客户端显示当前在线用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表: ")
	for id, _ := range onlineUsers {
		fmt.Printf("用户id: %d\n", id)
	}
	fmt.Println()
}

func updateUserStatus(notifyUserStatusMsg *message.NotifyUserStatusMsg) {

	user, ok := onlineUsers[notifyUserStatusMsg.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMsg.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMsg.Status
	// 更新客户端维护的map
	onlineUsers[notifyUserStatusMsg.UserId] = user

	outputOnlineUser()
}
