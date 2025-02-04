package socket

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
)

type ISocket interface{}
type SSocket struct{}

var (
	netDial        = net.Dial
	bufioNewReader = bufio.NewReader
)

func (s *SSocket) Connect() error {
	conn, err := netDial("unix", filepath.Join("/tmp", SockFileName))
	if err != nil {
		fmt.Println("Error connecting to Clerk-AWS:", err)
		os.Exit(1)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection:", err)
		}
	}(conn)

	// Send a request
	_, err = fmt.Fprintln(conn, "GET_JWT")
	if err != nil {
		return err
	}

	// Read response
	response, err := bufioNewReader(conn).ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Println("Received from Clerk-AWS:", response)

	return nil
}
