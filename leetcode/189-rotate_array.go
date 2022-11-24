/* Rotate Array
Source		: https://leetcode.com/problems/rotate-array/
Level		: Medium
Description	: Given an array, rotate the array to the right by k steps, where k is non-negative.

Example 1:
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]

Solution source: https://www.geeksforgeeks.org/array-rotation/
*/

package leetcode

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

// Rotate implements juggling algorithm to rotate array elements to the right
func Rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	numSet := gcd(k, n)
	for i := n - 1; i >= n-numSet; i-- {
		temp := nums[i]
		curr := i
		for {
			next := curr - k
			if next < 0 {
				next = next + n
			}
			if next == i {
				break
			}
			nums[curr] = nums[next]
			curr = next
		}
		nums[curr] = temp
	}
}
