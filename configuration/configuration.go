package configuration

import "github.com/spf13/viper"

// IConfiguration 🧩 Is the interface for the NewConfigurationService.
type IConfiguration interface {
	Instance() *viper.Viper
}

// SConfiguration 🏛️ Is the main structure for the NewConfigurationService.
type SConfiguration struct {
	// 🐍 viperInstance - Is an instance of Viper
	viperInstance *viper.Viper
}

// Instance 🗿 of Viper with all the settings set up in the new service.
func (s *SConfiguration) Instance() *viper.Viper {
	return s.viperInstance
}
