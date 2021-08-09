package reverseinteger

import "math"

func Reverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	reversed := 0
	for x != 0 {
		remainder := x % 10
		reversed = (reversed * 10) + remainder
		x /= 10
	}

	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	}

	return reversed
}
