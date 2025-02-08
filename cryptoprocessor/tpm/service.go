package tpm

import (
	"runtime"

	"github.com/AmadlaOrg/LibraryUtils/cryptoprocessor/tpm/unix"
	"github.com/AmadlaOrg/LibraryUtils/interconnection/server/windows"
)

// NewTPM creates the correct TPM implementation based on the OS.
func NewTPM() TPM {
	if runtime.GOOS == "windows" {
		return windows.NewWindowsTPM()
	}
	return unix.NewUnixTPM()
}
