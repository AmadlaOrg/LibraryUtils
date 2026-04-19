package remote

import (
	"github.com/AmadlaOrg/LibraryUtils/git/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("should return a new instance of GitRemote", func(t *testing.T) {
		gitRemoteService := New("git.local/repo", &config.Config{})
		assert.NotNil(t, gitRemoteService)
		assert.IsType(t, &remoteImpl{}, gitRemoteService)
	})
}
