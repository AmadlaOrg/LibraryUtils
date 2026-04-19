package git

import (
	"github.com/AmadlaOrg/LibraryUtils/git/config"
)

// New to set up the git service
func New(url, repositoryPath string, cnf *config.Config) Git {
	return &gitImpl{
		url:            url,
		repositoryPath: repositoryPath,
		config:         config.BuildDefaultConfig(cnf),
	}
}
