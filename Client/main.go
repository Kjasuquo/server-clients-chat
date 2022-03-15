package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func logFetal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	//dial into a network
	connection, err := net.Dial("tcp", ":8080")
	logFetal(err)

	//close at the end
	defer connection.Close()

	fmt.Println("Enter Username")

	//read message/username from the stdin (terminal)
	reader := bufio.NewReader(os.Stdin)
	//convert to a string
	username, err := reader.ReadString('\n')
	logFetal(err)

	username = strings.Trim(username, "\n")
	welcomeMessage := fmt.Sprintf("Welcome %s", username)
	fmt.Println(welcomeMessage)

	go read(connection)

	write(connection, username)

}

func read(connection net.Conn) {
	for {
		reader := bufio.NewReader(connection)
		message, err := reader.ReadString('\n')
		if err == io.EOF {
			connection.Close()
			fmt.Println("Connection closed")
			os.Exit(0)
		}
		fmt.Println(message)
		//fmt.Println("-----------------------------------")
	}
}

func write(connection net.Conn, username string) {
	for {
		//read message from the stdin (terminal)
		reader := bufio.NewReader(os.Stdin)
		//convert to a string
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		//to print out the username and the message sent, use sprintf to assign it to a variable
		mes := fmt.Sprintf("%s: %s\n", username, strings.Trim(message, "\n"))

		//Dial the written message for a server in the same port (:8080) to listen
		connection.Write([]byte(mes))

	}
}
