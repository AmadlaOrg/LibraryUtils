package git

import (
	"errors"
	"github.com/AmadlaOrg/LibraryUtils/git/config"
	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckoutTag(t *testing.T) {
	tests := []struct {
		name                 string
		inputTagName         string
		internalGitPlainOpen func(path string) (GoGitRepository, error)
		expectedError        error
		hasError             bool
	}{
		{
			name:         "Error: git.PlainOpen fails",
			inputTagName: "v1.0.0",
			internalGitPlainOpen: func(path string) (GoGitRepository, error) {
				return &git.Repository{}, errors.New("some error (git.PlainOpen)")
			},
			expectedError: errors.New("some error (git.PlainOpen)"),
			hasError:      true,
		},
		{
			name:         "Error: repo.Worktree fails",
			inputTagName: "v1.0.0",
			internalGitPlainOpen: func(path string) (GoGitRepository, error) {
				mockGoGitRepository := NewMockGoGitRepository(t)
				mockGoGitRepository.EXPECT().Worktree().Return(nil, errors.New("some error (repo.Worktree)"))
				return mockGoGitRepository, nil
			},
			expectedError: errors.New("some error (repo.Worktree)"),
			hasError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalGitPlainOpen := gitPlainOpen
			defer func() { gitPlainOpen = originalGitPlainOpen }()
			gitPlainOpen = tt.internalGitPlainOpen

			gitService := New("mock_repo_url", "mock_repo_local_path", &config.Config{})
			err := gitService.CheckoutTag(tt.inputTagName)
			if tt.hasError {
				assert.Error(t, err)
				assert.EqualError(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
