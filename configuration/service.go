package configuration

import (
	"github.com/spf13/viper"
	"strings"
)

var (
	viperNew = viper.New
)

// NewConfigurationService setups the configuration settings with Viper
//
// Params:
// - 📇 appName - Is the of the application (normally all lowercase)
func NewConfigurationService(appName string) IConfiguration {
	newViper := viperNew()
	newViper.SetEnvPrefix(strings.ToUpper(appName))
	newViper.AutomaticEnv()
	newViper.SetConfigName(appName)
	newViper.SetConfigType("yaml")

	newViper.AllKeys()

	// TODO:
	//newViper.AddConfigPath(".")

	return &SConfiguration{
		appName:       appName,
		viperInstance: newViper,
	}
}
