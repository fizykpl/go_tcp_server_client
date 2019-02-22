package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST      = "127.0.0.1"
	CONN_PORT_2011 = "2011"
	CONN_PORT_2012 = "2012"
	CONN_TYPE      = "tcp"
)

func main() {
	fmt.Println("Starting listening ... ")
	go listeningOnPort(CONN_PORT_2011)
	listeningOnPort(CONN_PORT_2012)

}

func listeningOnPort(port string) {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + port)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn, port)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn, port string) {

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Print("Message Received on port="+port+":", string(message))

	// Send a response back to person contacting us.
	conn.Write([]byte(message))
	// Close the connection when you're done with it.
	conn.Close()
}
