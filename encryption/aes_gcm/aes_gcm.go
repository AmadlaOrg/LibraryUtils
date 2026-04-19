package aes_gcm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

type AesGcm interface {
	Encrypt(plainText string) (string, error)
	Decrypt(cipherText string) (string, error)
}

type aesGcmImpl struct {
	key string
}

var (
	aesNewCipher      = aes.NewCipher
	ioReadFull        = io.ReadFull
	cipherNewGCM      = cipher.NewGCM
	base64StdEncoding = base64.StdEncoding
)

// Encrypt text using AES-GCM
func (s *aesGcmImpl) Encrypt(plainText string) (string, error) {
	block, err := aesNewCipher([]byte(s.key))
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
func (s *aesGcmImpl) Decrypt(cipherText string) (string, error) {
	data, err := base64StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aesNewCipher([]byte(s.key))
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
