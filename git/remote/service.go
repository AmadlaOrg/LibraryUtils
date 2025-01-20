package remote

// NewGitRemoteService to set up the Git Remote service
func NewGitRemoteService(url string) *SRemote {
	return &SRemote{
		url: url,
	}
}
