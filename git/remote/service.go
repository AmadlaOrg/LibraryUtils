package remote

import "github.com/AmadlaOrg/LibraryUtils/git/config"

// NewGitRemoteService to set up the Git Remote service
func NewGitRemoteService(url string, cnf *config.Config) IRemote {
	return &SRemote{
		url:    url,
		config: config.BuildDefaultConfig(cnf),
	}
}
