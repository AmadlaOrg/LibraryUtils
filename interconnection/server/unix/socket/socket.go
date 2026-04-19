//go:build !windows

package socket

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var (
	osRemove = os.Remove
	netListen = net.Listen
	osChmod  = os.Chmod
)

func Connect() error {
	// Remove old socket file if exists
	err := osRemove(socketPath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error removing old socket: %w", err)
	}

	// Create a Unix domain socket listener
	listener, err := netListen("unix", socketPath)
	if err != nil {
		return fmt.Errorf("error creating socket: %w", err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("Error closing socket:", err)
		}
	}(listener)

	// Set secure permissions: only owner can read/write
	err = osChmod(socketPath, 0600)
	if err != nil {
		return fmt.Errorf("error setting socket permissions: %w", err)
	}

	fmt.Println("doorman-aws listening on", socketPath)

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
