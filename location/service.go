// Package location 📍
package location

import (
	"fmt"
	"path/filepath"
)

// NewLocationService to set up the location service.
//
// Params:
// - 📇 appName - Is the of the application (normally all lowercase).
// - ♻️ appVersion - The version of the application.
// - 💊 pluginTypeName - The plugin type name (e.g.: HERY => entity, doorman => clerk).
func NewLocationService(appName, appVersion, pluginTypeName string) ILocation {
	serviceLocation := &SLocation{
		appName:    appName,
		appVersion: appVersion,
	}

	// Set
	serviceLocation.setSystemPaths()

	sysPaths := serviceLocation.paths.SystemPaths

	// TODO: What happens when there is no path set
	// TODO: There needs to be a check for each of these paths... Maybe use `mkdir -p`
	/*if sysPaths.DataHome == "" {
	}

	if sysPaths.ConfigHome == "" {
	}

	if sysPaths.StateHome == "" {
	}

	if sysPaths.CacheHome == "" {
	}

	if sysPaths.RuntimeDir == "" {
	}

	if sysPaths.BinHome == "" {
	}*/

	appDataHome := filepath.Join(sysPaths.DataHome, appName)
	appConfigHome := filepath.Join(sysPaths.ConfigHome, appName)
	appStateHome := filepath.Join(sysPaths.StateHome, appName)
	appCacheHome := filepath.Join(sysPaths.CacheHome, appName)
	appRuntimeDir := filepath.Join(sysPaths.RuntimeDir, appName)
	appBinHome := filepath.Join(sysPaths.BinHome, appName)

	// TODO: Validate them

	return &SLocation{
		paths: &Paths{
			SystemPaths: sysPaths,
			ThisApplicationPaths: &Application{
				Name: appName,
				Paths: &ApplicationPaths{
					DataHome:   appDataHome,
					ConfigHome: appConfigHome,
					StateHome:  appStateHome,
					CacheHome:  appCacheHome,
					RuntimeDir: appRuntimeDir,
					BinFile:    appBinHome,

					// Custom:
					ConfigFile: filepath.Join(appConfigHome, fmt.Sprintf("%s.json", appName)),

					// TODO:
					PluginsHome: filepath.Join(appDataHome, fmt.Sprintf("%s.d", pluginTypeName)),

					// TODO:
					CacheFile: filepath.Join(appDataHome, fmt.Sprintf("%s.cache", appName)),

					ApplicationsDesktopFile: filepath.Join(
						sysPaths.DataHome,
						"applications",
						fmt.Sprintf("%s.desktop", appName),
					),
				},
			},
		},
	}
}
