package stream

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"

	"text/template"
)

type Stream interface {
	Do() error
}

type streamImpl struct {
	tmplFile string
	input    io.Reader
	output   io.Writer
}

var (
	templateParseFiles = template.ParseFiles
)

// Do Function to process the template with streaming data
func (s *streamImpl) Do() error {
	// Open the template file
	tmpl, err := templateParseFiles(s.tmplFile)
	if err != nil {
		return err
	}

	// Parse and stream the input data
	dataStream, err := s.parse(s.input)
	if err != nil {
		return err
	}

	// Process each item in the stream
	count := 0
	for item := range dataStream {
		if err := tmpl.Execute(s.output, item); err != nil {
			return err
		}

		// Optionally add a newline or separator
		if _, err := s.output.Write([]byte("")); err != nil {
			return err
		}
		count++
	}

	if count == 0 {
		return errors.New("no valid data found in input")
	}

	return nil
}

// Detects and streams JSON or YAML content
func (s *streamImpl) parse(input io.Reader) (<-chan map[string]any, error) {
	// Ensure input is a ReadSeeker (required for Seek)
	var seekableInput io.ReadSeeker
	if rs, ok := input.(io.ReadSeeker); ok {
		seekableInput = rs
	} else {
		// Copy input into a buffer to enable seeking
		buf := new(bytes.Buffer)
		if _, err := io.Copy(buf, input); err != nil {
			return nil, fmt.Errorf("failed to copy input: %v", err)
		}
		seekableInput = bytes.NewReader(buf.Bytes())
	}

	// Channel to stream parsed data
	ch := make(chan map[string]any)

	// Try JSON decoding first
	decoder := json.NewDecoder(seekableInput)
	if _, err := decoder.Token(); err == nil { // Valid JSON array starts with '['
		go func() {
			defer close(ch)
			for decoder.More() {
				var item map[string]any
				if err := decoder.Decode(&item); err == nil {
					ch <- item
				}
			}
		}()
		return ch, nil
	}

	// Reset the reader for YAML decoding
	if _, err := seekableInput.Seek(0, io.SeekStart); err != nil {
		return nil, fmt.Errorf("failed to seek input: %v", err)
	}

	// Handle YAML decoding
	go func() {
		defer close(ch)
		yamlDecoder := yaml.NewDecoder(seekableInput)

		for {
			var item any
			if err := yamlDecoder.Decode(&item); err != nil {
				if err == io.EOF {
					break
				}
				fmt.Fprintf(os.Stderr, "Error decoding YAML: %v\n", err)
				break
			}

			switch value := item.(type) {
			case map[string]any:
				ch <- value
			case []any: // Handle sequences
				for _, v := range value {
					if m, ok := v.(map[string]any); ok {
						ch <- m
					}
				}
			default:
				fmt.Fprintf(os.Stderr, "Unsupported YAML structure: %v\n", value)
			}
		}
	}()

	return ch, nil
}
