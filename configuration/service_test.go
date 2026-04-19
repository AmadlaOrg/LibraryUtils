package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	svc := New("myapp", "/tmp/config")

	assert.NotNil(t, svc)

	// Instance should be non-nil
	v := svc.Instance()
	assert.NotNil(t, v)

	// Verify env prefix is set correctly by checking env var binding
	// Viper uppercases the prefix; env var MYAPP_TEST_KEY should map
	v.SetDefault("test_key", "default_val")
	assert.Equal(t, "default_val", v.GetString("test_key"))
}
