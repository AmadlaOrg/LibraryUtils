package configuration

import "github.com/spf13/viper"

type IConfiguration interface {
	AllPropertyNames() []string
	AllProperties() map[string]any
	Set(key string, defaultValue any)
}
type SConfiguration struct {
	// 📇 appName - Is the of the application (normally all lowercase)
	appName string

	// 🐍 viperInstance - Is an instance of Viper
	viperInstance *viper.Viper
}

// AllPropertyNames returns array of configuration property names
func (s *SConfiguration) AllPropertyNames() []string {
	return s.viperInstance.AllKeys()
}

// AllProperties all the configuration property name with their values
func (s *SConfiguration) AllProperties() map[string]any {
	return s.viperInstance.AllSettings()
}

// Set a configuration property
//
// Params:
// - 📇 name - For the configuration property name
// - 🌟 defaultValue - The configuration property default value
func (s *SConfiguration) Set(name string, defaultValue any) {
	s.viperInstance.Set(name, defaultValue)
}
