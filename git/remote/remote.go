package remote

import (
	"errors"
	"fmt"
	utilGitConfig "github.com/AmadlaOrg/LibraryUtils/git/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage"
	"github.com/go-git/go-git/v5/storage/memory"
)

// Remote to help with mocking
type Remote interface {
	Tags() ([]string, error)
	CommitHeadHash() (string, error)
}

type remoteImpl struct {
	url    string
	config *utilGitConfig.Config
}

var (
	gitNewRemote = func(s storage.Storer, c *config.RemoteConfig) GoGitRemote {
		return git.NewRemote(s, c)
	}
)

// Tags returns a list of tags for the repository at the specified URL.
func (s *remoteImpl) Tags() ([]string, error) {
	r := gitNewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{s.url},
	})

	refs, err := r.List(&git.ListOptions{
		Auth:            s.config.Auth,
		InsecureSkipTLS: *s.config.InsecureSkipTLS,
		CABundle:        s.config.CABundle,
		ProxyOptions:    s.config.ProxyOptions,
		Timeout:         s.config.Timeout,
		PeelingOption:   git.IgnorePeeled,
	})
	if err != nil {
		return nil, err
	}

	var tags []string
	for _, ref := range refs {
		if ref.Name().IsTag() {
			tags = append(tags, ref.Name().Short())
		}
	}

	return tags, nil
}

// CommitHeadHash retrieves the hash of the most recent commit
func (s *remoteImpl) CommitHeadHash() (string, error) {
	r := gitNewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{s.url},
	})

	// List all references from the remote repository
	refs, err := r.List(&git.ListOptions{
		Auth:            s.config.Auth,
		InsecureSkipTLS: *s.config.InsecureSkipTLS,
		CABundle:        s.config.CABundle,
		ProxyOptions:    s.config.ProxyOptions,
		Timeout:         s.config.Timeout,
		PeelingOption:   git.IgnorePeeled,
	})
	if err != nil {
		return "", err
	} else if refs == nil || len(refs) == 0 {
		return "", errors.New("no tags found")
	}

	var headRef *plumbing.Reference

	// Find the HEAD reference
	for _, ref := range refs {
		if ref.Name() == plumbing.HEAD {
			headRef = ref
			break
		}
	}

	if headRef == nil {
		return "", fmt.Errorf("HEAD reference not found")
	}

	var commitHash plumbing.Hash

	// Resolve symbolic reference if HEAD is symbolic
	if headRef.Type() == plumbing.SymbolicReference {
		// Find the reference that HEAD points to
		for _, ref := range refs {
			if ref.Name() == headRef.Target() {
				commitHash = ref.Hash()
				break
			}
		}
	} else {
		commitHash = headRef.Hash()
	}

	if commitHash.IsZero() {
		return "", fmt.Errorf("commit hash not found")
	}

	return commitHash.String(), nil
}
