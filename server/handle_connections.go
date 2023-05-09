package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var penguin = []string{
	"Welcome to TCP-Chat!",
	"	 _nnnn_ ",
	"        dGGGGMMb",
	"       @p~qp~~qMb",
	"       M|@||@) M|",
	"       @,----.JM|",
	"      JS^\\__/  qKL",
	"     dZP        qKRb",
	"    dZP          qKKb",
	"   fZP            SMMb",
	"   HZM            MMMM",
	"   FqM            MMMM",
	" __| \".        |\\dS\"qML",
	" |    `.       | `' \\Zq",
	"_)      \\.___.,|     .'",
	"\\____   )MMMMMP|   .'",
	"     `-'       `--'",
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for _, v := range penguin { // printing welcome and penguin
		fmt.Fprintf(conn, v)
		fmt.Fprintf(conn, "\n")
	}

	conn.Write([]byte("[ENTER YOUR NAME]:"))           // entering the name
	who, err := bufio.NewReader(conn).ReadString('\n') // reading the name from input
	if err != nil {                                    // error handling
		conn.Write([]byte("There is a problem with your Name <3"))
		return
	}
	who = strings.Trim(who, "\r\n") // error handling

	if !is_Valid_Name(who, conn) { // name error handling
		return
	}

	User := Clients{who, conn} // adding joined client to the list of clients (struct)
	mutex.Lock()
	open_connections[who] = User // adding client to open_connections map for further checking of the number of connected clients
	mutex.Unlock()
	mutex.Lock()
	if len(open_connections) > 10 { // check for number of clients in one chat, if more than 10 => return
		_, err := fmt.Fprintf(User.Conn, "Chat is full: no more than 10 users are allowed to be in the chat\n")
		if err != nil {
			log.Println(err)
		}
		delete(open_connections, who)
		mutex.Unlock()
		return
	}
	mutex.Unlock()

	join_msg := Message{
		message:   "has joined our chat...\n",
		time:      time.Now().Format("2006-01-02 15:04:05"),
		user_Name: who,
	}
	entering <- join_msg // sending message of joining the chat

	fmt.Fprintf(conn, "[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), who) // formating the typing input
	input := bufio.NewScanner(conn)
	for input.Scan() {
		text := strings.Trim(input.Text(), " ") // error handling
		if !is_Valid_Text(text) {               // text error handling
			fmt.Fprintln(conn, "Do not type empty messages")
			fmt.Fprintf(conn, "[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), who)
			continue
		}
		text_msg := Message{
			message:   text,
			time:      time.Now().Format("2006-01-02 15:04:05"),
			user_Name: who,
		}
		messages <- text_msg // sending typed texts from users
	}

	// if user wants to leave the chat and types ctrl + c => infinite loop is ended and the code below is executed for leaving...

	leaving_msg := Message{
		message:   "has left our chat...\n",
		time:      time.Now().Format("2006-01-02 15:04:05"),
		user_Name: who,
	}
	leaving <- leaving_msg // sending message for leaving the chat
	mutex.Lock()
	delete(open_connections, who) // deleting the client from map
	conn.Close()                  // closing the connection
	mutex.Unlock()
}
