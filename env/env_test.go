package env

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	// Create a temporary directory
	tmpDir := t.TempDir()

	// Create the env subdirectory within the temporary directory
	envDir := filepath.Join(tmpDir, "env")
	err := os.Mkdir(envDir, 0755)
	assert.NoError(t, err)

	// Create a temporary types.go file in the env subdirectory
	typesGoContent := `
	package env

	const (
		HeryStoragePath = "HERY_STORAGE_PATH"
		HeryCollection = "HERY_COLLECTION"
	)
	`
	typesGoPath := filepath.Join(envDir, "types.go")
	err = os.WriteFile(typesGoPath, []byte(typesGoContent), 0644)
	assert.NoError(t, err)

	// Override the filepath to point to the temporary directory
	originalDir, err := os.Getwd()
	assert.NoError(t, err)
	defer func() {
		err := os.Chdir(originalDir)
		if err != nil {
			t.Fatalf("Failed to change back to the original directory: %v", err)
		}
	}()
	err = os.Chdir(tmpDir)
	assert.NoError(t, err)

	// Call the List function and check the result
	expectedConstants := []string{"HERY_STORAGE_PATH", "HERY_COLLECTION"}
	envService := New(nil)
	actualConstants, err := envService.List()
	assert.NoError(t, err)
	assert.ElementsMatch(t, expectedConstants, actualConstants)

}
