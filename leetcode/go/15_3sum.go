package main

import "sort"

/**
 * LeetCode Problem : 3Sum
 * Topic            : Array, Two Pointers, Sorting
 * Level            : Medium
 * URL              : https://leetcode.com/problems/3sum
 * Description      :
 * Examples         :
 * 			Example 1:
 * 			Input: nums = [-1,0,1,2,-1,-4]
 * 			Output: [[-1,-1,2],[-1,0,1]]
 * 			Explanation:
 * 			nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
 * 			nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
 * 			nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
 * 			The distinct triplets are [-1,0,1] and [-1,-1,2].
 * 			Notice that the order of the output and the order of the triplets does not matter.
 *
 * 			Example 2:
 * 			Input: nums = [0,1,1]
 * 			Output: []
 * 			Explanation: The only possible triplet does not sum up to 0.
 */

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	len := len(nums)
	result := make([][]int, 0)

	for i := 0; i < len-2; i++ {
		n := nums[i]
		if i > 0 && n == nums[i-1] {
			continue
		}

		left := i + 1
		right := len - 1

		for left < right {
			sumResult := n + nums[left] + nums[right]
			if sumResult < 0 {
				left += 1
			} else if sumResult > 0 {
				right -= 1
			} else {
				result = append(result, []int{n, nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left += 1
				}
				left += 1
			}
		}
	}

	return result
}
