package main

import "fmt"

func fbn(n int8) []int {
	// n 代表下标，n从0开始
	fbn_slice := make([]int, n+1)
	if n == 0 || n == 1 {
		fbn_slice[n] = 1
		if n == 1 {
			fbn_slice[1-n] = 1
		}
		return fbn_slice
	}
	fbn_slice[0] = 1
	fbn_slice[1] = 1

	for i := int8(2); i <= n; i++ {
		fbn_slice[i] = fbn_slice[i-1] + fbn_slice[i-2]
	}
	return fbn_slice
}

func main() {
	res := fbn(10)
	fmt.Println(res)

}
