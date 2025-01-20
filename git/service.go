package git

import "github.com/AmadlaOrg/LibraryUtils/git/remote"

// NewGitService to set up the git service
func NewGitService(url, repositoryPath string) *SGit {
	return &SGit{
		url:              url,
		repositoryPath:   repositoryPath,
		serviceGitRemote: remote.NewGitRemoteService(url),
	}
}
