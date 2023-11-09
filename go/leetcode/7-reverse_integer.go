/* Reverse Integer
Source		: https://leetcode.com/problems/reverse-integer/
Level		: Medium
Description	: Given a signed 32-bit integer `x`, return `x` with its digits reversed. If reversing `x` causes the value
			to go outside the signed 32-bit integer range [-2^31, 2^31 - 1], then return `0`.

Example 1:
Input: x = 123
Output: 321

Example 2:
Input: x = -123
Output: -321

Example 3:
Input: x = 120
Output: 21
*/

package leetcode

import "math"

func (soln Solution) Reverse(x int) int {
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
