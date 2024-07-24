// 打印乘法表
package main

import "fmt"

func main() {
	zx()
	fmt.Println("==============================")
	zs()
}

func zx() {//左上
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v✕%v=%-2v ", i, j, i*j)
		}
		fmt.Println("")
	}
}

func zs() {//左下
	for i := 9; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v✕%v=%-2v ", i, j, i*j)
		}
		fmt.Println("")
	}
}
