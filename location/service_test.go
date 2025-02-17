package location

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLocationService(t *testing.T) {
	t.Run("should return a new instance of Location", func(t *testing.T) {
		service, err := NewLocationService("mockname", "1.0.0", []string{"nix"})
		assert.NoError(t, err)
		assert.NotNil(t, service)
		assert.IsType(t, &SLocation{}, service)
	})
}
