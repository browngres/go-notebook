package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	// 查看数据类型
	var n1 int = 666
	fmt.Printf("n1 的类型是 %T \n", n1)
	// 查看使用的字节数
	var n2 float64 = 3.14
	fmt.Printf("n2 的字节数 %d \n", unsafe.Sizeof(n2))

	// 浮点数
	var n3 float32 = -123.0000901
	var n4 float64 = -123.0000901
	fmt.Println("n3=", n3)
	fmt.Println("n4=", n4)

	data_type_trans()
	use_strconv()
}

func data_type_trans() {
	var n1 int = 99
	var f1 float64 = 23.456
	var b1 bool = true
	var c1 byte = 'h'
	var s1 string //默认值为空的str

	s1 = fmt.Sprintf("%d", n1) // 将n1转换为字符串给s1
	fmt.Printf("s1 type %T s1=%q\n", s1, s1)

	s1 = fmt.Sprintf("%f", f1)
	fmt.Printf("s1 type %T s1=%q\n", s1, s1)

	s1 = fmt.Sprintf("%t", b1)
	fmt.Printf("s1 type %T s1=%q\n", s1, s1)

	s1 = fmt.Sprintf("%c", c1)
	fmt.Printf("s1 type %T s1=%q\n", s1, s1)
}

func use_strconv() {

	fmt.Println("===value to string===")
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-42, 16)
	s4 := strconv.FormatUint(42, 16)
	fmt.Println(string(s1))
	fmt.Println(string(s2))
	fmt.Println(string(s3))
	fmt.Println(string(s4))

	fmt.Println("===string to value===")
	// 他们都会返回两个值，第二个是err。 如果不想要，可以这样写
	// b, _ := strconv.ParseBool("true")
	b, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Printf("%T, %v\n", b, b)
	}
	f, err := strconv.ParseFloat("3.1415", 64)
	if err == nil {
		fmt.Printf("%T, %v\n", f, f)
	}
	i, err := strconv.ParseInt("-42", 10, 64)
	if err == nil {
		fmt.Printf("%T, %v\n", i, i)
	}
	u, err := strconv.ParseUint("42", 10, 64)
	if err == nil {
		fmt.Printf("%T, %v\n", u, u)
	}

}
