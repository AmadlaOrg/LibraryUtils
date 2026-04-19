package stream

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func writeTempTemplate(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "test.tmpl")
	err := os.WriteFile(path, []byte(content), 0644)
	assert.NoError(t, err)
	return path
}

func TestDo(t *testing.T) {
	t.Run("template file missing", func(t *testing.T) {
		input := strings.NewReader(`[{"name": "test"}]`)
		var output bytes.Buffer

		s := &streamImpl{
			tmplFile: "/nonexistent/template.tmpl",
			input:    input,
			output:   &output,
		}

		err := s.Do()
		assert.Error(t, err)
	})

	t.Run("mock templateParseFiles fail", func(t *testing.T) {
		origTemplateParseFiles := templateParseFiles
		defer func() { templateParseFiles = origTemplateParseFiles }()

		templateParseFiles = func(filenames ...string) (*template.Template, error) {
			return nil, errors.New("parse files failed")
		}

		input := strings.NewReader(`[{"name": "test"}]`)
		var output bytes.Buffer

		s := &streamImpl{
			tmplFile: "dummy.tmpl",
			input:    input,
			output:   &output,
		}

		err := s.Do()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "parse files failed")
	})

	t.Run("empty input returns no valid data error", func(t *testing.T) {
		tmplPath := writeTempTemplate(t, "{{.name}}")
		input := strings.NewReader("")
		var output bytes.Buffer

		s := &streamImpl{
			tmplFile: tmplPath,
			input:    input,
			output:   &output,
		}

		err := s.Do()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no valid data found")
	})

	t.Run("JSON input with template", func(t *testing.T) {
		tmplPath := writeTempTemplate(t, "Name: {{.name}}\n")
		input := strings.NewReader(`[{"name": "Alice"}, {"name": "Bob"}]`)
		var output bytes.Buffer

		s := &streamImpl{
			tmplFile: tmplPath,
			input:    input,
			output:   &output,
		}

		err := s.Do()
		assert.NoError(t, err)
		assert.Equal(t, "Name: Alice\nName: Bob\n", output.String())
	})

	t.Run("YAML input with template", func(t *testing.T) {
		tmplPath := writeTempTemplate(t, "Host: {{.host}}\n")
		input := strings.NewReader("host: example.com\n")
		var output bytes.Buffer

		s := &streamImpl{
			tmplFile: tmplPath,
			input:    input,
			output:   &output,
		}

		err := s.Do()
		assert.NoError(t, err)
		assert.Equal(t, "Host: example.com\n", output.String())
	})

	t.Run("YAML array of maps", func(t *testing.T) {
		tmplPath := writeTempTemplate(t, "Item: {{.id}}\n")
		input := strings.NewReader("- id: 1\n- id: 2\n")
		var output bytes.Buffer

		s := &streamImpl{
			tmplFile: tmplPath,
			input:    input,
			output:   &output,
		}

		err := s.Do()
		assert.NoError(t, err)
		assert.Equal(t, "Item: 1\nItem: 2\n", output.String())
	})

	t.Run("template execution error", func(t *testing.T) {
		tmplPath := writeTempTemplate(t, "{{.name.Missing}}")
		input := strings.NewReader(`[{"name": "Alice"}]`)
		var output bytes.Buffer

		s := &streamImpl{
			tmplFile: tmplPath,
			input:    input,
			output:   &output,
		}

		err := s.Do()
		assert.Error(t, err)
	})
}

func TestParse_JSONArray(t *testing.T) {
	input := strings.NewReader(`[{"name": "Alice"}, {"name": "Bob"}]`)
	s := &streamImpl{}

	ch, err := s.parse(input)
	assert.NoError(t, err)
	assert.NotNil(t, ch)

	var results []map[string]any
	for item := range ch {
		results = append(results, item)
	}

	assert.Len(t, results, 2)
	assert.Equal(t, "Alice", results[0]["name"])
	assert.Equal(t, "Bob", results[1]["name"])
}

func TestParse_YAMLDoc(t *testing.T) {
	input := strings.NewReader("name: Bob\nage: 25\n")
	s := &streamImpl{}

	ch, err := s.parse(input)
	assert.NoError(t, err)

	var results []map[string]any
	for item := range ch {
		results = append(results, item)
	}

	assert.Len(t, results, 1)
	assert.Equal(t, "Bob", results[0]["name"])
}

func TestParse_YAMLArrayOfMaps(t *testing.T) {
	input := strings.NewReader("- id: 1\n- id: 2\n")
	s := &streamImpl{}

	ch, err := s.parse(input)
	assert.NoError(t, err)

	var results []map[string]any
	for item := range ch {
		results = append(results, item)
	}

	assert.Len(t, results, 2)
	assert.Equal(t, 1, results[0]["id"])
	assert.Equal(t, 2, results[1]["id"])
}

func TestParse_NonSeekableReader(t *testing.T) {
	// bytes.Buffer does not implement io.ReadSeeker, so parse() buffers it
	buf := bytes.NewBufferString("name: test\n")
	s := &streamImpl{}

	ch, err := s.parse(buf)
	assert.NoError(t, err)

	var results []map[string]any
	for item := range ch {
		results = append(results, item)
	}

	assert.Len(t, results, 1)
	assert.Equal(t, "test", results[0]["name"])
}

func TestParse_InvalidInput(t *testing.T) {
	// Scalar YAML string — not a map, goes to "Unsupported YAML structure" path
	input := strings.NewReader("just a plain string")
	s := &streamImpl{}

	ch, err := s.parse(input)
	assert.NoError(t, err)

	var results []map[string]any
	for item := range ch {
		results = append(results, item)
	}

	assert.Empty(t, results)
}

func TestParse_MultipleYAMLDocs(t *testing.T) {
	input := strings.NewReader("name: Alice\n---\nname: Bob\n")
	s := &streamImpl{}

	ch, err := s.parse(input)
	assert.NoError(t, err)

	var results []map[string]any
	for item := range ch {
		results = append(results, item)
	}

	assert.Len(t, results, 2)
	assert.Equal(t, "Alice", results[0]["name"])
	assert.Equal(t, "Bob", results[1]["name"])
}
