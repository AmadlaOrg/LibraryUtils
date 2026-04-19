//go:build !windows

package socket

import (
	"errors"
	"io"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSocketImpl_Connect_Success(t *testing.T) {
	originalDial := netDial
	defer func() { netDial = originalDial }()

	server, client := net.Pipe()

	netDial = func(network, address string) (net.Conn, error) {
		return client, nil
	}

	// Simulate server responding
	go func() {
		buf := make([]byte, 1024)
		n, _ := server.Read(buf)
		assert.Equal(t, "GET_JWT\n", string(buf[:n]))
		_, _ = server.Write([]byte("JWT_TOKEN_12345\n"))
		server.Close()
	}()

	svc := &socketImpl{}
	err := svc.Connect()
	assert.NoError(t, err)
}

func TestSocketImpl_Connect_DialError(t *testing.T) {
	originalDial := netDial
	defer func() { netDial = originalDial }()

	netDial = func(network, address string) (net.Conn, error) {
		return nil, errors.New("connection refused")
	}

	svc := &socketImpl{}
	err := svc.Connect()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error connecting to doorman-aws")
}

func TestSocketImpl_Connect_WriteError(t *testing.T) {
	originalDial := netDial
	defer func() { netDial = originalDial }()

	server, client := net.Pipe()

	netDial = func(network, address string) (net.Conn, error) {
		return client, nil
	}

	// Close server immediately to cause write error
	server.Close()

	svc := &socketImpl{}
	err := svc.Connect()
	assert.Error(t, err)
}

func TestSocketImpl_Connect_ReadError(t *testing.T) {
	originalDial := netDial
	defer func() { netDial = originalDial }()

	server, client := net.Pipe()

	netDial = func(network, address string) (net.Conn, error) {
		return client, nil
	}

	// Read the request then close without sending a newline-terminated response
	go func() {
		buf := make([]byte, 1024)
		_, _ = server.Read(buf)
		// Close without sending newline — causes ReadString('\n') to get EOF
		server.Close()
	}()

	svc := &socketImpl{}
	err := svc.Connect()
	assert.ErrorIs(t, err, io.EOF)
}
