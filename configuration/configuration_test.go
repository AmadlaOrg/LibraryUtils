package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstance(t *testing.T) {
	svc := New("testapp", "/tmp/config")
	instance := svc.Instance()

	assert.NotNil(t, instance)

	// Same instance returned each time
	assert.Same(t, instance, svc.Instance())
}
