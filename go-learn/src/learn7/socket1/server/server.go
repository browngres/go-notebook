package server

import (
	"fmt"
	_ "io"
	"net" //做网络 socket 开发时,net 包含有我们需要所有的方法和函数
)

func process(conn net.Conn) {
	//循环的接收客户端发送的数据
	defer conn.Close() //关闭 conn
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//1. 等待客户端通过 conn 发送信息
		//2. 如果客户端没有 wrtie[发送]，那么协程就阻塞在这里
		fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从 conn 读取
		if err != nil {
			fmt.Printf("客户端报错，退出。 err=%v", err)
			return
		}
		//3. 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func ServerListen(ip_port string) {
	fmt.Println("服务器开始监听....")
	listen, err := net.Listen("tcp", ip_port)
	if err != nil {
		fmt.Println("监听失败 err=", err)
		return
	}
	defer listen.Close() //延时关闭 listen
	// 循环等待客户端连接
	for {
		fmt.Println("监听成功，等待客户端....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept()失败 err=", err)
		} else {
			fmt.Printf("Accept()成功 con=%v 客户端 ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//准备一个协程，为客户端服务
		go process(conn)
	}
}
