//go:build !windows

package socket

import (
	"os"
	"path/filepath"
)

var socketPath = filepath.Join(os.TempDir(), "doorman.sock")
