package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheImpl_Open(t *testing.T) {
	svc := &cacheImpl{}
	err := svc.Open()
	assert.NoError(t, err)
}
