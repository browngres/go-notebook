// 使用8个线程，同时计算质数。
// 前面的示例是一读一算。写到最后一个写完完就关
// 难点是如何判断8个都已经完成（何时关闭计算结果通道）。
// 1000 以内的质数 有 168 个
package main

import (
	"fmt"
	// "time"
)

func putNum(intChan chan int, amount int) {
	// 放入指定个数的数字
	for i := 1; i <= amount; i++ {
		intChan <- i
	}
	close(intChan)
}

func prime(i int) bool {
	var j int
	if i == 1 {
		return false
	}
	for j = 2; j <= (i / j); j++ {
		if i%j == 0 {
			return false // 如果发现因子，则不是素数
		}
	}
	if j >= (i / j) {
		// fmt.Printf("%d  是素数\n", i)
		return true
	}
	return false
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		// time.Sleep(time.Millisecond * 10)    //不加这行会瞬间完成
		num, ok := <-intChan
		if !ok { //intChan 取不到..
			break
		}

		if prime(num) {
			//将这个数就放入到 primeChan
			primeChan <- num
		}
	}
	fmt.Println("有一个 primeNum 协程因为取不到数据，退出")
	//这里还不能关闭 primeChan
	//向 exitChan 写入 true
	exitChan <- true
}

func main() {
	intChan := make(chan int, 100)   // 放数字的管道只有100容量。放满就会阻塞，直到放完关闭。
	primeChan := make(chan int, 200) // 放质数的管道
	exitChan := make(chan bool, 8)   // //标识退出的管道，相当于 8 个 flag

	go putNum(intChan, 1000) //开启一个协程，向 intChan 放入 1000 个数

	//开启 8 个协程，从 intChan 取出数据，并判断是否为素数。如果是就放入到 primeChan
	for i := 0; i < 8; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	// 重点部分，从exitChan读到八个数据，就代表8个都已经完成。此时可以关闭 primeChan
	// 没读到就一直阻塞
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	// 最后遍历 primeChan， 取出结果。
	// 关闭通道，代表不能写入。不影响读取。
	//  “读关闭空零”——已经关闭的管道可以一直读空，再读就是零值。 未关闭的管道，读空，会阻塞。
	count := 0
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		count++
		// 将结果输出
		fmt.Printf("%4d 是质数\n", res)
	}
	fmt.Printf("统计到质数 %d 个\n", count)
	fmt.Println("main 线程退出")
}
