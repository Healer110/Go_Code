package utils

import (
	"chatRoom_system/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

// 将这些函数封装到结构体中, 并为结构体绑定方法，供外部函数调用
type Transfer struct {
	Conn net.Conn
	Buf  [8192]byte
}

// 放置工具类，代码通用的部分放置在这里，供其他函数调用

// 收包
func (trans *Transfer) ReadPkg() (msg message.Message, err error) {
	// 服务端处理接收到的包
	// buf := make([]byte, 8192)
	// Read是阻塞方法,只有在trans.Conn没有被关闭的情况下才会阻塞
	// 如果客户端关闭了，就不会阻塞了，就会发生异常
	_, err = trans.Conn.Read(trans.Buf[:4])
	if err != nil {
		return
	}

	// 根据buf[:4]转成一个uint32类型
	var pkgLen uint32 = binary.BigEndian.Uint32(trans.Buf[:4])

	// 根据上面计算的长度，读取消息
	// 解释，trans.Conn读取pkgLen个字节到buf中，最后返回的n不一定等于pkgLen，可以对返回的n做进一步判断
	n, err := trans.Conn.Read(trans.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("读取消息异常。。。", n, pkgLen, err)
		return
	}

	// 将得到的数据反序列化得到message类型，得到message中的data，然后将data再反序列化成LoginMsg
	err = json.Unmarshal(trans.Buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("反序列化失败...", err)
		err = errors.New("服务端读序列化消息体异常")
		return
	}
	return
}

// 发包
func (trans *Transfer) WritePkg(data []byte) (err error) {
	//  先发送长度给客户端
	var pkgLen uint32 = uint32(len(data))
	// var bytes [4]byte
	// 将一个数字转换为字节切片
	binary.BigEndian.PutUint32(trans.Buf[:4], pkgLen)
	n, err := trans.Conn.Write(trans.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("writePkg 发送数据长度失败：", err)
		return
	}

	n, err = trans.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("trans.Conn.Write 发送data失败：", err)
		return
	}
	return
}
