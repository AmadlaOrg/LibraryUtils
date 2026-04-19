package location

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os/user"
	"runtime"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("should return a new instance of Location", func(t *testing.T) {
		service, err := New("mockname", "1.0.0", []string{"nix"})
		assert.NoError(t, err)
		assert.NotNil(t, service)
		assert.IsType(t, &locationImpl{}, service)

		if runtime.GOOS != "windows" &&
			runtime.GOOS != "darwin" &&
			runtime.GOOS != "plan9" {
			currentUser, err := user.Current()
			if err != nil {
				t.Fatalf("Error getting current user: %v", err)
			}

			homeDir := fmt.Sprintf("/home/%s", currentUser.Username)

			assert.Equal(t, homeDir+"/.local/share/mockname", service.ApplicationPaths().DataHome)
			assert.Equal(t, homeDir+"/.config/mockname", service.ApplicationPaths().ConfigHome)
			assert.Equal(t, homeDir+"/.local/state/mockname", service.ApplicationPaths().StateHome)
			assert.Equal(t, homeDir+"/.cache/mockname", service.ApplicationPaths().CacheHome)
			assert.Equal(t, "/run/user/1000/mockname", service.ApplicationPaths().RuntimeDir)
			assert.Equal(t, homeDir+"/.local/bin/mockname", service.ApplicationPaths().BinFile)
			assert.Equal(t, homeDir+"/.config/mockname/mockname.yaml", service.ApplicationPaths().ConfigFile)
			assert.Equal(t, homeDir+"/.local/share/nix.d", service.ApplicationPaths().PluginsHome["nix"])
			assert.Equal(t, homeDir+"/.cache/mockname/mockname.cache", service.ApplicationPaths().CacheFile)
			assert.Equal(t,
				homeDir+"/.local/share/applications/mockname.desktop",
				service.ApplicationPaths().ApplicationsDesktopFile)
			assert.Equal(t, "/run/user/1000/mockname/secrets", service.ApplicationPaths().SecretsHome)
			assert.Equal(t, "/run/user/1000/mockname/secrets/mTLS", service.ApplicationPaths().MTLSHome)
		}
	})
}
