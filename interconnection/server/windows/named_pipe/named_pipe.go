package named_pipe

import (
	"fmt"
	"github.com/Microsoft/go-winio"
	"net"
	"os"
)

func Connect() {
	listener, err := winio.ListenPipe(pipeName, nil)
	if err != nil {
		fmt.Println("Error creating pipe:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Clerk-AWS listening on", pipeName)

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
			fmt.Println("Error closing connection:", err)
		}
	}(conn)

	// Read request
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	request := string(buf[:n])
	fmt.Println("Received request:", request)

	// Validate request
	if request == "GET_JWT" {
		// Send mock JWT
		_, err := conn.Write([]byte("JWT_TOKEN_12345"))
		if err != nil {
			return
		}
	} else {
		_, err := conn.Write([]byte("INVALID_REQUEST"))
		if err != nil {
			return
		}
	}
}
