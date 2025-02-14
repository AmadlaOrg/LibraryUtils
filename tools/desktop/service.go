package desktop

// NewDesktopService
//
// Params:
// - 📇 appName - Is the of the application (normally all lowercase).
func NewDesktopService(appName string) IDesktop {
	return &SDesktop{
		appName: appName,
	}
}
