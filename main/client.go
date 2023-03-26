package main

import (
	connector "awesomeProject1/conn_stuff"
	"bufio"
	"fmt"
	//"io/ioutil"
	"net"
	"os"
)

func main() {
	// Connect to the server
	conn, err := net.Dial(connector.ServConfObj.Type, connector.ServConfObj.Host+":"+connector.ServConfObj.Port)
	if err != nil {
		panic(err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("Client connected on port:", conn.RemoteAddr())

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
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
	}()
	// start a loop to read input from the user and send messages to the server
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Received message from server:", message)
		fmt.Print("Enter message to send to client: ")
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
		return
	}
}
