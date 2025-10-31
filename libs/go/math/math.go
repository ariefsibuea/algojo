package math

import "math/bits"

func Abs(n int) int {
	// Get sign bit as a mask: all 0s for positive, all 1s for negative
	mask := n >> (bits.UintSize - 1)

	// Handle INT_MIN specially by using unsign arithmetic
	return int((uint(n) ^ uint(mask)) - uint(mask))
}
