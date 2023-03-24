package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Client connected...")
	// Connect to the server
	conn, err := net.Dial(serverConfig.Type, serverConfig.Host+":"+serverConfig.Port)
	if err != nil {
		panic(err.Error())
		return
	}
	defer conn.Close()
}
