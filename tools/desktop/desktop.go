// Package desktop is for generating a `.desktop` content and file
package desktop

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/AmadlaOrg/LibraryUtils/pointer"
)

// IDesktop 🧩 Is the interface for the NewDesktopService.
type IDesktop interface {
	Build(desktop *Desktop) (IDesktop, error)
	Save(dirPath string) error
}

// SDesktop 🏛️ Is the main structure for the NewDesktopService.
type SDesktop struct {
	// 📇 appName - Is the of the application (normally all lowercase).
	appName string

	builder strings.Builder
}

// For mocking 🥸.
var (
	osMkdirAll  = os.MkdirAll
	osWriteFile = os.WriteFile
)

// Build generates the content for the `.desktop` file format
//
// Params:
// - 🖥️ desktop:
// - 📍 location:
func (service *SDesktop) Build(desktop *Desktop) (IDesktop, error) {
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
			thisSectionName = fmt.Sprintf("[%s]\n", defaultDesktopSectionName)
		} else {
			thisSectionName = fmt.Sprintf("[%s]\n", *section.Title)
		}

		service.builder.WriteString(thisSectionName)
		sectionNames = append(sectionNames, thisSectionName)

		//
		// Name
		//
		err := service.processContent("Name", section.Names, true)
		if err != nil {
			return nil, err
		}

		//
		// GenericNames
		//
		_ = service.processContent("GenericName", section.GenericNames, false)

		//
		// Comments
		//
		_ = service.processContent("Comment", section.Comments, false)

		//
		// Keywords
		//
		_ = service.processContent("Keywords", section.Keywords, false)

		//
		// Version
		//
		service.processPropertyDefault("Version", section.Version, "1.0")

		//
		// X-AppVersion
		//
		err = service.processRequiredProperty("X-AppVersion", section.XAppVersion)
		if err != nil {
			return nil, err
		}

		//
		// Icon
		//
		service.processNotRequiredProperty("Icon", section.Icon)

		//
		// Categories
		//
		_ = service.processList("Categories", section.Categories, false)

		//
		// x-kde-protocols
		//
		_ = service.processCommaList("X-KDE-Protocols", section.XKDEProtocols, false)

		//
		// Encoding
		//
		service.processNotRequiredProperty("Encoding", section.Encoding)

		//
		// Terminal
		//
		service.processPropertyDefault(
			"Terminal",
			pointer.ToPtr(strconv.FormatBool(section.Terminal)),
			"true")

		//
		// Type
		//
		if section.Type == nil || *section.Type == "" {
			section.Type = (*Type)(pointer.ToPtr("Application"))
		}

		service.builder.WriteString(fmt.Sprintf("Type=%s\n", string(*section.Type)))
		if *section.Type == ApplicationType {
			//
			// Exec
			//
			err = service.processRequiredProperty("Exec", &section.Exec)
			if err != nil {
				return nil, err
			}
		} else if *section.Type == LinkType {
			//
			// Url
			//
			err = service.processRequiredProperty("Url", &section.URL)
			if err != nil {
				return nil, err
			}
		} else if *section.Type == DirectoryType {
			// TODO:
		}

		//
		// MimeType
		//
		_ = service.processList("MimeType", section.MimeType, false)
	}

	return service, nil
}

// Save to a `.desktop` file.
//
// Params:
// - dirPath:
func (service *SDesktop) Save(dirPath string) error {
	err := osMkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Construct the file path
	filePath := filepath.Join(dirPath, fmt.Sprintf("%s.desktop", service.appName))

	// Get the .desktop content
	desktopContent := service.builder.String()

	// Write to the file
	err = osWriteFile(filePath, []byte(desktopContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// processContent processes a slice of Content values and appends them to the builder,
// formatting them according to the presence of a language.
//
// Params:
// - 🏠 propertyName: The name of the property being processed.
// - 💎 contentValues: A pointer to a slice of Content structs.
// - 🙏 need: A flag indicating whether the property must be set.
//
// Returns:
// - 🚨 error: Returns an error if isNeeded is true and the property is not set properly.
func (service *SDesktop) processContent(propertyName string, contentValues *[]Content, isNeeded bool) error {
	var (
		isSet         = false
		isSetWithLang = false
	)

	for _, value := range *contentValues {
		if value.Language == nil {
			isSet = true
			service.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, value.Value))
		} else {
			isSetWithLang = true
			service.builder.WriteString(fmt.Sprintf("%s[%s]=%s\n", propertyName, *value.Language, value.Value))
		}
	}

	if isNeeded {
		if !isSet && !isSetWithLang {
			return fmt.Errorf("%s not set", propertyName)
		} else if !isSet {
			return fmt.Errorf(
				"%s is set with specific language but not set without specific language",
				propertyName,
			)
		}
	}

	return nil
}

// processPropertyDefault processes a single property and assigns it either the provided value
// or a default value if the provided value is nil or empty.
//
// Params:
// - 🏠 propertyName: The name of the property being processed.
// - 💎 contentValue: A pointer to the string value of the property (maybe nil).
// - ⚠️ defaultValue: The default value to use if contentValue is nil or empty.
func (service *SDesktop) processPropertyDefault(propertyName string, contentValue *string, defaultValue string) {
	if contentValue == nil || *contentValue == "" {
		service.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, defaultValue))
	} else {
		service.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, *contentValue))
	}
}

// processRequiredProperty processes a required property, ensuring that it is set.
// If the property is missing or empty, an error is returned.
//
// Params:
// - 🏠 propertyName: The name of the property being processed.
// - 💎 contentValue: A pointer to the string value of the property (maybe nil).
//
// Returns:
// - 🚨 error: Returns an error if contentValue is nil or empty.
func (service *SDesktop) processRequiredProperty(propertyName string, contentValue *string) error {
	if contentValue != nil && *contentValue != "" {
		service.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, *contentValue))
		return nil
	} else {
		return fmt.Errorf("the property %s is required", propertyName)
	}
}

// processNotRequiredProperty processes a property that is not required.
// If the property has a value, it is written to the builder.
// If the property is nil or empty, nothing is written.
//
// Params:
// - 🏠 propertyName: The name of the property being processed.
// - 💎 contentValue: A pointer to the string value of the property (maybe nil).
func (service *SDesktop) processNotRequiredProperty(propertyName string, contentValue *string) {
	if contentValue != nil && *contentValue != "" {
		service.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, *contentValue))
	}
}

// processList processes a list of items and appends them to the builder in a semicolon-separated format.
// If the list is required (`isNeeded` is true) and empty, an error is returned.
//
// Params:
// - 🏠 propertyName: The name of the property being processed.
// - 👑 items: A pointer to a slice of List items (can be nil or empty).
// - 🙏 need: A flag indicating whether the property must be set.
//
// Returns:
// - 🚨 error: Returns an error if `isNeeded` is true and the list is empty.
func (service *SDesktop) processList(propertyName string, items *[]List, isNeeded bool) error {
	if isNeeded && (items == nil || len(*items) == 0) {
		return fmt.Errorf("the property %s is empty", propertyName)
	}

	if items != nil && len(*items) > 0 {
		// Convert []List to []string
		stringItems := make([]string, len(*items))
		for i, item := range *items {
			stringItems[i] = string(item) // Convert List to string
		}

		service.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, strings.Join(stringItems, ";")))
	}

	return nil
}

// processCommaList processes a list of items and appends them to the builder in a comma-separated format.
// If the list is required (`isNeeded` is true) and empty, an error is returned.
//
// Params:
// - 🏠 propertyName: The name of the property being processed.
// - 👑 items: A pointer to a slice of CommaList items (can be nil or empty).
// - 🙏 need: A flag indicating whether the property must be set.
//
// Returns:
// - 🚨 error: Returns an error if `isNeeded` is true and the list is empty.
func (service *SDesktop) processCommaList(propertyName string, items *[]CommaList, isNeeded bool) error {
	if isNeeded && (items == nil || len(*items) == 0) {
		return fmt.Errorf("the property %s is empty", propertyName)
	}

	if items != nil && len(*items) > 0 {
		// Convert []List to []string
		stringItems := make([]string, len(*items))
		for i, item := range *items {
			stringItems[i] = string(item) // Convert List to string
		}

		service.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, strings.Join(stringItems, ",")))
	}

	return nil
}
