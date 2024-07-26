package main

import "fmt"

// 返回2个函数类型的返回值
func test01(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := test01(10)
	// base仍然存活
	fmt.Println(f1(1), f2(2))
	// 此时base是9
	fmt.Println(f1(3), f2(4))
}
