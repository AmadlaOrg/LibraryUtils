package desktop

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AmadlaOrg/LibraryUtils/pointer"
)

type IDesktop interface {
	Build(desktop *Desktop) (string, error)
}

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
func (service *SDesktop) Build(desktop *Desktop) (string, error) {
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
			return "", err
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
			return "", err
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
				return "", err
			}
		} else if *section.Type == LinkType {
			//
			// Url
			//
			err = service.processRequiredProperty("Url", &section.URL)
			if err != nil {
				return "", err
			}
		} else if *section.Type == DirectoryType {
			// TODO:
		}

		//
		// MimeType
		//
		_ = service.processList("MimeType", section.MimeType, false)
	}

	return service.builder.String(), nil
}

// processContent
//
// Params:
// - 🏠 propertyName:
// - 💎 contentValues:
// - 🙏 need:
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

// processPropertyDefault
//
// Params:
// - 🏠 propertyName:
// - 💎 contentValue:
// - ⚠️ defaultValue:
func (service *SDesktop) processPropertyDefault(propertyName string, contentValue *string, defaultValue string) {
	if contentValue == nil || *contentValue == "" {
		service.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, defaultValue))
	} else {
		service.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, *contentValue))
	}
}

// processRequiredProperty
//
// Params:
// - 🏠 propertyName:
// - 💎 contentValue:
func (service *SDesktop) processRequiredProperty(propertyName string, contentValue *string) error {
	if contentValue != nil && *contentValue != "" {
		service.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, *contentValue))
		return nil
	} else {
		return fmt.Errorf("%s is required", propertyName)
	}
}

// processNotRequiredProperty
//
// Params:
// - 🏠 propertyName:
// - 💎 contentValue:
func (service *SDesktop) processNotRequiredProperty(propertyName string, contentValue *string) {
	if contentValue != nil && *contentValue != "" {
		service.builder.WriteString(fmt.Sprintf("%s=%s\n", propertyName, *contentValue))
	}
}

// processList
//
// Params:
// - 🏠 propertyName:
// - 👑 items:
// - 🙏 need:
func (service *SDesktop) processList(propertyName string, items *[]List, isNeeded bool) error {
	if isNeeded && (items == nil || len(*items) == 0) {
		return fmt.Errorf("%s is empty", propertyName)
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

// processCommaList
//
// Params:
// - 🏠 propertyName:
// - 👑 items:
// - 🙏 need:
func (service *SDesktop) processCommaList(propertyName string, items *[]CommaList, isNeeded bool) error {
	if isNeeded && (items == nil || len(*items) == 0) {
		return fmt.Errorf("%s is empty", propertyName)
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
