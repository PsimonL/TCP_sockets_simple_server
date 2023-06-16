package conn_manager

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// Handles connection for server and client as well.
// Responsible for single connection
func HandleConnection(conn net.Conn) {
	// Close connection after function ends
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// Start a loop to send messages to server/client - new thread
	// This function will be executed in parallel with the main function, handleConnection
	go func() {
		// Read input from the standard input (os.Stdin), waits for user input
		scanner := bufio.NewScanner(os.Stdin)
		// Until false, until data flows out of server/client
		for scanner.Scan() {
			// Extract the text message
			message := scanner.Text()
			if message == "q!" {
				err := conn.Close()
				if err != nil {
					return
				}
			}
			// Send message to client/server
			conn.Write([]byte(message + "\n"))
		}
		if err := scanner.Err(); err != nil {
			panic(err.Error())
			return
		}
	}()
	// Read incoming messages from client/server
	scanner := bufio.NewScanner(conn)
	// Loop until false, until data flows to server/client
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Received message from client:", message)
		fmt.Println("Enter message to send to client: ")
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
		return
	}
}
