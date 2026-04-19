package dbus

import "github.com/AmadlaOrg/LibraryUtils/encryption/aes_gcm"

// New
func New(key string) DBus {
	return &dbusImpl{
		encryptionAesGcmService: aes_gcm.New(key),
	}
}
