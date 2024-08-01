package main

import "fmt"

// 已知一个升序数组，插入一个数字， 仍然升序。最后打印

func sorted_insert(num int64, arr *[10]int64) (place int64) {
	// 二分找到目标位置
	start := 0
	end := len(arr) - 1
	for {
		mid := (end + start) / 2
		fmt.Println("mid=", mid, "arr[mid]=", arr[mid])
		if arr[mid] < num {
			start = mid
			// fmt.Println("start=", start)
			// fmt.Println("end=", end)
		} else {
			end = mid
			// fmt.Println("start=", start)
			// fmt.Println("end=", end)
		}
		if (mid+1) == end || (start+1) == mid {
			// fmt.Println("start=", start)
			// fmt.Println("end=", end)
			fmt.Println("目标位置:", start)
			break
		}
	}
	return int64(start)

}
func main() {
	var arr = [10]int64{1, 3, 22, 32, 54, 61, 88, 100, 201, 408}
	var num int64 = -66
	fmt.Println(arr)
	fmt.Println("要插入的数字为", num)
	place := sorted_insert(num, &arr)
	if num < arr[0] {
		// 有一个bug，数字只会往后面插，如果数字比第一个还小，也后插第一个后面。
		// 如果比第一个还小，直接插在前面。
		fmt.Println("比第一个还小")
		slice := make([]int64, len(arr)+1)
		slice[0] = num
		copy(slice[1:], arr[:])
		fmt.Println("原始数组：", arr)
		fmt.Println("插入数字后", slice)
	} else {
		slice := make([]int64, place+1) // 创建一个新切片，长度+1
		//这里不能使用原始数组的切片，因为切片的 append 会破坏原数组。
		// 用append插入数字，数组中下个位置的数字就被顶掉了。
		// 因此必须 make 新的切片
		copy(slice, arr[:place+1]) // 拷贝前半部分
		// fmt.Println(slice)
		slice = append(slice, num)              // 插入数字
		slice = append(slice, arr[place+1:]...) //插入数组后半部分
		fmt.Println("原始数组：", arr)
		fmt.Println("插入数字后", slice)
	}
}

// 20240802 0109 耗时1个小时
