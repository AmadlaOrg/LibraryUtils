package document

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// Document 🧩 Is the interface for the NewDocumentService.
type Document interface {
	All(outputPath string) error
	Man(outputPath string) error
	Markdown(outputPath string) error
}

// documentImpl 🏛️ Is the main structure for the NewDocumentService.
type documentImpl struct {
	// 📇 appName - Is the of the application (normally all lowercase).
	appName string

	// 📜 appTitle - Is the name but with uppercase letters where need be.
	appTitle string

	// ♻️ appVersion - The version of the application.
	appVersion string

	// 🐍 rootCmd - Cobra command that contains all the application commands and flags.
	rootCmd *cobra.Command
}

// For mocking 🥸.
var (
	docGenManTree      = doc.GenManTree
	docGenMarkdownTree = doc.GenMarkdownTree
)

// All document format are generated from commands and flags.
//
// Params:
// - 📁 outputPath: Is the absolute path to the directory where the document will be output.
func (service *documentImpl) All(outputPath string) error {
	manErr := service.Man(outputPath)
	mkErr := service.Markdown(outputPath)

	if manErr != nil && mkErr != nil {
		return errors.Join(manErr, mkErr)
	} else if manErr != nil {
		return manErr
	} else if mkErr != nil {
		return mkErr
	}

	return nil
}

// Man document generated from commands and flags.
//
// Params:
// - 📁 outputPath: Is the absolute path to the directory where the document will be output.
func (service *documentImpl) Man(outputPath string) error {
	header := &doc.GenManHeader{
		Title:   service.appTitle,
		Section: "1",
		Source:  fmt.Sprintf("%s v%s", service.appTitle, service.appVersion),
		Manual:  fmt.Sprintf("%s Manual", service.appTitle),
	}

	err := docGenManTree(service.rootCmd, header, outputPath)
	if err != nil {
		return err
	}

	return nil
}

// Markdown document generated from commands and flags.
//
// Params:
// - 📁 outputPath: Is the absolute path to the directory where the document will be output.
func (service *documentImpl) Markdown(outputPath string) error {
	err := docGenMarkdownTree(service.rootCmd, outputPath)
	if err != nil {
		return err
	}

	return nil
}
