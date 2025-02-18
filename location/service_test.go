package location

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os/user"
	"testing"
)

func TestNewLocationService(t *testing.T) {
	t.Run("should return a new instance of Location", func(t *testing.T) {
		service, err := NewLocationService("mockname", "1.0.0", []string{"nix"})
		assert.NoError(t, err)
		assert.NotNil(t, service)
		assert.IsType(t, &SLocation{}, service)

		currentUser, err := user.Current()
		if err != nil {
			t.Fatalf("Error getting current user: %v", err)
		}

		homeDir := fmt.Sprintf("/home/%s", currentUser.Username)

		assert.Equal(t, homeDir+"/.local/share/mockname", service.ThisAppPaths().DataHome)
		assert.Equal(t, homeDir+"/.config/mockname", service.ThisAppPaths().ConfigHome)
		assert.Equal(t, homeDir+"/.local/state/mockname", service.ThisAppPaths().StateHome)
		assert.Equal(t, homeDir+"/.cache/mockname", service.ThisAppPaths().CacheHome)
		assert.Equal(t, "/run/user/1000/mockname", service.ThisAppPaths().RuntimeDir)
		assert.Equal(t, homeDir+"/.local/bin/mockname", service.ThisAppPaths().BinFile)
		assert.Equal(t, homeDir+"/.config/mockname/mockname.yaml", service.ThisAppPaths().ConfigFile)
		assert.Equal(t, homeDir+"/.local/share/nix.d", service.ThisAppPaths().PluginsHome["nix"])
		assert.Equal(t, homeDir+"/.cache/mockname/mockname.cache", service.ThisAppPaths().CacheFile)
		assert.Equal(t,
			homeDir+"/.local/share/applications/mockname.desktop",
			service.ThisAppPaths().ApplicationsDesktopFile)
		assert.Equal(t, "/run/user/1000/mockname/secrets", service.ThisAppPaths().SecretsHome)
		assert.Equal(t, "/run/user/1000/mockname/secrets/mTLS", service.ThisAppPaths().MTLSHome)
	})
}
