package location

import "testing"

func TestPaths(t *testing.T) {
	locationService := NewLocationService()
	locationService.Paths()
}
