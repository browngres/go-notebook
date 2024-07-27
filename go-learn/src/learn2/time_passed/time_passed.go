package main

import (
	"fmt"
	"strconv"
	"time"
)


func test() {

	str := ""
	for i := 0; i < 10000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	start := time.Now().UnixMilli()
	test()
	end := time.Now().UnixMilli()
	fmt.Printf("耗费时间为%v毫秒\n", end-start)
}