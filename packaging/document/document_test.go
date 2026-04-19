package document

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/stretchr/testify/assert"
)

func newTestCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "testapp",
		Short: "A test application",
	}
}

func TestMan(t *testing.T) {
	tests := []struct {
		name       string
		mockManErr error
		wantErr    bool
	}{
		{
			name:       "happy path",
			mockManErr: nil,
		},
		{
			name:       "error from GenManTree",
			mockManErr: errors.New("man generation failed"),
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			origDocGenManTree := docGenManTree
			defer func() { docGenManTree = origDocGenManTree }()

			var capturedHeader *doc.GenManHeader
			docGenManTree = func(cmd *cobra.Command, header *doc.GenManHeader, dir string) error {
				capturedHeader = header
				return tt.mockManErr
			}

			svc := &documentImpl{
				appName:    "testapp",
				appTitle:   "TestApp",
				appVersion: "1.0.0",
				rootCmd:    newTestCmd(),
			}

			err := svc.Man("/tmp/docs")

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "man generation failed")
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, capturedHeader)
			assert.Equal(t, "TestApp", capturedHeader.Title)
			assert.Equal(t, "1", capturedHeader.Section)
			assert.Equal(t, "TestApp v1.0.0", capturedHeader.Source)
			assert.Equal(t, "TestApp Manual", capturedHeader.Manual)
		})
	}
}

func TestMarkdown(t *testing.T) {
	tests := []struct {
		name    string
		mockErr error
		wantErr bool
	}{
		{
			name: "happy path",
		},
		{
			name:    "error from GenMarkdownTree",
			mockErr: errors.New("markdown generation failed"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			origDocGenMarkdownTree := docGenMarkdownTree
			defer func() { docGenMarkdownTree = origDocGenMarkdownTree }()

			docGenMarkdownTree = func(cmd *cobra.Command, dir string) error {
				return tt.mockErr
			}

			svc := &documentImpl{
				appName:    "testapp",
				appTitle:   "TestApp",
				appVersion: "1.0.0",
				rootCmd:    newTestCmd(),
			}

			err := svc.Markdown("/tmp/docs")

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "markdown generation failed")
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		name       string
		manErr     error
		mkErr      error
		wantErr    bool
		errContain string
	}{
		{
			name: "both succeed",
		},
		{
			name:       "man only fails",
			manErr:     errors.New("man error"),
			wantErr:    true,
			errContain: "man error",
		},
		{
			name:       "markdown only fails",
			mkErr:      errors.New("markdown error"),
			wantErr:    true,
			errContain: "markdown error",
		},
		{
			name:       "both fail",
			manErr:     errors.New("man error"),
			mkErr:      errors.New("markdown error"),
			wantErr:    true,
			errContain: "man error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			origDocGenManTree := docGenManTree
			origDocGenMarkdownTree := docGenMarkdownTree
			defer func() {
				docGenManTree = origDocGenManTree
				docGenMarkdownTree = origDocGenMarkdownTree
			}()

			docGenManTree = func(cmd *cobra.Command, header *doc.GenManHeader, dir string) error {
				return tt.manErr
			}
			docGenMarkdownTree = func(cmd *cobra.Command, dir string) error {
				return tt.mkErr
			}

			svc := &documentImpl{
				appName:    "testapp",
				appTitle:   "TestApp",
				appVersion: "1.0.0",
				rootCmd:    newTestCmd(),
			}

			err := svc.All("/tmp/docs")

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContain)
				return
			}

			assert.NoError(t, err)
		})
	}
}
