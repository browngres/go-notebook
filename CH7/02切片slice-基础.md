# 02 切片 slice(基础)

#### Markdown Notes 创建于 2024-07-31T16:34:05.062Z

slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。

-   切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
-   切片的长度可以改变。
-   切片遍历方式和数组一样，可以用 `len()` 求长度。表示可用元素数量，读写操作不能超过该限制。
-   cap 可以求出 slice 最大扩张容量，不能超出数组限制。`0 <= len(slice) <= len(array)`，其中 array 是 slice 引用的数组。
-   总是有 len(s) <= cap(s)
-   切片的定义：`var 变量名 []类型`，比如 `var str []string`、`var arr []int`。
-   如果 `slice == nil`，那么 len、cap 结果都等于 0。
-   左闭右开，和 Python 语言索引一样
-   同样可以使用 for-range 遍历

```go
func main() {
    // 声明切片
    var s1 []int
    s2 := []int{}
    var s3 []int = make([]int, 0)
    s4 := []int{1, 2, 3}
    // 从数组切片
   arr := [5]int{1, 2, 3, 4, 5}
   var s5 []int
   s5 = arr[1:4]
}
```

### 切片使用

```go
// 全局：
var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arr[start:end]
var slice1 []int = arr[:end]
var slice2 []int = arr[start:]
var slice3 []int = arr[:]
var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素
// 局部：  同理使用 := 声明变量
```

`[:]` 的使用方法和 Python 几乎相同

-   读写操作实际目标是底层数组，只需注意索引号的差别。

```go
func main() {
    data := [...]int{0, 1, 2, 3, 4, 5}

    s := data[2:4]
    s[0] += 100
    s[1] += 200

    fmt.Println(s)
    fmt.Println(data)
}
```

-   可直接创建 slice 对象，自动分配底层数组。

```go
s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
fmt.Println(s1, len(s1), cap(s1))
// [0 1 2 3 0 0 0 0 100] 9 9
```

### 使用 make 创建切片

```go
// cap 可以省略
var slice []type = make([]type, len)
slice  := make([]type, len)
slice  := make([]type, len, cap)
```

![slice](https://www.topgoer.com/static/3.8/1.jpg)
![切片的内存布局](https://www.topgoer.com/static/3.8/2.jpg)

-   可用指针直接访问底层数组，退化成普通数组操作。
-   make 创建的切片对应的数组是由 make 底层维护，对外不可见，即只能通过 slice 访问各个元素.

```go
s := []int{0, 1, 2, 3}
p := &s[2] // *int, 获取底层数组元素指针。
*p += 100
//  [0 1 102 3]
```

### 多维

-   切片里面类型也是切片

```go
data := [][]int{
    []int{1, 2, 3},
    []int{100, 200},
    []int{11, 22, 33, 44},
}
```

-   可直接修改 struct array/slice 成员

```go
d := [5]struct {
    x int
}{}

s := d[:]

d[1].x = 10
s[2].x = 20

fmt.Println(d)
fmt.Printf("%p, %p\n", &d, &d[0])
// 输出结果
// [{0} {10} {20} {0} {0}]
// 0xc4200160f0, 0xc4200160f0
```
