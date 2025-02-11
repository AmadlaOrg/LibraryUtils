package location

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLocationService(t *testing.T) {
	t.Run("should return a new instance of Location", func(t *testing.T) {
		service := NewLocationService("mockname", "1.0.0", "nix")
		assert.NotNil(t, service)
		assert.IsType(t, &SLocation{}, service)
	})
}
