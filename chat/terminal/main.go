package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// We create a listener, with the protocol type TCP and address ":8081". In this case, it is just a localhost.
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error creating listener:", err)
		os.Exit(1)
	}

	// Defer close and checks errors.
	defer func(listener net.Listener) {
		err = listener.Close()
		if err != nil {
			fmt.Println("Error closing listener:", err)
		}
	}(listener)

	fmt.Println("Terminal started! Waiting for connections...")

	for {
		// We then set the listener to wait for the server to connect with the "Accept()" method.
		// The listener blocks until the server connects.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}
		// When a client connects, they will then be handled by "handleServerConnection".
		go handleServerConnection(conn)
	}
}

type User struct {
	Name string
}

func handleServerConnection(server net.Conn) {
	// Defer close() and checks for errors.
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(server)

	// We get the name of the connected client, with the first initial message sent from the server, containing the clients name.
	forwardedClientName, err := bufio.NewReader(server).ReadString('.')
	if err != nil {
		fmt.Println("Error reading client name:", err)
	}

	// Since it includes the delimiter, we trim it off.
	forwardedClientName = strings.Trim(forwardedClientName, ".")

	// This reader will receive the input of clients that went through the server, to here.
	reader := bufio.NewReader(server)

	for {
		// Read data from the connection
		message, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Client disconnected:", server.RemoteAddr().String())
			} else {
				fmt.Println("Error reading data:", err.Error())
			}
			return
		}

		// Print the received message
		fmt.Printf("%s: %s", forwardedClientName, message)
	}
}
