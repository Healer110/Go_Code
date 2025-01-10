package main

import "fmt"

func main() {
	account := Account{
		AccountNo: "JS62270001",
		Pwd:       "000000",
		Balance:   100,
	}

	account.BalanceInquiry("000000")
	account.Deposite(250, "000000")
	account.Withdraw(30, "000000")
	account.BalanceInquiry("000000")

}

// 定义结构体
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// 存款
func (account *Account) Deposite(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码错误...")
		return
	}

	if money <= 0 {
		fmt.Println("存款金额异常...")
		return
	}

	account.Balance += money
	fmt.Printf("存款%.2f成功...目前余额为：%.2f\n", money, account.Balance)
}

// 取款
func (account *Account) Withdraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码错误...")
		return
	}

	if money <= 0 {
		fmt.Println("取款金额异常...")
		return
	} else if money > account.Balance {
		fmt.Println("余额不足...")
		return
	}

	account.Balance -= money
	fmt.Printf("取款%.2f成功...目前余额为：%.2f\n", money, account.Balance)
}

// 余额查询
func (account *Account) BalanceInquiry(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码错误...")
		return
	}

	fmt.Printf("账号：%v目前余额为：%.2f\n", account.AccountNo, account.Balance)
}
