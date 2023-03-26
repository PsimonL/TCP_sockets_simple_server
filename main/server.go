package main

import (
	connector "awesomeProject1/conn_stuff"
	"bufio"
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

	// Accept connections in loop
	for {
		// net.Conn object - connection for request from client
		conn, err := server.Accept()
		if err != nil {
			panic(err.Error())
			return
		}
		fmt.Println("Client connected:", conn.RemoteAddr())
		// Ensures that we can handle couple connections at one time (go concurrency - small number of operating system threads)
		go handleConnection(conn)
	}
}

// Responsible for single connection
func handleConnection(conn net.Conn) {
	defer conn.Close()
	// start a loop to send messages to client
	go send_message(conn)
	// read incoming
	//go receive_message(conn)
	scanner := bufio.NewScanner(conn)
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

func send_message(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		if message == "q!" {
			conn.Close()
		}
		conn.Write([]byte(message + "\n"))
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
		return
	}
}
