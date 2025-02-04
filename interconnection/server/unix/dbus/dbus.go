package dbus

import (
	"fmt"
	"github.com/AmadlaOrg/LibraryUtils/encryption"
	"github.com/godbus/dbus/v5"
	"log"
)

type IDBus interface{}
type SDBus struct{}

const secretKey = "32byte-long-secret-key!!!!!" // Ensure key is 32 bytes

// Handle incoming requests
func getJWTToken() (string, *dbus.Error) {
	// Simulated JWT token
	jwtToken := "JWT_TOKEN_123456"

	// Encrypt the token before sending
	encryptedToken, err := encryption.Encrypt(jwtToken, secretKey)
	if err != nil {
		log.Println("Encryption error:", err)
		return "", dbus.NewError("com.clerk.aws.EncryptionError", []interface{}{"Encryption failed"})
	}

	return encryptedToken, nil
}

func Connect() {
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Fatal("Failed to connect to D-Bus:", err)
	}

	// Register D-Bus service
	replyObj, err := conn.RequestName("com.clerk.aws", dbus.NameFlagDoNotQueue)
	if err != nil {
		log.Fatal("Failed to query D-Bus:", err)
	}
	if replyObj != dbus.RequestNameReplyPrimaryOwner {
		log.Fatal("D-Bus name already taken")
	}

	// Register object & method
	obj := conn.Export(getJWTToken, "/com/clerk/aws", "com.clerk.aws")
	if obj != nil {
		log.Fatal("Failed to export method:", obj)
	}

	fmt.Println("Clerk-AWS running on D-Bus...")
	select {} // Keep running
}
