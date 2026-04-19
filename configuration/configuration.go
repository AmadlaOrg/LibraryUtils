package configuration

import "github.com/spf13/viper"

// Configuration 🧩 Is the interface for the NewConfigurationService.
type Configuration interface {
	Instance() *viper.Viper
}

// configImpl 🏛️ Is the main structure for the NewConfigurationService.
type configImpl struct {
	// 🐍 viperInstance - Is an instance of Viper
	viperInstance *viper.Viper
}

// Instance 🗿 of Viper with all the settings set up in the new service.
func (s *configImpl) Instance() *viper.Viper {
	return s.viperInstance
}
