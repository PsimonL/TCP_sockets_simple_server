package main

import (
	"fmt"
	"net"
)

var serverConfig = ServerConfig{
	Host: "localhost",
	Port: "8080",
	Type: "tcp",
}

func main() {
	fmt.Println("main")
	// TCP listener that can accept incoming connections on port 8080:
	ln, err := net.Listen(serverConfig.Type, serverConfig.Host+":"+serverConfig.Port) // net.Dial() <- remote
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
}
