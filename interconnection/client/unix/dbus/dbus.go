package dbus

import (
	"fmt"
	"github.com/AmadlaOrg/LibraryUtils/encryption/aes_gcm"
	"github.com/godbus/dbus/v5"
)

type DBus interface{}

type dbusImpl struct {
	encryptionAesGcmService aes_gcm.AesGcm
}

const secretKey = "32byte-long-secret-key!!!!!" // Ensure key matches doorman-aws

// Connect
func (s *dbusImpl) Connect() error {
	conn, err := dbus.SessionBus()
	if err != nil {
		return fmt.Errorf("failed to connect to D-Bus: %v", err)
	}

	obj := conn.Object("com.doorman.aws", "/com/doorman/aws")
	var encryptedJWT string

	// Call GetJWT method
	err = obj.Call("com.doorman.aws.GetJWT", 0).Store(&encryptedJWT)
	if err != nil {
		return fmt.Errorf("error calling doorman-aws: %v", err)
	}

	// Decrypt received JWT
	decryptedJWT, err := s.encryptionAesGcmService.Decrypt(encryptedJWT)
	if err != nil {
		return fmt.Errorf("decryption error: %v", err)
	}

	fmt.Println("Retrieved JWT:", decryptedJWT)

	return nil
}
