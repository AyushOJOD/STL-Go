package utils

// IsValidIndex checks if the given index is valid for a slice of given length
func IsValidIndex[T any](index, length int) bool {
	return index >= 0 && index < length
}

// Equal compares two values of the same type
func Equal[T comparable](a, b T) bool {
	return a == b
}