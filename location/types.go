// Package location 📍
package location

import (
	"github.com/adrg/xdg"
)

// PluginTypeNames
// TODO: Should be moved to the plugin package
type PluginTypeNames []string
type AppName string
type AppVersion string

// Paths 🛣️ contains all the paths
type Paths struct {
	// 🤖 SystemPaths contains the struct of the system paths
	SystemPaths *SystemPaths

	// 🚩 ThisApplicationPaths
	ThisApplicationPaths *Application

	// 🖧 Applications
	Applications *[]Application
}

// SystemPaths contains the system directories paths
type SystemPaths struct {
	// Home contains the path of the user's home directory.
	//
	// Example:
	// - 🐧 Linux: /home/username
	// - 🍎 Mac OS X: /Users/username ($HOME)
	// - 🪟 Windows: C:\Users\Username (%USERPROFILE%)
	Home string

	// UserLocalHome contains the path of the user's .local directory.
	//
	// Example:
	// - 🐧 Linux: /home/username/.local
	// - 🍎 Mac OS X: None
	// - 🪟 Windows: None
	UserLocalHome string

	// DataHome defines the base directory relative to which user-specific
	// data files should be stored. This directory is defined by the
	// $XDG_DATA_HOME environment variable. If the variable is not set,
	// a default equal to $HOME/.local/share should be used.
	//
	// Example:
	// - 🐧 Linux: /home/username/.local/share
	// - 🍎 Mac OS X: /Users/username/Library/Application Support
	// - 🪟 Windows: C:\Users\Username\AppData\Local (%LOCALAPPDATA%)
	DataHome string

	// DataDirs defines the preference-ordered set of base directories to
	// search for data files in addition to the DataHome base directory.
	// This set of directories is defined by the $XDG_DATA_DIRS environment
	// variable. If the variable is not set, the default directories
	// to be used are /usr/local/share and /usr/share, in that order. The
	// DataHome directory is considered more important than any of the
	// directories defined by DataDirs. Therefore, user data files should be
	// written relative to the DataHome directory, if possible.
	//
	// Example:
	// - 🐧 Linux: /usr/local/share and /usr/share
	// - 🍎 Mac OS X: /Library/Application Support and /System/Library/Application Support
	// - 🪟 Windows: C:\ProgramData
	DataDirs []string

	// ConfigHome defines the base directory relative to which user-specific
	// configuration files should be written. This directory is defined by
	// the $XDG_CONFIG_HOME environment variable. If the variable is
	// not set, a default equal to $HOME/.config should be used.
	//
	// Example:
	// - 🐧 Linux: /home/username/.config
	// - 🍎 Mac OS X: /Users/username/Library/Preferences
	// - 🪟 Windows: C:\Users\Username\AppData\Roaming (%APPDATA%)
	ConfigHome string

	// ConfigDirs defines the preference-ordered set of base directories to
	// search for configuration files in addition to the ConfigHome base
	// directory. This set of directories is defined by the $XDG_CONFIG_DIRS
	// environment variable. If the variable is not set, a default equal
	// to /etc/xdg should be used. The ConfigHome directory is considered
	// more important than any of the directories defined by ConfigDirs.
	// Therefore, user config files should be written relative to the
	// ConfigHome directory, if possible.
	//
	// Example:
	// - 🐧 Linux: /etc/xdg and more
	// - 🍎 Mac OS X: /Library/Preferences
	// - 🪟 Windows: C:\ProgramData
	ConfigDirs []string

	// StateHome defines the base directory relative to which user-specific
	// state files should be stored. This directory is defined by the
	// $XDG_STATE_HOME environment variable. If the variable is not set,
	// a default equal to ~/.local/state should be used.
	//
	// Example:
	// - 🐧 Linux: /home/username/.local/state
	// - 🍎 Mac OS X: /Users/username/Library/Application Support
	// - 🪟 Windows: C:\Users\Username\AppData\Local
	StateHome string

	// CacheHome defines the base directory relative to which user-specific
	// non-essential (cached) data should be written. This directory is
	// defined by the $XDG_CACHE_HOME environment variable. If the variable
	// is not set, a default equal to $HOME/.cache should be used.
	//
	// Example:
	// - 🐧 Linux: /home/username/.cache
	// - 🍎 Mac OS X: /Users/username/Library/Caches
	// - 🪟 Windows: C:\Users\Username\AppData\Local\cache
	CacheHome string

	// RuntimeDir defines the base directory relative to which user-specific
	// non-essential runtime files and other file objects (such as sockets,
	// named pipes, etc.) should be stored. This directory is defined by the
	// $XDG_RUNTIME_DIR environment variable. If the variable is not set,
	// applications should fall back to a replacement directory with similar
	// capabilities. Applications should use this directory for communication
	// and synchronization purposes and should not place larger files in it,
	// since it might reside in runtime memory and cannot necessarily be
	// swapped out to disk.
	//
	// Example:
	// - 🐧 Linux: /run/username/1000
	// - 🍎 Mac OS X: /var/folders/.../T (Temporary directory)
	// - 🪟 Windows: Usually not set (empty string)
	RuntimeDir string

	// BinHome defines the base directory relative to which user-specific
	// binary files should be written. This directory is defined by
	// the non-standard $XDG_BIN_HOME environment variable. If the variable is
	// not set, a default equal to $HOME/.local/bin should be used.
	//
	// Example:
	// - 🐧 Linux: /home/username/.local/bin
	// - 🍎 Mac OS X: /usr/local/bin, /Users/username/bin (if exists)
	// - 🪟 Windows: C:\Users\Username\AppData\Local\Microsoft\WindowsApps
	BinHome string

	// UserDirs defines the locations of well known user directories.
	//
	// Example:
	// - 🐧 Linux: Uses ~/Documents, ~/Downloads, etc.
	// - 🍎 Mac OS X: Uses ~/Documents, ~/Downloads, etc.
	// - 🪟 Windows: Uses SHGetKnownFolderPath, e.g., C:\Users\Username\Documents, C:\Users\Username\Downloads
	UserDirs xdg.UserDirectories

	// FontDirs defines the common locations where font files are stored.
	//
	// Example:
	// - 🐧 Linux:
	// - 🍎 Mac OS X: /Library/Fonts, ~/Library/Fonts
	// - 🪟 Windows: C:\Windows\Fonts, C:\Users\Username\AppData\Local\Microsoft\Windows\Fonts
	FontDirs []string

	// ApplicationDirs defines the common locations of applications.
	//
	// Example:
	// - 🐧 Linux: /home/username/.local/share/applications, /usr/local/share/applications, /usr/share/applications, etc.
	// - 🍎 Mac OS X: /Applications, ~/Applications
	// - 🪟 Windows: C:\ProgramData\Microsoft\Windows\Start Menu\Programs
	ApplicationDirs []string

	// UserApplicationsHome is a specific
	//
	// E.g.:
	// - 🐧 Linux: /home/username/.local/share/applications/
	// - 🍎 Mac OS X: None
	// - 🪟 Windows: None
	UserApplicationsHome string
}

// AmadlaPaths contains all the amadla root paths since all the amadla
// TODO: Might remove
type AmadlaPaths struct {
	/*Storage    string // e.g.: /home/user/.hery/
	Catalog    string // e.g.: /home/user/.hery/collection/
	Collection string // e.g.: /home/user/.hery/collection/amadla/
	Entities   string // e.g.: /home/user/.hery/collection/amadla/entity/
	Cache      string // e.g.: /home/user/.hery/collection/amadla/amadla.cache
	*/
}

// Application contains the specific paths and name/title of an application
type Application struct {
	Name            AppName
	PluginTypeNames PluginTypeNames
	Paths           *ApplicationPaths
}

// ApplicationPaths contains all the path
type ApplicationPaths struct {
	//
	// Section: Basic
	//

	// DataHome User-specific data storage for applications.
	//
	// Example:
	// - 🐧 Linux: /home/username/.local/share/{appName}
	// - 🍎 Mac OS X: /Users/username/Library/Application Support
	// - 🪟 Windows: C:\Users\Username\AppData\Local (%LOCALAPPDATA%)
	DataHome string

	// ConfigHome User-specific configuration files directory
	//
	// E.g.:
	// - 🐧 Linux: /home/username/.config/{appName}
	// - 🍎 Mac OS X: /Users/username/Library/Preferences
	// - 🪟 Windows: C:\Users\Username\AppData\Roaming (%APPDATA%)
	ConfigHome string

	// StateHome User-specific state files (runtime data, logs, etc.)
	// E.g.:
	// - 🐧 Linux: /home/username/.local/state/{appName}/
	// - 🍎 Mac OS X: /Users/username/Library/Application Support
	// - 🪟 Windows: C:\Users\Username\AppData\Local
	StateHome string

	// CacheHome User-specific cache storage (temporary files)
	//
	// E.g.:
	// - 🐧 Linux: /home/username/.cache/{appName}/
	// - 🍎 Mac OS X: /Users/username/Library/Caches
	// - 🪟 Windows: C:\Users\Username\AppData\Local\cache
	CacheHome string

	// RuntimeDir Temporary runtime files (e.g., sockets, PID files)
	//
	// E.g.:
	// - 🐧 Linux: /run/username/1000/{appName}/
	// - 🍎 Mac OS X: /var/folders/.../T (Temporary directory)
	// - 🪟 Windows: Usually not set (empty string)
	RuntimeDir string

	// BinFile Directory for user-installed executable binaries
	//
	// E.g.:
	// - 🐧 Linux: /home/username/.local/bin/{appName}
	// - 🍎 Mac OS X: /usr/local/bin, /Users/username/bin (if exists)
	// - 🪟 Windows: C:\Users\Username\AppData\Local\Microsoft\WindowsApps
	BinFile string

	// ConfigFile User-specific configuration files directory
	//
	// E.g.:
	// - 🐧 Linux: /home/username/.config/{appName}/{appName}.yml
	// - 🍎 Mac OS X: /Users/username/Library/Preferences
	// - 🪟 Windows: C:\Users\Username\AppData\Roaming (%APPDATA%)
	ConfigFile string

	// PluginsHome
	//
	// E.g.:
	// - 🐧 Linux: /home/username/.local/share/{appName}/{pluginTypeName}.d/
	// - 🍎 Mac OS X:
	// - 🪟 Windows:
	PluginsHome map[string]string

	// CacheFile is a SQLite3 file
	//
	// E.g.:
	// - 🐧 Linux: /home/username/.config/{appName}/{appName}.cache
	// - 🍎 Mac OS X:
	// - 🪟 Windows:
	CacheFile string

	// ApplicationsDesktopFile
	//
	// E.g.:
	// - 🐧 Linux: /home/username/.local/share/applications/{appName}.desktop
	// - 🍎 Mac OS X:
	// - 🪟 Windows:
	ApplicationsDesktopFile string

	//
	// Secrets
	//

	// Home contains the directories and the files for all the secrets the application might need
	// TODO: What about the tmp?
	// TODO: Is this the best place what about run directory for tmp secrets?
	// TODO: Is there other propositions for this?
	// E.g.:
	// - 🐧 Linux: /home/username/.config/{appName}/secrets/
	// - 🍎 Mac OS X:
	// - 🪟 Windows:
	SecretsHome string

	// MTLSHome to be able to connect to `doorman` you need mTLS certification that are assigned temporally by `doorman`
	//
	// Note:
	//
	// chmod 600 ~/.config/amadla/private/*.key ~/.config/amadla/mTLS/*.key
	// chmod 644 ~/.config/amadla/certs/*.crt ~/.config/amadla/mTLS/*.crt
	//
	// E.g.:
	// - 🐧 Linux: /home/user/.config/{appName}/secrets/mTLS/
	// - 🍎 Mac OS X:
	// - 🪟 Windows:
	MTLSHome string

	// ln -s ~/.local/lib/amadla/amadla ~/.local/bin/amadla

}
