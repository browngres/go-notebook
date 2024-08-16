# 04select

#### Markdown Notes 创建于 2024-08-16T02:09:15.097Z

某些场景下我们需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以接收将会发生阻塞。
Go 内置了 select 关键字，可以同时响应多个通道的操作。
select 的使用类似于 switch 语句，它有一系列 case 分支和一个默认的分支。每个 case 会对应一个通道的通信（接收或发送）过程。select 会一直等待，直到某个 case 的通信操作完成时，就会执行 case 分支对应的语句。具体格式如下：

```go
select {
    case <- chan1:
        // 如果chan1成功读到数据，则进行该case处理语句
    case chan2 <- 1:
        // 如果成功向chan2写入数据，则进行该case处理语句
    default:
        // 如果上面都没有成功，则进入default处理流程
}
```

-   select 可以同时监听一个或多个 channel，直到其中一个 channel ready。
-   如果多个 channel 同时 ready，则随机选择一个执行。

```go
go test1(output1)
go test2(output2)
// 用select监控
select {
    case s1 := <- output1:
        fmt.Println("s1=", s1)
    case s2 := <- output2:
        fmt.Println("s2=", s2)
}
```

-   结合 default 可以用于判断管道是否存满（阻塞）

```go
// 判断管道有没有存满
func main() {
	output1 := make(chan string, 10)
	// 子协程写数据
	go write(output1)
	// 取数据
	for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}
func write(ch chan string) {
	for {
		select {
		// 写数据
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")   // 上面的 case 都阻塞，就会执行 default
		}
		time.Sleep(time.Millisecond * 500)
	}
}

```
