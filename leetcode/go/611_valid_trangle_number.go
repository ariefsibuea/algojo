package main

import "sort"

/**
 * LeetCode Problem : Valid Triangle Number
 * Topic            : Array, Two Pointers, Binary Search, Greedy, Sorting
 * Level            : Medium
 * URL              : https://leetcode.com/problems/valid-triangle-number
 * Description      :
 * Examples         :
 * 			Example 1:
 * 			Input: nums = [2,2,3,4]
 * 			Output: 3
 * 			Explanation: Valid combinations are:
 * 			2,3,4 (using the first 2)
 * 			2,3,4 (using the second 2)
 * 			2,2,3
 *
 * 			Example 2:
 * 			Input: nums = [4,2,3,4]
 * 			Output: 4
 */

func triangleNumber(nums []int) int {
	sort.Ints(nums)

	len := len(nums)
	count := 0

	for i := len - 1; i > 1; i-- {
		n := nums[i]
		left := 0
		right := i - 1

		for left < right {
			if nums[left]+nums[right] > n {
				count += right - left
				right -= 1
			} else {
				left += 1
			}
		}
	}

	return count
}
