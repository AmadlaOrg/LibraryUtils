package desktop

import (
	"fmt"
	"github.com/AmadlaOrg/LibraryUtils/pointer"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	service := NewBuilder("mockname")
	assert.NotNil(t, service)
	assert.IsType(t, &Builder{}, service)
}

// FIXME:
/*func Test_Build(t *testing.T) {
	tests := []struct {
		name        string
		desktop     *Desktop
		expectErr   bool
		expectedStr string
	}{
		{
			name: "Valid desktop entry with application type",
			desktop: &Desktop{
				Groups: &[]Group{
					{
						Title:       pointer.ToPtr("Example"),
						Names:       &[]Content{{Value: "Example App"}},
						Version:     pointer.ToPtr("2.0"),
						XAppVersion: pointer.ToPtr("1.2"),
						Icon:        pointer.ToPtr("example.png"),
						Terminal:    false,
						Type:        pointer.ToPtr(ApplicationType),
						Exec:        "/usr/bin/example",
					},
				},
			},
			expectErr:   false,
			expectedStr: "[Example]\nName=Example App\nVersion=2.0\nX-AppVersion=1.2\nIcon=example.png\nTerminal=false\nType=Application\nExec=/usr/bin/example\n",
		},
		{
			name: "Missing required Name property",
			desktop: &Desktop{
				Groups: &[]Group{
					{
						Title:       pointer.ToPtr("InvalidApp"),
						Version:     pointer.ToPtr("1.0"),
						XAppVersion: pointer.ToPtr("1.0"),
						Icon:        pointer.ToPtr("invalid.png"),
						Terminal:    false,
						Type:        pointer.ToPtr(ApplicationType),
						Exec:        "/usr/bin/invalid",
					},
				},
			},
			expectErr:   true,
			expectedStr: "",
		},
		{
			name: "Valid desktop entry with link type",
			desktop: &Desktop{
				Groups: &[]Group{
					{
						Title:       pointer.ToPtr("WebLink"),
						Names:       &[]Content{{Value: "My Website"}},
						Version:     pointer.ToPtr("1.0"),
						XAppVersion: pointer.ToPtr("1.5"),
						Icon:        pointer.ToPtr("web.png"),
						Terminal:    false,
						Type:        pointer.ToPtr(LinkType),
						URL:         "https://example.com",
					},
				},
			},
			expectErr:   false,
			expectedStr: "[WebLink]\nName=My Website\nVersion=1.0\nX-AppVersion=1.5\nIcon=web.png\nTerminal=false\nType=Link\nUrl=https://example.com\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new SDesktop instance
			service := &SDesktop{builder: strings.Builder{}}

			// Execute Build function
			_, err := service.Build(tt.desktop)

			// Check for expected error condition
			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			}

			// Check for expected output
			if !tt.expectErr && service.builder.String() != tt.expectedStr {
				t.Errorf("Expected output:\n%q\nGot:\n%q", tt.expectedStr, service.builder.String())
			}
		})
	}
}*/

// TODO:
/*func Test_Save(t *testing.T) {
	// Define a temporary directory for testing
	tmpDir := t.TempDir()

	tests := []struct {
		name        string
		dirPath     string
		appName     string
		content     string
		expectErr   bool
		expectedStr string
	}{
		{
			name:        "Valid save",
			dirPath:     tmpDir,
			appName:     "testapp",
			content:     "[Desktop Entry]\nName=Test App\n",
			expectErr:   false,
			expectedStr: "[Desktop Entry]\nName=Test App\n",
		},
		{
			name:        "Empty content",
			dirPath:     tmpDir,
			appName:     "emptyapp",
			content:     "",
			expectErr:   false,
			expectedStr: "",
		},
		{
			name:        "Invalid directory",
			dirPath:     "/invalid/path",
			appName:     "invalidapp",
			content:     "Name=Invalid",
			expectErr:   true,
			expectedStr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create an instance of SDesktop with a builder containing test content
			service := &SDesktop{appName: tt.appName, builder: strings.Builder{}}
			service.builder.WriteString(tt.content)

			// Run Save function
			err := service.Save(tt.dirPath)

			// Check for expected error condition
			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			}

			// Verify file content if no error was expected
			if !tt.expectErr {
				filePath := filepath.Join(tt.dirPath, tt.appName+".desktop")
				data, err := os.ReadFile(filePath)
				if err != nil {
					t.Errorf("Failed to read file: %v", err)
				}
				if string(data) != tt.expectedStr {
					t.Errorf("Expected file content:\n%q\nGot:\n%q", tt.expectedStr, string(data))
				}
			}
		})
	}
}*/

func Test_processContent(t *testing.T) {
	tests := []struct {
		name          string
		propertyName  string
		contentValues []Content
		isNeeded      bool
		expectedErr   error
		expectedStr   string
	}{
		{
			name:          "No content and isNeeded is false",
			propertyName:  "TestProp",
			contentValues: []Content{},
			isNeeded:      false,
			expectedErr:   nil,
			expectedStr:   "",
		},
		{
			name:          "No content and isNeeded is true",
			propertyName:  "TestProp",
			contentValues: []Content{},
			isNeeded:      true,
			expectedErr:   fmt.Errorf("TestProp not set"),
			expectedStr:   "",
		},
		{
			name:         "Single content without language",
			propertyName: "TestProp",
			contentValues: []Content{
				{Language: nil, Value: "Hello"},
			},
			isNeeded:    false,
			expectedErr: nil,
			expectedStr: "TestProp=Hello\n",
		},
		{
			name:         "Single content with language",
			propertyName: "TestProp",
			contentValues: []Content{
				{Language: pointer.ToPtr("en"), Value: "Hello"},
			},
			isNeeded:    false,
			expectedErr: nil,
			expectedStr: "TestProp[en]=Hello\n",
		},
		{
			name:         "Multiple contents with and without language",
			propertyName: "TestProp",
			contentValues: []Content{
				{Language: nil, Value: "Default"},
				{Language: pointer.ToPtr("fr"), Value: "Bonjour"},
			},
			isNeeded:    false,
			expectedErr: nil,
			expectedStr: "TestProp=Default\nTestProp[fr]=Bonjour\n",
		},
		{
			name:         "Content only with language and isNeeded is true",
			propertyName: "TestProp",
			contentValues: []Content{
				{Language: pointer.ToPtr("fr"), Value: "Bonjour"},
			},
			isNeeded:    true,
			expectedErr: fmt.Errorf("TestProp is set with specific language but not set without specific language"),
			expectedStr: "TestProp[fr]=Bonjour\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Builder{builder: strings.Builder{}}
			err := service.processContent(tt.propertyName, &tt.contentValues, tt.isNeeded)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedStr, service.builder.String())
			}
		})
	}
}

func Test_processPropertyDefault(t *testing.T) {
	tests := []struct {
		name         string
		propertyName string
		contentValue *string
		defaultValue string
		expectedStr  string
	}{
		{
			name:         "Valid content value",
			propertyName: "Username",
			contentValue: pointer.ToPtr("JohnDoe"),
			defaultValue: "Guest",
			expectedStr:  "Username=JohnDoe\n",
		},
		{
			name:         "Nil content value (use default)",
			propertyName: "Username",
			contentValue: nil,
			defaultValue: "Guest",
			expectedStr:  "Username=Guest\n",
		},
		{
			name:         "Empty content value (use default)",
			propertyName: "Username",
			contentValue: pointer.ToPtr(""),
			defaultValue: "Guest",
			expectedStr:  "Username=Guest\n",
		},
		{
			name:         "Non-empty default value",
			propertyName: "Password",
			contentValue: nil,
			defaultValue: "1234",
			expectedStr:  "Password=1234\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Builder{builder: strings.Builder{}}
			service.processPropertyDefault(tt.propertyName, tt.contentValue, tt.defaultValue)

			assert.Equal(t, tt.expectedStr, service.builder.String())
		})
	}
}

func Test_processRequiredProperty(t *testing.T) {
	tests := []struct {
		name         string
		propertyName string
		contentValue *string
		expectedErr  error
		expectedStr  string
	}{
		{
			name:         "Valid content value",
			propertyName: "Username",
			contentValue: pointer.ToPtr("JohnDoe"),
			expectedErr:  nil,
			expectedStr:  "Username=JohnDoe\n",
		},
		{
			name:         "Nil content value (should error)",
			propertyName: "Username",
			contentValue: nil,
			expectedErr:  fmt.Errorf("the property Username is required"),
			expectedStr:  "",
		},
		{
			name:         "Empty content value (should error)",
			propertyName: "Username",
			contentValue: pointer.ToPtr(""),
			expectedErr:  fmt.Errorf("the property Username is required"),
			expectedStr:  "",
		},
		{
			name:         "Non-empty value for a different property",
			propertyName: "Password",
			contentValue: pointer.ToPtr("securePass"),
			expectedErr:  nil,
			expectedStr:  "Password=securePass\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new SDesktop with a buffer
			service := &Builder{builder: strings.Builder{}}
			err := service.processRequiredProperty(tt.propertyName, tt.contentValue)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedStr, service.builder.String())
			}
		})
	}
}

func Test_processNotRequiredProperty(t *testing.T) {
	tests := []struct {
		name         string
		propertyName string
		contentValue *string
		expectedStr  string
	}{
		{
			name:         "Valid content value",
			propertyName: "Username",
			contentValue: pointer.ToPtr("JohnDoe"),
			expectedStr:  "Username=JohnDoe\n",
		},
		{
			name:         "Nil content value (should write nothing)",
			propertyName: "Username",
			contentValue: nil,
			expectedStr:  "",
		},
		{
			name:         "Empty content value (should write nothing)",
			propertyName: "Username",
			contentValue: pointer.ToPtr(""),
			expectedStr:  "",
		},
		{
			name:         "Non-empty value for a different property",
			propertyName: "Password",
			contentValue: pointer.ToPtr("securePass"),
			expectedStr:  "Password=securePass\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Builder{builder: strings.Builder{}}
			service.processNotRequiredProperty(tt.propertyName, tt.contentValue)

			assert.Equal(t, tt.expectedStr, service.builder.String())
		})
	}
}

func Test_processList(t *testing.T) {
	tests := []struct {
		name         string
		propertyName string
		items        *[]List
		isNeeded     bool
		expectedErr  error
		expectedStr  string
	}{
		{
			name:         "Valid list",
			propertyName: "Fruits",
			items:        &[]List{"Apple", "Banana", "Cherry"},
			isNeeded:     false,
			expectedErr:  nil,
			expectedStr:  "Fruits=Apple;Banana;Cherry\n",
		},
		{
			name:         "Empty list but not required",
			propertyName: "Fruits",
			items:        &[]List{},
			isNeeded:     false,
			expectedErr:  nil,
			expectedStr:  "",
		},
		{
			name:         "Nil list but not required",
			propertyName: "Fruits",
			items:        nil,
			isNeeded:     false,
			expectedErr:  nil,
			expectedStr:  "",
		},
		{
			name:         "Empty list but required",
			propertyName: "Fruits",
			items:        &[]List{},
			isNeeded:     true,
			expectedErr:  fmt.Errorf("the property Fruits is empty"),
			expectedStr:  "",
		},
		{
			name:         "Nil list but required",
			propertyName: "Fruits",
			items:        nil,
			isNeeded:     true,
			expectedErr:  fmt.Errorf("the property Fruits is empty"),
			expectedStr:  "",
		},
		{
			name:         "Single item list",
			propertyName: "Colors",
			items:        &[]List{"Blue"},
			isNeeded:     false,
			expectedErr:  nil,
			expectedStr:  "Colors=Blue\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Builder{builder: strings.Builder{}}
			err := service.processList(tt.propertyName, tt.items, tt.isNeeded)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedStr, service.builder.String())
			}
		})
	}
}

func Test_processCommaList(t *testing.T) {
	tests := []struct {
		name         string
		propertyName string
		items        *[]CommaList
		isNeeded     bool
		expectedErr  error
		expectedStr  string
	}{
		{
			name:         "Valid list",
			propertyName: "Fruits",
			items:        &[]CommaList{"Apple", "Banana", "Cherry"},
			isNeeded:     false,
			expectedErr:  nil,
			expectedStr:  "Fruits=Apple,Banana,Cherry\n",
		},
		{
			name:         "Empty list but not required",
			propertyName: "Fruits",
			items:        &[]CommaList{},
			isNeeded:     false,
			expectedErr:  nil,
			expectedStr:  "",
		},
		{
			name:         "Nil list but not required",
			propertyName: "Fruits",
			items:        nil,
			isNeeded:     false,
			expectedErr:  nil,
			expectedStr:  "",
		},
		{
			name:         "Empty list but required",
			propertyName: "Fruits",
			items:        &[]CommaList{},
			isNeeded:     true,
			expectedErr:  fmt.Errorf("the property Fruits is empty"),
			expectedStr:  "",
		},
		{
			name:         "Nil list but required",
			propertyName: "Fruits",
			items:        nil,
			isNeeded:     true,
			expectedErr:  fmt.Errorf("the property Fruits is empty"),
			expectedStr:  "",
		},
		{
			name:         "Single item list",
			propertyName: "Colors",
			items:        &[]CommaList{"Blue"},
			isNeeded:     false,
			expectedErr:  nil,
			expectedStr:  "Colors=Blue\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &Builder{builder: strings.Builder{}}
			err := service.processCommaList(tt.propertyName, tt.items, tt.isNeeded)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedStr, service.builder.String())
			}
		})
	}
}
