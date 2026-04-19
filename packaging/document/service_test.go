package document

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	svc := New("myapp", "MyApp", "2.0.0", cmd)

	assert.NotNil(t, svc)
}
