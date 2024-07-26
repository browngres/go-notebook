# 09 延迟调用 defer

#### Markdown Notes 创建于 2024-07-26T08:16:22.918Z

**提醒：大量超纲内容，以后会明白**

## defer 特性：

1. 关键字 defer 用于注册延迟调用。
2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
3. 多个 defer 语句，按先进后出的方式执行。
4. defer 语句中的变量，在 defer 声明时就决定了。

## defer 用途：

1. 关闭文件句柄
2. 锁资源释放
3. 数据库连接释放

defer 功能强大，对于资源管理非常方便，但是如果没用好，也会有陷阱。defer 是先进后出。这个很自然,后面的语句会依赖前面的资源，因此如果先前面的资源先释放了，后面的语句就没法执行了。

```go
func main() {
    var whatever [5]struct{}

    for i := range whatever {
        defer fmt.Println(i)
    }
}
// 输出结果：    4    3    2    1    0
```

## defer 结合闭包

```go
func main() {
    var whatever [5]struct{}
    for i := range whatever {
        defer func() { fmt.Println(i) }()
    }
}
// 输出结果：    4    4    4    4    4
```

> Each time a "defer" statement executes, the function value and parameters to the call are evaluated as usualand saved anew but the actual function is not invoked.
> 每次执行“defer”语句时，函数值和调用的参数都会像往常一样进行计算并重新保存，但实际的函数不会被调用。

也就是说函数正常执行，由于闭包用到的变量 i 在执行的时候已经变成 4，所以输出全都是 4。

## defer close

犯错的例子：

```go
type Test struct {
    name string
}

func (t *Test) Close() {
    fmt.Println(t.name, " closed")
}
func main() {
    ts := []Test{{"a"}, {"b"}, {"c"}}
    for _, t := range ts {
        defer t.Close()
    }
}
// 输出结果：    c  closed    c  closed    c  closed
```

并不会像我们预计的输出 c b a，而是输出 c c c。换一种方式来调用：

```go
type Test struct {
    name string
}

func (t *Test) Close() {
    fmt.Println(t.name, " closed")
}
func Close(t Test) {
    t.Close()
}
func main() {
    ts := []Test{{"a"}, {"b"}, {"c"}}
    for _, t := range ts {
        defer Close(t)
    }
}
// 输出结果：    c  closed    b  closed    a  closed
```

这个时候输出的就是 c b a。当然，如果你不想多写一个函数可以像下面这样，也很简单，同样输出 c b a：

```go
type Test struct {
    name string
}

func (t *Test) Close() {
    fmt.Println(t.name, " closed")
}
func main() {
    ts := []Test{{"a"}, {"b"}, {"c"}}
    for _, t := range ts {
        t2 := t
        defer t2.Close()
    }
}
```

通过以上例子和英文引用的句子，可以得出下面的结论：
defer 后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行。也就是复制了一份。但是并没有说 struct 这里的 this 指针如何处理，通过这个例子可以看出 go 语言并没有把这个明确写出来的 this 指针当作参数来看待。
