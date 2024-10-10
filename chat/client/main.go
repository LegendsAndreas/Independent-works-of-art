package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// We connect to the server with the protocol TCP and the address "localhost:8080". Since this is hosted as a
	// localhost, we can just write ":8080"
	clientConn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Error connecting to the server:", err.Error())
		os.Exit(1)
	}

	// Defer close and checks for errors.
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(clientConn)

	// Sends the name to the server. Since the code on the server side is looking for the delimiter newline,
	// remember to include that at the end.
	_, err = clientConn.Write([]byte("Andreas."))
	if err != nil {
		fmt.Println("Error sending name:", err)
		os.Exit(1)
	}

	fmt.Println("Connected to the server. Type your messages below.")

	// We create a reader, that reads the users input.
	reader := bufio.NewReader(os.Stdin)

	for {
		// Asks for input and stores it in "message", with the delimiter being a newline. This allows us to include
		// spaces, since it only stops reading when it encounters a newline.
		fmt.Print("Enter message: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// We then send the message as a slice of bytes, to the server that we have connect to.
		_, err = clientConn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			continue
		}
	}
}
