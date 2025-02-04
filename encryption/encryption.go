package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

var (
	aesNewCipher      = aes.NewCipher
	ioReadFull        = io.ReadFull
	cipherNewGCM      = cipher.NewGCM
	base64StdEncoding = base64.StdEncoding
)

// Encrypt text using AES-GCM
func Encrypt(plainText, key string) (string, error) {
	block, err := aesNewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12) // 96-bit nonce
	if _, err := ioReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesGCM, err := cipherNewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt AES-GCM encrypted text
func Decrypt(cipherText, key string) (string, error) {
	data, err := base64StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aesNewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	if len(data) < 12 {
		return "", fmt.Errorf("invalid ciphertext")
	}

	nonce, ciphertext := data[:12], data[12:]
	aesGCM, err := cipherNewGCM(block)
	if err != nil {
		return "", err
	}

	plainText, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
