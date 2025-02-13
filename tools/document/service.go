// Package document is for generating in multiple formats of documents that can be used in different contexts.
package document

import "github.com/spf13/cobra"

// NewDocumentService populates the data required for generating documents.
//
// Params:
// - 📇 appName - Is the of the application (normally all lowercase).
// - 📜 appTitle - Is the name but with uppercase letters where need be.
// - ♻️ appVersion - The version of the application.
// - 🐍 rootCmd - Cobra command that contains all the application commands and flags.
func NewDocumentService(appName, appTitle, appVersion string, rootCmd *cobra.Command) IDocument {
	return &SDocument{
		appName:    appName,
		appTitle:   appTitle,
		appVersion: appVersion,
		rootCmd:    rootCmd,
	}
}
