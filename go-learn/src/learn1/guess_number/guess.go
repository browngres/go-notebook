package main

import "fmt"
import "math/rand"

func main() {
	Rand_num := int(rand.Int()) % 100
	var i int = -1
	// fmt.Printf("正确数字是：%v\n", Rand_num)
	for i!=Rand_num {
		fmt.Println("输入数字(0-99):")
		n, err := fmt.Scanln(&i)

		if err == nil && n == 1 {
			if i == Rand_num {
				fmt.Println("对了")
				fmt.Printf("正确数字是：%v\n", Rand_num)
			}
			if i > Rand_num {
				fmt.Println("大了")
				continue
			}
			if i < Rand_num {
				fmt.Println("小了")
				continue
			}
		}

	}

}
