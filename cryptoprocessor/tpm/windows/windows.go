//go:build windows
// +build windows

package windows

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// WindowsTPM implements the TPM interface for Windows.
type WindowsTPM struct{}

// NewWindowsTPM creates a new WindowsTPM instance.
func NewWindowsTPM() *WindowsTPM {
	return &WindowsTPM{}
}

// ListDevices checks for TPM presence using PowerShell.
func (t *WindowsTPM) ListDevices() (string, error) {
	cmd := exec.Command(
		"powershell",
		"-Command",
		"Get-WmiObject -Namespace 'Root\\CIMv2\\Security\\MicrosoftTpm' -Class Win32_Tpm")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error checking TPM: %v", err)
	}

	return strings.TrimSpace(out.String()), nil
}

// CheckLogs retrieves TPM logs from Windows Event Viewer.
func (t *WindowsTPM) CheckLogs() (string, error) {
	cmd := exec.Command(
		"powershell",
		"-Command",
		"Get-WinEvent -LogName 'Microsoft-Windows-Tpm*' | Select-Object TimeCreated, Message")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error retrieving TPM logs: %v", err)
	}

	return strings.TrimSpace(out.String()), nil
}
