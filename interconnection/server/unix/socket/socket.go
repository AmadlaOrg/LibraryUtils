package socket

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Connect() {
	// Remove old socket file if exists
	err := os.Remove(socketPath)
	if err != nil {
		return
	}

	// Create a Unix domain socket listener
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println("Error creating socket:", err)
		os.Exit(1)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("Error closing socket:", err)
		}
	}(listener)

	// Set secure permissions: only owner can read/write
	err = os.Chmod(socketPath, 0600)
	if err != nil {
		return
	}

	fmt.Println("Clerk-AWS listening on", socketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing socket:", err)
		}
	}(conn)
	reader := bufio.NewReader(conn)

	// Read request from client
	request, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println("Received request:", request)

	// Example: Validate request (Can implement authentication)
	if request == "GET_JWT\n" {
		// Respond with a mock JWT token
		_, err := conn.Write([]byte("JWT_TOKEN_12345\n"))
		if err != nil {
			return
		}
	} else {
		_, err := conn.Write([]byte("INVALID_REQUEST\n"))
		if err != nil {
			return
		}
	}
}
