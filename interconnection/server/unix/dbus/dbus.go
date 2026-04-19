package dbus

import (
	"fmt"
	"github.com/AmadlaOrg/LibraryUtils/encryption/aes_gcm"
	"github.com/godbus/dbus/v5"
	"log"
)

type DBus interface{}
type dbusImpl struct{}

const secretKey = "32byte-long-secret-key!!!!!" // Ensure key is 32 bytes

// Handle incoming requests
func getJWTToken() (string, *dbus.Error) {
	// Simulated JWT token
	jwtToken := "JWT_TOKEN_123456"

	// Encrypt the token before sending
	encryptionService := aes_gcm.New(secretKey)
	encryptedToken, err := encryptionService.Encrypt(jwtToken)
	if err != nil {
		log.Println("Encryption error:", err)
		return "", dbus.NewError("com.doorman.aws.EncryptionError", []interface{}{"Encryption failed"})
	}

	return encryptedToken, nil
}

func Connect() {
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Fatal("Failed to connect to D-Bus:", err)
	}

	// Register D-Bus service
	replyObj, err := conn.RequestName("com.doorman.aws", dbus.NameFlagDoNotQueue)
	if err != nil {
		log.Fatal("Failed to query D-Bus:", err)
	}
	if replyObj != dbus.RequestNameReplyPrimaryOwner {
		log.Fatal("D-Bus name already taken")
	}

	// Register object & method
	obj := conn.Export(getJWTToken, "/com/doorman/aws", "com.doorman.aws")
	if obj != nil {
		log.Fatal("Failed to export method:", obj)
	}

	fmt.Println("doorman-aws running on D-Bus...")
	select {} // Keep running
}
