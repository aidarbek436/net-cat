package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

var (
	open_connections = make(map[string]Clients)
	entering         = make(chan Message)
	leaving          = make(chan Message)
	messages         = make(chan Message)
	history          []string
	mutex            sync.Mutex
)

type Message struct {
	message   string
	time      string
	user_Name string
}

type Clients struct {
	Name string
	Conn net.Conn
}

func Server(port string) { // num of arguments handling
	listener, err := net.Listen("tcp", ":"+port) // creating and running a tcp server at specified port
	if err != nil {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	defer listener.Close()
	go broadcaster() // go routine for management of receiving messages for clients and history from server
	for {
		conn, err := listener.Accept() // accepting of connections via nc command
		if err != nil {
			log.Print(err)
			continue // not breaking the loop to keep accepting connections
		}
		go handleConn(conn) // go routine for sending required messages to the server from clients
	}
}
