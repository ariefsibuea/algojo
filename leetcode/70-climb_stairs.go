/* Climbing Stairs
Source		: https://leetcode.com/problems/climbing-stairs/
Level		: Easy
Description	: You are climbing a staircase. It takes n steps to reach the top. Each time you can either climb 1 or 2
			steps. In how many distinct ways can you climb to the top?

Example 1:
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/

package leetcode

func ClimbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	last1, last2, sum := 1, 1, 0
	for i := 2; i <= n; i++ {
		sum = last1 + last2
		last2 = last1
		last1 = sum
	}

	return sum
}
