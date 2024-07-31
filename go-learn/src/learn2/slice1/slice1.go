package main

import "fmt"

func main(){
    // 切片容量增加
    fmt.Println("切片容量增加")
    s := make([]int, 0, 1)
    c := cap(s)

    for i := 0; i < 50; i++ {
        s = append(s, i)
        if n := cap(s); n > c {
            fmt.Printf("cap: %d -> %d\n", c, n)
            c = n
        }
    }

    fmt.Println("======================")
    fmt.Println("索引切片 max")
    slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    d1 := slice[6:8]
    fmt.Println(d1, "len=",len(d1), "cap=",cap(d1))
    d2 := slice[2:6:8]
    fmt.Println(d2, "len=",len(d2), "cap=",cap(d2))


}



