package main

import "fmt"

// 定义一个接口
type Usb interface {
	Start()
	Stop()
}

// 定义另外一个接口，让自定义结构体同事实现多个接口
type Wifi interface {
	Current_use()
}

// 定义手机结构体
type Phone struct {
}

// 为手机结构体定义方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

// Phone机构提实现WiFi接口
func (p Phone) Current_use() {
	fmt.Println("正在使用手机热点联网...")
}

// 定义相机结构体
type Camera struct {
}

// 为相机结构体定义方法
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

// 定义一个计算机结构体
type Computer struct {
}

// 编写一个方法Working，接收一个Usb接口类型变相
// 只要是实现了Usb接口：即实现了usb接口声明的所有方法
func (computer Computer) Working(usb Usb) {
	// 通过usb接口变量调用Start跟Stop方法
	usb.Start()
	usb.Stop()
}

func (computer Computer) Using(wifi Wifi) {
	wifi.Current_use()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	computer.Working(phone)
	computer.Working(camera)

	computer.Using(phone) // 调用另外一个接口
	// 使用另外一种方式，调用接口中定义的方法
	var b2 Wifi = phone
	b2.Current_use()
}
