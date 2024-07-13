package main

import (
	"fmt"
	"math"
)

func main() {
	// 求`ax2+bx+c=0`方程的根。a,b,c分别为函数的参数，如果：`b^2-4ac>0`，则有两个解；
	// `b2-4ac=0`，则有一个解；`b2-4ac < 0`，则无解；
	// x1=(-b+sqrt(b2-4ac))/2a             x2=(-b-sqrt(b2-4ac))/2a

	var a float64 = 9.0
	var b float64 = -3.0
	var c float64 = -2.0
	// 9x^2-6x-2=0
	m := b*b - 4*a*c
	// 多分支判断
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		x2 := (-b - math.Sqrt(m)) / (2 * a)
		fmt.Printf("x1=%v x2=%v \n", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		fmt.Printf("x1=x2=%v \n", x1)
	} else {
		fmt.Println("无解")
	}

}
