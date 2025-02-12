package desktop

import (
	"errors"
	"fmt"
	"strings"

	"github.com/AmadlaOrg/LibraryUtils/location"
)

type IDesktop interface{}
type SDesktop struct {
	builder strings.Builder
}

const (
	defaultDesktopSectionName = "Desktop Entry"
)

// Build generates the content for the `.desktop` file format
//
// Params:
// - 🖥️ desktop:
// - 📍 location:
func (service *SDesktop) Build(desktop *Desktop, location *location.Application) (string, error) {
	var (
		sectionNames []string
	)

	sections := *desktop.Sections

	for _, section := range sections {
		//
		// Section
		//
		var thisSectionName string

		if section.Title == nil ||
			*section.Title == "" ||
			*section.Title == defaultDesktopSectionName {
			thisSectionName = fmt.Sprintf("[%s]", defaultDesktopSectionName)
		} else {
			thisSectionName = fmt.Sprintf("[%s]", *section.Title)
		}

		service.builder.WriteString(thisSectionName)
		sectionNames = append(sectionNames, thisSectionName)

		//
		// Name
		//
		if section.Names == nil {
			return "", errors.New("no section names")
		}

		service.processContent("Name", section.Names)

		//
		// GenericNames
		//
		service.processContent("GenericName", section.GenericNames)

		//
		// Comments
		//
		service.processContent("Comment", section.Comments)

		//
		// Keywords
		//
		service.processContent("Keyword", section.Keywords)

		//
		// Version
		//
		if section.Version == nil {
			service.builder.WriteString("Version=1.0")
		} else {
			service.builder.WriteString(fmt.Sprintf("Version=%s", *section.Version))
		}

		//
		// X-AppVersion
		//

	}

	// TODO: If Kali Linux: /usr/share/kali-menu/exec-in-shell "{appName} -h"

	/*
		cat /usr/share/kali-menu/exec-in-shell
		#!/usr/bin/env sh

		echo "$ $@"
		eval $@

		USER=${USER:-$(whoami)}
		SHELL=${SHELL:-$(getent passwd $USER | cut -d: -f7)}
		${SHELL:-bash} -i
	*/
	/*
		$ cat /usr/share/kali-menu/applications/kali-scapy.desktop

		[Desktop Entry]
		Name=scapy
		Comment=Interactive packet manipulation tool
		Encoding=UTF-8
		Exec=/usr/share/kali-menu/exec-in-shell "scapy"
		Icon=kali-scapy
		StartupNotify=false
		Terminal=true
		Type=Application
		Categories=09-sniffing-spoofing;
		X-Kali-Package=python3-scapy
	*/

	/*var
	b.WriteString("INSERT INTO ")
	b.WriteString(table.Name)
	b.WriteString(" (")
	b.WriteString(strings.Join(columnNames, ", "))
	b.WriteString(") VALUES (")
	b.WriteString(strings.Join(valuesPlaceholder, ", "))
	b.WriteString(");")
	b.String()

	execPath := filepath.Join(binHome, appName)

	// TODO: Maybe check if svg or png is supported
	iconPath := filepath.Join(dataHome, "icon.svg")

	d := `[Desktop Entry]
	Name=%s
	Name[fr]=%s
	GenericName=%s
	GenericName[fr]=%s
	Comment=%s
	Comment[fr]=%s
	Keywords=%s
	Keywords[fr]=%s
	Encoding=UTF-8
	Exec=%s
	Version=%s
	Icon=%s
	Terminal=true
	Type=Application
	Categories=Development;Utility;
	X-KDE-Protocols=%s`

	// TODO:
	// Version= ??

	// TODO:
	// X-KDE-Protocols=ftp,http,https,mms,rtmp

	// MimeType:
	// application/x-yaml
	// text/yaml
	// application/json
	//
	// Hery:
	// application/json+hery
	// application/x-yaml+hery
	// text/yaml_hery

	var buildDesktop = fmt.Sprintf(desktop, appName, comment, execPath, version, iconPath)
	if mimeType != nil {
		mimeTypeEntry := fmt.Sprintf("\nMimeType=%s;", *mimeType)
		buildDesktop = fmt.Appendln(buildDesktop, mimeTypeEntry)
	}*/

	return service.builder.String(), nil
}

// processContent
//
// Params:
// - 🏠 propertyName:
// - 💎 contentValues:
func (service *SDesktop) processContent(propertyName string, contentValues *[]Content) {
	for _, value := range *contentValues {
		if value.Language == nil {
			service.builder.WriteString(fmt.Sprintf("%s=%s", propertyName, value.Value))
		} else {
			service.builder.WriteString(fmt.Sprintf("%s[%s]=%s", propertyName, *value.Language, value.Value))
		}
	}
}
