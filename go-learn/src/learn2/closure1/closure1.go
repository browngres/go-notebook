package main

import "fmt"

func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func main() {
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	// 此时tmp1和tmp2不是一个实体了
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))
}
// 返回：
// 11 13
// 101 103