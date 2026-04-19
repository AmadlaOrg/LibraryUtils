//go:build windows

// Package named_pipe
//
// Doc: https://learn.microsoft.com/en-us/windows/win32/ipc/named-pipes
package named_pipe

import (
	"fmt"
	"github.com/Microsoft/go-winio"
	"os"
)

type NamedPipe interface{}

type namedPipeImpl struct{}

// Connect
func (s *namedPipeImpl) Connect() error {
	conn, err := winio.DialPipe(pipeName, nil)
	if err != nil {
		fmt.Println("Error connecting to doorman-aws:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Send a request
	// TODO: What about other types of secrets...
	conn.Write([]byte("GET_JWT"))

	// Read response
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println("Received from doorman-aws:", string(buf[:n]))

	return nil
}
