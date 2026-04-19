package aes_gcm

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	validKey16 := "1234567890123456"         // 16 bytes (AES-128)
	validKey32 := "12345678901234567890123456789012" // 32 bytes (AES-256)

	tests := []struct {
		name      string
		key       string
		plainText string
		wantErr   bool
	}{
		{
			name:      "valid 32-byte key",
			key:       validKey32,
			plainText: "hello world",
		},
		{
			name:      "valid 16-byte key",
			key:       validKey16,
			plainText: "hello world",
		},
		{
			name:      "empty plaintext",
			key:       validKey32,
			plainText: "",
		},
		{
			name:      "unicode plaintext",
			key:       validKey32,
			plainText: "こんにちは世界 🌍",
		},
		{
			name:    "invalid key size",
			key:     "short",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &aesGcmImpl{key: tt.key}
			result, err := s.Encrypt(tt.plainText)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, result)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, result)

			// Verify it's valid base64
			_, decErr := base64.StdEncoding.DecodeString(result)
			assert.NoError(t, decErr)
		})
	}
}

func TestEncrypt_TwoEncryptionsDiffer(t *testing.T) {
	key := "12345678901234567890123456789012"
	s := &aesGcmImpl{key: key}

	result1, err1 := s.Encrypt("same text")
	result2, err2 := s.Encrypt("same text")

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NotEqual(t, result1, result2, "two encryptions of the same text should differ due to random nonce")
}

func TestEncrypt_MockIoReadFullFail(t *testing.T) {
	origIoReadFull := ioReadFull
	defer func() { ioReadFull = origIoReadFull }()

	ioReadFull = func(_ io.Reader, _ []byte) (int, error) {
		return 0, errors.New("random read failed")
	}

	s := &aesGcmImpl{key: "12345678901234567890123456789012"}
	result, err := s.Encrypt("test")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "random read failed")
	assert.Empty(t, result)
}

func TestEncrypt_MockCipherNewGCMFail(t *testing.T) {
	origCipherNewGCM := cipherNewGCM
	defer func() { cipherNewGCM = origCipherNewGCM }()

	cipherNewGCM = func(_ cipher.Block) (cipher.AEAD, error) {
		return nil, errors.New("gcm creation failed")
	}

	s := &aesGcmImpl{key: "12345678901234567890123456789012"}
	result, err := s.Encrypt("test")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "gcm creation failed")
	assert.Empty(t, result)
}

func TestDecrypt(t *testing.T) {
	validKey := "12345678901234567890123456789012"

	// Pre-encrypt a value for round-trip tests
	s := &aesGcmImpl{key: validKey}
	encrypted, err := s.Encrypt("hello world")
	assert.NoError(t, err)

	tests := []struct {
		name       string
		key        string
		cipherText string
		wantText   string
		wantErr    bool
		errMsg     string
	}{
		{
			name:       "round trip decrypt",
			key:        validKey,
			cipherText: encrypted,
			wantText:   "hello world",
		},
		{
			name:       "invalid base64",
			key:        validKey,
			cipherText: "not-valid-base64!!!",
			wantErr:    true,
		},
		{
			name:       "ciphertext too short",
			key:        validKey,
			cipherText: base64.StdEncoding.EncodeToString([]byte("short")),
			wantErr:    true,
			errMsg:     "invalid ciphertext",
		},
		{
			name:       "wrong key",
			key:        "abcdefghijklmnopqrstuvwxyz123456",
			cipherText: encrypted,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &aesGcmImpl{key: tt.key}
			result, err := svc.Decrypt(tt.cipherText)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantText, result)
		})
	}
}

func TestDecrypt_MockAesNewCipherFail(t *testing.T) {
	origAesNewCipher := aesNewCipher
	defer func() { aesNewCipher = origAesNewCipher }()

	aesNewCipher = func(_ []byte) (cipher.Block, error) {
		return nil, errors.New("aes cipher failed")
	}

	// Need valid base64 that's long enough (>12 bytes decoded)
	fakeData := make([]byte, 32)
	cipherText := base64.StdEncoding.EncodeToString(fakeData)

	s := &aesGcmImpl{key: "12345678901234567890123456789012"}
	result, err := s.Decrypt(cipherText)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "aes cipher failed")
	assert.Empty(t, result)
}

func TestDecrypt_MockCipherNewGCMFail(t *testing.T) {
	origCipherNewGCM := cipherNewGCM
	defer func() { cipherNewGCM = origCipherNewGCM }()

	cipherNewGCM = func(_ cipher.Block) (cipher.AEAD, error) {
		return nil, errors.New("gcm creation failed")
	}

	fakeData := make([]byte, 32)
	cipherText := base64.StdEncoding.EncodeToString(fakeData)

	s := &aesGcmImpl{key: "12345678901234567890123456789012"}
	result, err := s.Decrypt(cipherText)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "gcm creation failed")
	assert.Empty(t, result)
}

func TestRoundTrip(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		plainText string
	}{
		{
			name:      "AES-128",
			key:       "1234567890123456",
			plainText: "hello AES-128",
		},
		{
			name:      "AES-192",
			key:       "123456789012345678901234",
			plainText: "hello AES-192",
		},
		{
			name:      "AES-256",
			key:       "12345678901234567890123456789012",
			plainText: "hello AES-256",
		},
		{
			name:      "long plaintext",
			key:       "12345678901234567890123456789012",
			plainText: string(make([]byte, 10000)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &aesGcmImpl{key: tt.key}

			encrypted, err := s.Encrypt(tt.plainText)
			assert.NoError(t, err)

			decrypted, err := s.Decrypt(encrypted)
			assert.NoError(t, err)
			assert.Equal(t, tt.plainText, decrypted)
		})
	}
}

// Ensure unused import doesn't cause issues
var _ = aes.BlockSize
