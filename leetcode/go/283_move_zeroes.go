package main

/**
 * LeetCode Problem : Move Zeroes
 * Topic            : Array, Two Pointers
 * Level            : Easy
 * URL              : https://leetcode.com/problems/move-zeroes
 * Description      :
 * Examples         :
 * 			Example 1:
 * 			Input: nums = [0,1,0,3,12]
 * 			Output: [1,3,12,0,0]
 *
 * 			Example 2:
 * 			Input: nums = [0]
 * 			Output: [0]
 */

func moveZeroes(nums []int) {
	indexNonZero := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[indexNonZero], nums[i] = nums[i], nums[indexNonZero]
			indexNonZero += 1
		}
	}
}
