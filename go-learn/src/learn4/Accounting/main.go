package main

import (
	"fmt"
)

var (
	balance uint64 = 88888
	expense uint64 = 0
	income  uint64 = 0
)

func printMenu() {
	fmt.Println("----------家庭收支记账软件----------")
	fmt.Println("         1. 收支明细")
	fmt.Println("         2. 添加收入")
	fmt.Println("         3. 添加支出")
	fmt.Println("         4. 退    出")
	fmt.Println("----------------------------------")
	fmt.Printf("余额：%d￥  总收入：%d￥  总支出：%d￥\n", balance, income, expense)
	fmt.Println("----------------------------------")
}

func main() {
	bills := []bill{}
	printMenu()
menu:
	for {
		var op uint8 = 0
		fmt.Println("输入选项：")
		fmt.Scanln(&op)
		switch op {
		case 1:
			PrintBills(&bills)
		case 2:
			NewIncome(&bills)
		case 3:
			NewExpense(&bills)
		case 4:
			fmt.Println("退出")
			break menu
		default:
			break menu
		}

	}

	// bills := []bill{}
	// bills = append(bills, bill{true, 100, "红包100"})
	// bills = append(bills, bill{true, 1000, "红包1000"})
	// fmt.Println(bills)
	// PrintBills(bills)
}
