package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func ClientConnect(ip_prot string) {
	conn, err := net.Dial("tcp", ip_prot)
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	// 客户端发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入[终端]
	// 从终端读取一行用户输入，并准备发送给服务器
	line, err := reader.ReadString('\n')
	line = strings.Trim(line, "\r\n")
	if err != nil {
		fmt.Println("readString err=", err)
	}
	// 将 line 发送给 服务器
	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("conn.Write err=", err)
	}
	fmt.Printf("客户端发送了 %d 字节的数据，并退出\n", n)
}

func main() {
	var ip_prot string = "127.0.0.1:17888"
	ClientConnect(ip_prot)
}
