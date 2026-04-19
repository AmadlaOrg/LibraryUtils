package stream

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	input := strings.NewReader("test")
	output := &bytes.Buffer{}

	svc := New("template.tmpl", input, output)
	assert.NotNil(t, svc)
}
