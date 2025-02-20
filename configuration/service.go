package configuration

import (
	"strings"

	"github.com/spf13/viper"
)

var (
	viperNew = viper.New
)

// NewConfigurationService setups the configuration settings with Viper
//
// Params:
// - 📇 appName - Is the of the application (normally all lowercase)
// - 📁 configDirPath - The configuration absolute path to its directory.
func NewConfigurationService(appName, configDirPath string) IConfiguration {
	newViper := viperNew()
	newViper.SetEnvPrefix(strings.ToUpper(appName))
	newViper.AutomaticEnv()
	newViper.SetConfigName(appName)
	newViper.SetConfigType("yaml")

	// Looks for config in the working directory
	newViper.AddConfigPath(".")
	newViper.AddConfigPath(configDirPath)

	return &SConfiguration{
		viperInstance: newViper,
	}
}
