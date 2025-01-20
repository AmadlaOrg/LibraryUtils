package file

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func Test_integration_OsOpen(t *testing.T) {
	// Use a temporary file for testing
	tempFileName := "test_integration_file.txt"
	tempFilePath := filepath.Join(os.TempDir(), tempFileName)
	tempContent := []byte("This is a test file.")

	// Create the temporary file
	err := os.WriteFile(tempFilePath, tempContent, 0644)
	assert.NoError(t, err, "Failed to create temporary test file")

	// Ensure the file is deleted after the test
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Errorf("Failed to remove temporary file %s", name)
		}
	}(tempFilePath)

	t.Run("should open existing file successfully", func(t *testing.T) {
		// Call osOpen (real os.Open is being used here)
		file, err := osOpen(tempFilePath)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, file)

		// Close the file after testing
		if file != nil {
			err := file.Close()
			assert.NoError(t, err, "Failed to close the file")
		}
	})

	t.Run("should return error for non-existent file", func(t *testing.T) {
		nonExistentFile := "non_existent_file.txt"

		// Call osOpen with a file that doesn't exist
		file, err := osOpen(nonExistentFile)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, file)
	})
}

func Test_integration_Exists(t *testing.T) {
	tmpDir := t.TempDir()

	t.Run("file exists", func(t *testing.T) {
		filePath := filepath.Join(tmpDir, "file.txt")
		err := os.WriteFile(filePath, []byte("test content"), 0644)
		assert.NoError(t, err)

		exists := Exists(filePath)
		assert.True(t, exists)
	})

	t.Run("file does not exist", func(t *testing.T) {
		filePath := filepath.Join(tmpDir, "nonexistent.txt")

		exists := Exists(filePath)
		assert.False(t, exists)
	})

	t.Run("directory exists", func(t *testing.T) {
		dirPath := filepath.Join(tmpDir, "subdir")
		err := os.Mkdir(dirPath, 0755)
		assert.NoError(t, err)

		exists := Exists(dirPath)
		assert.True(t, exists)
	})
}

func Test_integration_IsValidMagic(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		filePath, err := filepath.Abs("../test/fixture/db/VACUUM.cache")
		if err != nil {
			t.Error(err)
		}

		got, err := IsValidMagic(filePath, []byte("SQLite format 3"))
		if err != nil {
			t.Error(err)
		}
		assert.True(t, got)
	})

	t.Run("should return false", func(t *testing.T) {
		filePath, err := filepath.Abs("../test/fixture/db/NotSqliteFile.txt")
		if err != nil {
			t.Error(err)
		}

		got, err := IsValidMagic(filePath, []byte("SQLite"))
		assert.Error(t, err)
		assert.False(t, got)
	})

	t.Run("should return false even if the header string is the same as the one in the magic", func(t *testing.T) {
		filePath, err := filepath.Abs("../test/fixture/db/NotSqliteFile_with_SQLite_txt.txt")
		if err != nil {
			t.Error(err)
		}

		got, err := IsValidMagic(filePath, []byte("SQLite"))
		assert.NoError(t, err)
		assert.True(t, got)
	})
}
