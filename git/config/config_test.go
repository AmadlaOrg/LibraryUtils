package config

import (
	"github.com/AmadlaOrg/LibraryUtils/pointer"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBuildDefaultConfig_use_default(t *testing.T) {
	cnf := Config{}
	got := BuildDefaultConfig(&cnf)
	assert.Equal(t, &Config{
		InsecureSkipTLS: pointer.ToPtr(true),
		Timeout:         5,
		CloneOptions: &CloneOptions{
			Depth:             1,
			ShallowSubmodules: true,
			Progress:          os.Stdout,
		},
	}, got)
}

func TestBuildDefaultConfig_NO_usage_of_default(t *testing.T) {
	cnf := Config{
		InsecureSkipTLS: pointer.ToPtr(false),
	}
	got := BuildDefaultConfig(&cnf)
	assert.Equal(t, &Config{
		InsecureSkipTLS: pointer.ToPtr(false),
		Timeout:         5,
		CloneOptions: &CloneOptions{
			Depth:             1,
			ShallowSubmodules: true,
			Progress:          os.Stdout,
		},
	}, got)
}
