package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func logFetal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

var (
	OpenConnection = make(map[net.Conn]bool)
	NewConnection  = make(chan net.Conn)
	DeadConnection = make(chan net.Conn)
)

func main() {
	//To listen for a connection
	ln, err := net.Listen("tcp", ":8080")
	logFetal(err)

	//Ensure to close the connection at the end
	defer ln.Close()

	//using Goroutine
	go func() {

		//create an infinite loop to always accept a connection that has been discovered while listening
		for {
			conn, err := ln.Accept()
			logFetal(err)

			//If there is connection, and it accepts it, make it true
			OpenConnection[conn] = true

			//pass the connection through this channel
			NewConnection <- conn
		}
	}()
	connection := <-NewConnection
	reader := bufio.NewReader(connection)
	message, err := reader.ReadString('\n')
	logFetal(err)
	fmt.Println(message)

}
