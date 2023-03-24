package main

import (
	connector "awesomeProject1/connection_stuff"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Client connected...")
	// Connect to the server
	conn, err := net.Dial(connector.ServConfObj.Type, connector.ServConfObj.Host+":"+connector.ServConfObj.Port)
	if err != nil {
		panic(err.Error())
		return
	}
	defer conn.Close()
}
