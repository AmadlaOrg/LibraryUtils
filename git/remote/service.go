package remote

import "github.com/AmadlaOrg/LibraryUtils/git/config"

// New to set up the Git Remote service
func New(url string, cnf *config.Config) Remote {
	return &remoteImpl{
		url:    url,
		config: config.BuildDefaultConfig(cnf),
	}
}
