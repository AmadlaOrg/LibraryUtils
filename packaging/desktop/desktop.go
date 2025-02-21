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

// Builder 🏛️ Is the main structure for the NewDesktopService.
type Builder struct {
	// 📇 appName - Is the of the application (normally all lowercase).
	appName string

	builder strings.Builder
}

// For mocking 🥸.
var (
	osMkdirAll  = os.MkdirAll
	osWriteFile = os.WriteFile
)

// NewBuilder creates a new instance of DesktopBuilder.
//
// Params:
// - 📇 appName - Is the of the application (normally all lowercase).
func NewBuilder(appName string) Generator {
	return &Builder{
		appName: appName,
	}
}

// Generate generates the content for the `.desktop` file format by processing each group in the provided `Desktop` object.
// It builds sections and properties based on the `.desktop` specification.
//
// Params:
// - 🖥️ desktop: A pointer to the Desktop struct containing groups and their properties.
//
// Returns:
// - IDesktop: The constructed desktop representation.
// - 🚨 error: Returns an error if any required property is missing or incorrectly formatted.
func (b *Builder) Generate(desktop *Desktop) (Generator, error) {
	var (
		groupNames []string
	)

	groups := *desktop.Groups

	for _, group := range groups {
		var thisGroupName string

		if group.Title == nil ||
			*group.Title == "" ||
			*group.Title == defaultDesktopSectionName {
			thisGroupName = fmt.Sprintf("[%s]\n", defaultDesktopSectionName)
		} else {
			thisGroupName = fmt.Sprintf("[%s]\n", *group.Title)
		}

		b.builder.WriteString(thisGroupName)
		groupNames = append(groupNames, thisGroupName)

		err := b.processContent("Name", group.Names, true)
		if err != nil {
			return nil, err
		}
		_ = b.processContent("GenericName", group.GenericNames, false)
		_ = b.processContent("Comment", group.Comments, false)
		_ = b.processContent("Keywords", group.Keywords, false)
		b.processPropertyDefault("Version", group.Version, "1.0")
		err = b.processRequiredProperty("X-AppVersion", group.XAppVersion)
		if err != nil {
			return nil, err
		}
		b.processNotRequiredProperty("Icon", group.Icon)
		_ = b.processList("Categories", group.Categories, false)
		_ = b.processCommaList("X-KDE-Protocols", group.XKDEProtocols, false)
		b.processNotRequiredProperty("Encoding", group.Encoding)

		b.processPropertyDefault(
			"Terminal",
			pointer.ToPtr(strconv.FormatBool(group.Terminal)),
			"true")

		if group.Type == nil || *group.Type == "" {
			group.Type = (*Type)(pointer.ToPtr("Application"))
		}

		b.builder.WriteString(fmt.Sprintf("Type=%s\n", string(*group.Type)))
		if *group.Type == ApplicationType {
			//
			// Exec
			//
			err = b.processRequiredProperty("Exec", &group.Exec)
			if err != nil {
				return nil, err
			}
		} else if *group.Type == LinkType {
			//
			// Url
			//
			err = b.processRequiredProperty("Url", &group.URL)
			if err != nil {
				return nil, err
			}
		}

		_ = b.processList("MimeType", group.MimeType, false)
	}

	return b, nil
}

// WriteToFile writes the generated `.desktop` content to a file in the specified directory.
// It ensures the directory exists before writing.
//
// Params:
// - 📁 dirPath: The directory where the `.desktop` file should be saved.
//
// Returns:
// - 🚨 error: Returns an error if the directory cannot be created or if writing to the file fails.
func (b *Builder) WriteToFile(dirPath string) error {
	err := osMkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Construct the file path
	filePath := filepath.Join(dirPath, fmt.Sprintf("%s.desktop", b.appName))

	// Get the .desktop content
	desktopContent := b.builder.String()

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
func (b *Builder) processContent(propertyName string, contentValues *[]Content, isNeeded bool) error {
	var (
		isSet         = false
		isSetWithLang = false
	)

	for _, value := range *contentValues {
		if value.Language == nil {
			isSet = true
			b.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, value.Value))
		} else {
			isSetWithLang = true
			b.builder.WriteString(fmt.Sprintf("%s[%s]=%s\n", propertyName, *value.Language, value.Value))
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
func (b *Builder) processPropertyDefault(propertyName string, contentValue *string, defaultValue string) {
	if contentValue == nil || *contentValue == "" {
		b.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, defaultValue))
	} else {
		b.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, *contentValue))
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
func (b *Builder) processRequiredProperty(propertyName string, contentValue *string) error {
	if contentValue != nil && *contentValue != "" {
		b.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, *contentValue))
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
func (b *Builder) processNotRequiredProperty(propertyName string, contentValue *string) {
	if contentValue != nil && *contentValue != "" {
		b.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, *contentValue))
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
func (b *Builder) processList(propertyName string, items *[]List, isNeeded bool) error {
	if isNeeded && (items == nil || len(*items) == 0) {
		return fmt.Errorf("the property %s is empty", propertyName)
	}

	if items != nil && len(*items) > 0 {
		processedString := fmt.Sprintf(
			"%s=%s\n", propertyName, strings.Join(convertListToStrings(*items), ";"))
		b.builder.WriteString(processedString)
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
func (b *Builder) processCommaList(propertyName string, items *[]CommaList, isNeeded bool) error {
	if isNeeded && (items == nil || len(*items) == 0) {
		return fmt.Errorf("the property %s is empty", propertyName)
	}

	if items != nil && len(*items) > 0 {
		processedString := fmt.Sprintf(
			"%s=%s\n", propertyName, strings.Join(convertListToStrings(*items), ","))
		b.builder.WriteString(processedString)
	}

	return nil
}
