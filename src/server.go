package main


import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	fmt.Println("Connected server.")

	// run loop forever (or until ctrl-c)
	for {
		fmt.Println(conn.RemoteAddr().Network() ,conn.RemoteAddr().String())
		// will listen for message to process ending in newline (\n)
		messages, _ := bufio.NewReader(conn).ReadString('\n')

		// output mmessagesessage received
		splited := strings.Split(messages,"^")

		for _ , message := range splited {
			if len(message) == 0 || message =="\n"{
				continue
			}
			fmt.Println("Message Received:", string(message))
			// sample process for string received
			newmessage := strings.ToUpper(message)
			// send new string back to client
			conn.Write([]byte("message "))
			conn.Write([]byte("= " +newmessage+" "))
		}
		conn.Write([]byte("\n"))

	}
}

