// Package document is for generating in multiple formats of documents that can be used in different contexts.
package document

import "github.com/spf13/cobra"

// New populates the data required for generating documents.
//
// Params:
// - 📇 appName - Is the of the application (normally all lowercase).
// - 📜 appTitle - Is the name but with uppercase letters where need be.
// - ♻️ appVersion - The version of the application.
// - 🐍 rootCmd - Cobra command that contains all the application commands and flags.
func New(appName, appTitle, appVersion string, rootCmd *cobra.Command) Document {
	return &documentImpl{
		appName:    appName,
		appTitle:   appTitle,
		appVersion: appVersion,
		rootCmd:    rootCmd,
	}
}
