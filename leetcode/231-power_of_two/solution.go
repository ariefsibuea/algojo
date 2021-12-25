package poweroftwo

import "math"

/**
 * Problem source: https://leetcode.com/problems/power-of-two/
**/

func IsPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	if math.Ceil(math.Log2(float64(n))) == math.Floor(math.Log2(float64(n))) {
		return true
	}

	return false
}
