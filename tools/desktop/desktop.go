package desktop

import (
	"fmt"
	"path/filepath"
)

func New(appName, version, comment, binHome, dataHome string, mimeType *string) string {
	execPath := filepath.Join(binHome, appName)

	// TODO: Maybe check if svg or png is supported
	iconPath := filepath.Join(dataHome, "icon.svg")

	desktop := `[Desktop Entry]
	Name=%s
	Comment=%s
	Exec=%s
	Version=%s
	Icon=%s
	Terminal=true
	Type=Application
	Categories=Development;Utility;`

	var buildDesktop = fmt.Sprintf(desktop, appName, comment, execPath, version, iconPath)
	if mimeType != nil {
		mimeTypeEntry := fmt.Sprintf("\nMimeType=%s;", *mimeType)
		buildDesktop = fmt.Appendln(buildDesktop, mimeTypeEntry)
	}

	return buildDesktop
}
