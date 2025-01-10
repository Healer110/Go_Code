package process

import (
	"chatRoom_system/client/utils"
	"chatRoom_system/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

// 发送群聊的消息
func (sp *SmsProcess) SendGroupMes(content string) (err error) {

	var msg message.Message
	msg.Type = message.SmsMesType
	var smsMsg message.SmsMes
	smsMsg.Content = content
	smsMsg.UserId = CurUser.UserId
	smsMsg.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("[client/process/smsProcess.go] son.Marshal(smsMsg) err:", err)
		return
	}

	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("[client/process/smsProcess.go] son.Marshal(msg) err:", err)
		return
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("[client/process/smsProcess.go] conn.Write 发送data失败：", err)
		return
	}
	return
}
