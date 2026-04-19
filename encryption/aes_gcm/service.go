package aes_gcm

// New
func New(key string) AesGcm {
	return &aesGcmImpl{
		key: key,
	}
}
