# 10defer 特性

#### Markdown Notes 创建于 2024-07-26T09:27:37.558Z

-   多个 defer 注册，按 FILO 次序执行 ( 先进后出 )。哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行。

    ```go
    func test(x int) {
        defer println("a")
        defer println("b")

        defer func() {
            println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
        }()

        defer println("c")
    }

    func main() {
        test(0)
    }
    输出结果:    c    b    a    panic: runtime error: integer divide by zero
    ```

-   延迟调用参数在注册时求值或复制，可用指针或闭包 "延迟" 读取。

    ```go
    func test() {
        x, y := 10, 20

        defer func(i int) {
            println("defer:", i, y) // y 闭包引用
        }(x) // x 被复制

        x += 10
        y += 100
        println("x =", x, "y =", y)
    }

    func main() {
        test()
    }
    // 输出结果:    x = 20 y = 120    defer: 10 120
    ```

-   滥用 defer 可能会导致性能问题，尤其是在一个“大循环”里。

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

var lock sync.Mutex

func test() {
    lock.Lock()
    lock.Unlock()
}

func testdefer() {
    lock.Lock()
    defer lock.Unlock()
}

func main() {
    func() {
        t1 := time.Now()

        for i := 0; i < 10000; i++ {
            test()
        }
        elapsed := time.Since(t1)
        fmt.Println("test elapsed: ", elapsed)
    }()
    func() {
        t1 := time.Now()

        for i := 0; i < 10000; i++ {
            testdefer()
        }
        elapsed := time.Since(t1)
        fmt.Println("testdefer elapsed: ", elapsed)
    }()

}
// 输出结果:    test elapsed:  223.162µs    testdefer elapsed:  781.304µs
```
