package desktop

import (
	"fmt"
)

// convertListToStrings converts any slice of values to a slice of strings.
func convertListToStrings[T any](items []T) []string {
	stringItems := make([]string, len(items))
	for i, item := range items {
		stringItems[i] = fmt.Sprintf("%v", item) // Convert any type to string
	}
	return stringItems
}
