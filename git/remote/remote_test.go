package remote

import (
	"errors"
	utilGitConfig "github.com/AmadlaOrg/LibraryUtils/git/config"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestTags(t *testing.T) {
	tests := []struct {
		name                 string
		internalGitNewRemote func(s storage.Storer, c *config.RemoteConfig) IGoGitRemote
		expectedError        error
		hasError             bool
	}{
		{
			name: "Success",
			internalGitNewRemote: func(s storage.Storer, c *config.RemoteConfig) IGoGitRemote {
				mockGoGitRemote := NewMockGoGitRemote(t)
				mockGoGitRemote.EXPECT().List(mock.Anything).Return([]*plumbing.Reference{
					{},
				}, nil)
				return mockGoGitRemote
			},
			hasError: false,
		},
		//
		// Error
		//
		{
			name: "Error: fails at r.List",
			internalGitNewRemote: func(s storage.Storer, c *config.RemoteConfig) IGoGitRemote {
				mockGoGitRemote := NewMockGoGitRemote(t)
				mockGoGitRemote.EXPECT().List(mock.Anything).Return(nil, errors.New("some error (r.List)"))
				return mockGoGitRemote
			},
			expectedError: errors.New("some error (r.List)"),
			hasError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalGitNewRemote := gitNewRemote
			defer func() { gitNewRemote = originalGitNewRemote }()
			gitNewRemote = tt.internalGitNewRemote

			gitRemoteService := NewGitRemoteService("mock_repo_url", &utilGitConfig.Config{})
			tags, err := gitRemoteService.Tags()
			if tt.hasError {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, 0, len(tags))
		})
	}
}

func TestCommitHeadHash(t *testing.T) {
	tests := []struct {
		name                 string
		internalGitNewRemote func(s storage.Storer, c *config.RemoteConfig) IGoGitRemote
		expectedError        error
		hasError             bool
	}{
		//
		// Error
		//
		{
			name: "Error: fails at r.List",
			internalGitNewRemote: func(s storage.Storer, c *config.RemoteConfig) IGoGitRemote {
				mockGoGitRemote := NewMockGoGitRemote(t)
				mockGoGitRemote.EXPECT().List(mock.Anything).Return(nil, errors.New("some error (r.List)"))
				return mockGoGitRemote
			},
			expectedError: errors.New("some error (r.List)"),
			hasError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalGitNewRemote := gitNewRemote
			defer func() { gitNewRemote = originalGitNewRemote }()
			gitNewRemote = tt.internalGitNewRemote

			gitRemoteService := NewGitRemoteService("mock_repo_url", &utilGitConfig.Config{})
			tags, err := gitRemoteService.CommitHeadHash()
			if tt.hasError {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, 0, len(tags))
		})
	}
}
