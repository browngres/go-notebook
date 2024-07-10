package main

import (
	"fmt"
	"unsafe"
)

func main (){
	// 查看数据类型
	var n1 int = 666
	fmt.Printf("n1 的类型是 %T \n", n1)
	// 查看使用的字节数
	var n2 float64 = 3.14
	fmt.Printf("n2 的字节数 %d \n", unsafe.Sizeof(n2))

	// 浮点数
	var n3 float32 = -123.0000901
	var n4 float64 = -123.0000901
	fmt.Println("n3=",n3)
	fmt.Println("n4=",n4)
}