//go:build !windows

package socket

import (
	"bufio"
	"errors"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleClient_GetJWT(t *testing.T) {
	server, client := net.Pipe()
	defer client.Close()

	go handleClient(server)

	// Send GET_JWT request
	_, err := client.Write([]byte("GET_JWT\n"))
	assert.NoError(t, err)

	// Read response
	reader := bufio.NewReader(client)
	response, err := reader.ReadString('\n')
	assert.NoError(t, err)
	assert.Equal(t, "JWT_TOKEN_12345\n", response)
}

func TestHandleClient_InvalidRequest(t *testing.T) {
	server, client := net.Pipe()
	defer client.Close()

	go handleClient(server)

	// Send an invalid request
	_, err := client.Write([]byte("SOMETHING_ELSE\n"))
	assert.NoError(t, err)

	// Read response
	reader := bufio.NewReader(client)
	response, err := reader.ReadString('\n')
	assert.NoError(t, err)
	assert.Equal(t, "INVALID_REQUEST\n", response)
}

func TestConnect_ListenError(t *testing.T) {
	originalListen := netListen
	defer func() { netListen = originalListen }()

	// Ensure Remove succeeds (file not existing is fine)
	originalRemove := osRemove
	defer func() { osRemove = originalRemove }()
	osRemove = func(name string) error { return os.ErrNotExist }

	netListen = func(network, address string) (net.Listener, error) {
		return nil, errors.New("listen failed")
	}

	err := Connect()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error creating socket")
}

func TestConnect_RemoveError(t *testing.T) {
	originalRemove := osRemove
	defer func() { osRemove = originalRemove }()

	osRemove = func(name string) error {
		return errors.New("permission denied")
	}

	err := Connect()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error removing old socket")
}

func TestConnect_ChmodError(t *testing.T) {
	originalRemove := osRemove
	originalListen := netListen
	originalChmod := osChmod
	defer func() {
		osRemove = originalRemove
		netListen = originalListen
		osChmod = originalChmod
	}()

	osRemove = func(name string) error { return os.ErrNotExist }

	// Create a real listener on a temp path to satisfy the defer Close
	tmpDir := t.TempDir()
	tmpSock := tmpDir + "/test.sock"
	realListener, err := net.Listen("unix", tmpSock)
	assert.NoError(t, err)

	netListen = func(network, address string) (net.Listener, error) {
		return realListener, nil
	}

	osChmod = func(name string, mode os.FileMode) error {
		return errors.New("chmod failed")
	}

	err = Connect()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error setting socket permissions")
}
