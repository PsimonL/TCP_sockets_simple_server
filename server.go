package main

import (
	"fmt"
	"net"
)

type ServerConfig struct {
	Host string
	Port string
	Type string
}

var serverConfig = ServerConfig{
	Host: "localhost",
	Port: "8080",
	Type: "tcp",
}

func main() {
	fmt.Println("Running server...")
	// TCP listener that can accept incoming connections on port 8080:
	server, err := net.Listen(serverConfig.Type, serverConfig.Host+":"+serverConfig.Port) // net.Dial() <- remote
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer server.Close() // ensures that the network listener is properly closed when the function exits, even if an error occurred
	// Accept connections in a infinite loop
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Handle each connection in a new goroutine
		go handleConnection(conn)
	}
	fmt.Println("Server Stopped.")
}

func handleConnection(conn net.Conn) {

}
