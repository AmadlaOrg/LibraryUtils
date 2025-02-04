package dbus

import "github.com/AmadlaOrg/LibraryUtils/encryption/aes_gcm"

// NewDBusService
func NewDBusService(key string) IDBus {
	return &SDBus{
		encryptionAesGcmService: aes_gcm.NewAesGcmService(key),
	}
}
