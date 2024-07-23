# 04range

#### Markdown Notes 创建于 2024-07-23T11:51:12.511Z
Golang 提供 for-range 的方式，可以方便遍历字符串和数组(注: 数组的遍历，我们放到讲数组
的时候再讲解) ，案例说明如何遍历字符串。

```go
var str string = "hello,world!北京"
for i := 0; i < len(str); i++ {
	fmt.Printf("%c \n", str[i]) //使用到下标...
}
```
汉字出错。传统的对字符串的遍历是按照字节来遍历，而一个汉字在 utf8 编码是对应 3 个字节。
需要要将 str 转成 `[]rune` 切片，`str2 := []rune(str)`

使用range，按照字符方式遍历，汉字也可以。
```go
str = "abc~OK上海"
for index, val := range str {
    fmt.Printf("index=%d, val=%c \n", index, val)
}
```
- 数字数组
```go
numbers := [6]int{1, 2, 3, 5}

for i,x:= range numbers {
    fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
}
```