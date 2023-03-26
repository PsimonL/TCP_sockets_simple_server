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
	fmt.Println("Client connected on port " + connector.ServConfObj.Port + ":\n")
	defer conn.Close()
	go
	// start a loop to read input from the user and send messages to the server
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Fprintln(conn, message)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

