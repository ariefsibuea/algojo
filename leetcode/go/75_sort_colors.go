package main

/**
 * LeetCode Problem : Sort Colors
 * Topic            : Array, Two Pointers, Sorting
 * Level            : Medium
 * URL              : https://leetcode.com/problems/sort-colors
 * Description      :
 * Examples         :
 * 			Example 1:
 * 			Input: nums = [2,0,2,1,1,0]
 * 			Output: [0,0,1,1,2,2]
 *
 * 			Example 2:
 * 			Input: nums = [2,0,1]
 * 			Output: [0,1,2]
 */

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	index := 0

	for index <= right {
		switch nums[index] {
		case 0:
			nums[left], nums[index] = nums[index], nums[left]
			left += 1
			index += 1
		case 2:
			nums[right], nums[index] = nums[index], nums[right]
			right -= 1
		default:
			index += 1
		}
	}
}

func sortColors_prevSolution(nums []int) {
	index := 0

	// sort the red
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[index], nums[i] = nums[i], nums[index]
			index += 1
		}
	}

	// sort the white
	for i := index; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[index], nums[i] = nums[i], nums[index]
			index += 1
		}
	}
}
