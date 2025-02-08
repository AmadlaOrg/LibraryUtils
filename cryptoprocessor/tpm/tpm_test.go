package tpm

import (
	"testing"
)

// TestTPM runs cross-platform tests for TPM detection.
func TestTPM(t *testing.T) {
	tpm := NewTPM()

	devices, err := tpm.ListDevices()
	if err != nil {
		t.Errorf("Failed to list TPM devices: %v", err)
	}
	if devices == "" {
		t.Log("No TPM devices found, skipping test.")
	} else {
		t.Logf("TPM Devices:\n%s", devices)
	}

	logs, err := tpm.CheckLogs()
	if err != nil {
		t.Errorf("Failed to check TPM logs: %v", err)
	}
	if logs == "" {
		t.Log("No TPM logs found, skipping test.")
	} else {
		t.Logf("TPM Logs:\n%s", logs)
	}
}
