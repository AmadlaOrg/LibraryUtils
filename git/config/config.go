package config

import (
	"github.com/AmadlaOrg/LibraryUtils/pointer"
	"os"
)

// BuildDefaultConfig adds default values to the Go Git configurations
func BuildDefaultConfig(cnf *Config) *Config {
	if cnf != nil {
		//
		// Root configurations
		//
		if cnf.InsecureSkipTLS == nil {
			// MEMO: The default in the Go-Git library might be false and the best practice is to have as true
			cnf.InsecureSkipTLS = pointer.ToPtr(true)
		}
		if cnf.Timeout == 0 {
			// MEMO: The number is in seconds
			cnf.Timeout = 5
		}

		//
		// CloneOptions
		//
		if cnf.CloneOptions == nil {
			cnf.CloneOptions = &CloneOptions{
				Depth:             1,
				ShallowSubmodules: true,
				Progress:          os.Stdout,
			}
		}
	}
	return cnf
}
