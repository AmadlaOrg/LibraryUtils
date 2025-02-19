package location

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/AmadlaOrg/LibraryUtils/file"
	"github.com/adrg/xdg"
)

// ILocation 🧩 Is the interface for the NewLocationService.
type ILocation interface {
	SystemPaths() *SystemPaths
	ApplicationPaths() *ApplicationPaths
}

// SLocation 🏛️ Is the main structure for the NewLocationService.
type SLocation struct {
	appName         AppName
	appVersion      AppVersion
	pluginTypeNames PluginTypeNames
	paths           *Paths
}

// For mocking 🥸.
var (
	xdgCacheFile   = xdg.CacheFile
	xdgConfigFile  = xdg.ConfigFile
	xdgRuntimeFile = xdg.RuntimeFile
	xdgDataFile    = xdg.DataFile
)

// SystemPaths returns struct of all systems paths
func (service *SLocation) SystemPaths() *SystemPaths {
	return service.paths.SystemPaths
}

// ApplicationPaths returns the struct of all the main application paths
func (service *SLocation) ApplicationPaths() *ApplicationPaths {
	return service.paths.ApplicationPaths
}

// setAllPaths calls on all the private methods to set in the struct all the absolute paths.
func (service *SLocation) setAllPaths() error {
	relFilePath := fmt.Sprintf("%s/%s", service.appName, service.appName)

	err := service.setSystemPaths()
	if err != nil {
		return err
	}

	err = service.setDataPaths()
	if err != nil {
		return err
	}

	err = service.setConfigPaths(relFilePath)
	if err != nil {
		return err
	}

	err = service.setStatePaths()
	if err != nil {
		return err
	}

	err = service.setCachePaths(relFilePath)
	if err != nil {
		return err
	}

	service.setBinPath()

	err = service.setPluginPaths()
	if err != nil {
		return err
	}

	err = service.setSecretPaths()
	if err != nil {
		return err
	}

	return nil
}

// setSystemPaths sets the system paths.
func (service *SLocation) setSystemPaths() error {
	sysPaths := &SystemPaths{
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

	if runtime.GOOS != "windows" &&
		runtime.GOOS != "darwin" &&
		runtime.GOOS != "plan9" {
		sysPaths.UserLocalHome = filepath.Join(xdg.DataHome, ".local")

		var existApplicationsPath string
		for _, applicationsPath := range xdg.ApplicationDirs {
			if file.Exists(applicationsPath) {
				existApplicationsPath = applicationsPath
				break
			}
		}
		if existApplicationsPath != "" {
			sysPaths.UserApplicationsHome = existApplicationsPath
			service.paths.ApplicationPaths.ApplicationsDesktopFile = fmt.Sprintf(
				"%s/%s.desktop",
				existApplicationsPath,
				string(service.appName),
			)
		} else {
			return errors.New("applications directory not found")
		}
	}

	service.paths.SystemPaths = sysPaths

	return nil
}

// setDataPaths sets the application specific data directory path.
func (service *SLocation) setDataPaths() error {
	dataDir, err := xdgDataFile(string(service.appName))
	if err != nil {
		return errors.Join(fmt.Errorf(`xdg.DataFile was unable to set "%s" path`, dataDir), err)
	}
	service.paths.ApplicationPaths.DataHome = dataDir

	return nil
}

// setConfigPaths sets the application config directory and file YAML absolute paths.
func (service *SLocation) setConfigPaths(relFilePath string) error {
	configFile, err := xdgConfigFile(relFilePath + ".yaml")
	if err != nil {
		return errors.Join(fmt.Errorf(`xdg.ConfigFile was unable to set "%s" path`, configFile), err)
	}
	service.paths.ApplicationPaths.ConfigFile = configFile
	service.paths.ApplicationPaths.ConfigHome = filepath.Dir(configFile)

	return nil
}

// setStatePaths sets the absolute path to the state directory.
func (service *SLocation) setStatePaths() error {
	stateDir, err := xdg.StateFile(string(service.appName))
	if err != nil {
		return errors.Join(fmt.Errorf(`xdg.StateFile was unable to set "%s" path`, stateDir), err)
	}
	service.paths.ApplicationPaths.StateHome = stateDir

	return nil
}

// setCachePaths sets the to the cache directory and file absolute paths.
func (service *SLocation) setCachePaths(relFilePath string) error {
	cacheFilePath, err := xdgCacheFile(relFilePath + ".cache")
	if err != nil {
		return errors.Join(fmt.Errorf(`xdg.CacheFile was unable to set "%s" path`, cacheFilePath), err)
	}
	service.paths.ApplicationPaths.CacheHome = filepath.Dir(cacheFilePath)
	service.paths.ApplicationPaths.CacheFile = cacheFilePath

	return nil
}

// setBinPath sets the absolute bin path.
// The bin is the application main binary.
func (service *SLocation) setBinPath() {
	service.paths.ApplicationPaths.BinFile = fmt.Sprintf("%s/%s", xdg.BinHome, service.appName)
}

// setPluginPaths sets the path or paths to the plugin directory.
func (service *SLocation) setPluginPaths() error {
	pluginDirs := make(map[string]string)
	for _, pluginTypeName := range service.pluginTypeNames {
		pluginDirs[pluginTypeName] = fmt.Sprintf("%s/%s.d", xdg.DataHome, pluginTypeName)
	}
	service.paths.ApplicationPaths.PluginsHome = pluginDirs

	return nil
}

// setSecretPaths sets the secret absolute paths.
func (service *SLocation) setSecretPaths() error {
	secretsRelPath := fmt.Sprintf("%s/secrets", service.appName)

	tmpSecrets, err := xdgRuntimeFile(secretsRelPath)
	if err != nil {
		return errors.Join(fmt.Errorf(`xdg.RuntimeFile was unable to set "%s" path`, tmpSecrets), err)
	}
	service.paths.ApplicationPaths.RuntimeDir = filepath.Dir(tmpSecrets)
	service.paths.ApplicationPaths.SecretsHome = tmpSecrets

	tmpSecretsMTls, err := xdgRuntimeFile(secretsRelPath + "/mTLS")
	if err != nil {
		return errors.Join(fmt.Errorf(`xdg.RuntimeFile was unable to set "%s" path`, tmpSecrets), err)
	}
	service.paths.ApplicationPaths.MTLSHome = tmpSecretsMTls

	return nil
}
