package location

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

// NewLocationService to set up the location service
//
// Params:
// - 📇 appName - Is the of the application (normally all lowercase)
// - ♻️ version - The version of the application
func NewLocationService(appName, version string) ILocation {
	sysPaths := &System{
		Home:            xdg.Home,
		DataHome:        xdg.DataHome,
		DataDirs:        xdg.DataDirs,
		ConfigHome:      xdg.ConfigHome,
		ConfigDirs:      xdg.ConfigDirs,
		StateHome:       xdg.StateHome,
		CacheHome:       xdg.CacheHome,
		RuntimeDir:      xdg.RuntimeDir,
		BinHome:         xdg.BinHome,
		UserDirs:        xdg.UserDirs,
		FontDirs:        xdg.FontDirs,
		ApplicationDirs: xdg.ApplicationDirs,
	}

	// TODO: What happens when there is no path set
	// TODO: There needs to be a check for each of these paths... Maybe use `mkdir -p`
	if sysPaths.DataHome == "" {
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
	}

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
					BinHome:    appBinHome,
				},
			},
		},
	}
}
