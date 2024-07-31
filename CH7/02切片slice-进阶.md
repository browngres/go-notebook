# 02 切片 slice(进阶)

#### Markdown Notes 创建于 2024-07-31T17:06:34.014Z

### 索引切片 max

`s[low:high:max]` 从切片 s 的 low 到 high，索引出新的切片。
新的切片 len=high-low，cap=max-low

```go
slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
d1 := slice[6:8]
fmt.Println(d1, len(d1), cap(d1))
d2 := slice[:6:8]
fmt.Println(d2, len(d2), cap(d2))
```

## append 切片追加

append ：向 slice 尾部添加数据，返回新的 slice 对象。

```go
var a = []int{1, 2, 3}
var b = []int{4, 5, 6}
c := append(a, b...)
fmt.Printf("slice c : %v\n", c)
d := append(c, 7)
fmt.Printf("slice d : %v\n", d)
//    slice c : [1 2 3 4 5 6]
//    slice d : [1 2 3 4 5 6 7]
```

```go
s1 := make([]int, 0, 5)
fmt.Printf("%p\n", &s1)
s2 := append(s1, 1)
fmt.Printf("%p\n", &s2)
fmt.Println(s1, s2)
//     0xc42000a060
// 0xc42000a080
// [] [1]
```

-   超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满

```go
data := [...]int{0, 1, 2, 3, 4, 10: 0}
s := data[:2:3]

s = append(s, 100, 200) // 一次 append 两个值，超出 s.cap 限制。

fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。
// 输出结果 ：
//    [0 1 100 200] [0 1 2 3 4 0 0 0 0 0 0]
//    0xc4200160f0 0xc420070060
```

从输出结果可以看出，append 后的 s 重新分配了底层数组，并复制数据。如果只追加一个值，则不会超过 s.cap 限制，也就不会重新分配。 通常以 2 倍容量重新分配底层数组。在大批量添加数据时，建议一次性分配足够大的空间，以减少内存分配和数据复制开销。或初始化足够长的 len 属性，改用索引号进行操作。及时释放不再使用的 slice 对象，避免持有过期数组，造成 GC 无法回收。

-   重新分配的数组长度

```go
s := make([]int, 0, 1)
c := cap(s)

for i := 0; i < 50; i++ {
    s = append(s, i)
    if n := cap(s); n > c {
        fmt.Printf("cap: %d -> %d\n", c, n)
        c = n
    }
}
```

```
输出结果:
cap: 1 -> 2
cap: 2 -> 4
cap: 4 -> 8
cap: 8 -> 16
cap: 16 -> 32
cap: 32 -> 64
```

### 切片 copy

```go
s1 := []int{1, 2, 3, 4, 5}
fmt.Printf("slice s1 : %v\n", s1)
s2 := make([]int, 10)
fmt.Printf("slice s2 : %v\n", s2)
copy(s2, s1)
fmt.Printf("copied slice s1 : %v\n", s1)
fmt.Printf("copied slice s2 : %v\n", s2)
s3 := []int{1, 2, 3}
fmt.Printf("slice s3 : %v\n", s3)
s3 = append(s3, s2...)
fmt.Printf("appended slice s3 : %v\n", s3)
s3 = append(s3, 4, 5, 6)
fmt.Printf("last slice s3 : %v\n", s3)
```

```
输出结果：
slice s1 : [1 2 3 4 5]
slice s2 : [0 0 0 0 0 0 0 0 0 0]
copied slice s1 : [1 2 3 4 5]
copied slice s2 : [1 2 3 4 5 0 0 0 0 0]
slice s3 : [1 2 3]
appended slice s3 : [1 2 3 1 2 3 4 5 0 0 0 0 0]
last slice s3 : [1 2 3 1 2 3 4 5 0 0 0 0 0 4 5 6]
```

copy 在两个 slice 间复制数据，复制长度以 len 小的为准。两个 slice 可指向同一底层数组，允许元素区间重叠。

```go
data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
fmt.Println("array data : ", data)
s1 := data[8:]
s2 := data[:5]
fmt.Printf("slice s1 : %v\n", s1)
fmt.Printf("slice s2 : %v\n", s2)
copy(s2, s1)
fmt.Printf("copied slice s1 : %v\n", s1)
fmt.Printf("copied slice s2 : %v\n", s2)
fmt.Println("last array data : ", data)
```

```
输出结果:
array data :  [0 1 2 3 4 5 6 7 8 9]
slice s1 : [8 9]
slice s2 : [0 1 2 3 4]
copied slice s1 : [8 9]
copied slice s2 : [8 9 2 3 4]
last array data :  [8 9 2 3 4 5 6 7 8 9]
```

应及时将所需数据 copy 到较小的 slice，以便释放超大号底层数组内存。

### 重新索引，调整大小

```go
var a = []int{1, 3, 4, 5}
fmt.Printf("slice a : %v , len(a) : %v\n", a, len(a))
b := a[1:2]
fmt.Printf("slice b : %v , len(b) : %v\n", b, len(b))
c := b[0:3]
fmt.Printf("slice c : %v , len(c) : %v\n", c, len(c))
//    slice a : [1 3 4 5] , len(a) : 4
//    slice b : [3] , len(b) : 1
//    slice c : [3 4 5] , len(c) : 3
```

结合切片和数组内存中的布局就能明白。
![切片的内存布局](https://www.topgoer.com/static/3.8/2.jpg)

### 字符串切片

string 底层就是一个 byte 的数组，因此，也可以进行切片操作。

```go
str := "Hello world"
s := []byte(str) //中文字符需要用[]rune(str)
s[6] = 'G'
s = s[:8]
s = append(s, '!')
str = string(s)
fmt.Println(str)
//    Hello Go!
```

-   数组 or 切片 转字符串：
    `strings.Replace(strings.Trim(fmt.Sprint(array_or_slice), "[]"), " ", ",", -1)`
