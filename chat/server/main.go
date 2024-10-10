package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// We create a listener, with the protocol type TCP and address ":8080". In this case, it is just a localhost.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error creating listener:", err.Error())
		os.Exit(1)
	}

	// Defer close and checks errors.
	defer func(listener net.Listener) {
		err = listener.Close()
		if err != nil {
			fmt.Println("Error closing listener:", err)
		}
	}(listener)

	fmt.Println("Server started! Waiting for connections...")

	for {
		// We then set the listener to wait for a client to connect with the "Accept()" method.
		// The listener blocks until a client connects.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// When a client connects, they will then be handled by "handleConnection".
		go handleConnection(conn)
	}
}

func handleConnection(client net.Conn) {
	// Defer close() and checks for errors.
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection:", err)
		}
	}(client)

	// We print the clients network address to know where that little shit is.
	fmt.Println("Client connected:", client.RemoteAddr().String())

	// We get the name of the connected client, with the first initial message sent, containing the clients name.
	clientName, err := bufio.NewReader(client).ReadString('.')
	if err != nil {
		fmt.Println("Error reading client name:", err)
	}

	// We connect to the terminal, that can be viewed by all clients.
	terminalConn, err := net.Dial("tcp", ":8081")
	if err != nil {
		fmt.Println("Error connecting to terminal:", err)
	}

	// We send the clients name to the terminal.
	_, err = terminalConn.Write([]byte(clientName))
	if err != nil {
		fmt.Println("Error sending name:", err)
		os.Exit(1)
	}

	// We create a reader, that reads from the connected client
	reader := bufio.NewReader(client)

	defer func(terminalConn net.Conn) {
		err = terminalConn.Close()
		if err != nil {
			fmt.Println("Error closing terminal connection:", err)
		}
	}(terminalConn)

	for {
		// Read data from the connection
		message, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Client disconnected:", client.RemoteAddr().String())
			} else {
				fmt.Println("Error reading data:", err.Error())
			}
			return
		}
		_, err = terminalConn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			continue
		}

		// Print the received message
		fmt.Printf("Client name: %s | Message: %s", clientName, message)
	}
}
