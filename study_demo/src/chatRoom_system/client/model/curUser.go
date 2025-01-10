package model

import (
	"chatRoom_system/common/message"
	"net"
)

// 在客户端，很多地方会使用到curUser，这里我们做成全局变量, 放在process/userMgr.go中

type CurUser struct {
	Conn net.Conn
	message.User
}
