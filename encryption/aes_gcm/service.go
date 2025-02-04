package aes_gcm

// NewAesGcmService
func NewAesGcmService(key string) IAesGcm {
	return &SAesGcm{
		key: key,
	}
}
