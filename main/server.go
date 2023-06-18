package main

import (
	connector "awesomeProject1/conn_manager"
	"crypto/tls"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting server...")

	// Load certificate and private key
	cert, err := tls.LoadX509KeyPair("path/to/certificate.pem", "path/to/private_key.pem")
	if err != nil {
		panic(err.Error())
	}

	// TLS config
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

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

	// Create TLS listener using TCP listener and TLS configuration
	tlsListener := tls.NewListener(server, tlsConfig)

	// Accept connections in loop (separate goroutine for certain connection)
	for {
		// net.Conn object - connection for request from client
		conn, err := tlsListener.Accept()
		if err != nil {
			panic(err.Error())
			return
		}
		fmt.Println("Client connected:", conn.RemoteAddr())
		// Ensures that we can handle couple connections at one time - separate thread for connection
		go connector.HandleConnection(conn)
	}
}
