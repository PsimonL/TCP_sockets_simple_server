package main

import (
	connector "awesomeProject1/conn_stuff"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Running server...")
	// TCP listener on port 8080
	server, err := net.Listen(connector.ServConfObj.Type, connector.ServConfObj.Host+":"+connector.ServConfObj.Port) // net.Dial() <- remote
	// If error exists exit main
	if err != nil {
		panic(err.Error())
		return
	}
	// Ensures that listener is properly closed - always after main execution (even if an error occurred)
	defer server.Close()
	fmt.Printf("Server listening on port %s", connector.ServConfObj.Port)

	// Accept connections in loop
	for {
		// net.Conn object - connection for request from client
		conn, err := server.Accept()
		if err != nil {
			panic(err.Error())
			continue
		}
		// Ensures that we can handle couple connections at one time (go concurrency - small number of operating system threads)
		go handleConnection(conn)
	}
	fmt.Println("Server Stopped.")
}

// Responsible for single connection
func handleConnection(conn net.Conn) {
	// Make a buffer to hold incoming data - used to store the incoming data until the server has received the complete request
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer - ignore number of bytes read
	_, err := conn.Read(buf)
	if err != nil {
		panic(err.Error())
		return
	}
	// Send response to client - ignore number of bytes written
	response := "Hello, client!"
	_, err = conn.Write([]byte(response))
	if err != nil {
		panic(err.Error())
		return
	}
}
