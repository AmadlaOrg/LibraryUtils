//go:build !windows

package socket

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
)

type Socket interface{}
type socketImpl struct{}

var (
	netDial        = net.Dial
	bufioNewReader = bufio.NewReader
)

func (s *socketImpl) Connect() error {
	conn, err := netDial("unix", filepath.Join(os.TempDir(), SockFileName))
	if err != nil {
		return fmt.Errorf("error connecting to doorman-aws: %w", err)
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

	fmt.Println("Received from doorman-aws:", response)

	return nil
}
