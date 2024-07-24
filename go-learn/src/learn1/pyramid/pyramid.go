// 打印金字塔
package main

import "fmt"

func main() {
	var row int = 5
	var clo int = 9
	for i := 0; i < row; i++ {
		for j := 0; j < clo; j++ {
			if (clo/2-i) <= j && j <= (clo/2+i){ // 实心金字塔
				//(clo/2-i) == j || j == (clo/2+i) 空心金字塔
				fmt.Printf("* ")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println("")
	}
}
