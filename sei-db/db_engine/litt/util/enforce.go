package util

// True asserts a condition is true and panics with a message if the condition is false.
func True(condition bool, message string, args ...any) { _ = "STUB: not implemented"; return }

// MapDoesNotContainKey asserts that a map does not contain a specific key and panics
// with an error message if it does.
func MapDoesNotContainKey[K comparable, V any](m map[K]V, key K, message string, args ...any) {
	_ = "STUB: not implemented"
	return
}
