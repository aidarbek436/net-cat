package main

import (
	"fmt"
	"net-cat/server"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Println("Listening on the port :8989")
		server.Server("8989")
	case 2:
		fmt.Println("Listening on the port :", os.Args[1])
		server.Server(os.Args[1])
	default:
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
}
