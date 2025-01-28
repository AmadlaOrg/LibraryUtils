package file

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

func TestExists(t *testing.T) {
	tests := []struct {
		name                 string
		inputPath            string
		internalOsStat       func(name string) (os.FileInfo, error)
		internalOsIsNotExist func(err error) bool
		expectedExists       bool
	}{
		{
			name:      "exists",
			inputPath: "testdata/exists.txt",
			internalOsStat: func(name string) (os.FileInfo, error) {
				return nil, nil
			},
			expectedExists: true,
		},
		{
			name:      "does not exists",
			inputPath: "testdata/exists.txt",
			internalOsStat: func(name string) (os.FileInfo, error) {
				return nil, errors.New("test error (osStat)")
			},
			internalOsIsNotExist: func(err error) bool {
				return true
			},
			expectedExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalOsStat := osStat
			defer func() { osStat = originalOsStat }()
			osStat = func(name string) (os.FileInfo, error) {
				return tt.internalOsStat(name)
			}

			originalOsIsNotExist := osIsNotExist
			defer func() { osIsNotExist = originalOsIsNotExist }()
			osIsNotExist = func(err error) bool {
				return tt.internalOsIsNotExist(err)
			}

			got := Exists(tt.inputPath)
			assert.Equal(t, tt.expectedExists, got)
		})
	}
}

func TestIsFile(t *testing.T) {
	tests := []struct {
		name           string
		inputPath      string
		internalOsStat func(string) (os.FileInfo, error)
		expected       bool
		expectErr      error
		hasError       bool
	}{
		{
			name:      "file exists",
			inputPath: "./testdata/test.txt",
			internalOsStat: func(string) (os.FileInfo, error) {
				mockFileInfo := NewMockFileInfo(t)
				mockFileInfo.EXPECT().IsDir().Return(false)
				return mockFileInfo, nil
			},
			expected:  true,
			expectErr: nil,
			hasError:  false,
		},
		//
		// Error
		//
		{
			name:      "file exists",
			inputPath: "./testdata/test.txt",
			internalOsStat: func(string) (os.FileInfo, error) {
				mockFileInfo := NewMockFileInfo(t)
				return mockFileInfo, errors.New("stat ./testdata/test.txt: no such file or directory")
			},
			expected:  false,
			expectErr: errors.New("stat ./testdata/test.txt: no such file or directory"),
			hasError:  true,
		},
		{
			name:      "file exists",
			inputPath: "./testdata/test.txt",
			internalOsStat: func(string) (os.FileInfo, error) {
				mockFileInfo := NewMockFileInfo(t)
				mockFileInfo.EXPECT().IsDir().Return(true)
				return mockFileInfo, nil
			},
			expected:  false,
			expectErr: errors.New("not a file"),
			hasError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalOsStat := osStat
			defer func() { osStat = originalOsStat }()
			osStat = tt.internalOsStat

			got, err := IsFile(tt.inputPath)
			if tt.hasError {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.expectErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestIsFilePathDirExist(t *testing.T) {
	tests := []struct {
		name                string
		inputPath           string
		internalOsStat      func(string) (os.FileInfo, error)
		expected            bool
		expectedFileDirPath string
		expectErr           error
		hasError            bool
	}{
		{
			name:      "True: directory exists",
			inputPath: "./testdata/file.txt",
			internalOsStat: func(path string) (os.FileInfo, error) {
				mockFileInfo := NewMockFileInfo(t)
				mockFileInfo.EXPECT().IsDir().Return(true)
				return mockFileInfo, nil
			},
			expected:            true,
			expectedFileDirPath: "testdata",
			expectErr:           nil,
			hasError:            false,
		},
		{
			name:      "True: directory exists with full path",
			inputPath: "/home/user/testdata/file.txt",
			internalOsStat: func(path string) (os.FileInfo, error) {
				mockFileInfo := NewMockFileInfo(t)
				mockFileInfo.EXPECT().IsDir().Return(true)
				return mockFileInfo, nil
			},
			expected:            true,
			expectedFileDirPath: "/home/user/testdata",
			expectErr:           nil,
			hasError:            false,
		},
		{
			name:      "False: directory does not exist",
			inputPath: "./nonexistent/file.txt",
			internalOsStat: func(path string) (os.FileInfo, error) {
				return nil, os.ErrNotExist
			},
			expected:            false,
			expectedFileDirPath: "nonexistent",
			expectErr:           os.ErrNotExist,
			hasError:            true,
		},
		{
			name:      "True: input path is a directory",
			inputPath: "testdata",
			internalOsStat: func(path string) (os.FileInfo, error) {
				mockFileInfo := NewMockFileInfo(t)
				mockFileInfo.EXPECT().IsDir().Return(true)
				return mockFileInfo, nil
			},
			expected:            true,
			expectedFileDirPath: ".",
			expectErr:           nil,
			hasError:            false,
		},
		{
			name:      "True: file in current directory",
			inputPath: "file.txt",
			internalOsStat: func(path string) (os.FileInfo, error) {
				mockFileInfo := NewMockFileInfo(t)
				mockFileInfo.EXPECT().IsDir().Return(true)
				return mockFileInfo, nil
			},
			expected:            true,
			expectedFileDirPath: ".",
			expectErr:           nil,
			hasError:            false,
		},
		{
			name:      "False: empty input path",
			inputPath: "",
			internalOsStat: func(path string) (os.FileInfo, error) {
				return nil, os.ErrInvalid
			},
			expected:            false,
			expectedFileDirPath: ".",
			expectErr:           os.ErrInvalid,
			hasError:            true,
		},
		{
			name:      "False: invalid path format",
			inputPath: "/invalid|path/file.txt",
			internalOsStat: func(path string) (os.FileInfo, error) {
				return nil, os.ErrInvalid
			},
			expected:            false,
			expectedFileDirPath: "/invalid|path",
			expectErr:           os.ErrInvalid,
			hasError:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock osStat
			originalOsStat := osStat
			defer func() { osStat = originalOsStat }()
			osStat = tt.internalOsStat

			// Execute the function
			got, dirPath, err := IsFilePathDirExist(tt.inputPath)

			// Assertions
			if tt.hasError {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.expectErr.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedFileDirPath, dirPath)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestIsValidMagic(t *testing.T) {
	tests := []struct {
		name                string
		inputPath           string
		inputMagic          []byte
		internalOsStatError error
		internalOsOpen      func(name string) (IFile, error)
		internalBytesEqual  func(a, b []byte) bool
		expect              bool
		expectedErr         error
		hasError            bool
	}{
		{
			name:      "Success",
			inputPath: `testdata/valid.db`,
			internalOsOpen: func(name string) (IFile, error) {
				mockFile := NewMockFile(t)
				mockFile.EXPECT().Read(mock.Anything).Return(0, nil)
				mockFile.EXPECT().Close().Return(nil)
				return mockFile, nil
			},
			internalBytesEqual: func(a, b []byte) bool {
				return true
			},
			expect:   true,
			hasError: false,
		},
		//
		// Error
		//
		{
			name:                "Error: fails at IsFile",
			inputPath:           `testdata/valid.db`,
			internalOsStatError: errors.New("open ./testdata/valid.db: no such file or directory"),
			expect:              false,
			expectedErr:         errors.New("open ./testdata/valid.db: no such file or directory"),
			hasError:            true,
		},
		{
			name:      "Error: fails at os.Open",
			inputPath: `testdata/valid.db`,
			internalOsOpen: func(name string) (IFile, error) {
				mockFile := NewMockFile(t)
				return mockFile, errors.New("open ./testdata/valid.db: no such file or directory")
			},
			expect:      false,
			expectedErr: errors.New("open ./testdata/valid.db: no such file or directory"),
			hasError:    true,
		},
		{
			name:       "Error: fails at Read file",
			inputPath:  `testdata/valid.db`,
			inputMagic: []byte("SQLite"),
			internalOsOpen: func(name string) (IFile, error) {
				mockFile := NewMockFile(t)
				mockFile.EXPECT().Read(mock.Anything).Return(0, errors.New("some error (file.Read())"))
				mockFile.EXPECT().Close().Return(nil)
				return mockFile, nil
			},
			expect:      false,
			expectedErr: errors.New("some error (file.Read())"),
			hasError:    true,
		},
		{
			name:      "Error: fails at bytes.Equal",
			inputPath: `testdata/valid.db`,
			internalOsOpen: func(name string) (IFile, error) {
				mockFile := NewMockFile(t)
				mockFile.EXPECT().Read(mock.Anything).Return(0, nil)
				mockFile.EXPECT().Close().Return(nil)
				return mockFile, nil
			},
			internalBytesEqual: func(a, b []byte) bool {
				return false
			},
			expect:      false,
			expectedErr: errors.New("does not match magic header"),
			hasError:    true,
		},
		{
			name:      "Error: fails at closing file",
			inputPath: `testdata/valid.db`,
			internalOsOpen: func(name string) (IFile, error) {
				mockFile := NewMockFile(t)
				mockFile.EXPECT().Read(mock.Anything).Return(0, nil)
				mockFile.EXPECT().Close().Return(errors.New("some error (file.Close())"))
				return mockFile, nil
			},
			internalBytesEqual: func(a, b []byte) bool {
				return true
			},
			expect:   true,
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalOsStat := osStat
			defer func() { osStat = originalOsStat }()
			osStat = func(string) (os.FileInfo, error) {
				mockFileInfo := NewMockFileInfo(t)
				if tt.internalOsStatError == nil {
					mockFileInfo.EXPECT().IsDir().Return(false)
				}
				return mockFileInfo, tt.internalOsStatError
			}

			originalOsOpen := osOpen
			defer func() { osOpen = originalOsOpen }()
			osOpen = tt.internalOsOpen

			originalBytesEqual := bytesEqual
			defer func() { bytesEqual = originalBytesEqual }()
			bytesEqual = tt.internalBytesEqual

			got, err := IsValidMagic(tt.inputPath, tt.inputMagic)
			if tt.hasError {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expect, got)
		})
	}
}
