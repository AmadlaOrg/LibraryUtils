//go:build !windows
// +build !windows

package unix

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

// UnixTPM implements the TPM interface for Linux/macOS.
type UnixTPM struct{}

// NewUnixTPM creates a new UnixTPM instance.
func NewUnixTPM() *UnixTPM {
	return &UnixTPM{}
}

// ListDevices lists TPM devices without using `ls`.
func (t *UnixTPM) ListDevices() (string, error) {
	files, err := ioutil.ReadDir("/sys/class/tpm/")
	if err != nil {
		return "", fmt.Errorf("failed to list TPM devices: %v", err)
	}

	var devices []string
	for _, file := range files {
		devices = append(devices, file.Name())
	}
	return strings.Join(devices, "\n"), nil
}

// CheckLogs fetches TPM logs using dmesg.
func (t *UnixTPM) CheckLogs() (string, error) {
	cmd := exec.Command("sh", "-c", "dmesg | grep -i tpm2")

	var out strings.Builder
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error checking TPM logs: %v", err)
	}

	return strings.TrimSpace(out.String()), nil
}
