package main

import (
	"account/account"
	"fmt"
)

func main() {
	// 创建一个账户
	a1 := account.NewAccount("188888", "123456", 8888.8)
	if a1 != nil {
		fmt.Println("创建成功:", a1)
	} else {
		fmt.Println("创建失败")
	}
	a1.GetBalance("123456")
	// 存款
	a1.Deposit(999, "1234567")  // 密码输入错误
	a1.Deposit(10000, "123456") // 密码正确
	a1.GetBalance("123456")
	// fmt.Println(a1.pk) // 私有的，获取不到
}
