package server

import "fmt"

func broadcaster() {
	for {
		select {
		case msg := <-messages:
			mutex.Lock()
			for _, user := range open_connections {
				if user.Name != msg.user_Name {
					fmt.Fprintf(user.Conn, "\n[%s][%s]:%s\n", msg.time, msg.user_Name, msg.message) // receiving the messages from others
				} else {
					hst := fmt.Sprintf("\n[%s][%s]:%s", msg.time, msg.user_Name, msg.message) // saving the messages into history from clients
					history = append(history, hst)

				}
				fmt.Fprintf(user.Conn, "[%s][%s]:", msg.time, user.Name) // printing date and name after receiving the messages

			}
			mutex.Unlock()
		case cli := <-entering:
			mutex.Lock()
			for _, user := range open_connections {
				if user.Name != cli.user_Name {
					fmt.Fprintf(user.Conn, "\n%s %s", cli.user_Name, cli.message) // receiving the message of entering of someone
					fmt.Fprintf(user.Conn, "[%s][%s]:", cli.time, user.Name)      // printing date and name after joining message

				} else {
					for _, prev := range history {
						fmt.Fprintf(user.Conn, prev) // printing the history if necessary
					}
					if history != nil {
						fmt.Fprintf(user.Conn, "\n[%s][%s]:", cli.time, user.Name) // case for the first user (nil history)
					}

				}
			}
			mutex.Unlock()
		case cli := <-leaving:
			mutex.Lock()
			for _, user := range open_connections {
				if user.Name != cli.user_Name {
					fmt.Fprintf(user.Conn, "\n%s %s", cli.user_Name, cli.message) // receiving the message of someone leaving
					fmt.Fprintf(user.Conn, "[%s][%s]:", cli.time, user.Name)      // pringting date and name after leaving message
				}
			}
			mutex.Unlock()
		}
	}
}
