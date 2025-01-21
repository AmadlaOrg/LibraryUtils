package git

import (
	"github.com/AmadlaOrg/LibraryUtils/git/config"
)

// NewGitService to set up the git service
func NewGitService(url, repositoryPath string, cnf *config.Config) *SGit {
	return &SGit{
		url:            url,
		repositoryPath: repositoryPath,
		config:         config.BuildDefaultConfig(cnf),
	}
}
