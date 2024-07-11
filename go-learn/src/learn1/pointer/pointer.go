package main

import "fmt"

func main() {
	/*
	var a *int
    *a = 100
    fmt.Println(*a)
	// panic: runtime error: invalid memory address or nil pointer dereference
	*/
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(*a)

	/*
	var b map[string]int
    b["测试"] = 100
    fmt.Println(b)
	*/
	var b map[string]int
	b = make(map[string]int, 10)
	b["测试"] = 100
	b["go"] = 200
	fmt.Println(b)
}
