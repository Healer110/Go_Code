package main

import (
	"fmt"
	"io"
	"net"
	"server/common/message"
	"server/process"
	"server/utils"
)

// 主控制器，一般放在跟main.go同一个目录下
type Processor struct {
	Conn net.Conn
}

// 编写一个ServerProcessMsg函数
// 功能：根据客户端发送的消息种类不同，决定调用哪个函数来处理
func (processor *Processor) ServerProcessMsg(msg *message.Message) (err error) {
	switch msg.Type {
	case message.LoginMsgType:
		// 处理登录的逻辑
		up := &process.UserProcess{
			Conn: processor.Conn,
		}
		err = up.ServerProcessLogin(msg)
		// err = ServerProcessLogin(processor.Conn, msg)
	case message.RegisterMsgType:
		// 处理注册逻辑
		up := &process.UserProcess{
			Conn: processor.Conn,
		}
		err = up.ServerProcessRegister(msg)
	case message.SmsMesType:
		smsProcess := &process.SmsProcess{}
		smsProcess.SendGroupMsg(msg)

	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

// 主控制器，main函数调用主控制器，主控制器调用ServerProcessMsg()，
// 然后判断不同的消息类型，调用不同的模块进行处理
func (processor *Processor) ProcessController() (err error) {
	for {
		// 创建transfer实例，完成包的读取操作
		tf := &utils.Transfer{
			Conn: processor.Conn,
		}
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("Read EOF...")
			} else {
				fmt.Println("readPkg err: ", err)
			}
			return err
		}

		// fmt.Println(msg)

		err = processor.ServerProcessMsg(&msg)
		if err != nil {
			fmt.Println("process err: ", err)
			return err
		}
	}
}
