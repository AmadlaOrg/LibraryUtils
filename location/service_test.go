package location

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLocationService(t *testing.T) {
	t.Run("should return a new instance of Location", func(t *testing.T) {
		service := NewLocationService()
		assert.NotNil(t, service)
		assert.IsType(t, &SLocation{}, service)
	})
}
