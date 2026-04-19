package git

import (
	"github.com/AmadlaOrg/LibraryUtils/git/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("should return a new instance of Git", func(t *testing.T) {
		gitService := New("git.local/repo", "/home/user/repos", &config.Config{})
		assert.NotNil(t, gitService)
		assert.IsType(t, &gitImpl{}, gitService)
	})
}
