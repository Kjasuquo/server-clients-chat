package main

import (
	"bufio"
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

	for {
		select {
		// if there is something in the channel, invoke broadcast message (broadcast to other connections)
		case conn := <-NewConnection:
			go broadcastMessage(conn)

		//Delete the connection if nothing
		case conn := <-DeadConnection:
			for item := range OpenConnection {
				if item == conn {
					break
				}
			}
			delete(OpenConnection, conn)
		}

	}

}

func broadcastMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		//loop through all the open connections and send messages to these connections
		//except the connections that sent the message
		for item := range OpenConnection {
			if item != conn {
				item.Write([]byte(message))
			}
		}
	}

	DeadConnection <- conn
}
