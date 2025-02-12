package desktop

// Desktop contains a list of sections that are in `.desktop`
type Desktop struct {
	Sections *[]Section
}

// Section in the `.desktop` file
type Section struct {
	// Title (default: [Desktop Entry] is there is more than one section,
	// and they don't have a title an error is thrown) is the name/title of the section
	//
	// E.g.: [Desktop Entry]
	Title *string `json:"title,omitempty"`

	// Names (default: throws error) is a list of the application/service name in different languages
	//
	// E.g.: Name= / Name[fr]=
	Names *[]Content `json:"names"`

	// GenericNames (default: none) is similar to Name but is more generic
	//
	// E.g.:
	// Name=VLC Media Player
	// GenericName=Media Player
	GenericNames *[]Content `json:"genericNames,omitempty"`

	// Comments (default: throws error) is a list of comment (similar to a description) in different languages
	//
	// E.g.: Comment= / Comment[fr]=
	Comments *[]Content `json:"comment"`

	// Keywords (default: none) is a list of Keywords sets in different languages
	//
	// E.g.: Development;Utility;
	Keywords *[]Content `json:"keywords,omitempty"`

	// Version (default: 1.0) is the desktop entry specification version
	// This field is used to indicate which version of the desktop entry specification the file adheres to.\
	// Currently, the recommended value is `1.0`, as defined by the `freedesktop.org` specification.
	//
	// E.g.: 1.0
	Version *string `json:"version,omitempty"`

	// XAppVersion (`X-AppVersion`) (default: throws error) is the actual version of the application
	//
	// E.g.: X-AppVersion=1.0.0
	XAppVersion *string `json:"x-app-version"`

	// Icon (default: none) is the absolute file path to the icon of the application
	//
	// E.g.: /home/user/.local/share/appName/icon.svg
	Icon *string `json:"icon,omitempty"`

	// Categories (default: Development;Utility;) is an optional list of categories that represent the application
	// This is for internal purposes so there is no need for multi-language support
	//
	// E.g.: Categories=Development;Utility;
	Categories *[]List `json:"categories"`

	// XKDEProtocols (default: none) is a KDE-specific key that lists the supported URL protocols for an application
	// It is primarily used for applications that handle specific network or file protocols,
	// such as http, ftp, smb, or mailto.
	//
	// For the support of other GUIs:
	// MimeType=x-scheme-handler/http;x-scheme-handler/https;x-scheme-handler/mailto;
	//
	// E.g.: X-KDE-Protocols=ftp,http,https,mms,rtmp
	XKDEProtocols *[]CommaList `json:"x-kde-protocols"`

	// Encoding (default: none) is the text encoding
	//
	// E.g.:
	// Encoding=UTF-8
	Encoding *string `json:"encoding,omitempty"`

	// Exec (default: /home/user/.local/bin/{appName}) contains the absolute path the application
	//
	// E.g.: Exec=/home/user/.local/bin/appName
	//
	// Example in Kali Linux:
	// /usr/share/kali-menu/exec-in-shell "tcpdump -h"
	Exec string `json:"exec"`

	// Terminal (default: true) is a boolean key that determines whether the application should be launched in a
	// terminal emulator
	Terminal bool `json:"terminal"`

	// Type (default: Application) specifies the type of the entry, defining how it should be handled by the system
	//
	// Supported types:
	// - Application: Represents a program or script. Requires `Exec=` to specify the command.
	// - Link: Represents a web link (URL). Requires `URL=` to specify the destination.
	// - Directory: Represents a virtual folder (used in application menus, not for real directories).
	Type *Type `json:"type,omitempty"`

	// URL (default: none unless Type=Link then it throws an error) specifies the web address that should be opened
	// when the shortcut is clicked
	URL string `json:"url"`

	// MimeType (default: none)
	MimeType *[]List `json:"mime-type"`
}

// Content for content that can be set in multiple languages
type Content struct {
	// Language is the language code ISO 639-1
	// Default is nil, if nil then it won't have the bracket with the language code (it is the `.desktop` default)
	Language *string `json:"language"`

	// Value is the value
	Value string `json:"value"`
}

// List is for standard value list that are seperated by `;`
type List string

// CommaList is for none-standard value list e.g.: X-KDE-Protocols
type CommaList string

type Type string

const (
	// ApplicationType represents a program or script. Requires `Exec=` to specify the command
	ApplicationType Type = "Application"

	// LinkType represents a web link (URL). Requires `URL=` to specify the destination
	LinkType Type = "Link"

	// DirectoryType represents a virtual folder (used in application menus, not for real directories)
	DirectoryType Type = "Directory"
)
