package dbus

import (
	"fmt"
	"github.com/AmadlaOrg/LibraryUtils/encryption"
	"github.com/godbus/dbus/v5"
)

type IDBus interface{}

type SDBus struct{}

const secretKey = "32byte-long-secret-key!!!!!" // Ensure key matches Clerk-AWS

// Connect
func (s *SDBus) Connect() error {
	conn, err := dbus.SessionBus()
	if err != nil {
		return fmt.Errorf("failed to connect to D-Bus: %v", err)
	}

	obj := conn.Object("com.clerk.aws", "/com/clerk/aws")
	var encryptedJWT string

	// Call GetJWT method
	err = obj.Call("com.clerk.aws.GetJWT", 0).Store(&encryptedJWT)
	if err != nil {
		return fmt.Errorf("error calling Clerk-AWS: %v", err)
	}

	// Decrypt received JWT
	decryptedJWT, err := encryption.Decrypt(encryptedJWT, secretKey)
	if err != nil {
		return fmt.Errorf("decryption error: %v", err)
	}

	fmt.Println("Retrieved JWT:", decryptedJWT)

	return nil
}
