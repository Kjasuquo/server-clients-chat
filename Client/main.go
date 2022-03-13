package main

import (
	"bufio"
	"fmt"
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

	reader := bufio.NewReader(os.Stdin)
	message, err := reader.ReadString('\n')

	Msg := fmt.Sprintf(message)
	connection.Write([]byte(Msg))

}
