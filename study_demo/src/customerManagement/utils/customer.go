package utils

import "fmt"

// 定义客户所需信息的结构体
type CustomerInfo struct {
	Name      string
	Gender    string
	Age       uint8
	Telephone string
	Mail      string
}

type customerInfoSlice []CustomerInfo

// 格式化工厂返回切片
func NewCustomerInfoSlice() *customerInfoSlice {
	return new(customerInfoSlice)
}

// 主菜单
func (ci *customerInfoSlice) MainMenu() {
	fmt.Println("-----------------客户信息管理软件-----------------")
	fmt.Println("                  1 添加客户")
	fmt.Println("                  2 修改客户")
	fmt.Println("                  3 删除客户")
	fmt.Println("                  4 客户列表")
	fmt.Println("                  5 退    出")
	fmt.Println()
	fmt.Print("请选择(1-5): ")
}

// 添加客户信息
func (ci *customerInfoSlice) AddCustomer() {
	var tmp CustomerInfo
	fmt.Println("-----------------添加客户-----------------")
	fmt.Print("姓名: ")
	fmt.Scanln(&(tmp.Name))
	fmt.Print("性别: ")
	fmt.Scanln(&(tmp.Gender))
	fmt.Print("年龄: ")
	fmt.Scanln(&(tmp.Age))
	fmt.Print("电话: ")
	fmt.Scanln(&(tmp.Telephone))
	fmt.Print("邮箱: ")
	fmt.Scanln(&(tmp.Mail))
	*ci = append(*ci, tmp)
	fmt.Printf("-----------------添加完成-----------------\n\n")
}

// 修改客户信息
func (ci *customerInfoSlice) ModifyCustomer() {
	var tmpInput int
	fmt.Println("-----------------修改客户-----------------")
	for {
		fmt.Print("请选择待修改客户编号(-1退出): ")
		fmt.Scanln(&tmpInput)
		if tmpInput == -1 {
			return
		}

		if tmpInput < 1 || tmpInput > len(*ci) {
			fmt.Println("未查询到该编码对应的客户...")
			continue
		}

		var name, gender, telephone, mail string
		var age uint8
		customer := (*ci)[tmpInput-1]

		fmt.Printf("姓名(%v): ", customer.Name)
		fmt.Scanln(&name)
		if name != "" {
			customer.Name = name
		}

		fmt.Printf("性别(%v): ", customer.Gender)
		fmt.Scanln(&gender)
		if gender != "" {
			customer.Gender = gender
		}

		fmt.Printf("年龄(%v): ", customer.Age)
		fmt.Scanln(&age)
		if age != 0 {
			customer.Age = age
		}

		fmt.Printf("电话(%v): ", customer.Telephone)
		fmt.Scanln(&telephone)
		if telephone != "" {
			customer.Telephone = telephone
		}

		fmt.Printf("邮箱(%v): ", customer.Mail)
		fmt.Scanln(&mail)
		if mail != "" {
			customer.Mail = mail
		}
		(*ci)[tmpInput-1] = customer
		fmt.Println("-----------------修改完成-----------------")
		fmt.Println()
	}

}

// 删除客户信息
func (ci *customerInfoSlice) DeleteCustomer() {
	var tmpInput int
	fmt.Println("-----------------删除客户-----------------")
	for {
		fmt.Print("请选择待删除客户编号(-1退出): ")
		fmt.Scanln(&tmpInput)
		if tmpInput == -1 {
			return
		}
		if tmpInput > 0 && tmpInput <= len(*ci) {
			fmt.Print("请确认是否删除(y/n): ")
			var ok string
			fmt.Scanln(&ok)
			if ok == "n" {
				continue
			}
			tmpInput--
			if tmpInput == 0 {
				*ci = (*ci)[1:]
			} else if tmpInput == len(*ci)-1 {
				*ci = (*ci)[:len(*ci)-1]
			} else {
				*ci = append((*ci)[:tmpInput], (*ci)[tmpInput+1:]...)
			}
			fmt.Println("-----------------删除完成-----------------")
			fmt.Println()
		} else {
			fmt.Println("输入编号错误...")
			fmt.Println()
		}
	}
}

// 展示客户信息
func (ci customerInfoSlice) ShowCustomer() {
	if len(ci) == 0 {
		fmt.Printf("客户资源为空...\n\n")
		return
	}
	fmt.Println("-----------------客户列表-----------------")
	fmt.Printf("编号\t姓名\t性别\t年龄\t电话\t邮箱\n")
	for idx, cus := range ci {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", idx, cus.Name, cus.Gender, cus.Age, cus.Telephone, cus.Mail)
	}
	fmt.Printf("---------------客户列表完成-----------------\n\n")
}
