package location

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/AmadlaOrg/LibraryUtils/file"
	"github.com/adrg/xdg"
)

type ILocation interface {
	SystemPaths() *SystemPaths
	ThisAppPaths() *ApplicationPaths
}
type SLocation struct {
	appName         AppName
	appVersion      AppVersion
	pluginTypeNames PluginTypeNames
	paths           *Paths
}

var (
	xdgCacheFile   = xdg.CacheFile
	xdgConfigFile  = xdg.ConfigFile
	xdgRuntimeFile = xdg.RuntimeFile
	xdgDataFile    = xdg.DataFile
)

func (service *SLocation) setAll() error {
	service.paths.ThisApplicationPaths.Name = service.appName
	service.paths.ThisApplicationPaths.PluginTypeNames = service.pluginTypeNames

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

	service.setBinPaths()

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

// setSystemPaths sets the basic system paths
func (service *SLocation) setSystemPaths() error {
	//
	// System
	//
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

		// TODO:
		//sysPaths.UserApplicationsHome = filepath.Join(xdg.DataHome, "applications")

		var existApplicationsPath string
		for _, applicationsPath := range xdg.ApplicationDirs {
			if file.Exists(applicationsPath) {
				existApplicationsPath = applicationsPath
				break
			}
		}
		if existApplicationsPath != "" {
			sysPaths.UserApplicationsHome = existApplicationsPath
			service.paths.ThisApplicationPaths.Paths.ApplicationsDesktopFile = fmt.Sprintf(
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

// setDataPaths
func (service *SLocation) setDataPaths() error {
	dataDir, err := xdgDataFile(string(service.appName))
	if err != nil {
		return errors.Join(fmt.Errorf(`xdg.DataFile was unable to set "%s" path`, dataDir), err)
	}
	service.paths.ThisApplicationPaths.Paths.DataHome = dataDir

	return nil
}

// setConfigPaths
func (service *SLocation) setConfigPaths(relFilePath string) error {
	configFile, err := xdgConfigFile(relFilePath + ".yaml")
	if err != nil {
		return errors.Join(fmt.Errorf(`xdg.ConfigFile was unable to set "%s" path`, configFile), err)
	}
	service.paths.ThisApplicationPaths.Paths.ConfigFile = configFile
	service.paths.ThisApplicationPaths.Paths.ConfigHome = filepath.Dir(configFile)

	return nil
}

// setStatePaths
func (service *SLocation) setStatePaths() error {
	stateDir, err := xdg.StateFile(string(service.appName))
	if err != nil {
		return errors.Join(fmt.Errorf(`xdg.StateFile was unable to set "%s" path`, stateDir), err)
	}
	service.paths.ThisApplicationPaths.Paths.StateHome = stateDir

	return nil
}

// setCachePaths
func (service *SLocation) setCachePaths(relFilePath string) error {
	cacheFilePath, err := xdgCacheFile(relFilePath + ".cache")
	if err != nil {
		return errors.Join(fmt.Errorf(`xdg.CacheFile was unable to set "%s" path`, cacheFilePath), err)
	}
	service.paths.ThisApplicationPaths.Paths.CacheHome = filepath.Dir(cacheFilePath)
	service.paths.ThisApplicationPaths.Paths.CacheFile = cacheFilePath

	return nil
}

// setBinPaths
func (service *SLocation) setBinPaths() {
	service.paths.ThisApplicationPaths.Paths.BinFile = fmt.Sprintf("%s/%s", xdg.BinHome, service.appName)
}

// setPluginPaths
func (service *SLocation) setPluginPaths() error {
	pluginDirs := make(map[string]string)
	for _, pluginTypeName := range service.pluginTypeNames {
		pluginDirs[pluginTypeName] = fmt.Sprintf("%s/%s.d", xdg.DataHome, pluginTypeName)
	}
	service.paths.ThisApplicationPaths.Paths.PluginsHome = pluginDirs

	return nil
}

// setSecretPaths
func (service *SLocation) setSecretPaths() error {
	secretsRelPath := fmt.Sprintf("%s/secrets", service.appName)

	tmpSecrets, err := xdgRuntimeFile(secretsRelPath)
	if err != nil {
		return errors.Join(fmt.Errorf(`xdg.RuntimeFile was unable to set "%s" path`, tmpSecrets), err)
	}
	service.paths.ThisApplicationPaths.Paths.RuntimeDir = filepath.Dir(tmpSecrets)
	service.paths.ThisApplicationPaths.Paths.SecretsHome = tmpSecrets

	tmpSecretsMTls, err := xdgRuntimeFile(secretsRelPath + "/mTLS")
	if err != nil {
		return errors.Join(fmt.Errorf(`xdg.RuntimeFile was unable to set "%s" path`, tmpSecrets), err)
	}
	service.paths.ThisApplicationPaths.Paths.MTLSHome = tmpSecretsMTls

	return nil
}

// SystemPaths returns struct of all systems paths
func (service *SLocation) SystemPaths() *SystemPaths {
	return service.paths.SystemPaths
}

// ThisAppPaths returns the struct of all the main application paths
func (service *SLocation) ThisAppPaths() *ApplicationPaths {
	return service.paths.ThisApplicationPaths.Paths
}

// Paths return the absolute paths for the different parts of storage
/*func (d *AbsPaths) Paths(collectionName string) (*AbsPaths, error) {
	mainPath, err := d.Main()
	if err != nil {
		return &AbsPaths{}, err
	}

	return d.paths(mainPath, collectionName)
}

// Main returns the main path for `.hery` storage path
// TODO: Maybe name it Root
func (d *AbsPaths) Main() (string, error) {
	//
	// Using env var
	//

	envStoragePathValue := os.Getenv(HeryStoragePath)

	if envStoragePathValue != "" {
		envStoragePath, err := filepathAbs(envStoragePathValue)
		if err != nil {
			return "", err
		}
		return envStoragePath, nil
	}

	//
	// Using current location
	//

	cwd, err := osGetwd()
	if err != nil {
		return "", err
	}

	localStoragePath := filepathJoin(cwd, ".hery")

	if fileExists(localStoragePath) {
		return localStoragePath, nil
	}

	//
	// Default
	//

	var mainDir string
	switch runtime.GOOS {
	case "windows":
		appDataDir := os.Getenv("APPDATA")
		mainDir = filepathJoin(appDataDir, "Hery")
	default: // "linux" and "darwin" (macOS)
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("error getting home directory: %s", err)
		}
		mainDir = filepathJoin(homeDir, ".hery")
	}

	return mainDir, nil
}

// EntityPath returns the absolute path to a specific entity
func (d *AbsPaths) EntityPath(entitiesPath, entityRelativePath string) string {
	return filepathJoin(entitiesPath, entityRelativePath)
}

// TmpPaths returns the tmp absolute paths for the different parts of storage
func (d *AbsPaths) TmpPaths(collectionName string) (*AbsPaths, error) {
	mainTmpPath, err := d.TmpMain()
	if err != nil {
		return &AbsPaths{}, err
	}

	return d.paths(mainTmpPath, collectionName)
}

// TmpMain returns the tmp main path for `.hery` storage path
func (d *AbsPaths) TmpMain() (string, error) {
	tempDir, err := osMkdirTemp("", "hery_*")
	if err != nil {
		return "", err
	}

	storageTmpPath := filepath.Join(tempDir, ".hery")
	err = osMkdirAll(storageTmpPath, perm)
	if err != nil {
		return "", err
	}

	return storageTmpPath, nil
}

// MakePaths makes all the storage subdirectories
func (d *AbsPaths) MakePaths(paths AbsPaths) error {
	err := osMkdirAll(paths.Storage, perm)
	if err != nil {
		return err
	}

	err = osMkdirAll(paths.Catalog, perm)
	if err != nil {
		return err
	}

	err = osMkdirAll(paths.Collection, perm)
	if err != nil {
		return err
	}

	err = osMkdirAll(paths.Entities, perm)
	if err != nil {
		return err
	}

	return nil
}

// paths internal function to generate the paths for different part of storage based on main path and collection name
func (d *AbsPaths) paths(mainPath, collectionName string) (*AbsPaths, error) {
	catalogPath := d.catalogPath(mainPath)
	collectionPath := d.collectionPath(catalogPath, collectionName)
	entityPath := d.entitiesPath(collectionPath)
	cachePath := d.cachePath(collectionName, collectionPath)

	return &AbsPaths{
		Storage:    mainPath,
		Catalog:    catalogPath,
		Collection: collectionPath,
		Entities:   entityPath,
		Cache:      cachePath,
	}, nil
}

// catalogPath returns the catalog absolute path
func (d *AbsPaths) catalogPath(mainPath string) string {
	return filepathJoin(mainPath, "collection")
}

// collectionPath returns the collection absolute path
func (d *AbsPaths) collectionPath(mainPath, collectionName string) string {
	return filepathJoin(mainPath, collectionName)
}

// entityPath returns the entity absolute path
func (d *AbsPaths) entitiesPath(collectionPath string) string {
	return filepathJoin(collectionPath, "entity")
}

// cachePath returns the collection cache absolute path
func (d *AbsPaths) cachePath(collectionName, collectionPath string) string {
	return filepathJoin(collectionPath, fmt.Sprintf("%s.cache", collectionName))
}*/
