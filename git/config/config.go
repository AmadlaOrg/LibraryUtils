package config

import "github.com/AmadlaOrg/LibraryUtils/pointer"

// BuildDefaultConfig adds default values to the Go Git configurations
func BuildDefaultConfig(cnf *Config) *Config {
	if cnf != nil {
		if cnf.InsecureSkipTLS == nil {
			cnf.InsecureSkipTLS = pointer.ToPtr(true)
		}
	}
	return cnf
}
