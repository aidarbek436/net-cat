package server

import (
	"fmt"
	"net"
)

func is_Valid_Name(user_name string, connection net.Conn) bool {
	if user_name == "" || len(user_name) == 0 { // empty name handling
		fmt.Fprint(connection, "The usernime must not be empty to enter the chat\n")
		connection.Close()
		return false
	}

	for _, symbol := range user_name { // ascii table handling
		if symbol < 32 || symbol > 127 {
			fmt.Fprintln(connection, "Incorrect input by ascii table\n")
			connection.Close()
			return false
		}
	}
	for names, _ := range open_connections { // taken names handling
		if user_name == names {
			fmt.Fprintln(connection, "Username is already taken, try to connect to server again and use another name\n")
			connection.Close()
			return false
		}
	}
	return true
}
