package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type ServerConfig struct {
	Host string
	Port string
	Type string
}

var serverConfig ServerConfig

func init() {
	file, err := os.Open("connection_params.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var myList []string
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ";")
		for _, word := range words {
			myList = append(myList, word)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	serverConfig.Host = myList[0]
	serverConfig.Port = myList[1]
	serverConfig.Type = myList[2]
}

func main() {
	fmt.Println("Running server...")
	// TCP listener on port 8080
	server, err := net.Listen(serverConfig.Type, serverConfig.Host+":"+serverConfig.Port) // net.Dial() <- remote
	fmt.Println("FLAG1")
	// If error exists exit main
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	// Ensures that listener is properly closed - always after main execution (even if an error occurred)
	defer server.Close()
	fmt.Println("Server listening on port 8080")

	// Accept connections in loop
	for {
		// net.Conn object - connection for request from client
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Ensures that we can handle couple connections at one time (go concurrency - small number of operating system threads)
		go handleConnection(conn)
	}
	fmt.Println("Server Stopped.")
}

// responsible for single connection
func handleConnection(conn net.Conn) {
	// Make a buffer to hold incoming data - used to store the incoming data until the server has received the complete request
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer - ignore number of bytes read
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	// Send response to client - ignore number of bytes written
	response := "Hello, client!"
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
		return
	}
}
