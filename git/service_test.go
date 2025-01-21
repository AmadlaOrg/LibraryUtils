package git

import (
	"github.com/AmadlaOrg/LibraryUtils/git/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGitService(t *testing.T) {
	t.Run("should return a new instance of Git", func(t *testing.T) {
		gitService := NewGitService("git.local/repo", "/home/user/repos", &config.Config{})
		assert.NotNil(t, gitService)
		assert.IsType(t, &SGit{}, gitService)
	})
}
