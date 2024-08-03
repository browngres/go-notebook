package account

import (
	"fmt"
)

// 定义一个结构体 account
type account struct {
	pk      string
	pwd     string
	balance float64
}

// 构造函数
func NewAccount(pk string, pwd string, balance float64) *account {

	if len(pk) < 6 || len(pk) > 10 {
		fmt.Println("账号的长度不对...")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码的长度不对...")
		return nil
	}
	if balance < 0 {
		fmt.Println("余额数目不对...")
		return nil
	}
	return &account{
		pk:      pk,
		pwd:     pwd,
		balance: balance,
	}
}

// 方法
// 存款
func (a *account) Deposit(money float64, pwd string) {
	if pwd != a.pwd {
		fmt.Println("密码不正确")
		return
	}
	if money <= 0 {
		fmt.Println("输入金额不正确")
		return
	}
	a.balance += money
	fmt.Println("存款成功")
}

// 取款
func (a *account) WithDraw(money float64, pwd string) {
	if pwd != a.pwd {
		fmt.Println("密码不正确")
		return
	}
	if money <= 0 || money > a.balance {
		fmt.Println("输入金额不正确")
		return
	}
	a.balance -= money
	fmt.Println("取款成功")
}

// 查询余额
func (a *account) GetBalance(pwd string) {
	if pwd != a.pwd {
		fmt.Println("密码不正确")
		return
	}
	fmt.Printf("账号为=%v 余额=%v \n", a.pk, a.balance)
}
