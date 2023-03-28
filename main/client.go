package main

import (
	connector "awesomeProject1/conn_manager"
	"fmt"
	//"io/ioutil"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial(connector.ServConfObj.Type, connector.ServConfObj.Host+":"+connector.ServConfObj.Port)
	if err != nil {
		panic(err.Error())
		return
	}
	// Close connection after function ends
	defer conn.Close()
	fmt.Println("Client connected on port:", conn.RemoteAddr())

	go connector.HandleConnection(conn)

	// Program will wait until any of cases can proceed - cases do not exist so program will not proceed, won't finish up (making kind of "infinite loop" with that operation as I understand)
	select {}
}
