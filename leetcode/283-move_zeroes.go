/* Move Zeroes
Source		: https://leetcode.com/problems/move-zeroes/
Level		: Easy
Description	: Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the
			non-zero elements. Note that you must do this in-place without making a copy of the array.

Example 1:
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:
Input: nums = [0]
Output: [0]
*/

package leetcode

// MoveZeroes implements two pointers technique to solve move zeroes problem
func MoveZeroes(nums []int) {
	lastNonZero := 0
	for index := range nums {
		if nums[index] == 0 {
			continue
		}
		nums[lastNonZero] = nums[index]
		lastNonZero += 1
	}
	for i := lastNonZero; i < len(nums); i++ {
		nums[i] = 0
	}
}
