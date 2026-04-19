//go:build !windows

package tpm

import (
	"github.com/AmadlaOrg/LibraryUtils/cryptoprocessor/tpm/unix"
)

// NewTPM creates the correct TPM implementation based on the OS.
func NewTPM() TPM {
	return unix.NewUnixTPM()
}
