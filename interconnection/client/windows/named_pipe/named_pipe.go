// Package named_pipe
//
// On environments that are not windows this part of the code will show errors and issues.
//
// Doc: https://learn.microsoft.com/en-us/windows/win32/ipc/named-pipes
//
// TODO: Find a way to run all the test but bypass the Windows section
package named_pipe

import (
	"fmt"
	"github.com/Microsoft/go-winio"
	"os"
)

type INamedPipe interface{}

type SNamedPipe struct{}

// Connect
func (s *SNamedPipe) Connect() error {
	conn, err := winio.DialPipe(pipeName, nil)
	if err != nil {
		fmt.Println("Error connecting to Clerk-AWS:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Send a request
	// TODO: What about other types of secrets...
	conn.Write([]byte("GET_JWT"))

	// Read response
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println("Received from Clerk-AWS:", string(buf[:n]))

	return nil
}
