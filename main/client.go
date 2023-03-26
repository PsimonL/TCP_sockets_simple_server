package main

import (
	connector "awesomeProject1/conn_stuff"
	"fmt"
	"io/ioutil"
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

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}
