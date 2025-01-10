package process

import (
	"chatRoom_system/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMsg(msg *message.Message) {
	// 显示消息
	var smsMsg message.SmsMes
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("[client/process/smsMgr.go] error:", err)
		return
	}

	info := fmt.Sprintf("用户id: %d, 对大家说：%s", smsMsg.UserId, smsMsg.Content)
	fmt.Println(info)
	fmt.Println()

}
