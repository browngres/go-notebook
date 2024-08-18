package main

import (
	"socket1/server"
)

func main() {
	var ip_prot string = "127.0.0.1:17888"
	server.ServerListen(ip_prot)
}
