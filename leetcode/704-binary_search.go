/* Binary Search
Source		: https://leetcode.com/problems/binary-search/
Level		: Easy
Description	: Given an array of integers nums which is sorted in ascending order, and an integer target, write a
			function to search target in nums. If target exists, then return its index. Otherwise, return -1. You must
			write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1

Solution source: https://www.programiz.com/dsa/binary-search
*/

package leetcode

// Search implements binary search to find target and returns the target index
func Search(nums []int, target int) int {
	// list of index
	low, high, mid := 0, len(nums)-1, 0

	// repeat until the pointers low and high meet each other
	for low <= high {
		mid = (low + high) / 2
		switch {
		case target == nums[mid]:
			return mid
		case target > nums[mid]:
			low = mid + 1
		default:
			high = mid - 1
		}
	}

	return -1
}
