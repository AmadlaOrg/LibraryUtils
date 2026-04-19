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
