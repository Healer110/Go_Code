package main

import "fmt"

func firstMenu() {
	fmt.Println("------------------家庭收支项目------------------")
	fmt.Println("                  1 收支明细")
	fmt.Println("                  2 登记收入")
	fmt.Println("                  3 登记支出")
	fmt.Println("                  4 退出软件")
	fmt.Println("-----------------------------------------------")
	fmt.Print("请选择(1-4): ")
}

type Account struct {
	Amount      float64 // 账户金额
	SzAmount    float64 // 收支金额
	description string  // 收支说明
	balance     float64 // 余额
}

// 收支明细菜单
func showAccountInfo(slice []Account) {
	fmt.Println()
	fmt.Println("账户金额\t收支金额\t收支说明\t账户余额")
	for _, v := range slice {
		fmt.Printf("%-7.2f\t%-7.2f\t%-7s\t%-7.2f\n", v.Amount, v.SzAmount, v.description, v.balance)
	}
}

// 收入菜单; 支出菜单。flag=true 收入，flag=false 支出
func option23Menu(flag bool) (float64, string) {
	var amount float64
	var descripton string
	if flag {
		fmt.Print("本次收入金额: ")
	} else {
		fmt.Print("本次支出金额: ")
	}
	fmt.Scanln(&amount)
	if flag {
		fmt.Print("本次收入说明: ")
	} else {
		fmt.Print("本次支出说明: ")
	}
	fmt.Scanln(&descripton)
	return amount, descripton
}

/*
家庭收支机长软件
*/
func main() {
	var optionNum int8 // 记录输入的选项
	var initAccount Account = Account{
		Amount:      10000,
		SzAmount:    0,
		description: "",
		balance:     10000,
	}

	var sliceAcount []Account
	sliceAcount = append(sliceAcount, initAccount)
	for {
		firstMenu()
		fmt.Scanln(&optionNum)
		tmp := sliceAcount[len(sliceAcount)-1].balance // 上一次记录的余额
		if optionNum == 4 {
			fmt.Println("退出.....")
			return
		} else if optionNum == 2 {
			amount, desp := option23Menu(true)
			sliceAcount = append(sliceAcount, Account{tmp, amount, desp, tmp + amount})
			fmt.Println("------------------登记完成------------------")
		} else if optionNum == 3 {
			amount, desp := option23Menu(false)
			if amount > tmp {
				fmt.Println("余额不足...")
				continue
			}
			sliceAcount = append(sliceAcount, Account{tmp, amount, desp, tmp - amount})
			fmt.Println("------------------登记完成------------------")
		} else if optionNum == 1 {
			showAccountInfo(sliceAcount)
		} else {
			fmt.Println("输入异常............")
		}
	}
}
