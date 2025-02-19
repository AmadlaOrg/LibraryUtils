// Package location 📍 is a utility to get all the paths that are required by one or many applications.
package location

// NewLocationService to set up the location service.
//
// Params:
// - 📇 appName - Is the of the application (normally all lowercase).
// - ♻️ appVersion - The version of the application.
// - 💊 pluginTypeName - The plugin type names (e.g.: HERY => entity, doorman => clerk).
func NewLocationService(appName AppName, appVersion AppVersion, pluginTypeNames PluginTypeNames) (ILocation, error) {
	serviceLocation := &SLocation{
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
