package main

import (
	"customerManagement/utils"
	"fmt"
)

// 客户信息管理系统
// 实现对客户对象的插入、修改、删除(用切片实现)，并能够打印客户明细
// 多对象协同工作原理
func main() {
	var inputNumber int8
	// 获取存放客户信息的切片
	customerSlice := utils.NewCustomerInfoSlice()
	for {
		customerSlice.MainMenu()
		switch fmt.Scanln(&inputNumber); inputNumber {
		case 1:
			// 添加
			customerSlice.AddCustomer()
		case 2:
			// 修改
			customerSlice.ModifyCustomer()
		case 3:
			// 删除
			customerSlice.DeleteCustomer()
		case 4:
			// 客户列表
			customerSlice.ShowCustomer()
		case 5:
			fmt.Println("退出系统...")
			return
		default:
			fmt.Println("输入异常...")
		}
	}
}
