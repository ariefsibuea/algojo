/* Power of Two
Source		: https://leetcode.com/problems/power-of-two/
Level		: Easy
Description	: Given an integer n, return true if it is a power of two. Otherwise, return false.

Example 1:
Input: n = 1
Output: true
Explanation: 2^0 = 1

Example 2:
Input: n = 16
Output: true
Explanation: 2^4 = 16

Example 3:
Input: n = 3
Output: false
*/

package leetcode

import "math"

func IsPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	if math.Ceil(math.Log2(float64(n))) == math.Floor(math.Log2(float64(n))) {
		return true
	}

	return false
}
