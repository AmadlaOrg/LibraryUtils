package stream

import "io"

// New
func New(tmplFile string, input io.Reader, output io.Writer) Stream {
	return &streamImpl{
		tmplFile: tmplFile,
		input:    input,
		output:   output,
	}
}
