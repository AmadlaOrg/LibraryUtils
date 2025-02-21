package desktop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convertListToStrings(t *testing.T) {
	t.Run("Convert slice of strings", func(t *testing.T) {
		input := []string{"hello", "world"}
		expected := []string{"hello", "world"}
		result := convertListToStrings(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Convert slice of integers", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []string{"1", "2", "3"}
		result := convertListToStrings(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Convert slice of floats", func(t *testing.T) {
		input := []float64{1.1, 2.2, 3.3}
		expected := []string{"1.1", "2.2", "3.3"}
		result := convertListToStrings(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Convert slice of bools", func(t *testing.T) {
		input := []bool{true, false, true}
		expected := []string{"true", "false", "true"}
		result := convertListToStrings(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Convert empty slice", func(t *testing.T) {
		input := []string{}
		expected := []string{}
		result := convertListToStrings(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Convert slice of custom struct", func(t *testing.T) {
		type Example struct {
			ID   int
			Name string
		}
		input := []Example{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
		}
		expected := []string{"{1 Alice}", "{2 Bob}"} // Default Go formatting
		result := convertListToStrings(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Convert slice of interfaces", func(t *testing.T) {
		input := []any{"string", 42, true, 3.14}
		expected := []string{"string", "42", "true", "3.14"}
		result := convertListToStrings(input)
		assert.Equal(t, expected, result)
	})
}
