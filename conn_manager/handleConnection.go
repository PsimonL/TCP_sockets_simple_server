package conn_manager

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Handles connection for server and client as well.
// Responsible for single connection
func HandleConnection(conn net.Conn) {
	// Close connection after function ends
	defer conn.Close()

	// Perform authentication
	if !authenticate(conn) {
		return
	}

	// Start a loop to send messages to server/client - new thread
	// This function will be executed in parallel with the main function, handleConnection
	go func() {
		// Read input from the standard input (os.Stdin), waits for user input
		scanner1 := bufio.NewScanner(os.Stdin)
		// Until false, until data flows out of server/client
		for scanner1.Scan() {
			// Extract the text message
			message := scanner1.Text()
			if message == "q!" {
				conn.Close()
			}
			// Send message to client/server
			conn.Write([]byte(message + "\n"))
		}
		if err := scanner1.Err(); err != nil {
			panic(err.Error())
			return
		}
	}()

	// Read incoming messages from client/server
	scanner2 := bufio.NewScanner(conn) // Reuse the scanner variable
	// Loop until false, until data flows to server/client
	for scanner2.Scan() {
		message := scanner2.Text()
		fmt.Println("Received message from client:", message)
		fmt.Println("Enter message to send to client: ")
	}
	if err := scanner2.Err(); err != nil {
		panic(err.Error())
		return
	}
}

func authenticate(conn net.Conn) bool {
	conn.Write([]byte("Enter username: "))
	username, _ := bufio.NewReader(conn).ReadString('\n')
	conn.Write([]byte("Enter password: "))
	password, _ := bufio.NewReader(conn).ReadString('\n')

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	if username != ServConfObj.Credentials.Username || password != ServConfObj.Credentials.Password {
		conn.Write([]byte("Invalid username or password. Permission denied, connection closed.\n"))
		return false
	}

	conn.Write([]byte("Authentication successful. Connected successfully.\n"))
	return true
}
