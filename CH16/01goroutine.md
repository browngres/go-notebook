# 01goroutine

#### Markdown Notes 创建于 2024-08-11T09:43:06.493Z

## 并发和并行

并发主要由切换时间片来实现"同时"运行，并行则是直接利用多核实现多线程的运行，go 可以设置使用核数，以发挥多核计算机的能力。

-   进程和线程
    -   进程是程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位。
    -   线程是进程的一个执行实体,是 CPU 调度和分派的基本单位,它是比进程更小的能独立运行的基本单位。
    -   一个进程可以创建和撤销多个线程；同一个进程中的多个线程之间可以并发执行。

多线程程序在一个核的 cpu 上运行，就是并发。多线程程序在多个核的 cpu 上运行，就是并行。

-   协程和线程

协程(routine)：独立的栈空间，共享堆空间，调度由用户自己控制，有点类似于用户级线程，这些用户级线程的调度也是自己实现的。
一个线程上可以跑多个协程，协程是轻量级的线程。

goroutine 是由官方实现的超级"线程池"。
每个实例 4~5KB 的栈内存占用和由于实现机制而大幅减少的创建和销毁开销是 go 高并发的根本原因。
操作系统线程一般都有固定的栈内存（通常为 2MB）,一个 goroutine 的栈在其生命周期开始时只有很小的栈（典型情况下 2KB），goroutine 的栈不是固定的，可以按需增大和缩小，大小可以达到 1GB，虽然极少会用到这么大。所以在 Go 语言中一次创建十万左右的 goroutine 也是可以的。
**goroutine 奉行通过通信来共享内存，而不是共享内存来通信。**

## goroutine

在 java/c++中我们要实现并发编程的时候，我们通常需要自己维护一个线程池，并且需要自己去包装一个又一个的任务，同时需要自己去调度线程执行任务并维护上下文切换，这一切通常会耗费程序员大量的心智。那么能不能有一种机制，程序员只需要定义很多个任务，让系统去帮助我们把这些任务分配到 CPU 上实现并发执行呢？
goroutine 的概念类似于线程，但 goroutine 是由 Go 的运行时（runtime）调度和管理的。Go 程序会智能地将 goroutine 中的任务合理地分配给每个 CPU。Go 语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。
在 Go 语言编程中你不需要去自己写进程、线程、协程。当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个 goroutine 去执行这个函数就可以了，就是这么简单。

## 使用 goroutine

    只需要在调用函数的时候在前面加上 go 关键字，就可以为一个函数创建一个 goroutine。一个 goroutine 必定对应一个函数，可以创建多个 goroutine 去执行相同的函数。

## 启动单个 goroutine

```go
func hello() {
    fmt.Println("Hello Goroutine!")
}
func main() {
    hello()
    fmt.Println("main goroutine done!")
}
```

语句是串行的，执行的结果是打印完 Hello Goroutine!后打印 main goroutine done!。

```go
func main() {
    go hello() // 启动另外一个goroutine去执行hello函数
    fmt.Println("main goroutine done!")
}
```

这一次的执行结果只打印了 main goroutine done!，并没有打印 Hello Goroutine!。为什么呢？
创建新的 goroutine 的时候需要花费一些时间，而此时 main 函数所在的 goroutine 是继续执行的。当 main()函数返回的时候该 goroutine 就结束了，所有在 main()函数中启动的 goroutine 会一同结束。
要想办法让 main 函数等一等 hello 函数。

```go
func main() {
    go hello() // 启动另外一个goroutine去执行hello函数
    fmt.Println("main goroutine done!")
    time.Sleep(time.Second)
}
```

这一次先打印 main goroutine done!，然后紧接着打印 Hello Goroutine!。

## 启动多个 goroutine

再来一个例子： （这里使用了 `sync.WaitGroup` 来实现 goroutine 的同步）

```go
var wg sync.WaitGroup

func hello(i int) {
    defer wg.Done() // goroutine结束就登记-1
    fmt.Println("Hello Goroutine!", i)
}
func main() {

    for i := 0; i < 10; i++ {
        wg.Add(1) // 启动一个goroutine就登记+1
        go hello(i)
    }
    wg.Wait() // 等待所有登记的goroutine都结束
}
```

多次执行上面的代码，会发现每次打印的数字的顺序都不一致。这是因为 10 个 goroutine 是并发执行的，而 goroutine 的调度是随机的。
