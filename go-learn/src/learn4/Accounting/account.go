package main

import (
	"fmt"
)

type bill struct {
	direction bool // true 表示支出， false 表示收入
	money     uint64
	note      string
}

func (b *bill) Print(index int) {
	// bill 拥有打印自己的方法
	if b.direction {
		fmt.Printf("第%d条：  支出  %5d元   %v\n", index+1, b.money, b.note)
	} else {
		fmt.Printf("第%d条：  收入  %5d元   %v\n", index+1, b.money, b.note)
	}
}

func PrintBills(b *[]bill) {
	// 打印所有账单
	fmt.Println("         收支   金额    说明")
	for index, v := range *b {
		v.Print(index)
	}
	fmt.Println("########################")
}

func newBill(direction bool, money uint64, note string) bill {
	n := bill{direction, money, note}
	return n
}

func NewIncome(b *[]bill) {
	fmt.Println("输入收入金额：")
	var money uint64
	fmt.Scanln(&money)
	fmt.Println("输入备注：")
	var note string
	fmt.Scanln(&note)
	*b = append(*b, newBill(false, money, note))
	income += money
	balance += money
	fmt.Println("添加完成")
}

func NewExpense(b *[]bill) {
	fmt.Println("输入支出金额：")
	var money uint64
	fmt.Scanln(&money)
	fmt.Println("输入备注：")
	var note string
	fmt.Scanln(&note)
	*b = append(*b, newBill(true, money, note))
	expense += money
	balance -= money
	fmt.Println("添加完成")

}
