package main

/**
 * LeetCode Problem : Sort Colors
 * Topic            : Array, Two Pointers, Sorting
 * Level            : Medium
 * URL              : https://leetcode.com/problems/sort-colors
 * Description      : Given an array nums with n objects colored red, white, or blue, sort them in-place so that
 * 						objects of the same color are adjacent, with the colors in the order red, white, and blue. We
 * 						will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
 * 						You must solve this problem without using the library's sort function.
 * Examples         :
 * 					Example 1:
 * 					Input: nums = [2,0,2,1,1,0]
 * 					Output: [0,0,1,1,2,2]
 *
 * 					Example 2:
 * 					Input: nums = [2,0,1]
 * 					Output: [0,1,2]
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

func sortColorsSolutionI(nums []int) {
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
