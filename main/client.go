package main

import (
	connector "awesomeProject1/conn_manager"
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"os"
)

func main() {
	// TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // MUST HAVE!!! -> if using self-signed certificate
	}

	// Connect to the server
	conn, err := tls.Dial(connector.ServConfObj.Type, connector.ServConfObj.Host+":"+connector.ServConfObj.Port, tlsConfig)
	if err != nil {
		panic(err.Error())
		return
	}
	// Close connection after function ends
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	fmt.Println("Client connected on port:", conn.RemoteAddr())

	// Authenticate user
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Enter password: ")
	scanner.Scan()
	password := scanner.Text()

	if username == connector.ServConfObj.Credentials.Username && password == connector.ServConfObj.Credentials.Password {
		fmt.Println("Authentication successful. Permission granted.")
	} else {
		fmt.Println("Authentication failed. Permission denied.")
		return
	}

	go connector.HandleConnection(conn)

	// Program will wait until any of cases can proceed - cases do not exist so program will not proceed, won't finish up (making kind of "infinite loop" with that operation as I understand)
	select {}
}
