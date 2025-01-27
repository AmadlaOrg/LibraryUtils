package stream

import "io"

// NewStreamService
func NewStreamService(tmplFile string, input io.Reader, output io.Writer) IStream {
	return &SStream{
		tmplFile: tmplFile,
		input:    input,
		output:   output,
	}
}
