// Package location 📍 is a utility to get all the paths that are required by one or many applications.
package location

// New to set up the location service.
//
// Params:
// - 📇 appName - Is the of the application (normally all lowercase).
// - ♻️ appVersion - The version of the application.
// - 💊 pluginTypeName - The plugin type names (e.g.: HERY => entity, doorman => plugin).
func New(appName AppName, appVersion AppVersion, pluginTypeNames PluginTypeNames) (Location, error) {
	serviceLocation := &locationImpl{
		appName:         appName,
		appVersion:      appVersion,
		pluginTypeNames: pluginTypeNames,
		paths: &Paths{
			SystemPaths: &SystemPaths{},
			ApplicationPaths: &ApplicationPaths{
				PluginsHome: make(map[string]string),
			},
		},
	}

	err := serviceLocation.setAllPaths()
	if err != nil {
		return nil, err
	}

	return serviceLocation, nil
}
