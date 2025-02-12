package pointer

// ToPtr is a helper function that creates a pointer to a value of any type.
//
// Params:
// - 🫶 v: The value of any type that will be returned as a pointer to its type.
func ToPtr[T any](v T) *T {
	return &v
}
