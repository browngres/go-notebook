# 11defer 陷阱

#### Markdown Notes 创建于 2024-07-26T09:41:54.842Z

## defer 与 闭包

```go
package main

import (
    "errors"
    "fmt"
)

func foo(a, b int) (i int, err error) {
    defer fmt.Printf("first defer err %v\n", err)
    defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
    defer func() { fmt.Printf("third defer err %v\n", err) }()
    if b == 0 {
        err = errors.New("divided by zero!")
        return
    }

    i = a / b
    return
}

func main() {
    foo(2, 0)
}
```

输出结果：

```
third defer err divided by zero!
second defer err <nil>
first defer err <nil>
```

解释：如果 defer 后面跟的不是一个 closure 最后执行的时候我们得到的并不是最新的值。

## defer 与 return

```go
func foo() (i int) {

    i = 0
    defer func() {
        fmt.Println(i)
    }()

    return 2
}
// 输出结果：    2
```

解释：在有具名返回值的函数中（这里指明返回值为 i），执行 return 2 的时候实际上已经将 i 的值重新赋值为 2。所以输出结果为 2 而不是 0。

## defer nil 函数

```go
func test() {
    var run func() = nil
    defer run()
    fmt.Println("runs")
}

func main() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()
    test()
}
```

输出结果：

```
runs
runtime error: invalid memory address or nil pointer dereference
```

解释：名为 test 的函数一直运行至结束，然后 defer 函数会被执行且会因为值为 nil 而产生 panic 异常。然而值得注意的是，run() 的声明是没有问题，因为在 test 函数运行完成后它才会被调用。

## 在错误的位置使用 defer

当 http.Get 失败时会抛出异常。

```go
package main

import "net/http"

func do() error {
    res, err := http.Get("http://www.google.com")
    defer res.Body.Close()
    if err != nil {
        return err
    }

    // ..code...

    return nil
}

func main() {
    do()
}
// 输出结果 panic: runtime error: invalid memory address or nil pointer dereference
```

因为在这里我们并没有检查我们的请求是否成功执行，当它失败的时候，我们访问了 Body 中的空变量 res ，因此会抛出异常。
解决方案：总是在一次成功的资源分配下面使用 defer ，对于这种情况来说意味着：当且仅当 http.Get 成功执行时才使用 defer

```go
if res != nil {
        defer res.Body.Close()
    }
```

在上述的代码中，当有错误的时候，err 会被返回，否则当整个函数返回的时候，会关闭 res.Body 。

解释：在这里，你同样需要检查 res 的值是否为 nil ，这是 http.Get 中的一个警告。通常情况下，出错的时候，返回的内容应为空并且错误会被返回，可当你获得的是一个重定向 error 时， res 的值并不会为 nil ，但其又会将错误返回。上面的代码保证了无论如何 Body 都会被关闭，如果你没有打算使用其中的数据，那么你还需要丢弃已经接收的数据。
