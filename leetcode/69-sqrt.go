/* Sqrt(x)
Source		: https://leetcode.com/problems/sqrtx/
Level		: Easy
Description	: Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The
			returned integer should be non-negative as well. You must not use any built-in exponent function or operator.

Example 1:
Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.

Example 2:
Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.
*/

package leetcode

func (soln Solution) MySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x
	for left < right {
		mid := left + ((right - left) / 2)

		if mid*mid == x {
			return mid
		}

		if mid*mid > x {
			right = mid
			continue
		}
		left = mid + 1
	}

	return left - 1
}
