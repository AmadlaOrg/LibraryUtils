package plugin

// New creates a new plugin service.
func New() Plugin {
	return &pluginImpl{}
}
