package main

import (
	connector "awesomeProject1/conn_manager"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting server...")
	// TCP listener on port 8080
	server, err := net.Listen(connector.ServConfObj.Type, connector.ServConfObj.Host+":"+connector.ServConfObj.Port) // net.Dial() <- remote
	// If error exists exit main
	if err != nil {
		panic(err.Error())
		return
	}
	// Ensures that listener is properly closed - always after main execution (even if an error occurred)
	// Close connection after function ends
	defer server.Close()
	fmt.Println("Server listening on: " + connector.ServConfObj.Port)

	// Accept connections in loop (separate goroutine for certain connection)
	for {
		// net.Conn object - connection for request from client
		conn, err := server.Accept()
		if err != nil {
			panic(err.Error())
			return
		}
		fmt.Println("Client connected:", conn.RemoteAddr())
		// Ensures that we can handle couple connections at one time - separate thread for connection
		go connector.HandleConnection(conn)
	}
}
