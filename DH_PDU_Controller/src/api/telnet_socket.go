package api

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

type PduSocket struct {
	ConnetStatus bool
	Conn         *net.Conn
}

func (pduSocket *PduSocket) Connect(ip string, port int) bool {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, time.Second*5)
	if err != nil {
		PrintErrors(fmt.Sprintf("failed to connect to %s:%d: %s", ip, port, err))
		return false
	}
	pduSocket.Conn = &conn
	pduSocket.ConnetStatus = true
	return true
}

// 发送命令
func (pduSocket *PduSocket) Send(cmd string) bool {
	if pduSocket.Conn == nil {
		PrintErrors("connection is nil")
		return false
	}
	_, err := (*pduSocket.Conn).Write([]byte(cmd + "\n"))
	if err != nil {
		pduSocket.ConnetStatus = false
		PrintErrors(fmt.Sprintf("failed to send command %s: %s", cmd, err))
		return false
	}
	return true
}

// 接收数据
func (pduSocket *PduSocket) Receive() (string, error) {
	if pduSocket.Conn == nil {
		PrintErrors("connection is nil")
		return "", fmt.Errorf("connection is nil")
	}
	reader := bufio.NewReader(*pduSocket.Conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		if err != io.EOF {
			pduSocket.ConnetStatus = false
			PrintErrors(fmt.Sprintf("failed to receive data: %s", err))
			return "", err
		}
	}
	return line, nil
}

func (pduSocket *PduSocket) Disconnect() {
	if pduSocket.Conn == nil {
		PrintErrors("connection is nil")
		return
	}
	(*pduSocket.Conn).Close()
	pduSocket.ConnetStatus = false
}
