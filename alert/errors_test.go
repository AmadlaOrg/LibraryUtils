package alert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNotFound(t *testing.T) {
	assert.EqualError(t, ErrorNotFound, "not found")
}

func TestErrorMultipleFound(t *testing.T) {
	assert.EqualError(t, ErrorMultipleFound, "multiple found")
}

func TestErrorsAreDistinct(t *testing.T) {
	assert.NotEqual(t, ErrorNotFound, ErrorMultipleFound)
}
