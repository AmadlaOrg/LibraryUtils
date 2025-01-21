package config

import (
	"github.com/AmadlaOrg/LibraryUtils/pointer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildDefaultConfig_use_default(t *testing.T) {
	cnf := Config{}
	got := BuildDefaultConfig(&cnf)
	assert.Equal(t, &Config{
		InsecureSkipTLS: pointer.ToPtr(true),
	}, got)
}

func TestBuildDefaultConfig_NO_usage_of_default(t *testing.T) {
	cnf := Config{
		InsecureSkipTLS: pointer.ToPtr(false),
	}
	got := BuildDefaultConfig(&cnf)
	assert.Equal(t, &Config{
		InsecureSkipTLS: pointer.ToPtr(false),
	}, got)
}
