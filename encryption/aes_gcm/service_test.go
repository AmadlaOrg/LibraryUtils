package aes_gcm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	key := "12345678901234567890123456789012"
	svc := New(key)

	assert.NotNil(t, svc)

	// Functional check: can encrypt and decrypt
	encrypted, err := svc.Encrypt("test")
	assert.NoError(t, err)
	assert.NotEmpty(t, encrypted)

	decrypted, err := svc.Decrypt(encrypted)
	assert.NoError(t, err)
	assert.Equal(t, "test", decrypted)
}
