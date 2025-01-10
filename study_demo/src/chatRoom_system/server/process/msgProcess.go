package process

import (
	"encoding/json"
	"fmt"
	"net"
	"server/common/message"
	"server/utils"
)

// 处理消息相关的流程
type SmsProcess struct {
	// 暂时不添加字段
}

// 转发消息
func (sp *SmsProcess) SendGroupMsg(msg *message.Message) {
	// 遍历服务器端map[int]*UserProcess，并将消息转发出去

	// 取出消息内容
	var smsMsg message.SmsMes
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("[server/process/msgProcess.go] SendGroupMsg() error:", err)
		return
	}

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("[server/process/msgProcess.go] SendGroupMsg() error:", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		// 这里过滤掉自己
		if id == smsMsg.UserId {
			continue
		}
		sp.sendMsgToEachOnlineUser(data, up.Conn)
	}

}

// 执行发送动作
func (sp *SmsProcess) sendMsgToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("[server/process/msgProcess.go] sendMsgToEachOnlineUser() error:", err)
		return
	}

}
