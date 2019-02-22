package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("TCP port is required in args.")
		os.Exit(1)
	}

	port := os.Args[1]
	fmt.Println("Client try connect on " + port + " port.")

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Text to send: ")
		text, _ := reader.ReadString('\n')
		splited := strings.Split(text[:len(text)-1], "^")

		// send to socket
		sendMessage(port, splited)
	}
}

func sendMessage(port string, message []string) {
	for _, element := range message {
		conn := getConnection(port)
		//if len(element)==0{
		//	continue
		//}
		fmt.Fprintf(conn, element+"\n")
		// listen for reply
		reply, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + reply)
		conn.Close()
	}

}

func getConnection(port string) net.Conn {
	host := "127.0.0.1:" + port
	// connect to this socket
	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Client connected to " + host)
	return conn
}
