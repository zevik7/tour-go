package main

// =============================================================================
// SWAP FUNCTION
// =============================================================================

// swap demonstrates multiple return values in Go
// This is a fundamental Go feature that allows functions to return more than
// one value, which is commonly used for returning both a result and an error
//
// This demonstrates:
// - Multiple return values (string, string)
// - Simple value swapping without temporary variables
// - Clean syntax for returning multiple values
//
// Parameters:
//   x - first string to swap
//   y - second string to swap
//
// Returns:
//   Two strings in swapped order (y, x)
//
// Example:
//   first, second := swap("hello", "world")
//   // first = "world", second = "hello"
func swap(x, y string) (string, string) {
	return y, x
}
