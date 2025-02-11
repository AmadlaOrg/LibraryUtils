package location

import (
	"github.com/adrg/xdg"
)

// Paths contains all the paths
type Paths struct {
	SystemPaths          *SystemPaths
	ThisApplicationPaths *Application
	Applications         *[]Application
}

// SystemPaths contains the system directories paths
type SystemPaths struct {
	// Home contains the path of the user's home directory.
	Home string

	// DataHome defines the base directory relative to which user-specific
	// data files should be stored. This directory is defined by the
	// $XDG_DATA_HOME environment variable. If the variable is not set,
	// a default equal to $HOME/.local/share should be used.
	DataHome string

	// DataDirs defines the preference-ordered set of base directories to
	// search for data files in addition to the DataHome base directory.
	// This set of directories is defined by the $XDG_DATA_DIRS environment
	// variable. If the variable is not set, the default directories
	// to be used are /usr/local/share and /usr/share, in that order. The
	// DataHome directory is considered more important than any of the
	// directories defined by DataDirs. Therefore, user data files should be
	// written relative to the DataHome directory, if possible.
	DataDirs []string

	// ConfigHome defines the base directory relative to which user-specific
	// configuration files should be written. This directory is defined by
	// the $XDG_CONFIG_HOME environment variable. If the variable is
	// not set, a default equal to $HOME/.config should be used.
	ConfigHome string

	// ConfigDirs defines the preference-ordered set of base directories to
	// search for configuration files in addition to the ConfigHome base
	// directory. This set of directories is defined by the $XDG_CONFIG_DIRS
	// environment variable. If the variable is not set, a default equal
	// to /etc/xdg should be used. The ConfigHome directory is considered
	// more important than any of the directories defined by ConfigDirs.
	// Therefore, user config files should be written relative to the
	// ConfigHome directory, if possible.
	ConfigDirs []string

	// StateHome defines the base directory relative to which user-specific
	// state files should be stored. This directory is defined by the
	// $XDG_STATE_HOME environment variable. If the variable is not set,
	// a default equal to ~/.local/state should be used.
	StateHome string

	// CacheHome defines the base directory relative to which user-specific
	// non-essential (cached) data should be written. This directory is
	// defined by the $XDG_CACHE_HOME environment variable. If the variable
	// is not set, a default equal to $HOME/.cache should be used.
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
	RuntimeDir string

	// BinHome defines the base directory relative to which user-specific
	// binary files should be written. This directory is defined by
	// the non-standard $XDG_BIN_HOME environment variable. If the variable is
	// not set, a default equal to $HOME/.local/bin should be used.
	BinHome string

	// UserDirs defines the locations of well known user directories.
	UserDirs xdg.UserDirectories

	// FontDirs defines the common locations where font files are stored.
	FontDirs []string

	// ApplicationDirs defines the common locations of applications.
	ApplicationDirs []string
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
	Name  string
	Paths *ApplicationPaths
}

// ApplicationPaths contains all the path
type ApplicationPaths struct {

	//
	// Section: Basic
	//

	// E.g.: /home/user/.local/share/{appName}/
	DataHome string

	// E.g.: /home/user/.config/{appName}/
	ConfigHome string

	// E.g.: /home/user/.local/state/{appName}/
	StateHome string

	// E.g.: /home/user/.cache/{appName}/
	CacheHome string

	// E.g.: /run/user/1000/{appName}/
	RuntimeDir string

	// E.g.: /home/user/.local/bin/{appName}
	BinHome string

	// E.g.: /home/user/.config/{appName}/{appName}.yml
	ConfigFile string

	// TODO:
	// E.g.: /home/user/.config/{appName}/plugins.d/
	PluginsHome string

	// CacheFile is a SQLite3 file
	// E.g.: /home/user/.config/{appName}/{appName}.cache
	CacheFile string

	// E.g.: /home/user/.local/share/applications/
	ApplicationsHome string

	// E.g.: /home/user/.local/share/applications/{appName}.desktop
	ApplicationsDesktopFile string

	//
	// Section: Secrets
	//

	// SecretsHome contains the directories and the files for all the secrets the application might need
	// TODO: What about the tmp?
	// TODO: Is this the best place what about run directory for tmp secrets?
	// TODO: Is there other propositions for this?
	// E.g.: /home/user/.config/{appName}/secrets/
	SecretsHome string

	/*
		chmod 600 ~/.config/amadla/private/*.key ~/.config/amadla/mTLS/*.key
		chmod 644 ~/.config/amadla/certs/*.crt ~/.config/amadla/mTLS/*.crt
	*/

	// MTLSHome to be able to connect to `doorman` you need mTLS certification that are assigned temporally by `doorman`
	// E.g.: /home/user/.config/{appName}/secrets/mTLS/
	MTLSHome string

	//
	// Temporary
	//

	// ln -s ~/.local/lib/amadla/amadla ~/.local/bin/amadla

}
