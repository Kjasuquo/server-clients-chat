package main

import (
	"bufio"
	"log"
	"net"
	"os"
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

	//read message from the stdin (terminal)
	reader := bufio.NewReader(os.Stdin)
	//read every line as a string
	message, err := reader.ReadString('\n')
	logFetal(err)

	//Msg := fmt.Sprintf(message)

	//Dial the written message for a server in the same port (:8080) to listen
	connection.Write([]byte(message))

}
