package tpm

// TPM provides a cross-platform interface for TPM operations.
type TPM interface {
	ListDevices() (string, error)
	CheckLogs() (string, error)
}

/*var (
	osStat      = os.Stat
	tpm2OpenTPM = tpm2.OpenTPM
)

// IsAvailable
func IsAvailable() (bool, error) {
	tpmPath := "/dev/tpm0" // TPM device path (Linux)
	_, err := osStat(tpmPath)
	if os.IsNotExist(err) {
		return false, fmt.Errorf("TPM not found")
	} else if err != nil {
		return false, err
	}

	rwc, err := tpm2OpenTPM(tpmPath)
	if err != nil {
		return false, fmt.Errorf("failed to open TPM: %v", err)
	}
	defer func(rwc io.ReadWriteCloser) {
		err := rwc.Close()
		if err != nil {
			return
		}
	}(rwc)

	return true, nil
}*/
