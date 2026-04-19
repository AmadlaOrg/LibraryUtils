package sqlite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("should return a new instance of Database", func(t *testing.T) {
		service := New("/home/user/")
		assert.NotNil(t, service)
		assert.IsType(t, &databaseImpl{}, service)
	})
}
