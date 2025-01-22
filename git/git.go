package git

import (
	"fmt"
	"github.com/AmadlaOrg/LibraryUtils/git/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// IGit to help with mocking
type IGit interface {
	Clone() error
	CommitHeadHash() (string, error)
	CheckoutTag(tagName string) error
}

type SGit struct {
	url            string
	repositoryPath string
	config         *config.Config
}

var (
	gitPlainOpen  = git.PlainOpen
	gitPlainClone = git.PlainClone
)

// Clone clones the repository from the given URL to the specified destination
func (s *SGit) Clone() error {
	_, err := gitPlainClone(s.repositoryPath, false, &git.CloneOptions{
		URL:               s.url,
		RemoteName:        s.config.RemoteName, // TODO: How to handle default, maybe add origin as the default
		ShallowSubmodules: s.config.CloneOptions.ShallowSubmodules,
		Depth:             s.config.CloneOptions.Depth,
		Auth:              s.config.Auth,
		RecurseSubmodules: git.SubmoduleRescursivity(s.config.CloneOptions.RecurseSubmodules),
		Progress:          s.config.CloneOptions.Progress,
		InsecureSkipTLS:   *s.config.InsecureSkipTLS,
		CABundle:          s.config.CABundle,
		ProxyOptions:      s.config.ProxyOptions,
	})
	return err
}

// CommitHeadHash retrieves the hash of the most recent commit
func (s *SGit) CommitHeadHash() (string, error) {
	repo, err := gitPlainOpen(s.repositoryPath)
	if err != nil {
		return "", err
	}

	// Get the HEAD reference
	ref, err := repo.Head()
	if err != nil {
		return "", err
	}

	// Get the commit object
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return "", err
	}

	return commit.Hash.String(), nil
}

// CheckoutTag checks out the specified branch or tag in the repository.
func (s *SGit) CheckoutTag(tagName string) error {
	repo, err := gitPlainOpen(s.repositoryPath)
	if err != nil {
		return err
	}

	// Get the working tree
	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	// Attempt to check out the reference as a branch
	err = worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(fmt.Sprintf("refs/tags/%s", tagName)), //plumbing.NewBranchReferenceName(refName),
		Force:  true,
	})
	if err != nil {
		return err
	}

	return nil
}
